from abc import ABC
from typing import Any

from django.db.models import Model


class ExportModelToFormat(ABC):
    def export(self, model: Model) -> Any: ...
