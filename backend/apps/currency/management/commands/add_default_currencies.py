import logging

from django.contrib.auth.models import User
from django.core.management.base import BaseCommand

from apps.currency.models import Currency


def create_default_currencies(user: User):
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
                name=name, representation=representation, user=user
            )
        except Exception as err:
            logging.warning(
                "Could not create default currency: {} ({})".format(
                    name, representation
                )
            )
            logging.warning(str(err))


class Command(BaseCommand):
    help = "Adds default currencies for a specific user"

    def add_arguments(self, parser):
        parser.add_argument(
            "user_id", type=int, help="Specify which user to create default currencies"
        )

    def handle(self, *args, **kwargs):
        user = User.objects.filter(id=kwargs["user_id"]).first()
        if user is None:
            raise Exception("Could not find user")

        create_default_currencies(user)
