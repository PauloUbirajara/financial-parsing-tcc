from rest_framework import serializers


class BulkDeleteSerializer(serializers.Serializer):
    ids = serializers.ListField(required=True)
