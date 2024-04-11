from abc import ABC

from django.contrib.auth.models import User


class SendPasswordReset(ABC):
    def send(self, user: User): ...
