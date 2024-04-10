from domain.usecases.send_password_reset import SendPasswordReset

from django.contrib.auth.models import User
from django.core.mail import send_mail

from typing import Optional


class DjangoEmailSendPasswordReset(SendPasswordReset):
    subject = "Financial Parsing - Reset Password"
    message_template = """
        Your account password was reset.

        New password:  {temporary_password}

        To access the Financial Parsing application, use the following link:
        {application_link}
    """
    application_link: str

    def __init__(self, application_link: Optional[str]) -> None:
        if application_link is None:
            raise ValueError("Missing application link")

        self.application_link = application_link

    def send(self, user: User, temporary_password: str):
        message = self.message_template.format(
            temporary_password=temporary_password,
            application_link=self.application_link
        )
        send_mail(
            subject=self.subject,
            message=message,
            from_email=None,
            recipient_list=[user.email],
            fail_silently=False,
        )
