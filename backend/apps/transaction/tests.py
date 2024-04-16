import datetime
from http import HTTPStatus

from django.apps import apps
from django.contrib.auth.models import User
from django.test import TestCase
from rest_framework.test import APIClient

from .models import Transaction

Wallet = apps.get_model("wallet", "Wallet")
Currency = apps.get_model("currency", "Currency")
Transaction = apps.get_model("transaction", "Transaction")
Category = apps.get_model("category", "Category")


class TransactionTestCase(TestCase):
    def setUp(self):
        self.user = User.objects.create_user(
            username="testuser", email="test@example.com", password="testpassword"
        )
        self.currency = Currency.objects.create(
            user=self.user, name="Test Currency", representation="TST"
        )
        self.wallet = Wallet.objects.create(
            user=self.user,
            name="Test Wallet",
            description="Test Wallet Description",
            currency=self.currency,
        )
        self.category = Category.objects.create(
            user=self.user,
            name="Test Category",
        )
        self.client = APIClient()
        self.client.force_authenticate(user=self.user)

    def test_create_transaction(self):
        data = {
            "wallet": self.wallet.id,
            "categories": [self.category.id],
            "name": "Test Transaction",
            "description": "Test Transaction Description",
            "transaction_date": datetime.date.today(),
            "value": "100.00",
        }
        response = self.client.post("/api/transactions/", data, format="json")
        self.assertEqual(response.status_code, HTTPStatus.CREATED)
        self.assertEqual(Transaction.objects.count(), 1)
        self.assertEqual(Transaction.objects.get().name, "Test Transaction")

    def test_read_transaction(self):
        transaction = Transaction.objects.create(
            user=self.user,
            wallet=self.wallet,
            name="Test Transaction",
            description="Test Transaction Description",
            transaction_date=datetime.date.today(),
            value="100.00",
        )
        response = self.client.get(f"/api/transactions/{transaction.id}/")
        self.assertEqual(response.status_code, HTTPStatus.OK)
        self.assertEqual(response.data["name"], "Test Transaction")

    def test_update_transaction(self):
        transaction = Transaction.objects.create(
            user=self.user,
            wallet=self.wallet,
            name="Test Transaction",
            description="Test Transaction Description",
            transaction_date=datetime.date.today(),
            value="100.00",
        )
        data = {
            "wallet": self.wallet.id,
            "categories": [self.category.id],
            "name": "Updated Test Transaction",
            "description": "Updated Test Transaction Description",
            "transaction_date": datetime.date.today(),
            "value": "200.00",
        }
        response = self.client.put(
            f"/api/transactions/{transaction.id}/", data, format="json"
        )
        self.assertEqual(response.status_code, HTTPStatus.OK)
        transaction.refresh_from_db()
        self.assertEqual(transaction.name, "Updated Test Transaction")
        self.assertEqual(transaction.value, 200.00)

    def test_delete_transaction(self):
        transaction = Transaction.objects.create(
            user=self.user,
            wallet=self.wallet,
            name="Test Transaction",
            description="Test Transaction Description",
            transaction_date=datetime.date.today(),
            value="100.00",
        )
        response = self.client.delete(f"/api/transactions/{transaction.id}/")
        self.assertEqual(response.status_code, HTTPStatus.NO_CONTENT)
        self.assertEqual(Transaction.objects.count(), 0)
