from rest_framework import serializers


class CurrencySerializer(serializers.Serializer):
    id = serializers.UUIDField(required=True)
    name = serializers.CharField(required=True)
    representation = serializers.CharField(required=True)

class CreateCurrencySerializer(serializers.Serializer):
    name = serializers.CharField(required=True)
    representation = serializers.CharField(required=True)

class UpdateCurrencySerializer(serializers.Serializer):
    name = serializers.CharField(required=True)
    representation = serializers.CharField(required=True)
