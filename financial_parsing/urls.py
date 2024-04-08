"""
URL configuration for financial_parsing project.

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/5.0/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include

from apps.currency.views import CurrencyViewSet
from apps.wallet.views import WalletViewSet
from apps.transaction.views import TransactionViewSet
from apps.category.views import CategoryViewSet

from rest_framework_extensions.routers import ExtendedSimpleRouter
from rest_framework_simplejwt.views import (
    TokenObtainPairView,
    TokenRefreshView,
    TokenVerifyView,
)
from apps.user_activation.views import UserActivationRegisterView, UserActivationConfirmView, UserActivationResendActivationView


router = ExtendedSimpleRouter()
router.register(
    r'currencies',
    CurrencyViewSet,
    basename='currency'
)
router.register(
    r'categories',
    CategoryViewSet,
    basename='category'
)
router.register(
    r'wallets',
    WalletViewSet,
    basename='wallet'
)
router.register(
    r'transactions',
    TransactionViewSet,
    basename='transaction'
)

urlpatterns = [
    path('admin/', admin.site.urls),

    # API endpoints
    path('api/', include(router.urls)),

    # User Management
    path('auth/register', UserActivationRegisterView.as_view(), name='user-register'),
    path('auth/activate/<uuid:token>', UserActivationConfirmView.as_view(), name='user-activate'),
    path('auth/activate/resend', UserActivationResendActivationView.as_view(), name='user-resend-activation'),

    # JWT
    path('auth/login', TokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('auth/refresh', TokenRefreshView.as_view(), name='token_refresh'),
    path('auth/validate', TokenVerifyView.as_view(), name='token_verify'),
]
