from django.contrib.auth.models import User
from django.db.utils import IntegrityError
from django.core.mail import send_mail
from django.core.management import call_command
from django.conf import settings

from apps.user_activation.models import UserActivation
from apps.user_activation.serializers import UserActivationRegisterSerializer, UserActivationResendActivationSerializer

from rest_framework.views import APIView
from rest_framework.reverse import reverse
from rest_framework.response import Response
from rest_framework.request import Request
from http import HTTPStatus
from urllib.parse import urljoin
from datetime import datetime, timedelta, timezone
import uuid


def setup_account_activation(user: User, request: Request) -> Response:
    # Create/Update user activation
    user_activation, _ = UserActivation.objects.get_or_create(user=user)
    user_activation.activation_token = str(uuid.uuid4())
    user_activation.created_at = datetime.now(tz=timezone.utc)
    user_activation.save()

    # Setup link for activating
    activation_link = urljoin(
        request.build_absolute_uri('/'),
        reverse(
            'user-activate',
            kwargs={
                "token": user_activation.activation_token
            }
        ),
    )
    send_mail(
        subject="Financial Parsing - Account Activation",
        message="Use this link to active your account on the Financial Parsing platform:\n{}".format(activation_link),
        from_email=None,
        recipient_list=[user.email],
        fail_silently=False,
    )
    return Response(status=HTTPStatus.OK)


class UserActivationResendActivationView(APIView):
    def post(self, request):
        if request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        if request.user.is_active:
            data = {"message": "User already active"}
            return Response(status=HTTPStatus.OK, data=data)

        serializer = UserActivationResendActivationSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        user = User.objects.filter(username=serializer.validated_data['username']).first()

        if user is None:
            error = {"error": "Could not find registered user"}
            return Response(status=HTTPStatus.NOT_FOUND, data=error)

        return setup_account_activation(user=user, request=request)


class UserActivationRegisterView(APIView):
    def post(self, request):
        if request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        serializer = UserActivationRegisterSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        try:
            User.objects.create_user(
                username=serializer.validated_data['username'],
                email=serializer.validated_data['email'],
                password=serializer.validated_data['password'],
                is_active=False
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


class UserActivationConfirmView(APIView):
    def get(self, request, token=None):
        if request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        if token is None:
            return Response(status=HTTPStatus.BAD_REQUEST)

        pending_user_activation = (
            UserActivation.objects
            .filter(activation_token=token)
            .first()
        )
        if pending_user_activation is None:
            error = {"error": "There were no pending account activation requests"}
            return Response(status=HTTPStatus.NOT_FOUND, data=error)

        if pending_user_activation.user.is_active:
            error = {"error": "User already active"}
            return Response(status=HTTPStatus.CONFLICT, data=error)

        # Check if activation is not expired
        if datetime.now(tz=timezone.utc) - pending_user_activation.created_at > timedelta(minutes=settings.ACTIVATION_EXPIRATION_TIME_IN_MINUTES):
            error = {"error": "Activation link expired, please request a new one"}
            return Response(status=HTTPStatus.GONE, data=error)

        # Activate user
        pending_user = pending_user_activation.user
        pending_user.is_active = True
        pending_user.save()

        # Create default currencies
        call_command('apps.currency.add_default_currencies')

        # Remove activation
        pending_user_activation.delete()

        return Response(status=HTTPStatus.OK)
