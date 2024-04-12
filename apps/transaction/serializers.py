from django.apps import apps
from rest_framework import serializers

Transaction = apps.get_model("transaction", "Transaction")
Wallet = apps.get_model("wallet", "Wallet")
Category = apps.get_model("category", "Category")


class TransactionSerializer(serializers.ModelSerializer):
    class Meta:
        model = Transaction
        exclude = ["user", "created_at", "updated_at"]
        depth = 1
        ordering = ["transaction_date"]


class CreateTransactionSerializer(serializers.ModelSerializer):
    class Meta:
        model = Transaction
        exclude = ["id", "user", "created_at", "updated_at"]


class UpdateTransactionSerializer(serializers.ModelSerializer):
    class Meta:
        model = Transaction
        exclude = ["id", "user", "created_at", "updated_at"]
