#!/usr/bin/python3
import requests
import os
import time
import datetime
import json
from django import setup
import asyncio
from aiohttp import ClientSession, FormData

async def update_sub(sub, headers):
	date = sub["next_billing"]
	date = date.split("-")
	if int(date[1]) != 12:
		date[1] = str(int(date[1])+1)
	else:
		date[1] = "1"
	date = "-".join(date)
	data = FormData()
	data.add_field("status", False)
	data.add_field("use_in_this_month", False)
	data.add_field("next_billing", date)
	async with ClientSession() as session:
		async with session.patch(f"http://subscriptions:8080/api/subscriptions/update/{sub["subscription_id"]}", headers=headers, data=data): pass

def notify():
	os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'backend.settings')
	setup()
	from django.core.mail import send_mail
	from  user.models import User
	from django.conf import settings
	while True:
			for user in User.objects.all():
				if user.notifications:
					header = {"X-User-Id": str(user.id)}
					response = requests.get("http://subscriptions:8080/api/subscriptions/all", headers=header)
					subscriptions = response.json()
					if len(subscriptions) == 0:
						continue
					for sub in subscriptions:
						nb = sub["next_billing"].split("-")
						nb = list(map(int, nb))
						nb_date = datetime.datetime(nb[0], nb[1], nb[2])
						now = datetime.datetime.now()
						delta = nb_date - now
						if delta.days == 1 or delta.days == 0:
							text = f"Здравствуйте, {user.username}, напоминаем вам о предстоящем списании по подписке {sub["name"]}: {sub["url_service"]}"
							send_mail(f"Наипоминание о списании по подписке {sub["name"]}",
							text,
							settings.DEFAULT_FROM_EMAIL,
       						  	[user.email],
       						       	fail_silently=False)
						elif delta.days <= -1:
							asyncio.run(update_sub(sub, header))
       							
			time.sleep(86400)
			
			
notify()
