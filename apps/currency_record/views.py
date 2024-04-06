from apps.currency_record.models import CurrencyRecord
from apps.currency_record.serializers import CurrencyRecordSerializer

from rest_framework import viewsets
from rest_framework_extensions.mixins import NestedViewSetMixin


class CurrencyRecordViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    queryset = CurrencyRecord.objects.all()
    serializer_class = CurrencyRecordSerializer

    def list(self, request, *args, **kwargs):
        return super().list(request, *args, **kwargs)

    def retrieve(self, request, *args, **kwargs):
        return super().retrieve(request, *args, **kwargs)

    def update(self, request, *args, **kwargs):
        return super().update(request, *args, **kwargs)

    def create(self, request, *args, **kwargs):
        return super().create(request, *args, **kwargs)
