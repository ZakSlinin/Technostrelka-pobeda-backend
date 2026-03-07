import aiohttp
import os
import uuid
import asyncio
from aiohttp import web
from aiohttp import FormData
import json

async def get_user_id(headers):
	async with aiohttp.ClientSession() as session:
        		async with session.get(f'http://{os.environ.get("user_service_host")}:8080/api/user', headers=headers) as response:
        			data = await response.json()
        			return data["user"]["id"]


async def user_get_handler(request):
	async with aiohttp.ClientSession() as session:
        		async with session.get(f'http://{os.environ.get("user_service_host")}:8080/api/user', headers=request.headers) as response:
        			return web.json_response(await response.json(), status=response.status)
			
			
async def user_register_handler(request):
	reader = await request.multipart()
	request_data = FormData()
	field = await reader.next()
	while field:
		if field.name == 'username':
			request_data.add_field('username',  await field.text())
		elif field.name == 'fullname':
			request_data.add_field('fullname',  await field.text())
		elif field.name == 'email':
			request_data.add_field('email',  await field.text())
		elif field.name == 'password':
			request_data.add_field('password',  await field.text())
		elif field.name == 'avatar':
			request_data.add_field('avatar',  await field.read(),  filename='image.jpg')
		elif field.name == 'notifications':
			request_data.add_field('notifications',  await field.text())
		field = await reader.next()
				
	async with aiohttp.ClientSession() as session:
        		async with session.post(f'http://{os.environ.get("user_service_host")}:8080/api/user/register', data=request_data) as response:
        			return web.json_response(await response.json(), status=response.status)
	
async def user_login_handler(request):
	async with aiohttp.ClientSession() as session:
		async with session.post(f'http://{os.environ.get("user_service_host")}:8080/api/user/login', data=await request.json()) as response:
			return web.json_response(await response.json(), status=response.status)
			
async def refresh_token_handler(request):
	async with aiohttp.ClientSession() as session:
		async with session.post(f'http://{os.environ.get("user_service_host")}:8080/api/user/refresh', data=await request.json()) as response:
			return web.json_response(await response.json(), status=response.status)
			
async def user_delete_handler(request):
	async with aiohttp.ClientSession() as session:
		async with session.delete(f'http://{os.environ.get("user_service_host")}:8080/api/user/delete', headers=request.headers) as response:
			return web.json_response(await response.json(), status=response.status)

async def user_update_handler(request):
	reader = await request.multipart()
	request_data = FormData()
	field = await reader.next()
	headers = {}
	headers['Authorization'] = request.headers['Authorization']
	while field:
		if field.name == 'username':
			request_data.add_field('username',  await field.text())
		elif field.name == 'fullname':
			request_data.add_field('fullname',  await field.text())
		elif field.name == 'email':
			request_data.add_field('email',  await field.text())
		elif field.name == 'password':
			request_data.add_field('password',  await field.text())
		elif field.name == 'avatar':
			request_data.add_field('avatar',  await field.read(),  filename='image.jpg')
		elif field.name == 'notifications':
			request_data.add_field('notifications',  await field.text())
		else:
			web.json(f"Undefined field: {field.name}", status=400)
		field = await reader.next()
				
	async with aiohttp.ClientSession() as session:
        		async with session.patch(f'http://{os.environ.get("user_service_host")}:8080/api/user/update', data=request_data, headers=headers) as response:
        			return web.json_response(await response.json(), status=response.status)
        			
        			
async def get_avatar_handler(request):
	file_name =request.match_info.get('file', "file.jpg")
	async with aiohttp.ClientSession() as session:
		async with session.get(f'http://{os.environ.get("avatars_host")}/avatars/{file_name}') as response:
			return web.Response(body=await response.read(), \
			headers={'Content-Disposition': f'attachment; filename={file_name}', 'Content-Type': 'application/octet-stream'})
			
		
