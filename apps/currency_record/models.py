from django.db import models

import uuid


class CurrencyRecord(models.Model):
    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False
    )
    created_at = models.DateField(auto_now_add=True)
    updated_at = models.DateField(auto_now=True)

    value = models.FloatField()
    record_date = models.DateField()
