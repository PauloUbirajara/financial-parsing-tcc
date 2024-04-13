from abc import ABC

from django.contrib.auth.models import User


class SendDeleteAccount(ABC):
    def send(self, user: User): ...
