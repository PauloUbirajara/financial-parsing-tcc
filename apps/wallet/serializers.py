from django.apps import apps
from rest_framework import serializers

Currency = apps.get_model("currency", "Currency")
Wallet = apps.get_model("wallet", "Wallet")


class WalletSerializer(serializers.ModelSerializer):
    class Meta:
        model = Wallet
        exclude = ["user", "created_at", "updated_at"]
        depth = 1


class CreateWalletSerializer(serializers.ModelSerializer):
    class Meta:
        model = Wallet
        exclude = ["id", "user", "created_at", "updated_at"]


class UpdateWalletSerializer(serializers.ModelSerializer):
    class Meta:
        model = Wallet
        exclude = ["id", "user", "created_at", "updated_at"]
