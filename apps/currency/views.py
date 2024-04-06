from apps.currency.models import Currency
from apps.currency import serializers

import logging
from rest_framework import viewsets
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin


class CurrencyViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    queryset = Currency.objects.all()
    
    def get_serializer_class(self):
        supported_serializers = {
            "create": serializers.CreateCurrencySerializer,
            "update": serializers.UpdateCurrencySerializer,
        }
        serializer_class = supported_serializers.get(self.action, serializers.CurrencySerializer)

        return serializer_class

    def list(self, request):
        currencies = self.queryset.all()
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(currencies, many=True)
        return Response(data=serializer.data)

    def retrieve(self, request, pk):
        currency = self.queryset.filter(id=pk).first()

        if currency is None:
            return Response(status=404)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(currency)

        return Response(data=serializer.data)

    def update(self, request, pk, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=400, data=serializer.error_messages)

        currency: Currency = self.queryset.filter(id=pk).first()

        if currency is None:
            return Response(status=404)

        currency.name = serializer.validated_data.get('name', currency.name)
        currency.representation = serializer.validated_data.get('representation', currency.representation)

        currency.save()

        return Response(data=serializer.data)

    def create(self, request, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=400, data=serializer.error_messages)

        self.queryset.create(**serializer.data)
        return Response(data=serializer.data)
