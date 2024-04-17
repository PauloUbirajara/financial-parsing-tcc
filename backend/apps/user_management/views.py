from datetime import datetime, timedelta, timezone

from django.conf import settings
from django.contrib.auth import get_user_model
from django.contrib.auth.hashers import make_password
from django.contrib.auth.models import User
from django.core.management import call_command
from django.utils import timezone
from rest_framework import status
from rest_framework.response import Response
from rest_framework.reverse import reverse
from rest_framework.views import APIView

from apps.user_management.helpers import is_active
from apps.user_management.models import UserManagement
from apps.user_management.serializers import UserRegistrationSerializer
from data.usecases.send_activate_account.django_email_send_activate_account import (
    DjangoEmailSendActivateAccount,
)
from domain.usecases.send_activate_account import SendActivateAccount
from protocols.send_email import SendEmail
from utils.send_email.django_send_email import DjangoSendEmail

User = get_user_model()


class UserActivationView(APIView):
    def get(self, request, activation_token: str = None):
        user_management = UserManagement.objects.filter(token=activation_token).first()
        if user_management is None:
            return Response(
                {"error": "Token de ativação de conta inválido."},
                status=status.HTTP_400_BAD_REQUEST,
            )
        if not is_active(user_management):
            return Response(
                {"error": "Token de ativação de conta expirado."},
                status=status.HTTP_400_BAD_REQUEST,
            )

        # Activate user account and remove token
        user = user_management.user
        user.is_active = True
        user.save()
        user_management.delete()

        # Create default categories and wallet per user
        call_command("add_default_categories", user.id)
        call_command("add_default_wallet", user.id)

        return Response(
            {"detail": "Ativação de conta realizada com sucesso."},
            status=status.HTTP_200_OK,
        )


class UserRegistrationView(APIView):
    def post(self, request):
        serializer = UserRegistrationSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(data=serializer.errors, status=status.HTTP_400_BAD_REQUEST)

        username = serializer.validated_data.get("username")
        email = serializer.validated_data.get("email")
        # Check if username or email already exists
        if User.objects.filter(username=username).exists():
            return Response(
                {"error": "Apelido já cadastrado."},
                status=status.HTTP_400_BAD_REQUEST,
            )

        if User.objects.filter(email=email).exists():
            return Response(
                {"error": "E-mail já cadastrado."}, status=status.HTTP_400_BAD_REQUEST
            )

        # Create the user
        password = make_password(serializer.validated_data.get("password"))
        user = User.objects.create(username=username, email=email, password=password)

        # Send activation email
        user_management = UserManagement.objects.create(user=user)
        activation_token = user_management.token
        activation_link = request.build_absolute_uri(
            reverse("user-activate", kwargs={"activation_token": activation_token})
        )

        send_email: SendEmail = DjangoSendEmail(user=user)
        send_activate_account: SendActivateAccount = DjangoEmailSendActivateAccount(
            activation_link=activation_link, send_email=send_email
        )
        send_activate_account.send()

        return Response(
            {
                "detail": "Usuário cadastrado com sucesso. Confirme a sua conta através do link que enviamos para o seu e-mail."
            },
            status=status.HTTP_201_CREATED,
        )
