start-network:
	docker network create wisdom-net

build-server-docker-image:
	docker build -f Dockerfile-server -t wisdom-server-image .

build-client-docker-image:
	docker build -f Dockerfile-client -t wisdom-client-image .

start-server:
	docker run --network wisdom-net --name wisdom-server -p 8080:8080 wisdom-server-image

start-client:
	docker run --network wisdom-net --name wisdom-client wisdom-client-image

stop-server:
	docker stop wisdom-server

stop-client:
	docker stop wisdom-client