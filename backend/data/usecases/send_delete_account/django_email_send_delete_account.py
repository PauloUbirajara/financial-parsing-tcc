from inspect import cleandoc

from domain.usecases.send_delete_account import SendDeleteAccount
from protocols.send_email import SendEmail


class DjangoEmailSendDeleteAccount(SendDeleteAccount):
    subject = "Financial Parsing - Account Deletion"
    message_template = cleandoc(
        """
        To permanently delete your account from our Financial Parsing application, you just need to confirm through the following link:

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
