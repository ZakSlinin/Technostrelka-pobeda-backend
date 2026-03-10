FROM alt:p11

COPY backend /backend
WORKDIR /backend
EXPOSE 8080

RUN apt-get update
RUN apt-get install python3 pip postgresql17 python3-module-psycopg2 python3-module-django python3-module-djangorestframework   -y
RUN pip install django-filter django-cors-headers djangorestframework-simplejwt tzdata requests aiohttp

RUN adduser service-user

USER service-user