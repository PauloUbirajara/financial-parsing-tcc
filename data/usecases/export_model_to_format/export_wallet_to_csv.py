import csv
from typing import Any

from django.apps import apps
from django.db.models import Model
from django.http import HttpResponse

from apps.transaction.serializers import TransactionSerializer
from apps.wallet.serializers import WalletSerializer
from domain.usecases.export_model_to_format import ExportModelToFormat

Transaction = apps.get_model("transaction", "Transaction")


class ExportWalletToCSV(ExportModelToFormat):
    csv_filename: str

    def __init__(self, csv_filename: str):
        self.csv_filename = csv_filename

    def export(self, model: Model) -> Any:
        transactions = Transaction.objects.filter(wallet=model)

        context = {
            "wallet": WalletSerializer(model).data,
            "transactions": TransactionSerializer(transactions, many=True).data,
            "total": sum(t.value for t in transactions),
        }
        response = HttpResponse(
            content_type="text/csv",
            headers={
                "Content-Disposition": 'attachment; filename="{}.csv"'.format(
                    self.csv_filename
                )
            },
        )
        writer = csv.writer(response)

        # Wallet info
        writer.writerow(["Nome", "Descrição", "Moeda", "Total"])
        writer.writerow(
            [
                context["wallet"]["name"],
                context["wallet"]["description"],
                "{} ({})".format(
                    context["wallet"]["currency"]["name"],
                    context["wallet"]["currency"]["representation"],
                ),
                context["total"],
            ]
        )
        writer.writerow([])

        # Transactions info
        writer.writerow(["Data", "Nome", "Descrição", "Categorias", "Valor"])
        for trx in context["transactions"]:
            writer.writerow(
                [
                    trx["transaction_date"],
                    trx["name"],
                    trx["description"],
                    " ".join(c["name"] for c in trx["categories"]),
                    trx["value"],
                ]
            )
        return response
