import logging

from apps.currency.models import Currency
from django.contrib.auth.models import User
from django.core.management.base import BaseCommand


def create_default_currencies(user: User):
    # Values from py-money python package
    default_currencies_dict = {
        "ARS": "Argentine Peso",
        "AUD": "Australian Dollar",
        "BRL": "Brazilian Real",
        "CAD": "Canadian Dollar",
        "EUR": "Euro",
        "GBP": "Pound Sterling",
        "JPY": "Yen",
        "MXN": "Mexican Peso",
        "RUB": "Russian Ruble",
        "USD": "US Dollar",
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
