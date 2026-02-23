from rest_framework.serializers import ModelSerializer
from .models import User


class UserSerializer(ModelSerializer):
	
	class Meta:
		model = User
		fields = "__all__"
		write_only_fields = ('password')