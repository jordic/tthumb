

VERSION = R1


build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tthumb github.com/jordic/tthumb


docker_build:
	docker build -t jordic/tthumb:$(VERSION) .


docker_push:
	docker push jordic/tthumb:$(VERSION)
