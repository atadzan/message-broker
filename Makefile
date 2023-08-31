run-rabbit:
	docker run --rm -it -p 15672:15672 -p 5672:5672 -d rabbitmq:3-management