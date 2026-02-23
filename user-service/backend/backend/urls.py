from django.contrib import admin
from django.urls import path
from user.views import UserLoginView, UserRegisterView, RefreshTokenAPIView


urlpatterns = [
    path('admin/', admin.site.urls),
    path('api/user/login', UserLoginView.as_view()),
    path('api/user/register', UserRegisterView.as_view()),
    path('api/user/refresh', RefreshTokenAPIView.as_view(), name='token_refresh'),
]
