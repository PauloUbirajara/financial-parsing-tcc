import os
import uuid
from datetime import datetime, timedelta, timezone
from http import HTTPStatus
from urllib.parse import urljoin

from django.conf import settings
from django.contrib.auth import logout
from django.contrib.auth.models import User
from django.core.management import call_command
from django.db.utils import IntegrityError
from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.reverse import reverse
from rest_framework.views import APIView

from apps.user_management.models import UserManagement
from apps.user_management.serializers import (
    UserManagementChangeEmailSerializer,
    UserManagementChangePasswordSerializer,
    UserManagementRegisterSerializer,
    UserManagementResendActivationSerializer,
    UserManagementResendDeletionSerializer,
    UserManagementResetPasswordSerializer,
)
from data.usecases.send_activate_account.django_email_send_activate_account import (
    DjangoEmailSendActivateAccount,
)
from data.usecases.send_delete_account.django_email_send_delete_account import (
    DjangoEmailSendDeleteAccount,
)
from data.usecases.send_password_reset.django_email_send_password_reset import (
    DjangoEmailSendPasswordReset,
)
from domain.usecases.send_activate_account import SendActivateAccount
from domain.usecases.send_delete_account import SendDeleteAccount
from domain.usecases.send_password_reset import SendPasswordReset
from protocols.send_email import SendEmail
from utils.send_email.django_send_email import DjangoSendEmail


def setup_account_activation(user: User, request: Request) -> Response:
    # Create/Update user activation
    user_management, _ = UserManagement.objects.get_or_create(user=user)
    user_management.token = str(uuid.uuid4())
    user_management.created_at = datetime.now(tz=timezone.utc)
    user_management.save()

    # Send activation link to email
    send_email: SendEmail = DjangoSendEmail(user=user)

    send_activate_account: SendActivateAccount = DjangoEmailSendActivateAccount(
        activation_link=urljoin(
            request.build_absolute_uri("/"),
            reverse("user-activate", kwargs={"token": user_management.token}),
        ),
        send_email=send_email,
    )
    send_activate_account.send()

    return Response(status=HTTPStatus.OK)


def setup_account_deletion(user: User, request: Request) -> Response:
    # Create/Update user activation
    user_management, _ = UserManagement.objects.get_or_create(user=user)
    user_management.token = str(uuid.uuid4())
    user_management.created_at = datetime.now(tz=timezone.utc)
    user_management.save()

    # Send deletion link to email
    send_email: SendEmail = DjangoSendEmail(user=user)

    send_delete_account: SendDeleteAccount = DjangoEmailSendDeleteAccount(
        deletion_link=urljoin(
            request.build_absolute_uri("/"),
            reverse("user-delete-account", kwargs={"token": user_management.token}),
        ),
        send_email=send_email,
    )
    send_delete_account.send()

    return Response(status=HTTPStatus.OK)


def setup_password_reset(user: User) -> Response:
    # Change user password
    temporary_password = str(uuid.uuid4())
    user.set_password(temporary_password)
    user.save()

    # Send password reset email
    send_email: SendEmail = DjangoSendEmail(user=user)
    send_password_reset: SendPasswordReset = DjangoEmailSendPasswordReset(
        application_link=os.getenv("FRONTEND_LOGIN_URL"),
        temporary_password=temporary_password,
        send_email=send_email,
    )
    send_password_reset.send()

    return Response(status=HTTPStatus.OK)


class UserManagementResendActivationView(APIView):
    def post(self, request):
        if request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        if request.user.is_active:
            data = {"message": "User already active"}
            return Response(status=HTTPStatus.OK, data=data)

        serializer = UserManagementResendActivationSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        user = User.objects.filter(
            username=serializer.validated_data["username"]
        ).first()

        if user is None:
            error = {"error": "Could not find registered user"}
            return Response(status=HTTPStatus.NOT_FOUND, data=error)

        return setup_account_activation(user=user, request=request)


class UserManagementResendDeletionView(APIView):
    def post(self, request):
        if not request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        if request.user is None:
            data = {"message": "User already deleted"}
            return Response(status=HTTPStatus.OK, data=data)

        serializer = UserManagementResendDeletionSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        return setup_account_deletion(user=request.user, request=request)


