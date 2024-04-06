from django.db import models

import uuid


class Currency(models.Model):
    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False
    )

    user = models.ForeignKey(
        to='auth.User',
        on_delete=models.CASCADE
    )

    name = models.TextField()
    representation = models.TextField()

    created_at = models.DateField(auto_now_add=True)
    updated_at = models.DateField(auto_now=True)
