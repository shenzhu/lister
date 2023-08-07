Commands
===
```
docker build -t lister:0.1.2 .
docker tag 6d57b9f8d84f shenzhuinsta/lister:0.1.2
docker push shenzhuinsta/lister:0.1.2

k create role poddepl --resource pods,deployments --verb list

k create rolebinding poddepl --role poddepl --serviceaccount default:default
```

Notes
===

Kubernetes Objects
---
Any go struct that implements
- `GetObjectKind() GroupVersionKind`
- `SetGroupVersionKind(kind GroupVersionKind)`
- `DeepCopyObject() runtime.Object`

can be called a Kubernetes Object.

Pod uses typeMeta to implement the first two methods, and implements the third method on its own.

```
k8s object
    typeMeta
        kind
        apiversion
    objectmeta
    spec
    status
        true, done
```

```
apiVersion: apps/v1
kind: Deployment
```

API Machinery
---

## Kind
Mostly singlar nouns, but not always, use camel case

Pod\
Deployment


Deployment\
apps/v1

GroupVersionKind

Kinds are not mapped 1-1 to http endpoints

## Resource
Samller case, plural nouns, mapped 1-1 to HTTP endpoints

deployments\
apps/v1


apis/apps/v1/namespaces/default/deployments

replicasets\
apis/apps/v1/namespaces/default/replicasets?limit=500

GroupVersionResource

## RestMapping
https://pkg.go.dev/k8s.io/apimachinery/pkg/api/meta#RESTMapper

## Scheme

Convert go struct to GroupVersionKinds: https://pkg.go.dev/k8s.io/apimachinery@v0.27.4/pkg/runtime#Scheme.ObjectKinds, this is only going to work if the Kubernetes object is registered, for which AddKnownTypes can be used

