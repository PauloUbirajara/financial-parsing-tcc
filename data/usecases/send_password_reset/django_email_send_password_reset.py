from inspect import cleandoc

from domain.usecases.send_password_reset import SendPasswordReset
from protocols.send_email import SendEmail


class DjangoEmailSendPasswordReset(SendPasswordReset):
    subject = "Financial Parsing - Redefinição de senha"
    message_template = cleandoc(
        """
        Uma solicitação de redefinição de senha foi solicitada.

        Caso não tenha sido você quem solicitou, favor ignorar o e-mail.

        Entre no link a seguir para redefinir a sua senha:

        {reset_link}
    """
    )
    reset_link: str
    send_email: SendEmail

    def __init__(
        self,
        reset_link: str,
        send_email: SendEmail,
    ):
        self.reset_link = reset_link
        self.send_email = send_email

    def send(self):
        message = self.message_template.format(
            reset_link=self.reset_link,
        )
        self.send_email.send(subject=self.subject, message=message)
