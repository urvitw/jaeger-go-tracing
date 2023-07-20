all: start

.PHONY: service-a
service-a:
	@docker build -t service-a -f service-a/Dockerfile .

.PHONY: service-b
service-b:
	@docker build -t service-b -f service-b/Dockerfile .

.PHONY: service-c
service-c:
	@docker build -t service-c -f service-c/Dockerfile .

.PHONY: service-d
service-d:
	@docker build -t service-d -f service-d/Dockerfile .

.PHONY: service-e
service-e:
	@docker build -t service-e -f service-e/Dockerfile .

.PHONY: start
start: service-a service-b service-c service-d service-e
	@docker-compose up -d --remove-orphans

.PHONY: stop
stop:
	@docker-compose down
	docker rmi service-a service-b service-c service-d service-e
