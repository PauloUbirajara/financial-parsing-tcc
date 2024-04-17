from inspect import cleandoc
from typing import Optional

from domain.usecases.send_password_reset import SendPasswordReset
from protocols.send_email import SendEmail


class DjangoEmailSendPasswordReset(SendPasswordReset):
    subject = "Financial Parsing - Redefinição de senha"
    message_template = cleandoc(
        """
        Uma solicitação de redefinição de senha foi solicitada.

        Caso não tenha sido você quem solicitou, favor ignorar o e-mail.

        Entre no link a seguir para redefinir a sua senha:

        {application_link}
    """
    )
    application_link: str
    send_email: SendEmail

    def __init__(
        self,
        application_link: Optional[str],
        send_email: SendEmail,
    ) -> None:
        if application_link is None:
            raise ValueError("Missing application link")

        self.application_link = application_link
        self.send_email = send_email

    def send(self):
        message = self.message_template.format(
            application_link=self.application_link,
        )
        self.send_email.send(subject=self.subject, message=message)
