from abc import ABC


class SendEmail(ABC):
    def send(self, subject: str, message: str): ...
