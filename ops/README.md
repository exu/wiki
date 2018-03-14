Devops ([back to index](../README.md))

# Kubernetes

## Community

### Apps SIG

> "Covers deploying and operating applications in Kubernetes. We focus on the developer and devops experience of running applications in Kubernetes. We discuss how to define and run apps in Kubernetes, demo relevant tools and projects, and discuss areas of friction that can lead to suggesting improvements or feature requests."

- [github](https://github.com/kubernetes/community/tree/master/sig-apps)
- [recordings on YT](https://www.youtube.com/watch?v=hn23Z-vL_cM&list=PL69nYSiGNLP2LMq7vznITnpd2Fk1YIZF3)


## Naming

### Pods

> *“Pods are the smallest deployable units of computing that can be created and managed in Kubernetes.”*

say the official Kubernetes docs for pods. While pods can contain one single container, they are not limited to one and can contain as many containers as needed.

What makes these containers a pod, is that **all containers in a pod run as if they would have been running on a single host in pre-container world**. Thus, they share a set of Linux namespaces and do not run isolated from each other. This results in them sharing an IP address and port space, and being able to find each other over localhost or communicate over the IPC namespace. Further, all containers in a pod have access to shared volumes, that is they can mount and work on the same volumes if needed.

In order to gain all this functionality a pod is a single deployable unit. Each single instance of the pod (with all its containers) is always scheduled together.

[source](https://blog.giantswarm.io/understanding-basic-kubernetes-concepts-i-introduction-to-pods-labels-replicas/)


### Labels & Selectors

Labels are key/value pairs that can be attached to objects, such as pods, but also any other object in Kubernetes, even nodes.

Labels can be used to organize and select subsets of objects. They are often used for example to identify releases (beta, stable), environments (dev, prod), or tiers (frontend, backend).

Using label **selectors** a client or user can identify and subsequently manage a group of objects. This is the core grouping primitive of Kubernetes and used in many places. One example of its use is working with replica sets.


### Replica Sets (and Replication Controllers)

As mentioned above a pod by itself is ephemeral and won’t be rescheduled if the node it is running on goes down. This is where the **replica set** comes in and ensures that a specific number of pod instances (or replicas) are running at any given time. The replica set then **takes care of (re)scheduling your instances** for you.

There is a higher level concept called a deployment, which manages replica sets. Therefore, you usually won’t need to create or manipulate replica set objects directly.


### Deployments

Before deployments, there were replication controllers, which managed pods and ensured a certain number of them were running. Now with deployments we move to replica sets, which are basically the next-generation of replication controllers. Only this time we don’t manage them, but they get managed by the deployments we define.

Thus, the chain is like following:

    Deployment -> Replica Set -> Pod(s).

And we only have to take care of the first.


Additional to what replication controllers (or replica sets) offer, deployments give you **declarative control** over the update strategy used for the deployment. This replaces the old kubectl rolling-update way of updating, but offers the same flexibility in terms of defining `maxSurge` and `maxUnavailable`, i.e. how many additional and how many unavailable pods are allowed. Defining this in a deployment enables you to “spec once use many times”, which helps even more when working in teams or managing a multitude of deployments.

Resources:
* [source](https://blog.giantswarm.io/understanding-basic-kubernetes-concepts-using-deployments-manage-services-declaratively/)
* [creating deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#use-case)


### Services

Services work by defining a **logical set of pods and a policy by which to access them**. The selection of pods is based on label selectors (which we talked about in the first blog post). In case you select multiple pods, the service automatically takes care of load balancing and assigns them a single (virtual) service IP (which you can also set manually).

You can use the selector to choose a group of pods and define a `targetPort` on which to access them. Further helping with abstraction, this `targetPort` can also refer to the name of a port, which gives you more freedom to implement the actual pods behind the service. Even the port of each pod could be different as long as they carry the same name.

Additionally, services **can abstract away other kinds of backends that are not Kubernetes pods**. You can for example abstract away an external database cluster behind a service. This way you can for example use a simple local database for your development environment and a professionally managed database cluster in production without having to change the way that your other services talk to that database service.



## Tools

- **Minikube** is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a VM on your laptop for users looking to try out Kubernetes or develop with it day-to-day. https://github.com/kubernetes/minikube


- **Telepresence** - fast, local development for kubernetes and openshift microservices - proxy for remote kubernetes  - https://www.telepresence.io/

- **Helm** - Package manager for "Kubernetes Applications" https://helm.sh/ "Helm helps you manage Kubernetes applications — Helm Charts helps you define, install, and upgrade even the most complex Kubernetes application. Charts are easy to create, version, share, and publish — so start using Helm and stop the copy-and-paste madness."

## Articles

- Development Workflows:
    - https://dzone.com/articles/a-development-workflow-for-kubernetes-services
    - https://dev.to/datawireio/fast-development-workflow-with-docker-and-kubernetes-1if



## Snippets

### Deployments


#### Creating deployments

Example deployment.yml file for 3 nginx httpd server pods

```yml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
    name: nginx-deployment
    labels:
        app: nginx
    spec:
    replicas: 3
    selector:
        matchLabels:
        app: nginx
    template:
        metadata:
        labels:
            app: nginx
        spec:
        containers:
        - name: nginx
            image: nginx:1.7.9
            ports:
            - containerPort: 80
```

creating new deployment on kubernetes cluster:

    kubectl create -f https://raw.githubusercontent.com/kubernetes/website/master/docs/concepts/workloads/controllers/nginx-deployment.yaml

To get deployments

    kubectl get deployments


To see the Deployment rollout status, run

    kubectl  rollout status deployment/nginx-deployment.

To see labels automatically generated for each pod use:

    kubectl get pods --show-labels



#### Changing deployments

    $ kubectl set image deployment/nginx-deployment nginx=nginx:1.9.1
    deployment "nginx-deployment" image updated

Alternatively, we can edit the Deployment and change `.spec.template.spec.containers[0].image` from `nginx:1.7.9` to `nginx:1.9.1`:

    $ kubectl edit deployment/nginx-deployment
    deployment "nginx-deployment" edited

We can run `kubectl get rs` to see that the Deployment updated the Pods by creating a new `ReplicaSet` and scaling it up to 3 replicas, as well as scaling down the old ReplicaSet to 0 replicas.

    $ kubectl get rs
    NAME                          DESIRED   CURRENT   READY   AGE
    nginx-deployment-1564180365   3         3         3       6s
    nginx-deployment-2035384211   0         0         0       36s

#### Rolling back deployments

When something will fuck up (i.e. invalid `nginx 1.91` version deployed previously):

    $ kubectl rollout history deployment/nginx-deployment
    deployments "nginx-deployment"
    REVISION    CHANGE-CAUSE
    1           kubectl create -f docs/user-guide/nginx-deployment.yaml --record
    2           kubectl set image deployment/nginx-deployment nginx=nginx:1.9.1
    3           kubectl set image deployment/nginx-deployment nginx=nginx:1.91

We can roll back changes to previous revision:

    $ kubectl rollout undo deployment/nginx-deployment
    deployment "nginx-deployment" rolled back

Alternatively, you can rollback to a specific revision by specify that in --to-revision:

    $ kubectl rollout undo deployment/nginx-deployment --to-revision=2
    deployment "nginx-deployment" rolled back


#### Scaling deployment

    $ kubectl scale deployment nginx-deployment --replicas=10
    deployment "nginx-deployment" scaled


Assuming [horizontal pod autoscaling](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/) is enabled in your cluster, you can setup an autoscaler for your Deployment and choose the minimum and maximum number of Pods you want to run based on the CPU utilization of your existing Pods.

    $ kubectl autoscale deployment nginx-deployment --min=10 --max=15 --cpu-percent=80
    deployment "nginx-deployment" autoscaled





# Docker

## Go and docker
- building go based images https://tachingchen.com/blog/building-minimal-docker-image-for-go-applications/ with multistage build example
