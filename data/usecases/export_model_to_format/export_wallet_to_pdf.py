from typing import Any

from django.apps import apps
from django.db.models import Model

from domain.usecases.export_model_to_format import ExportModelToFormat

Transaction = apps.get_model("transaction", "Transaction")


class ExportWalletToPDF(ExportModelToFormat):
    def export(self, model: Model) -> Any:
        raise NotImplementedError()
