from apps.currency.models import Currency
from apps.currency.serializers import CurrencySerializer

from rest_framework import viewsets


class CurrencyViewSet(viewsets.ModelViewSet):
    queryset = Currency.objects.all()
    serializer_class = CurrencySerializer
