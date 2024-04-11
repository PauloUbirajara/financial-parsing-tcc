from http import HTTPStatus

from rest_framework import viewsets
from rest_framework.exceptions import NotAuthenticated
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin

from apps.category import serializers
from apps.category.models import Category


class CategoryViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        return Category.objects.filter(user=self.request.user)

    def get_serializer_class(self):
        supported_serializers = {
            "create": serializers.CreateCategorySerializer,
            "update": serializers.UpdateCategorySerializer,
        }
        serializer_class = supported_serializers.get(
            self.action, serializers.CategorySerializer
        )

        return serializer_class

    def list(self, request):
        categories = self.get_queryset()
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(categories, many=True)
        return Response(data=serializer.data)

    def retrieve(self, request, pk):
        category = self.get_queryset().filter(id=pk).first()

        if category is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(category)

        return Response(data=serializer.data)

    def update(self, request, pk, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        category: Category = self.get_queryset().filter(id=pk).first()

        if category is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        category.name = serializer.validated_data.get("name", category.name)

        category.save()

        return Response(data=serializer.data)

    def create(self, request, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        category = {**serializer.validated_data, "user": request.user}
        self.get_queryset().create(**category)

        return Response(data=serializer.data)
