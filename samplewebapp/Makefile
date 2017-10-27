.PHONY: docker docker-run clean

samplewebapp: main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o samplewebapp .

docker: Dockerfile samplewebapp
	docker build -t riccardomc/samplewebapp:0.1 .

docker-run: docker
	docker run -p 9990:8080 -ti riccardomc/samplewebapp:0.1

clean:
	rm -rf samplewebapp
	docker rmi -f riccardomc/samplewebapp:0.1 | true
