from http import HTTPStatus

from django.apps import apps
from rest_framework import viewsets
from rest_framework.decorators import action
from rest_framework.exceptions import NotAuthenticated
from rest_framework.response import Response
from rest_framework_extensions.mixins import NestedViewSetMixin

from apps.wallet import serializers
from apps.wallet.models import Wallet
from data.usecases.export_model_to_format.export_wallet_to_csv import ExportWalletToCSV
from data.usecases.export_model_to_format.export_wallet_to_html import (
    ExportWalletToHTML,
)
from data.usecases.export_model_to_format.export_wallet_to_pdf import ExportWalletToPDF
from domain.models.bulk_delete_serializer import BulkDeleteSerializer
from domain.models.model_pagination import ModelPagination
from domain.usecases.export_model_to_format import ExportModelToFormat

Currency = apps.get_model("currency", "Currency")


class WalletViewSet(viewsets.ModelViewSet, NestedViewSetMixin):
    pagination_class = ModelPagination

    def get_queryset(self):
        if not self.request.user.is_authenticated:
            raise NotAuthenticated()

        queryset = Wallet.objects.filter()
        return queryset

    def get_serializer_class(self):
        supported_serializers = {
            "list": serializers.ListWalletSerializer,
            "retrieve": serializers.ListWalletSerializer,
            "bulk_delete": BulkDeleteSerializer,
        }
        serializer_class = supported_serializers.get(
            self.action, serializers.WalletSerializer
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

    def retrieve(self, request, pk):
        wallet = self.get_queryset().filter(id=pk).first()

        if wallet is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(wallet)

        return Response(data=serializer.data)

    def update(self, request, pk, *args, **kwargs):
        wallet: Wallet = self.get_queryset().filter(id=pk).first()

        if wallet is None:
            return Response(status=HTTPStatus.NOT_FOUND)

        serializer_class = self.get_serializer_class()
        serializer = serializer_class(wallet, data=request.data, partial=True)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        wallet.name = serializer.validated_data.get("name", wallet.name)
        wallet.description = serializer.validated_data.get(
            "description", wallet.description
        )
        wallet.currency = serializer.validated_data.get("currency", wallet.currency)

        wallet.save()

        return Response(data=serializer.data)

    def create(self, request, *args, **kwargs):
        serializer_class = self.get_serializer_class()
        serializer = serializer_class(data=request.data)

        if not serializer.is_valid():
            return Response(status=HTTPStatus.BAD_REQUEST, data=serializer.errors)

        wallet = {
            **serializer.validated_data,
            "user": request.user,
        }
        self.get_queryset().create(**wallet)

        return Response(data=serializer.data, status=HTTPStatus.CREATED)

    @action(methods=["POST"], detail=True)
    def export(self, request, pk=None, *args, **kwargs):
        wallet = self.get_queryset().filter(id=pk).first()
        export_format = request.data.get("format")

        if wallet is None:
            error = {"error": "Não foi possível exportar carteira."}
            return Response(status=HTTPStatus.NOT_FOUND, data=error)

        supported_formats: dict[str, ExportModelToFormat] = {
            "csv": ExportWalletToCSV(csv_filename="wallet_{}".format(wallet.name)),
            "html": ExportWalletToHTML(template_name="export_html/index.html"),
            "pdf": ExportWalletToPDF(
                template_name="export_pdf/index.html",
                pdf_filename="wallet_{}".format(wallet.name),
            ),
        }

        wallet_export_usecase = supported_formats.get(export_format)

        if wallet_export_usecase is None:
            error = {"error": "Formato inválido de exportação"}
            return Response(status=HTTPStatus.BAD_REQUEST, data=error)

        return wallet_export_usecase.export(model=wallet)

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
