start:
	docker-compose up -d

start-build:
	docker-compose up --build -d

stop:
	docker-compose down 

restart:
	stop start