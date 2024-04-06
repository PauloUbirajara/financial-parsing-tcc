from rest_framework.exceptions import NotAuthenticated
from apps.currency.models import Currency
from apps.currency_record.models import CurrencyRecord
from apps.currency_record import serializers

from rest_framework import viewsets
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin
from typing import Optional
from http import HTTPStatus


class CurrencyRecordViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        return CurrencyRecord.objects.all()

    def get_currency(self) -> Optional[Currency]:
        currency_id = self.get_parents_query_dict().get('currency')
        currency = Currency.objects.filter(id=currency_id).first()

        return currency
    
    def get_serializer_class(self):
        supported_serializers = {
            "create": serializers.CreateCurrencyRecordSerializer,
            "update": serializers.UpdateCurrencyRecordSerializer,
        }
        serializer_class = supported_serializers.get(self.action, serializers.CurrencyRecordSerializer)

        return serializer_class

    def list(self, request, *args, **kwargs):
        currency = self.get_currency()

        if currency is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        records = self.get_queryset().filter(currency=currency)

        if not records.exists():
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(records, many=True)

        return Response(data=serializer.data)

    def retrieve(self, request, pk, *args, **kwargs):
        currency = self.get_currency()

        if currency is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        record = self.get_queryset().filter(currency=currency, id=pk).first()

        if record is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(record)

        return Response(data=serializer.data)

    def update(self, request, pk, *args, **kwargs):
        currency = self.get_currency()

        if currency is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        record: CurrencyRecord = self.get_queryset().filter(currency=currency,id=pk).first()

        if record is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        record.value = serializer.validated_data.get('value', record.value)
        record.record_date = serializer.validated_data.get('record_date', record.record_date)

        record.save()

        return Response(data=serializer.data)

    def create(self, request, *args, **kwargs):
        currency = self.get_currency()

        if currency is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        record = {
            **serializer.validated_data,
            "currency": currency
        }

        self.get_queryset().create(**record)

        return Response(data=serializer.data)
