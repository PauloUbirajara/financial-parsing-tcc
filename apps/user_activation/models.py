from django.db import models

import uuid


class UserActivation(models.Model):
    activation_token = models.UUIDField(
        default=uuid.uuid4,
        editable=False
    )
    user = models.ForeignKey(
        to='auth.User',
        on_delete=models.CASCADE
    )
