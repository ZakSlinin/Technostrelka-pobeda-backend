from rest_framework_simplejwt.tokens import RefreshToken
from rest_framework.response import Response
from rest_framework.decorators import api_view, authentication_classes
from .serializers import UserSerializer, UserUpdateSerializer
from .models import User


@api_view(['POST'])
@authentication_classes(())
def user_login(request):
	try:
		user = User.objects.get(email=request.data["email"], password=request.data["password"])
	except Exception as e:
		return Response("No such user", status=401)
	user_serializer = UserSerializer(user)
	token = RefreshToken.for_user(user)
	user_data = dict(user_serializer.data)
	user_data["avatar_url"] = user_data["avatar"]
	user_data.pop("avatar")
	res_obj = {
		"user": user_data,
		"tokens": {
		"acces_token": str(token.access_token),
		"refresh_token": str(token),
		"token_type": "Bearer",
		"expires_in": 86400,
		}
	}
	return Response(res_obj,  status=200)
	
@api_view(['POST'])
@authentication_classes(())
def user_registration(request):
	serializer = UserSerializer(data=request.data)
	if serializer.is_valid():
		user = serializer.save()
	else:
		return Response(serializer.errors, status=400)
	token = RefreshToken.for_user(user)
	user_data = dict(serializer.data)
	user_data["avatar_url"] = user_data["avatar"]
	user_data.pop("avatar")
	res_obj = {
		"user": user_data,
		"tokens": {
		"acces_token": str(token.access_token),
		"refresh_token": str(token),
		"token_type": "Bearer",
		"expires_in": 86400,
		}}
	return Response(res_obj, status=201)


@api_view(['GET'])
def get_user(request):
	user =  request.user
	if not user.username:
		return Response("Not authorized", status=401) 
	serializer = UserSerializer(user)
	user_data = dict(serializer.data)
	user_data["avatar_url"] = user_data["avatar"]
	user_data.pop("avatar")
	return Response({"user": user_data}, status=200)
	
	

@api_view(['DELETE'])
def delete_user(request):
	request.user.delete()
	return Response("User deleted successfully", status=200)


@api_view(['PATCH'])
def update_user(request):
	serializer = UserUpdateSerializer(request.user, data=request.data)
	if serializer.is_valid():
		serializer.save()
	else:
		return Response(serializer.errors, status=400)
	return Response("User updated successfully", status=200)


@api_view(['POST'])
@authentication_classes(())
def refresh_token(request):
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
