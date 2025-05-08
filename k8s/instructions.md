# Kubernetes Deployment Guide 🚀

This guide provides step-by-step instructions for deploying the Pinned Place application to a Kubernetes cluster.

## Prerequisites

Before you begin, ensure you have:
- A running Kubernetes cluster
- `kubectl` configured and connected to your cluster
- Docker registry credentials (if using a private registry)
- Basic understanding of Kubernetes concepts

## 1. Environment Configuration

### 1.1 Create Environment File

1. Navigate to the `k8s` directory:
```bash
cd k8s
```

2. Create your environment file from the template:
```bash
cp .env.example .env
```

3. Edit the `.env` file with your specific configuration values.

### 1.2 Create Kubernetes Secrets

Create a secret containing your environment variables:

```bash
kubectl create secret generic place-secret \
  --from-env-file k8s/.env \
  --namespace=default
```

> **Note**: Replace `default` with your target namespace if different.

## 2. Docker Registry Configuration

If you're using a private Docker registry, you'll need to create a registry secret:

```bash
kubectl create secret docker-registry regsecret \
  --docker-server=$DOCKER_REGISTRY_SERVER \
  --docker-username=$DOCKER_USER \
  --docker-password=$DOCKER_PASSWORD \
  --docker-email=$DOCKER_EMAIL \
  --namespace=default
```

### Required Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `DOCKER_REGISTRY_SERVER` | Docker registry URL | `docker.io` or `ghcr.io` |
| `DOCKER_USER` | Registry username | `your-username` |
| `DOCKER_PASSWORD` | Registry password/token | `your-password` |
| `DOCKER_EMAIL` | Email (optional) | `your-email@example.com` |

## 3. Deployment

### 3.1 Deploy All Resources

To deploy all Kubernetes resources:

```bash
kubectl apply -f ./k8s
```

### 3.2 Verify Deployment

Check the status of your deployment:

```bash
# Check pods
kubectl get pods

# Check services
kubectl get services

# Check deployments
kubectl get deployments
```

## 4. Troubleshooting

### Common Issues

1. **Secret Creation Fails**
   - Ensure all required environment variables are set
   - Verify you have the necessary permissions

2. **Image Pull Errors**
   - Verify registry credentials
   - Check image name and tag
   - Ensure network connectivity to registry

3. **Pod Startup Issues**
   - Check pod logs: `kubectl logs <pod-name>`
   - Describe pod: `kubectl describe pod <pod-name>`

## 5. Maintenance

### Updating Secrets

To update existing secrets:

```bash
# Delete existing secret
kubectl delete secret place-secret

# Create new secret
kubectl create secret generic place-secret --from-env-file k8s/.env
```

### Scaling

To scale your deployment:

```bash
kubectl scale deployment place --replicas=<number>
```

## 6. Cleanup

To remove all deployed resources:

```bash
kubectl delete -f ./k8s
```

## Additional Resources

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Kubernetes Secrets Management](https://kubernetes.io/docs/concepts/configuration/secret/)

---

For more information or support, please open an issue in the repository.
