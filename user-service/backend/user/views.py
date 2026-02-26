from rest_framework.generics import CreateAPIView
from rest_framework_simplejwt.tokens import RefreshToken
from rest_framework.response import Response
from rest_framework.decorators import api_view, authentication_classes
from .serializers import UserSerializer
from .models import User
# Create your views here.



@api_view(['POST'])
@authentication_classes(())
def user_login(request, *args, **kwargs):
	user = User.objects.get(email=request.data["email"])
	token = RefreshToken.for_user(user)
	res_obj = {
		"user": {
		"id": str(user.user_id),
		"username": str(user.username),
		"fullname": str(user.fullname),
		"email": str(user.email),
		"avatar_url": str(user.avatar),
		},
		"tokens": {
		"acces_token": str(token.access_token),
		"refresh_token": str(token),
		"token_type": "Bearer",
		"expires_in": 86400,
		}
	}
	return Response(res_obj,  status=200)


class UserRegisterView(CreateAPIView):
	serializer_class = UserSerializer
	authentication_classes = ()
	def post(self, request, *args, **kwargs):
		response = super().post(request, *args, **kwargs)
		user = User.objects.get(email=response.data["email"])
		token = RefreshToken.for_user(user)
		res_obj = {
			"user": {
			"id": str(user.user_id),
			"username": str(user.username),
			"fullname": str(user.fullname),
			"email": str(user.email),
			"avatar_url": str(user.avatar),
			},
			"tokens": {
			"acces_token": str(token.access_token),
			"refresh_token": str(token),
			"token_type": "Bearer",
			"expires_in": 86400,
			}
		}
		return Response(res_obj, status=201)


@api_view(['GET'])
def get_user(request, *args, **kwargs):
	user =  request.user 
	res_obj = {
		"user": {
		"id": str(user.user_id),
		"username": str(user.username),
		"fullname": str(user.fullname),
		"email": str(user.email),
		"avatar_url": str(user.avatar),
		}}
	return Response(res_obj, status=200)


@api_view(['DELETE'])
def delete_user(request, *args, **kwargs):
	request.user.delete()
	return Response("User deleted successfully", status=200)


@api_view(['PATCH'])
def update_user(request, *args, **kwargs):
	user = User.objects.get(id=request.user.id)
	if request.data["password"]:
		user.password = request.data["password"]
	if request.data["email"]:
		user.email = request.data["email"]
	if request.data["username"] :
		user.username = request.data["username"]
	if request.data["fullname"]:
		user.fullname = request.data["fullname"]
	if request.data["id"]:
		user.user_id = request.data["id"]
	if request.data["avatar"]:
		user.avatar = request.data["avatar"]
	user.save()

	res_obj = {
	"user": {
	"id": str(user.user_id),
	"username": str(user.username),
	"fullname": str(user.fullname),
	"email": str(user.email),
	"avatar_url": str(user.avatar),
	}}
	return Response("User updated successfully", status=200)

@api_view(['POST'])
@authentication_classes(())
def refresh_token(request, *args, **kwargs):
	token = RefreshToken(request.data["refresh_token"])
	res_obj = {
		"tokens": {
			"acces_token": str(token.access_token),
			"refresh_token": str(token),
			"token_type": "Bearer",
			"expires_in": 86400,
			}
		}
	return Response(res_obj, status=200)
