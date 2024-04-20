from rest_framework import serializers


class CategorySerializer(serializers.Serializer):
    id = serializers.UUIDField(read_only=True)
    name = serializers.CharField(required=True)


class BulkDeleteCategorySerializer(serializers.Serializer):
    ids = serializers.ListField(required=True)
