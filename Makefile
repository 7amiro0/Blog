up:
	docker-compose -f ./deployments/docker-compose.yaml -f ./deployments/docker-compose-prod.yaml --env-file ./deployments/.env up --build

down:
	docker-compose -f ./deployments/docker-compose.yaml -f ./deployments/docker-compose-prod.yaml --env-file ./deployments/.env down