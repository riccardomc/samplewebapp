# Sample Web App

A simple web application that I am going to use in examples and tests.

## Build and Run

Use the Makefile target `docker-run` to build and run the application:
```
$ make docker-run
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o samplewebapp .
docker build -t riccardomc/samplewebapp:0.1 .
Sending build context to Docker daemon  8.145MB
Step 1/4 : FROM scratch
 ---> 
Step 2/4 : EXPOSE 8080
 ---> Running in 74951f38410e
 ---> ba14e2de0e94
Removing intermediate container 74951f38410e
Step 3/4 : ADD samplewebapp /
 ---> 23f2ef7e257d
Step 4/4 : CMD /samplewebapp
 ---> Running in 0b12a5ad4bc1
 ---> e0cca6661182
Removing intermediate container 0b12a5ad4bc1
Successfully built e0cca6661182
Successfully tagged riccardomc/samplewebapp:0.1
docker run -p 8080:8080 -ti riccardomc/samplewebapp:0.1

```
the application will be available on `localhost:8080`:
```
$ curl localhost:8080/something
Hi there, I love something!
```
