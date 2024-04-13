from http import HTTPStatus

from django.apps import apps
from django.contrib.auth.models import User
from django.test import TestCase
from rest_framework.test import APIClient

Category = apps.get_model("category", "Category")


class CategoryTestCase(TestCase):
    def setUp(self):
        self.user = User.objects.create_user(
            username="testuser", email="test@example.com", password="testpassword"
        )
        self.client = APIClient()
        self.client.force_authenticate(user=self.user)

    def test_create_category(self):
        data = {"name": "Test Category"}
        response = self.client.post("/api/categories/", data, format="json")
        self.assertEqual(response.status_code, HTTPStatus.CREATED)
        self.assertEqual(Category.objects.count(), 1)
        self.assertEqual(Category.objects.get().name, "Test Category")

    def test_read_category(self):
        category = Category.objects.create(user=self.user, name="Test Category")
        response = self.client.get(f"/api/categories/{category.id}/")
        self.assertEqual(response.status_code, HTTPStatus.OK)
        self.assertEqual(response.data["name"], "Test Category")

    def test_update_category(self):
        category = Category.objects.create(user=self.user, name="Test Category")
        data = {"name": "Updated Test Category"}
        response = self.client.put(
            f"/api/categories/{category.id}/", data, format="json"
        )
        self.assertEqual(response.status_code, HTTPStatus.OK)
        category.refresh_from_db()
        self.assertEqual(category.name, "Updated Test Category")

    def test_delete_category(self):
        category = Category.objects.create(user=self.user, name="Test Category")
        response = self.client.delete(f"/api/categories/{category.id}/")
        self.assertEqual(response.status_code, HTTPStatus.NO_CONTENT)
        self.assertEqual(Category.objects.count(), 0)
