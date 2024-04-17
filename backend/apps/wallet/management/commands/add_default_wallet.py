import logging

from django.apps import apps
from django.contrib.auth.models import User
from django.core.management.base import BaseCommand

from apps.wallet.models import Wallet

Currency = apps.get_model("currency", "Currency")


def create_default_wallet(user: User):
    default_currency = Currency.objects.filter(representation="USD").first()
    try:
        Wallet.objects.create(
            name="Geral",
            description="Esta carteira serve para qualquer transação",
            currency=default_currency,
            user=user,
        )
    except Exception as err:
        logging.warning("Could not create default wallet")
        logging.warning(str(err))


class Command(BaseCommand):
    help = "Adds default wallet for a specific user"

    def add_arguments(self, parser):
        parser.add_argument(
            "user_id", type=int, help="Specify which user to create default wallet"
        )

    def handle(self, *args, **kwargs):
        user = User.objects.filter(id=kwargs["user_id"]).first()
        if user is None:
            raise Exception("Could not find user")

        if Wallet.objects.filter(user=user).count() > 0:
            logging.warning("User already has one or more wallets")
            return

        create_default_wallet(user)
