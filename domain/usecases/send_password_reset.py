from django.contrib.auth.models import User

from abc import ABC


class SendPasswordReset(ABC):
    def send(self, user: User):
        ...