from typing import Any

from django.apps import apps
from django.db.models import Model
from django.http import HttpResponse
from django.template.loader import render_to_string

from apps.transaction.serializers import TransactionSerializer
from apps.wallet.serializers import WalletSerializer
from domain.usecases.export_model_to_format import ExportModelToFormat

Transaction = apps.get_model("transaction", "Transaction")


class ExportWalletToHTML(ExportModelToFormat):
    def export(self, model: Model) -> Any:
        transactions = Transaction.objects.filter(wallet=model)

        context = {
            "wallet": WalletSerializer(model).data,
            "transactions": TransactionSerializer(transactions, many=True).data,
        }
        return HttpResponse(
            render_to_string(context=context, template_name="export/index.html")
        )
