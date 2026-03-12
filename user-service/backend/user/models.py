import uuid
from django.db import models
from django.contrib.auth.models import AbstractUser

# Create your models here.

class User(AbstractUser):
	id =  models.UUIDField(default=uuid.uuid4, null=False, blank=False, primary_key=True)
	fullname = models.CharField(max_length=100, blank=False, null=False)
	avatar = models.FileField(upload_to="avatars", blank=False, null=False)
	notifications = models.BooleanField(default=True)
	is_connected_email =models.BooleanField(default=False)
	is_active = True

	def __str__(self):
		return self.username
		
