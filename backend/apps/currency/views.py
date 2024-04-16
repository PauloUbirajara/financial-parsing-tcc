from http import HTTPStatus

from rest_framework import viewsets
from rest_framework.exceptions import NotAuthenticated
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin

from apps.currency import serializers
from apps.currency.models import Currency


class CurrencyViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    serializer_class = serializers.CurrencySerializer

    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        return Currency.objects.filter(user=self.request.user)

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
