.PHONY: docker clean

docker: Dockerfile default.conf
	docker build -t riccardomc/webserver:0.1 .

clean:
	docker rmi -f riccardomc/webserver:0.1 | true
