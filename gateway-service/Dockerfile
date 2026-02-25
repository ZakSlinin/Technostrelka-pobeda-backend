FROM alt:p11

COPY ./code /code
WORKDIR /code
EXPOSE 8080

RUN apt-get update
RUN apt-get install python3 pip -y
RUN pip install aiohttp aiohttp-cors

RUN adduser service-user
USER service-user
CMD ["python3", "gateway.py"]