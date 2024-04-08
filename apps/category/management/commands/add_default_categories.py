from django.core.management.base import BaseCommand
from django.apps import apps
from django.contrib.auth.models import User

from apps.category.models import Category
import logging


Currency = apps.get_model('currency', 'Currency')


def create_default_categories(user: User):
    default_categories = (
        "Income",
        "Expense",
        "Budget",
        "Investment",
        "Debt",
        "Recurrent",
        "Transfer",
        "Miscellaneous",
    )
    for category in default_categories:
        try:
            Category.objects.create(
                name=category,
                user=user
            )
        except Exception as err:
            logging.warning("Could not create default category: {}".format(category))
            logging.warning(str(err))


class Command(BaseCommand):
    help = 'Adds default categories for a specific user'

    def add_arguments(self, parser):
        parser.add_argument(
            'user_id',
            type=int,
            help='Specify which user to create default categories'
        )

    def handle(self, *args, **kwargs):
        user = User.objects.filter(id=kwargs['user_id']).first()
        if user is None:
            raise Exception("Could not find user")

        if Category.objects.filter(user=user).count() > 0:
            logging.warning("User already has one or more categories")
            return

        create_default_categories(user)
