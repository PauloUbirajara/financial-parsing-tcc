from inspect import cleandoc

from domain.usecases.send_activate_account import SendActivateAccount
from protocols.send_email import SendEmail


class DjangoEmailSendActivateAccount(SendActivateAccount):
    subject = "Financial Parsing - Ativação de conta"
    message_template = cleandoc(
        """
        Para ativar a sua conta em nossa plataforma Financial Parsing, basta que você confirme através do link a seguir:

        {activation_link}
    """
    )

    activation_link: str
    send_email: SendEmail

    def __init__(self, activation_link: str, send_email: SendEmail) -> None:
        self.activation_link = activation_link
        self.send_email = send_email

    def send(self):
        message = self.message_template.format(activation_link=self.activation_link)
        self.send_email.send(subject=self.subject, message=message)
