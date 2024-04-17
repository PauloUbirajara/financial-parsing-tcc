from datetime import datetime, timedelta, timezone

from django.conf import settings

from apps.user_management.models import UserManagement


def is_token_active(user_management: UserManagement) -> bool:
    created_at: datetime = user_management.created_at
    time_elapsed = datetime.now(tz=timezone.utc) - created_at

    return time_elapsed <= timedelta(
        minutes=settings.ACTIVATION_EXPIRATION_TIME_IN_MINUTES
    )
