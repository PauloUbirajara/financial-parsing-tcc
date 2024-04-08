from django.apps import apps

from rest_framework import serializers


Wallet = apps.get_model('wallet', 'Wallet')


class TransactionSerializer(serializers.Serializer):
    id = serializers.UUIDField(required=True)
    name = serializers.CharField(required=True)
    description = serializers.CharField(allow_blank=True)
    wallet = serializers.PrimaryKeyRelatedField(
        required=True,
        queryset=Wallet.objects.all()
    )
    value = serializers.DecimalField(required=True, max_digits=10, decimal_places=2)
    transaction_date = serializers.DateField(required=True)

class CreateTransactionSerializer(serializers.Serializer):
    name = serializers.CharField(required=True)
    description = serializers.CharField(allow_blank=True)
    wallet = serializers.PrimaryKeyRelatedField(
        required=True,
        queryset=Wallet.objects.all()
    )
    value = serializers.DecimalField(required=True, max_digits=10, decimal_places=2)
    transaction_date = serializers.DateField(required=True)

class UpdateTransactionSerializer(serializers.Serializer):
    name = serializers.CharField(required=True)
    description = serializers.CharField(allow_blank=True)
    wallet = serializers.PrimaryKeyRelatedField(
        required=True,
        queryset=Wallet.objects.all()
    )
    value = serializers.DecimalField(required=True, max_digits=10, decimal_places=2)
    transaction_date = serializers.DateField(required=True)
