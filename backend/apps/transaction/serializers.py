from typing import Optional

from django.apps import apps
from django.contrib.auth.models import User
from rest_framework import serializers

from apps.category.serializers import CategorySerializer
from apps.wallet.serializers import WalletSerializer

Transaction = apps.get_model("transaction", "Transaction")
Wallet = apps.get_model("wallet", "Wallet")
Category = apps.get_model("category", "Category")


def get_transaction_serializer(user: User):
    class TransactionSerializer(serializers.ModelSerializer):
        id = serializers.UUIDField(read_only=True)
        wallet = serializers.PrimaryKeyRelatedField(
            required=True, queryset=Wallet.objects.filter(user=user)
        )
        categories = serializers.PrimaryKeyRelatedField(
            required=True, queryset=Category.objects.filter(user=user), many=True
        )

        class Meta:
            model = Transaction
            exclude = ["user", "updated_at"]
            depth = 1
            ordering = ["transaction_date"]

    return TransactionSerializer


class ListTransactionSerializer(serializers.ModelSerializer):
    id = serializers.UUIDField(read_only=True)
    wallet = WalletSerializer(required=True)
    categories = CategorySerializer(required=True, many=True)

    class Meta:
        model = Transaction
        exclude = ["user", "updated_at"]
        depth = 1
        ordering = ["transaction_date"]
