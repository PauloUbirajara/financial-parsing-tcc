from django.db import models

import uuid


class CurrencyRecord(models.Model):
    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False
    )

    user = models.ForeignKey(
        to='auth.User',
        on_delete=models.CASCADE
    )
    currency = models.ForeignKey(
        to='currency.Currency',
        on_delete=models.CASCADE
    )

    value = models.FloatField()
    record_date = models.DateField()

    created_at = models.DateField(auto_now_add=True)
    updated_at = models.DateField(auto_now=True)
