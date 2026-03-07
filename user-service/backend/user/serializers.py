from rest_framework.serializers import ModelSerializer, CharField
from .models import User


class UserSerializer(ModelSerializer):
	password = CharField(max_length=300, write_only=True, required=False)
	
	class Meta:
		model = User
		fields = "__all__"