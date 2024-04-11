from abc import ABC

from django.contrib.auth.models import User


class SendActivateAccount(ABC):
    def send(self, user: User): ...
