from rest_framework.pagination import PageNumberPagination


class ModelPagination(PageNumberPagination):
    page_size = 10
    page_query_param = "page"
    page_size_query_param = None
    max_page_size = 1000
