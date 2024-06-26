import uuid

from django.db import models


class Wallet(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)

    user = models.ForeignKey(to="auth.User", on_delete=models.CASCADE)
    currency = models.ForeignKey(to="currency.Currency", on_delete=models.CASCADE)

    name = models.TextField(unique=True)
    description = models.TextField(blank=True, default="")

    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self) -> str:
        return "{} ({})".format(self.name, self.id)

    class Meta:
        ordering = ["-updated_at"]
