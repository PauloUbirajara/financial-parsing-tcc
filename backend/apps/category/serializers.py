from rest_framework import serializers


class CategorySerializer(serializers.Serializer):
    id = serializers.UUIDField(required=True)
    name = serializers.CharField(required=True)


class CreateCategorySerializer(serializers.Serializer):
    name = serializers.CharField(required=True)


class UpdateCategorySerializer(serializers.Serializer):
    name = serializers.CharField(required=True)
