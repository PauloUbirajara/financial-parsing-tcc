from django.db import models

import uuid


class Transaction(models.Model):
    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False
    )

    user = models.ForeignKey(
        to='auth.User',
        on_delete=models.CASCADE
    )
    wallet = models.ForeignKey(
        to='wallet.Wallet',
        on_delete=models.CASCADE
    )

    name = models.TextField()
    description = models.TextField(blank=True, default="")
    transaction_date = models.DateField()
    value = models.FloatField()

    created_at = models.DateField(auto_now_add=True)
    updated_at = models.DateField(auto_now=True)