class UserManagementRegisterView(APIView):
    def post(self, request):
        if request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        serializer = UserManagementRegisterSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        try:
            User.objects.create_user(
                username=serializer.validated_data["username"],
                email=serializer.validated_data["email"],
                password=serializer.validated_data["password"],
                is_active=False,
            )
        except IntegrityError as e:
            error = {
                "error": e.__str__(),
            }
            return Response(status=HTTPStatus.CONFLICT, data=error)

        except Exception as e:
            error = {
                "error": e.__str__(),
            }
            return Response(status=HTTPStatus.INTERNAL_SERVER_ERROR, data=error)

        return Response(status=HTTPStatus.OK)


class UserManagementConfirmView(APIView):
    def get(self, request, token=None):
        if request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        if token is None:
            return Response(status=HTTPStatus.BAD_REQUEST)

        pending_user_management = UserManagement.objects.filter(token=token).first()
        if pending_user_management is None:
            error = {"error": "There were no pending account activation requests"}
            return Response(status=HTTPStatus.NOT_FOUND, data=error)

        if pending_user_management.user.is_active:
            error = {"error": "User already active"}
            return Response(status=HTTPStatus.CONFLICT, data=error)

        # Check if activation is not expired
        if datetime.now(
            tz=timezone.utc
        ) - pending_user_management.created_at > timedelta(
            minutes=settings.ACTIVATION_EXPIRATION_TIME_IN_MINUTES
        ):
            error = {"error": "Activation link expired, please request a new one"}
            return Response(status=HTTPStatus.GONE, data=error)

        # Activate user
        pending_user = pending_user_management.user
        pending_user.is_active = True
        pending_user.save()

        # Create default models
        # TODO ADD TRANSACTION TO AVOID ISSUES ON SOME POINTS AND HAVING IT DOING CHANGES
        call_command("add_default_currencies", pending_user.id)
        call_command("add_default_wallet", pending_user.id)
        call_command("add_default_categories", pending_user.id)

        # Remove activation
        pending_user_management.delete()

        return Response(status=HTTPStatus.OK)


class UserManagementDeleteView(APIView):
    def get(self, request, token=None):
        if not request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        if token is None:
            return Response(status=HTTPStatus.BAD_REQUEST)

        pending_user_management = UserManagement.objects.filter(token=token).first()
        if pending_user_management is None:
            error = {"error": "There were no pending deletion requests"}
            return Response(status=HTTPStatus.NOT_FOUND, data=error)

        if pending_user_management.user is None:
            error = {"error": "User already deleted"}
            return Response(status=HTTPStatus.CONFLICT, data=error)

        # Check if deletion is not expired
        if datetime.now(
            tz=timezone.utc
        ) - pending_user_management.created_at > timedelta(
            minutes=settings.DELETION_EXPIRATION_TIME_IN_MINUTES
        ):
            error = {"error": "Deletion link expired, please request a new one"}
            return Response(status=HTTPStatus.GONE, data=error)

        # Delete user
        pending_user_management.user.delete()
        logout(request)

        # Remove deletion request
        pending_user_management.delete()

        return Response(status=HTTPStatus.OK)


class UserManagementChangeEmailView(APIView):
    def post(self, request):
        if not request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        serializer = UserManagementChangeEmailSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        if request.user.email != serializer.validated_data["old_email"]:
            error = {"error": "Invalid email"}
            return Response(status=HTTPStatus.BAD_REQUEST, data=error)

        request.user.email = serializer.validated_data["new_email"]
        request.user.save()

        return Response(status=HTTPStatus.OK)


class UserManagementChangePasswordView(APIView):
    def post(self, request):
        if not request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        serializer = UserManagementChangePasswordSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        if not request.user.check_password(serializer.validated_data["old_password"]):
            error = {"error": "Invalid password"}
            return Response(status=HTTPStatus.BAD_REQUEST, data=error)

        request.user.set_password(serializer.validated_data["new_password"])
        request.user.save()

        return Response(status=HTTPStatus.OK)


class UserManagementResetPasswordView(APIView):
    def post(self, request):
        if request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        serializer = UserManagementResetPasswordSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        user = User.objects.filter(
            username=serializer.validated_data["username"]
        ).first()

        if user is None:
            error = {"error": "Invalid username"}
            return Response(status=HTTPStatus.BAD_REQUEST, data=error)

        return setup_password_reset(user=user)
