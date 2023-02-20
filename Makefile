up:
	docker-compose -f ./deployments/docker-compose.yaml up --build

down:
	docker-compose -f ./deployments/docker-compose.yaml down

k8s:
	kubectl apply -f ./deployments/k8s-deployment.yml