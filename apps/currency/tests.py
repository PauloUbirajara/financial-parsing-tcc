from http import HTTPStatus

from django.apps import apps
from django.contrib.auth.models import User
from django.test import TestCase
from rest_framework.test import APIClient

Currency = apps.get_model("currency", "Currency")


class CurrencyTestCase(TestCase):
    def setUp(self):
        self.user = User.objects.create_user(
            username="testuser", email="test@example.com", password="testpassword"
        )
        self.client = APIClient()
        self.client.force_authenticate(user=self.user)

    def test_read_currency(self):
        currency = Currency.objects.create(
            user=self.user, name="Test Currency", representation="TST"
        )
        response = self.client.get(f"/api/currencies/{currency.id}/")
        self.assertEqual(response.status_code, HTTPStatus.OK)
        self.assertEqual(response.data["name"], "Test Currency")

    def test_read_currencies(self):
        Currency.objects.create(
            user=self.user, name="Test Currency", representation="TST"
        )
        Currency.objects.create(
            user=self.user, name="Test Currency 2", representation="TST2"
        )
        response = self.client.get(f"/api/currencies/")
        self.assertEqual(response.status_code, HTTPStatus.OK)
        self.assertEqual(len(response.data), 2)
