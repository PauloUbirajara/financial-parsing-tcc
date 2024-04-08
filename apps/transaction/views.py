from django.apps import apps

from apps.transaction.models import Transaction
from apps.transaction import serializers

from rest_framework.exceptions import NotAuthenticated
from rest_framework import viewsets
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin
from http import HTTPStatus


Wallet = apps.get_model('wallet', 'wallet')


class TransactionViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        return Transaction.objects.filter(user=self.request.user)

    def get_serializer_class(self):
        supported_serializers = {
            "create": serializers.CreateTransactionSerializer,
            "update": serializers.UpdateTransactionSerializer,
        }
        serializer_class = supported_serializers.get(self.action, serializers.TransactionSerializer)

        return serializer_class

    def list(self, request):
        transactions = self.get_queryset()
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(transactions, many=True)
        return Response(data=serializer.data)

    def retrieve(self, request, pk):
        transaction = self.get_queryset().filter(id=pk).first()

        if transaction is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(transaction)

        return Response(data=serializer.data)

    def update(self, request, pk, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)


        transaction: Transaction = self.get_queryset().filter(id=pk).first()

        if transaction is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        transaction.name = serializer.validated_data.get('name', transaction.name)
        transaction.description = serializer.validated_data.get('description', transaction.description)
        transaction.transaction_date = serializer.validated_data.get('transaction_date', transaction.transaction_date)
        transaction.value = serializer.validated_data.get('value', transaction.value)

        if not serializer.validated_data.get('wallet_id'):
            transaction.save()
            return Response(data=serializer.data)

        wallet = (
            Wallet.objects
            .filter(id=serializer.validated_data.get('wallet_id'))
            .first()
        )

        if wallet is None:
            return Response(status=HTTPStatus.BAD_REQUEST)

        transaction.wallet = wallet

        transaction.save()

        return Response(data=serializer.data)

    def create(self, request, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        transaction = {
            **serializer.validated_data,
            "user": request.user
        }
        self.get_queryset().create(**transaction)

        return Response(
            data=serializer.data,
            status=HTTPStatus.CREATED
        )
