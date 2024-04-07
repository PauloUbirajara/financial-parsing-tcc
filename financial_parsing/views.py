from django.contrib.auth.models import User
from django.db.utils import IntegrityError

from financial_parsing.serializers import RegisterSerializer

from rest_framework.views import APIView
from rest_framework.response import Response
from http import HTTPStatus


class RegisterView(APIView):
    def post(self, request):
        if request.user.is_authenticated:
            return Response(status=HTTPStatus.UNAUTHORIZED)

        serializer = RegisterSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        try:
            User.objects.create_user(
                username=serializer.validated_data['username'],
                email=serializer.validated_data['email'],
                password=serializer.validated_data['password'],
                is_active=False
            )
            # TODO Send email about confirmation
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

        return Response(status=HTTPStatus.CREATED)
