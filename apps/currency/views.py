from rest_framework.exceptions import NotAuthenticated
from apps.currency.models import Currency
from apps.currency import serializers

from rest_framework import viewsets
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin
from http import HTTPStatus


class CurrencyViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        return Currency.objects.all()
    
    def get_serializer_class(self):
        supported_serializers = {
            "create": serializers.CreateCurrencySerializer,
            "update": serializers.UpdateCurrencySerializer,
        }
        serializer_class = supported_serializers.get(self.action, serializers.CurrencySerializer)

        return serializer_class

    def list(self, request):
        currencies = self.get_queryset()
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(currencies, many=True)
        return Response(data=serializer.data)

    def retrieve(self, request, pk):
        currency = self.get_queryset().filter(id=pk).first()

        if currency is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(currency)

        return Response(data=serializer.data)

    def update(self, request, pk, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        currency: Currency = self.get_queryset().filter(id=pk).first()

        if currency is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        currency.name = serializer.validated_data.get('name', currency.name)
        currency.representation = serializer.validated_data.get('representation', currency.representation)

        currency.save()

        return Response(data=serializer.data)

    def create(self, request, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        Currency.objects.create(**serializer.validated_data)

        return Response(data=serializer.data)