async def create_subscription_handler(request):
	request_data = FormData() 
	reader = await request.multipart()
	field = await reader.next()
	while field:
		if field.name == "name":
			request_data.add_field("name", await field.text())
		elif field.name == "cost":
			request_data.add_field("cost", await field.text())
		elif field.name == "next_billing":
			request_data.add_field("next_billing", await field.text())
		elif field.name == "status":
			request_data.add_field("status", await field.text())
		elif field.name == "subscription_avatar":
			request_data.add_field("subscription_avatar", await field.read())
		elif field.name == "category":
			request_data.add_field("category", await field.text())
		elif field.name == "url_service":
			request_data.add_field("url_service", await field.text())
		elif field.name == "use_in_this_month":
			request_data.add_field("use_in_this_month", await field.text())
		elif field.name == "cancellation_link":
			request_data.add_field("cancellation_link", await field.text())
		field = await reader.next()
	headers = {}
	user_id = await get_user_id(request.headers)
	headers["X-User-Id"] = user_id
	async with aiohttp.ClientSession() as session:
		async with session.post(f'http://{os.environ.get("subscriptions_host")}:8080/api/subscriptions/create', data=request_data, headers=headers) as response:
			return web.json_response(await response.json(), status=response.status)
			

async def update_subscription_handler(request):
	request_data = FormData()
	id = request.match_info.get('id', None) 
	if not id:
		web.json_response("subscription id is not provided", status=400)
	reader = await request.multipart()
	field = await reader.next()
	while field:
		if field.name == "name":
			request_data.add_field("name", await field.text())
		elif field.name == "cost":
			request_data.add_field("cost", await field.text())
		elif field.name == "next_billing":
			request_data.add_field("next_billing", await field.text())
		elif field.name == "status":
			request_data.add_field("status", await field.text())
		elif field.name == "subscription_avatar":
			request_data.add_field("subscription_avatar", await field.read())
		elif field.name == "category":
			request_data.add_field("category", await field.text())
		elif field.name == "url_service":
			request_data.add_field("url_service", await field.text())
		elif field.name == "use_in_this_month":
			request_data.add_field("use_in_this_month", await field.text())
		elif field.name == "cancellation_link":
			request_data.add_field("cancellation_link", await field.text())
		field = await reader.next()
	headers = {}
	user_id = await get_user_id(request.headers)
	headers["X-User-Id"] = user_id
	async with aiohttp.ClientSession() as session:
		async with session.patch(f'http://{os.environ.get("subscriptions_host")}:8080/api/subscriptions/update/{id}', data=request_data, headers=headers) as response:
			return web.json_response(await response.json(), status=response.status)

async def delete_subscription_handler(request):
	headers = {}
	id = request.match_info.get('id', None) 
	if not id:
		web.json_response("subscription id is not provided", status=response.status)
	user_id = await get_user_id(request.headers)
	headers["X-User-Id"] = user_id
	async with aiohttp.ClientSession() as session:
		async with session.delete(f'http://{os.environ.get("subscriptions_host")}:8080/api/subscriptions/delete/{id}', headers=headers) as response:
			return web.json_response(await response.json(), status=response.status)
			
async def get_subscription_handler(request):
	headers = {}
	user_id = await get_user_id(request.headers)
	headers["X-User-Id"] = user_id
	async with aiohttp.ClientSession() as session:
		async with session.get(f'http://{os.environ.get("subscriptions_host")}:8080/api/subscriptions/all', headers=headers) as response:
			return web.json_response(await response.json(), status=response.status)

async def get_subavatar_handler(request):
	file_name =request.match_info.get('file', "file.jpg")
	async with aiohttp.ClientSession() as session:
		async with session.get(f'http://{os.environ.get("subavatars_host")}/uploads/{file_name}') as response:
			return web.Response(body=await response.read(), \
			headers={'Content-Disposition': f'attachment; filename={file_name}', 'Content-Type': 'application/octet-stream'})
			
async def getid_subscription_handler(request):
	id = request.match_info.get('id', None) 
	headers = {}
	user_id = await get_user_id(request.headers)
	headers["X-User-Id"] = user_id
	async with aiohttp.ClientSession() as session:
		async with session.get(f'http://{os.environ.get("subscriptions_host")}:8080/api/subscriptions/{id}', headers=headers) as response:
			return web.json_response(await response.json(), status=response.status)

