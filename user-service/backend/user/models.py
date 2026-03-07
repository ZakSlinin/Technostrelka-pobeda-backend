import uuid
from django.db import models
from django.contrib.auth.models import AbstractUser

# Create your models here.

class User(AbstractUser):
	user_id =  models.UUIDField(default=uuid.uuid4, null=False, blank=True)
	fullname = models.CharField(max_length=100, blank=False, null=False)
	avatar = models.FileField(upload_to="avatars", blank=True, null=True)
	notifications = models.BooleanField(default=True)
	is_active = True

	def __str__(self):
		return self.username
		
