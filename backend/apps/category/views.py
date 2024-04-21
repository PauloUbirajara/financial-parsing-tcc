from http import HTTPStatus

from rest_framework import viewsets
from rest_framework.decorators import action
from rest_framework.exceptions import NotAuthenticated
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin

from apps.category import serializers
from apps.category.models import Category
from domain.models.bulk_delete_serializer import BulkDeleteSerializer
from domain.models.model_pagination import ModelPagination


class CategoryViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    pagination_class = ModelPagination

    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        queryset = Category.objects.filter(user=self.request.user)
        return queryset

    def get_serializer_class(self):
        supported_serializers = {"bulk_delete": BulkDeleteSerializer}
        serializer_class = supported_serializers.get(
            self.action, serializers.CategorySerializer
        )
        return serializer_class

    def list(self, request, *args, **kwargs):
        queryset = self.get_queryset()

        search_term = self.request.query_params.get("search")
        if search_term:
            queryset = queryset.filter(name__icontains=search_term)

        paginated_queryset = self.paginate_queryset(queryset)
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(paginated_queryset, many=True)
        return self.paginator.get_paginated_response(serializer.data)

    def retrieve(self, request, *args, **kwargs):
        pk = kwargs.get("pk")
        category = self.get_queryset().filter(id=pk).first()

        if category is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(category)

        return Response(data=serializer.data)

    def update(self, request, *args, **kwargs):
        pk = kwargs.get("pk")
        category: Category = self.get_queryset().filter(id=pk).first()

        if category is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(category, data=request.data, partial=True)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

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

        return Response(data=serializer.data, status=HTTPStatus.CREATED)

    @action(methods=["POST"], detail=False)
    def bulk_delete(self, request, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        self.get_queryset().filter(
            id__in=serializer.validated_data.get("ids", [])
        ).delete()

        return Response(status=HTTPStatus.OK)
