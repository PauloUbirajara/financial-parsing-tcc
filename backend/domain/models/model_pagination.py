from rest_framework.pagination import PageNumberPagination
from rest_framework.response import Response


class ModelPagination(PageNumberPagination):
    page_size = 10
    page_query_param = "page"
    page_size_query_param = None
    max_page_size = 1000

    def get_paginated_response(self, data):
        return Response(
            {
                "links": {
                    "next": self.get_next_link(),
                    "previous": self.get_previous_link(),
                },
                "count": self.page.paginator.count,
                "num_pages": self.page.paginator.num_pages,
                "results": data,
            }
        )
