import logging

from django.contrib.auth.models import User
from django.core.management.base import BaseCommand

from apps.category.models import Category


def create_default_categories(user: User):
    default_categories = (
        "Renda",
        "Despesa",
        "Alimentação",
        "Lazer",
        "Compras",
        "Outros",
    )
    for category in default_categories:
        try:
            Category.objects.get_or_create(name=category, user=user)
        except Exception as err:
            logging.warning("Could not create default category: {}".format(category))
            logging.warning(str(err))


class Command(BaseCommand):
    help = "Adds default categories for a specific user"

    def add_arguments(self, parser):
        parser.add_argument(
            "user_id", type=int, help="Specify which user to create default categories"
        )

    def handle(self, *args, **kwargs):
        user = User.objects.filter(id=kwargs["user_id"]).first()
        if user is None:
            raise Exception("Could not find user")

        create_default_categories(user)
