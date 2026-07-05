# Deployment Guide

Applications are deployed to Kubernetes using Deployments.

## Deploy

```bash
kubectl apply -f deployment.yaml
```

## Rolling Update

Deployments use rolling updates by default.

## Rollback

```bash
kubectl rollout undo deployment/<deployment-name>
```

Always verify the rollout:

```bash
kubectl rollout status deployment/<deployment-name>
```