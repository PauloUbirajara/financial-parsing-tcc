from rest_framework import serializers


class CurrencySerializer(serializers.Serializer):
    id = serializers.UUIDField(read_only=True)
    name = serializers.CharField(required=True)
    representation = serializers.CharField(required=True)
