import logging
from typing import Any

from django.apps import apps
from django.db.models import Model
from django.http import HttpResponse
from django.template.loader import get_template
from xhtml2pdf import pisa

from apps.transaction.serializers import ListTransactionSerializer
from apps.wallet.serializers import ListWalletSerializer
from domain.usecases.export_model_to_format import ExportModelToFormat

Transaction = apps.get_model("transaction", "Transaction")


class ExportWalletToPDF(ExportModelToFormat):
    template_name: str
    pdf_filename: str

    def __init__(self, template_name: str, pdf_filename: str):
        self.template_name = template_name
        self.pdf_filename = pdf_filename

    def export(self, model: Model) -> Any:
        template = get_template(self.template_name)
        transactions = Transaction.objects.filter(wallet=model)

        context = {
            "wallet": ListWalletSerializer(model).data,
            "transactions": ListTransactionSerializer(transactions, many=True).data,
            "total": sum(t.value for t in transactions),
        }
        html = template.render(context)
        response = HttpResponse(content_type="application/pdf")
        response["Content-Disposition"] = (
            'attachment; filename="{pdf_filename}.pdf"'.format(
                pdf_filename=self.pdf_filename
            )
        )
        pdf_status = pisa.CreatePDF(html, dest=response)
        if pdf_status.err:
            error = {"error": "Erro ao gerar relatório PDF de carteira."}
            logging.warning(error, pdf_status.err)
            return HttpResponse(data=error)

        return response
