import uuid

from django.db import models


class UserManagement(models.Model):
    token = models.UUIDField(default=uuid.uuid4, editable=False)
    user = models.ForeignKey(to="auth.User", on_delete=models.CASCADE)
    created_at = models.DateTimeField(auto_now_add=True)
