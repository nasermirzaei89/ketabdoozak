# Ketabdoozak Staging Infrastructure

Use project root as working directory.

## Connect to Cluster

For staging, we used [k3s](https://docs.k3s.io/quick-start)

```shell
kubectl config use-context <fixme>
```

## Create Backend Image

Login to `ghcr.io`.

Then:

```shell
DOCKER_DEFAULT_PLATFORM=linux/amd64 make docker-build
make docker-push
```

## Create terraform variables

Create a `terraform.tfvars` file in the `infra/staging` directory with the desired values.

## Provision

```shell
cd infra/staging
terraform init
```

```shell
terraform apply
```
