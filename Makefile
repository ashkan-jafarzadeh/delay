up:
	docker-compose up -d;
down:
	docker-compose stop;

migrate:
	docker-compose exec app ./migrate
seed:
	docker-compose exec app ./seed
test:
	go test -v ./...
