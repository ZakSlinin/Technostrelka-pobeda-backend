from rest_framework.serializers import Serializer, ModelSerializer, CharField, UUIDField, BooleanField, IntegerField, FileField
from .models import User


class UserSerializer(ModelSerializer):
	password = CharField(max_length=300, write_only=True, required=True)

	class Meta:
		model = User
		fields = ("password", "username", "id", "fullname", "notifications", "email", "avatar", "is_connected_email")

		
		
		
		
class UserUpdateSerializer(ModelSerializer):
	password = CharField(max_length=300, write_only=True, required=False)
	username = CharField(max_length=300, required=False)
	fullname = CharField(max_length=100, required=False)
	notifications = BooleanField(required=False)
	avatar = FileField(required=False)
	is_connected_email =BooleanField(required=False)
	
	class Meta:
		model = User
		fields = ("password", "username", "id", "fullname", "notifications", "email", "avatar", "is_connected_email")
		
		