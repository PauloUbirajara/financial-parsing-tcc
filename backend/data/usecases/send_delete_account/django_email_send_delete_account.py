from inspect import cleandoc

from domain.usecases.send_delete_account import SendDeleteAccount
from protocols.send_email import SendEmail


class DjangoEmailSendDeleteAccount(SendDeleteAccount):
    subject = "Financial Parsing - Remoção definitiva de conta"
    message_template = cleandoc(
        """
        Para apagar permanentemente a sua conta de nossa plataforma Financial Parsing, basta que você confirme através do link a seguir:

        {deletion_link}
    """
    )

    deletion_link: str
    send_email: SendEmail

    def __init__(self, deletion_link: str, send_email: SendEmail) -> None:
        self.deletion_link = deletion_link
        self.send_email = send_email

    def send(self):
        message = self.message_template.format(deletion_link=self.deletion_link)
        self.send_email.send(subject=self.subject, message=message)
