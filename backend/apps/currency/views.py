from http import HTTPStatus

from rest_framework import viewsets
from rest_framework.exceptions import NotAuthenticated
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin

from apps.currency import serializers
from apps.currency.models import Currency
from domain.models.model_pagination import ModelPagination


class CurrencyViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    serializer_class = serializers.CurrencySerializer
    pagination_class = ModelPagination

    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        queryset = Currency.objects.filter()
        return queryset

    def list(self, request, *args, **kwargs):
        queryset = self.get_queryset()

        search_term = self.request.query_params.get("search")
        if search_term:
            queryset = queryset.filter(name__icontains=search_term)

        serializer_class = self.get_serializer_class()
        page = self.request.query_params.get("page")
        if not page:
            serializer = serializer_class(queryset, many=True)

            return Response(
                {
                    "links": {
                        "next": None,
                        "previous": None,
                    },
                    "count": queryset.count(),
                    "num_pages": 1,
                    "results": serializer.data,
                }
            )

        paginated_queryset = self.paginate_queryset(queryset)
        serializer = serializer_class(paginated_queryset, many=True)
        return self.paginator.get_paginated_response(serializer.data)

    def retrieve(self, request, *args, **kwargs):
        pk = kwargs.get("pk")
        currency = self.get_queryset().filter(id=pk).first()

        if currency is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(currency)

        return Response(data=serializer.data)

    def update(self, request, *args, **kwargs):
        return Response(status=HTTPStatus.UNAUTHORIZED)

    def delete(self, request, *args, **kwargs):
        return Response(status=HTTPStatus.UNAUTHORIZED)
