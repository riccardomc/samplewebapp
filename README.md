# Deploy

```
$ kubectl create -f multi-container.yaml
$ kubectl describe pod multiwebz
[ ... ]
$ kubectl proxy
Starting to serve on 127.0.0.1:8001
$ curl http://localhost:8001/api/v1/proxy/namespaces/default/pods/multiwebz/
<html>
  <body>
    <h1>Yo!</h1>
    This is some static content
  </body>
</html>
$ curl http://localhost:8001/api/v1/proxy/namespaces/default/pods/multiwebz/app/this
$ Hi there, got 1 requests, running on multiwebz, I love this!% 
```
