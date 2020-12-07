all: login build push

login:
	docker login

build:
	docker build -t happilymarrieddadudemy/node-web-app .

run:
	docker run -p 8000:8000 -d happilymarrieddadudemy/node-web-app

push:
	docker push happilymarrieddadudemy/node-web-app
