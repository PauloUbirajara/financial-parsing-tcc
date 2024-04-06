from apps.currency_record.models import CurrencyRecord
from apps.currency_record.serializers import CurrencyRecordSerializer

from rest_framework import viewsets


class CurrencyRecordViewSet(viewsets.ModelViewSet):
    queryset = CurrencyRecord.objects.all()
    serializer_class = CurrencyRecordSerializer
