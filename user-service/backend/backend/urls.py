from django.contrib import admin
from django.urls import path
from user.views import user_login, refresh_token, get_user, delete_user, update_user, user_registration


urlpatterns = [
    path('admin/', admin.site.urls),
    path('api/user/login', user_login),
    path('api/user/register', user_registration),
    path('api/user/refresh', refresh_token),
    path('api/user', get_user),
    path('api/user/update', update_user),
    path('api/user/delete', delete_user)
]
