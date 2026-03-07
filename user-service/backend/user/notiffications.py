import requests
import os
import time
import datetime
import django

def notify():
	os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'backend.settings')
	django.setup()
	from django.core.mail import send_mail
	from  .models import User
	while True:
			for user in User.objects.all():
				if user.notifications:
					header = {"X-User-Id": str(user.user_id)}
					response = requests.get("http://subscriptions:8080/api/subscriptions/all", headers=header)
					subscriptions = response.json()
					if len(subscriptions) == 0:
						continue
					for sub in subscriptions:
						nb = sub["next_billing"].split("-")
						nb = list(map(int, nb))
						try:
							nb_date = datetime.datetime(nb[0], nb[1], nb[2])
							now = datetime.datetime.now()
							delta = now - nb_date
							if delta.days <= 1:
								text = f"Здравствуйте, {user.username}, напоминаем вам о предстоящем списании по подписке {sub["name"]}: {sub["url_service"]}"
								send_mail(f"Наипоминание о списании по подписке {sub["name"]}",
								text,
								env("EMAIL"),
       						  		[user.email],
       						       		fail_silently=False)
						except Exception as e:
							print(e)
			time.sleep(3600)
			