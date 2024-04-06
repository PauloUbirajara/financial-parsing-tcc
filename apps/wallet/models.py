from django.db import models

import uuid


class Wallet(models.Model):
    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False
    )

    currency = models.ForeignKey(
        to='currency.Currency',
        on_delete=models.CASCADE
    )

    name = models.TextField()
    description = models.TextField()

    created_at = models.DateField(auto_now_add=True)
    updated_at = models.DateField(auto_now=True)
