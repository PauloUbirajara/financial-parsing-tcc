from apps.currency.models import Currency
from apps.currency.serializer import CurrencySerializer

from rest_framework import viewsets
from rest_framework.response import Response


class CurrencyViewSet(viewsets.ModelViewSet):
    queryset = Currency.objects.all()
    serializer_class = CurrencySerializer
