# HELM the "Package Manager" for Kubernetes apps [back to index](/)

https://helm.sh/

## Running

    helm init

will install Helm in current Kubernetes cluster

    helm create mychart

will create new directories and sample files for your chart.

Next remove all files from templates directory (BTW they're go `text/template` based) and start simplier

```yml
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
data:
  myvalue: "Hello World"
  drink: {{ .Values.favorite.drink }}
  drink: {{ .Values.favorite.food }}
```

We've created sample Kubernetes configMap for use in our cluster (it looks like the one from simpliest cluster parts in Kubernetes)

We can also add some values inside our `values.yml` file.

```yml
favorite:
  drink: coffee
  food: pizza
```

We can make a dry run first before running it on our cluster.

    helm install --debug --dry-run --set favorite.food=beer ./mychart
