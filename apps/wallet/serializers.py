from django.apps import apps
from rest_framework import serializers

Currency = apps.get_model("currency", "Currency")


class WalletSerializer(serializers.Serializer):
    id = serializers.UUIDField(required=True)
    name = serializers.CharField(required=True)
    description = serializers.CharField()
    currency = serializers.PrimaryKeyRelatedField(
        required=True, queryset=Currency.objects.all()
    )


class CreateWalletSerializer(serializers.Serializer):
    name = serializers.CharField(required=True)
    description = serializers.CharField(allow_blank=True)
    currency = serializers.PrimaryKeyRelatedField(
        required=True, queryset=Currency.objects.all()
    )


class UpdateWalletSerializer(serializers.Serializer):
    name = serializers.CharField(required=True)
    description = serializers.CharField()
    currency = serializers.PrimaryKeyRelatedField(
        required=True, queryset=Currency.objects.all()
    )
