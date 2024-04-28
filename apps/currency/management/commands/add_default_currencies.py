import logging

from apps.currency.models import Currency
from django.contrib.auth.models import User
from django.core.management.base import BaseCommand


def create_default_currencies():
    # Values from py-money python package
    default_currencies_dict = {
        "ARS": "Peso Argentino",
        "AUD": "Dólar Australiano",
        "BRL": "Real Brasileiro",
        "CAD": "Dólar Canadense",
        "EUR": "Euro",
        "GBP": "Libra Esterlina",
        "JPY": "Iene",
        "MXN": "Peso Mexicano",
        "RUB": "Rublo Russo",
        "USD": "Dólar Americano",
    }

    for representation, name in default_currencies_dict.items():
        try:
            Currency.objects.get_or_create(
                name=name, representation=representation
            )
        except Exception as err:
            logging.warning(
                "Could not create default currency: {} ({})".format(
                    name, representation
                )
            )
            logging.warning(str(err))


class Command(BaseCommand):
    help = "Adds default currencies"

    def handle(self, *args, **kwargs):
        create_default_currencies()
