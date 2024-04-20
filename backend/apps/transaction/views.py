from http import HTTPStatus

from django.apps import apps
from rest_framework import viewsets
from rest_framework.decorators import action
from rest_framework.exceptions import NotAuthenticated
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin

from apps.transaction import serializers
from apps.transaction.models import Transaction
from domain.models.bulk_delete_serializer import BulkDeleteSerializer
from domain.models.model_pagination import ModelPagination

Wallet = apps.get_model("wallet", "wallet")


class TransactionViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    pagination_class = ModelPagination

    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        queryset = Transaction.objects.filter(user=self.request.user)
        return queryset

    def get_serializer_class(self):
        supported_serializers = {
            "list": serializers.ListTransactionSerializer,
            "bulk_delete": BulkDeleteSerializer,
        }
        serializer_class = supported_serializers.get(
            self.action, serializers.get_transaction_serializer(user=self.request.user)
        )
        return serializer_class

    def list(self, request, *args, **kwargs):
        queryset = self.get_queryset()
        paginated_queryset = self.paginate_queryset(queryset)
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(paginated_queryset, many=True)
        return self.paginator.get_paginated_response(serializer.data)

    def retrieve(self, request, pk):
        transaction = self.get_queryset().filter(id=pk).first()

        if transaction is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(transaction)

        return Response(data=serializer.data)

    def update(self, request, pk, *args, **kwargs):
        transaction: Transaction = self.get_queryset().filter(id=pk).first()

        if transaction is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(transaction, data=request.data, partial=True)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        transaction.name = serializer.validated_data.get("name", transaction.name)
        transaction.description = serializer.validated_data.get(
            "description", transaction.description
        )
        transaction.transaction_date = serializer.validated_data.get(
            "transaction_date", transaction.transaction_date
        )
        transaction.value = serializer.validated_data.get("value", transaction.value)
        transaction.categories.set(
            serializer.validated_data.get("categories", transaction.categories)
        )
        transaction.wallet = serializer.validated_data.get("wallet", transaction.wallet)

        transaction.save()

        return Response(data=serializer.data)

    def create(self, request, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        # Create transaction with no categories
        transaction = Transaction(
            **{
                "name": serializer.validated_data["name"],
                "description": serializer.validated_data["description"],
                "transaction_date": serializer.validated_data["transaction_date"],
                "wallet": serializer.validated_data["wallet"],
                "value": serializer.validated_data["value"],
                "user": request.user,
            }
        )
        transaction.save()

        # Add categories once the transaction was created due to ID
        transaction.categories.set(serializer.validated_data["categories"])
        transaction.save()

        return Response(data=serializer.data, status=HTTPStatus.CREATED)

    @action(methods=["POST"], detail=False)
    def bulk_delete(self, request, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        self.get_queryset().filter(
            id__in=serializer.validated_data.get("ids", [])
        ).delete()

        return Response(status=HTTPStatus.OK)
