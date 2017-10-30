# Deploy


```
$ kubectl create -f multi-container.yaml
$ kubectl describe pod multiwebz
[ ... ]
```

## Proxy
```
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

## Service
```
$ kubectl expose pod multiwebz --type=NodePort --port=80
$ kubectl get services
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
kubernetes   ClusterIP   10.0.0.1     <none>        443/TCP        6d
multiwebz    NodePort    10.0.0.250   <none>        80:32491/TCP   1m
$ curl http://192.168.99.100:32491/        
<html>
  <body>
    <h1>Yo!</h1>
    This is some static content
  </body>
</html>
$ curl http://192.168.99.100:32491/app/this
Hi there, got 1 requests, running on multiwebz, I love this!% 
```
