import re

from django.contrib.auth.password_validation import validate_password
from django.utils.translation import gettext as _
from rest_framework import serializers


class UserRegistrationSerializer(serializers.Serializer):
    username = serializers.CharField(required=True)
    email = serializers.EmailField(required=True)
    password = serializers.CharField(required=True)

    def validate(self, attrs):
        username = attrs.get("username")
        password = attrs.get("password")

        # Validate username
        if not re.match(r"^[a-zA-Z0-9-]{3,32}$", username):
            raise serializers.ValidationError(
                _(
                    "Apelido deve conter entre 3 até 32 caracteres alfanuméricos, separados por hífens."
                )
            )

        # Validate password
        validate_password(password)

        return attrs


class UserPasswordResetSerializer(serializers.Serializer):
    password = serializers.CharField(required=True)

    def validate(self, attrs):
        password = attrs.get("password")

        validate_password(password)
