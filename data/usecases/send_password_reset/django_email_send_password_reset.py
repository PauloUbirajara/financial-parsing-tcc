from inspect import cleandoc
from typing import Optional

from django.contrib.auth.models import User
from django.core.mail import send_mail

from domain.usecases.send_password_reset import SendPasswordReset


class DjangoEmailSendPasswordReset(SendPasswordReset):
    subject = "Financial Parsing - Reset Password"
    message_template = cleandoc(
        """
        Your account password was reset.

        New password:  {temporary_password}

        To access the Financial Parsing application, use the following link:
        {application_link}
    """
    )
    application_link: str
    temporary_password: str

    def __init__(
        self, application_link: Optional[str], temporary_password: str
    ) -> None:
        if application_link is None:
            raise ValueError("Missing application link")

        self.application_link = application_link
        self.temporary_password = temporary_password

    def send(self, user: User):
        message = self.message_template.format(
            temporary_password=self.temporary_password,
            application_link=self.application_link,
        )
        send_mail(
            subject=self.subject,
            message=message,
            from_email=None,
            recipient_list=[user.email],
            fail_silently=False,
        )
