import smtplib
from email.message import Message

from django.conf import settings
from django.contrib.auth.models import User

from protocols.send_email import SendEmail


class SMTPLibSendEmail(SendEmail):
    user: User

    def __init__(self, user: User) -> None:
        self.user = user

    def send(self, subject: str, message: str):
        # Setup SMTP connection to server
        host = getattr(settings, "EMAIL_HOST")
        server = smtplib.SMTP(host=host, port=25)
        server.starttls()
        server.ehlo()

        # Login
        username = getattr(settings, "EMAIL_HOST_USER")
        password = getattr(settings, "EMAIL_HOST_PASSWORD")
        server.login(user=username, password=password)

        # Setup message object
        msg = Message()
        msg["From"] = username
        msg["To"] = self.user.email
        msg["Subject"] = subject
        msg.set_payload(message, "utf8")

        # Send
        server.send_message(from_addr=username, to_addrs=self.user.email, msg=msg)
