from rest_framework import serializers


class CurrencyRecordSerializer(serializers.Serializer):
    id = serializers.UUIDField(required=True)
    value = serializers.FloatField(required=True)
    record_date = serializers.DateField(required=True)
