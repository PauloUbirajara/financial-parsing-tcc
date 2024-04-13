from abc import ABC


class SendPasswordReset(ABC):
    def send(self): ...
