from http import HTTPStatus

from django.apps import apps
from django.contrib.auth.models import User
from django.test import TestCase
from rest_framework.test import APIClient

Currency = apps.get_model("currency", "Currency")
Wallet = apps.get_model("wallet", "Wallet")


class WalletTestCase(TestCase):
    def setUp(self):
        self.user = User.objects.create_user(
            username="testuser", email="test@example.com", password="testpassword"
        )
        self.currency = Currency.objects.create(
            name="USD", representation="$", user=self.user
        )
        self.client = APIClient()
        self.client.force_authenticate(user=self.user)

    def test_create_wallet(self):
        data = {
            "name": "Test Wallet",
            "description": "Test Wallet Description",
            "currency": self.currency.id,
        }
        response = self.client.post("/api/wallets/", data, format="json")
        self.assertEqual(response.status_code, HTTPStatus.CREATED)
        self.assertEqual(Wallet.objects.count(), 1)
        self.assertEqual(Wallet.objects.get().name, "Test Wallet")

    def test_read_wallet(self):
        wallet = Wallet.objects.create(
            user=self.user,
            currency=self.currency,
            name="Test Wallet",
            description="Test Wallet Description",
        )
        response = self.client.get(f"/api/wallets/{wallet.id}/")
        self.assertEqual(response.status_code, HTTPStatus.OK)
        self.assertEqual(response.data["name"], "Test Wallet")

    def test_update_wallet(self):
        wallet = Wallet.objects.create(
            user=self.user,
            currency=self.currency,
            name="Test Wallet",
            description="Test Wallet Description",
        )
        data = {
            "name": "Updated Test Wallet",
            "description": "Updated Test Wallet Description",
            "currency": self.currency.id,
        }
        response = self.client.put(f"/api/wallets/{wallet.id}/", data, format="json")
        self.assertEqual(response.status_code, HTTPStatus.OK)
        wallet.refresh_from_db()
        self.assertEqual(wallet.name, "Updated Test Wallet")
        self.assertEqual(wallet.description, "Updated Test Wallet Description")

    def test_delete_wallet(self):
        wallet = Wallet.objects.create(
            user=self.user,
            currency=self.currency,
            name="Test Wallet",
            description="Test Wallet Description",
        )
        response = self.client.delete(f"/api/wallets/{wallet.id}/")
        self.assertEqual(response.status_code, HTTPStatus.NO_CONTENT)
        self.assertEqual(Wallet.objects.count(), 0)
