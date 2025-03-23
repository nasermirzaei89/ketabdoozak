# Development

Use project root as working directory.

## Create Cluster

```shell
kind create cluster --name ketabdoozak-dev --image kindest/node:v1.32.0
```

## Create Image

```shell
make tailwindcss-build
make npm-install
make npm-build
make templ-generate
make generate-docs
IMAGE_TAG=dev make docker-build
```

load image into kind cluster

```shell
kind load docker-image ghcr.io/nasermirzaei89/ketabdoozak:dev --name ketabdoozak-dev
```

## Create Namespace

```shell
kubectl create namespace ketabdoozak
```

## Load Env

Copy `.env.example` to `.env` and update variables.

Then load environment variables to the shell.

```shell
[ -f .env ] && source .env
```

## Run Postgres

```shell
helm upgrade --install postgres oci://registry-1.docker.io/bitnamicharts/postgresql \
  --namespace postgres --create-namespace \
  --set auth.username=$POSTGRES_USERNAME \
  --set auth.password=$POSTGRES_PASSWORD \
  --set auth.database=$POSTGRES_DBNAME \
  --version 16.4.3
```

## Run Redis

```shell
helm upgrade --install redis oci://registry-1.docker.io/bitnamicharts/redis \
  --namespace redis --create-namespace \
  --set auth.password=$REDIS_PASSWORD \
  --set architecture=standalone \
  --version 20.11.3
```

## Run Minio

```shell
helm upgrade --install minio oci://registry-1.docker.io/bitnamicharts/minio \
  --namespace minio --create-namespace \
  --set auth.rootUser=$AWS_ACCESS_KEY_ID \
  --set auth.rootPassword=$AWS_SECRET_ACCESS_KEY \
  --version 14.10.3
```

## Run Jaeger

```shell
helm upgrade --install jaeger oci://registry-1.docker.io/bitnamicharts/jaeger \
  --namespace jaeger --create-namespace \
  --version 5.1.2
```

## Run Keycloak

```shell
helm upgrade --install keycloak oci://registry-1.docker.io/bitnamicharts/keycloak \
  --namespace keycloak --create-namespace \
  --values infra/dev/keycloak-values.yaml \
  --version 24.4.10
```

## Create Secrets

```shell
kubectl create secret generic authentication \
  --from-literal=AUTHENTICATION_OIDC_ISSUER_URL=http://keycloak.keycloak/realms/ketabdoozak \
  --from-literal=AUTHENTICATION_OIDC_CLIENT_ID=api \
  --from-literal=AUTHENTICATION_OIDC_USERNAME_CLAIM=$AUTHENTICATION_OIDC_USERNAME_CLAIM \
  --from-literal=AUTHENTICATION_OIDC_GROUPS_CLAIM=$AUTHENTICATION_OIDC_GROUPS_CLAIM \
  --namespace ketabdoozak \
  --dry-run=client -o yaml | kubectl apply -f -
```

```shell
kubectl create secret generic www \
  --from-literal=WWW_COOKIE_STORE_KEY=$WWW_COOKIE_STORE_KEY \
  --from-literal=WWW_CSRF_AUTH_KEY=$WWW_CSRF_AUTH_KEY \
  --from-literal=WWW_OIDC_ISSUER_URL=http://keycloak.keycloak/realms/ketabdoozak \
  --from-literal=WWW_OIDC_CLIENT_ID=www \
  --from-literal=WWW_OIDC_CLIENT_SECRET=$WWW_OIDC_CLIENT_SECRET \
  --from-literal=WWW_OIDC_REDIRECT_URL=http://api-dev.ketabdoozak/www/auth/callback \
  --from-literal=WWW_OIDC_LOGOUT_REDIRECT_URL=http://api-dev.ketabdoozak \
  --namespace ketabdoozak \
  --dry-run=client -o yaml | kubectl apply -f -
```

```shell
kubectl create secret generic minio --namespace ketabdoozak \
  --from-literal=username=$AWS_ACCESS_KEY_ID \
  --from-literal=password=$AWS_SECRET_ACCESS_KEY \
  --dry-run=client -o yaml | kubectl apply -f -
```

```shell
kubectl create secret generic file-manager --namespace ketabdoozak \
  --from-literal=FILE_MANAGER_FILES_BUCKET_URL=$FILE_MANAGER_FILES_BUCKET_URL \
  --dry-run=client -o yaml | kubectl apply -f -
```

```shell
kubectl create secret generic database --namespace ketabdoozak \
  --from-literal=dsn=$DATABASE_DSN \
  --dry-run=client -o yaml | kubectl apply -f -
```

```shell
kubectl create secret generic redis --namespace ketabdoozak \
  --from-literal=url=redis-master.redis:6379 \
  --from-literal=password=$REDIS_PASSWORD \
  --from-literal=db=$REDIS_DB \
  --dry-run=client -o yaml | kubectl apply -f -
```

```shell
kubectl create secret generic otel --namespace ketabdoozak \
  --from-literal=OTEL_SERVICE_NAME=$OTEL_SERVICE_NAME \
  --from-literal=OTEL_TRACES_EXPORTER=$OTEL_TRACES_EXPORTER \
  --from-literal=OTEL_METRICS_EXPORTER=$OTEL_METRICS_EXPORTER \
  --from-literal=OTEL_LOGS_EXPORTER=$OTEL_LOGS_EXPORTER \
  --from-literal=OTEL_TRACES_SAMPLER=$OTEL_TRACES_SAMPLER \
  --from-literal=OTEL_EXPORTER_OTLP_ENDPOINT=http://jaeger-collector.jaeger:4318 \
  --from-literal=OTEL_EXPORTER_OTLP_INSECURE=$OTEL_EXPORTER_OTLP_INSECURE \
  --dry-run=client -o yaml | kubectl apply -f -
```

## Deploy

```shell
helm upgrade --install backend ./infra/helm \
  --namespace ketabdoozak \
  --values infra/dev/values.yaml
```

## Access with Telepresence

Install from https://www.telepresence.io/docs/install/client

Install the Traffic Manager

```shell
telepresence helm install --request-timeout 1m
```

Connect to the cluster

```shell
telepresence connect
```

Visit website: http://api-dev.ketabdoozak

Visit swagger: http://api-dev.ketabdoozak/swagger/index.html

Get client id for swagger from: https://manage.auth0.com/dashboard/eu/ketabdoozak/applications/<APP_ID>/settings

Visit jaeger: http://jaeger-query.jaeger:16686

Visit minio: http://minio.minio:9001

## Cleanup

To remove everything, run:

```shell
kind delete cluster --name ketabdoozak-dev
```
