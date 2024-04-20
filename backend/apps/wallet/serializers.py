from django.apps import apps
from rest_framework import serializers

from apps.currency.serializers import CurrencySerializer

Wallet = apps.get_model("wallet", "Wallet")
Currency = apps.get_model("currency", "Currency")


class WalletSerializer(serializers.ModelSerializer):
    id = serializers.UUIDField(read_only=True)
    currency = serializers.PrimaryKeyRelatedField(
        required=True, queryset=Currency.objects.all()
    )

    class Meta:
        model = Wallet
        exclude = ["user", "updated_at"]
        depth = 1


class ListWalletSerializer(serializers.ModelSerializer):
    id = serializers.UUIDField(read_only=True)

    class Meta:
        model = Wallet
        exclude = ["user", "updated_at"]
        depth = 1
