from apps.currency.models import Currency
from apps.currency_record.models import CurrencyRecord
from apps.currency_record import serializers

from rest_framework import viewsets
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin
from typing import Optional


class CurrencyRecordViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    queryset = CurrencyRecord.objects.all()

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
            return Response(status=404)

        records = self.get_queryset().filter(currency=currency)

        if not records.exists():
            return Response(status=404)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(records, many=True)

        return Response(data=serializer.data)

    def retrieve(self, request, pk, *args, **kwargs):
        currency = self.get_currency()

        if currency is None:
            return Response(status=404)

        record = self.get_queryset().filter(currency=currency, id=pk).first()

        if record is None:
            return Response(status=404)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(record)

        return Response(data=serializer.data)

    def update(self, request, pk, *args, **kwargs):
        currency = self.get_currency()

        if currency is None:
            return Response(status=404)

        record: CurrencyRecord = self.get_queryset().filter(currency=currency,id=pk).first()

        if record is None:
            return Response(status=404)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=400, data=serializer.error_messages)

        record.value = serializer.validated_data.get('value', record.value)
        record.record_date = serializer.validated_data.get('record_date', record.record_date)

        record.save()

        return Response(data=serializer.data)

    def create(self, request, *args, **kwargs):
        currency = self.get_currency()

        if currency is None:
            return Response(status=404)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=400, data=serializer.error_messages)

        record = {
            **serializer.validated_data,
            "currency": currency
        }

        self.get_queryset().create(**record)

        return Response(data=serializer.data)
