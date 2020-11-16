
provision:
	@echo "Provisioning PNCounter Cluster"	
	bash scripts/provision.sh

pncounter-build:
	@echo "Building PNCounter Docker Image"	
	docker build -t pncounter -f Dockerfile .

pncounter-run:
	@echo "Running Single PNCounter Docker Container"
	docker run -p 8080:8080 -d pncounter

info:
	echo "PNCounter Cluster Nodes"
	docker ps | grep 'pncounter'
	docker network ls | grep pncounter_network

clean:
	@echo "Cleaning PNCounter Cluster"
	docker ps -a | awk '$$2 ~ /pncounter/ {print $$1}' | xargs -I {} docker rm -f {}
	docker network rm pncounter_network

build:
	@echo "Building PNCounter Server"	
	go build -o bin/pncounter main.go

fmt:
	@echo "go fmt PNCounter Server"	
	go fmt ./...

test:
	@echo "Testing PNCounter"	
	go test -v --cover ./...