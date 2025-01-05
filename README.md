### K8s Chart playground

### How to deploy the application to the Kubernetes cluster:
1. Create namespace (if you don't want to use default namespace)
```bash
kubectl create namespace <namespace>
```

2. Install Helm chart
```bash
cd k8s/chart
helm install <release-name> ./example -n <namespace>
```

You can override default values by providing your own values file or using `set` flag.
