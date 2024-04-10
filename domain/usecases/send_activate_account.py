from django.contrib.auth.models import User

from abc import ABC


class SendActivateAccount(ABC):
    def send(self, user: User):
        ...
