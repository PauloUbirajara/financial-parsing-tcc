from django.apps import apps

from apps.wallet.models import Wallet
from apps.wallet import serializers

from rest_framework.exceptions import NotAuthenticated
from rest_framework import viewsets
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin
from http import HTTPStatus


Currency = apps.get_model('currency', 'Currency')


class WalletViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        return Wallet.objects.filter(user=self.request.user)
    
    def get_serializer_class(self):
        supported_serializers = {
            "create": serializers.CreateWalletSerializer,
            "update": serializers.UpdateWalletSerializer,
        }
        serializer_class = supported_serializers.get(self.action, serializers.WalletSerializer)

        return serializer_class

    def list(self, request):
        wallets = self.get_queryset()
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(wallets, many=True)
        return Response(data=serializer.data)

    def retrieve(self, request, pk):
        wallet = self.get_queryset().filter(id=pk).first()

        if wallet is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(wallet)

        return Response(data=serializer.data)

    def update(self, request, pk, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)


        wallet: Wallet = self.get_queryset().filter(id=pk).first()

        if wallet is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        wallet.name = serializer.validated_data.get('name', wallet.name)
        wallet.description = serializer.validated_data.get('description', wallet.description)

        if not serializer.validated_data.get('currency_id'):
            wallet.save()
            return Response(data=serializer.data)

        currency = (
            Currency.objects
            .filter(id=serializer.validated_data.get('currency_id'))
            .first()
        )

        if currency is None:
            return Response(status=HTTPStatus.BAD_REQUEST)

        wallet.currency = currency

        wallet.save()

        return Response(data=serializer.data)

    def create(self, request, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        currency = (
            Currency.objects
            .filter(id=serializer.validated_data.get('currency_id'))
            .first()
        )

        if currency is None:
            return Response(status=HTTPStatus.BAD_REQUEST)

        wallet = {
            **serializer.validated_data,
            "currency": currency,
            "user": request.user
        }
        self.get_queryset().create(**wallet)

        return Response(
            data=serializer.data,
            status=HTTPStatus.CREATED
        )
