# How to Deploy

## Create a K3D cluster
Cluster name `gtr`
```
k3d cluster create gtr -p "8081:80@loadbalancer" 
```
## Deploy PostgreSQL

```
helm install postgresql bitnami/postgresql
```
Get PostgreSQL password with:
```
kubectl get secret --namespace default postgresql -o jsonpath="{.data.postgresql-password}" | base64 --decode)
```
## Deploy Auth service
(For development only)

Go to `release/auth/templates/deployment.yaml`

Change `DB_PWD`

Then do
```
helm install --debug gtr-auth release/auth --set controller.publishService.enabled=true
```
