postgres-start:
	docker run --name django_acs -p 5700:5700 -e POSTGRES_USER=root -e POSTGRES_PASSWORD_FILE=12345 -d postgres:14-bullseye

postgres-stop:
	docker stop django_acs
	
createdb:
	docker exec -it django_acs createdb --username=root --owner=root teachers_pool

dropdb:
	docker exec -it django_acs dropdb teachers_pool

.PHONY: postgres createdb dropdb