from inspect import cleandoc

from django.contrib.auth.models import User
from django.core.mail import send_mail

from domain.usecases.send_delete_account import SendDeleteAccount


class DjangoEmailSendDeleteAccount(SendDeleteAccount):
    subject = "Financial Parsing - Account Deletion"
    message_template = cleandoc(
        """
        To permanently delete your account from our Financial Parsing application, you just need to confirm through the following link:

        {deletion_link}
    """
    )

    deletion_link: str

    def __init__(self, deletion_link: str) -> None:
        self.deletion_link = deletion_link

    def send(self, user: User):
        message = self.message_template.format(deletion_link=self.deletion_link)
        send_mail(
            subject=self.subject,
            message=message,
            from_email=None,
            recipient_list=[user.email],
            fail_silently=False,
        )
