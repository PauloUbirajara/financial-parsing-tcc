from domain.usecases.send_activate_account import SendActivateAccount

from django.contrib.auth.models import User
from django.core.mail import send_mail
from inspect import cleandoc


class DjangoEmailSendActivateAccount(SendActivateAccount):
    subject = "Financial Parsing - Account Activation"
    message_template = cleandoc("""
        To activate your account in our Financial Parsing application, you just need to confirm through the following link:

        {activation_link}
    """)

    activation_link: str

    def __init__(self, activation_link: str) -> None:
        self.activation_link = activation_link

    def send(self, user: User):
        message = self.message_template.format(
            activation_link=self.activation_link
        )
        send_mail(
            subject=self.subject,
            message=message,
            from_email=None,
            recipient_list=[user.email],
            fail_silently=False,
        )
