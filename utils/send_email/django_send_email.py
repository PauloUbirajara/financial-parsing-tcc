from django.contrib.auth.models import User
from django.core.mail import send_mail

from protocols.send_email import SendEmail


class DjangoSendEmail(SendEmail):
    user: User

    def __init__(self, user: User) -> None:
        self.user = user

    def send(self, subject: str, message: str):
        send_mail(
            subject=subject,
            message=message,
            from_email=None,
            recipient_list=[self.user.email],
            fail_silently=False,
        )
