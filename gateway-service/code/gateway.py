#!/usr/bin/python3
import aiohttp
import asyncio
from aiohttp import web
import aiohttp_cors
import handlers


app = web.Application()

cors = aiohttp_cors.setup(app, defaults={
    "*": aiohttp_cors.ResourceOptions(
        allow_credentials=True,
        expose_headers="*",
        allow_headers="*",
    )
})
resource = cors.add(app.router.add_resource("/api/user"))
cors.add(resource.add_route("GET", handlers.user_get_handler))
resource = cors.add(app.router.add_resource("/api/user/register"))
cors.add(resource.add_route("POST", handlers.user_register_handler))
resource = cors.add(app.router.add_resource("/api/user/login"))
cors.add(resource.add_route("POST", handlers.user_login_handler))
resource = cors.add(app.router.add_resource("/api/user/refresh"))
cors.add(resource.add_route("POST", handlers.refresh_token_handler))
resource = cors.add(app.router.add_resource("/api/user/delete"))
cors.add(resource.add_route("DELETE", handlers.user_delete_handler))
resource = cors.add(app.router.add_resource("/api/user/update"))
cors.add(resource.add_route("PATCH", handlers.user_update_handler))
resource = cors.add(app.router.add_resource("/avatars/{file}"))
cors.add(resource.add_route("GET", handlers.get_avatar_handler))
resource = cors.add(app.router.add_resource("/api/subscriptions/create"))
cors.add(resource.add_route("POST", handlers.create_subscription_handler))

web.run_app(app, host="0.0.0.0", port=8080)