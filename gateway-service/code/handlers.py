import aiohttp
import asyncio
from aiohttp import web
from aiohttp import FormData


async def user_get_handler(request):
	async with aiohttp.ClientSession() as session:
        		async with session.get('http://localhost:8088/api/user', headers=request.headers) as response:
        			data = await response.json()
        			return web.json_response(data, status=response.status)
			
			
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
		field = await reader.next()
				
	async with aiohttp.ClientSession() as session:
        		async with session.post('http://localhost:8088/api/user/register', data=request_data) as response:
        			return web.json_response(await response.json(), status=response.status)
	
async def user_login_handler(request):
	async with aiohttp.ClientSession() as session:
		async with session.post('http://localhost:8088/api/user/login', data=await request.json()) as response:
			return web.json_response(await response.json(), status=200)
			
async def refresh_token_handler(request):
	async with aiohttp.ClientSession() as session:
		async with session.post('http://localhost:8088/api/user/refresh', data=await request.json()) as response:
			return web.json_response(await response.json(), status=200)
			
async def user_delete_handler(request):
	async with aiohttp.ClientSession() as session:
		async with session.delete('http://localhost:8088/api/user/delete', headers=request.headers) as response:
			return web.json_response(await response.json(), status=200)

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
		field = await reader.next()
				
	async with aiohttp.ClientSession() as session:
        		async with session.patch('http://localhost:8088/api/user/update', data=request_data, headers=headers) as response:
        			return web.json_response(await response.json(), status=response.status)
        			
        			
async def get_avatar_handler(request):
	file_name =request.match_info.get('file', "file.jpg")
	async with aiohttp.ClientSession() as session:
		async with session.post(f'http://localhost:8040/avatars/{file_name}') as response:
			return web.Response(body=await response.read(), \
			headers={'Content-Disposition': f'attachment; filename={file_name}', 'Content-Type': 'application/octet-stream'})