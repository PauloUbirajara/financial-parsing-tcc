from inspect import cleandoc
from typing import Optional

from domain.usecases.send_password_reset import SendPasswordReset
from protocols.send_email import SendEmail


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
    send_email: SendEmail

    def __init__(
        self,
        application_link: Optional[str],
        temporary_password: str,
        send_email: SendEmail,
    ) -> None:
        if application_link is None:
            raise ValueError("Missing application link")

        self.application_link = application_link
        self.temporary_password = temporary_password
        self.send_email = send_email

    def send(self):
        message = self.message_template.format(
            temporary_password=self.temporary_password,
            application_link=self.application_link,
        )
        self.send_email.send(subject=self.subject, message=message)
