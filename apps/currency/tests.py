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

    def test_create_currency(self):
        data = {"name": "Test Currency", "representation": "TST"}
        response = self.client.post("/api/currencies/", data, format="json")
        self.assertEqual(response.status_code, HTTPStatus.CREATED)
        self.assertEqual(Currency.objects.count(), 1)
        self.assertEqual(Currency.objects.get().name, "Test Currency")

    def test_read_currency(self):
        currency = Currency.objects.create(
            user=self.user, name="Test Currency", representation="TST"
        )
        response = self.client.get(f"/api/currencies/{currency.id}/")
        self.assertEqual(response.status_code, HTTPStatus.OK)
        self.assertEqual(response.data["name"], "Test Currency")

    def test_update_currency(self):
        currency = Currency.objects.create(
            user=self.user, name="Test Currency", representation="TST"
        )
        data = {"name": "Updated Test Currency", "representation": "UTS"}
        response = self.client.put(
            f"/api/currencies/{currency.id}/", data, format="json"
        )
        self.assertEqual(response.status_code, HTTPStatus.OK)
        currency.refresh_from_db()
        self.assertEqual(currency.name, "Updated Test Currency")
        self.assertEqual(currency.representation, "UTS")

    def test_delete_currency(self):
        currency = Currency.objects.create(
            user=self.user, name="Test Currency", representation="TST"
        )
        response = self.client.delete(f"/api/currencies/{currency.id}/")
        self.assertEqual(response.status_code, HTTPStatus.NO_CONTENT)
        self.assertEqual(Currency.objects.count(), 0)
