resource "kubernetes_namespace" "ketabdoozak-staging" {
  metadata {
    name = "ketabdoozak-staging"
  }
}

resource "kubernetes_secret" "ketabdoozak-staging-ghcr-registry-credentials" {
  metadata {
    name      = "ghcr-registry-credentials"
    namespace = kubernetes_namespace.ketabdoozak-staging.metadata[0].name
  }

  type = "kubernetes.io/dockerconfigjson"

  data = {
    ".dockerconfigjson" = jsonencode({
      auths = {
        (var.registry_server) = {
          "username" = var.registry_username
          "password" = var.registry_password
          "auth"     = base64encode("${var.registry_username}:${var.registry_password}")
        }
      }
    })
  }

  depends_on = [kubernetes_namespace.ketabdoozak-staging]
}

resource "kubernetes_secret" "ketabdoozak-staging-authentication" {
  metadata {
    name      = "authentication"
    namespace = kubernetes_namespace.ketabdoozak-staging.metadata[0].name
  }

  data = {
    AUTHENTICATION_OIDC_ISSUER_URL     = var.authentication_oidc_issuer_url
    AUTHENTICATION_OIDC_CLIENT_ID      = var.authentication_oidc_client_id
    AUTHENTICATION_OIDC_USERNAME_CLAIM = var.authentication_oidc_username_claim
    AUTHENTICATION_OIDC_GROUPS_CLAIM   = var.authentication_oidc_groups_claim
  }

  depends_on = [kubernetes_namespace.ketabdoozak-staging]
}

resource "kubernetes_secret" "ketabdoozak-staging-www" {
  metadata {
    name      = "www"
    namespace = kubernetes_namespace.ketabdoozak-staging.metadata[0].name
  }

  data = {
    WWW_COOKIE_STORE_KEY         = var.www_cookie_store_key
    WWW_OIDC_ISSUER_URL          = var.www_oidc_issuer_url
    WWW_OIDC_CLIENT_ID           = var.www_oidc_client_id
    WWW_OIDC_CLIENT_SECRET       = var.www_oidc_client_secret
    WWW_OIDC_REDIRECT_URL        = var.www_oidc_redirect_url
    WWW_OIDC_LOGOUT_REDIRECT_URL = var.www_oidc_logout_redirect_url
  }

  depends_on = [kubernetes_namespace.ketabdoozak-staging]
}

resource "kubernetes_secret" "ketabdoozak-staging-minio" {
  metadata {
    name      = "minio"
    namespace = kubernetes_namespace.ketabdoozak-staging.metadata[0].name
  }

  data = {
    username = var.aws_access_key_id
    password = var.aws_secret_access_key
  }

  depends_on = [kubernetes_namespace.ketabdoozak-staging]
}

resource "kubernetes_secret" "ketabdoozak-staging-file-manager" {
  metadata {
    name      = "file-manager"
    namespace = kubernetes_namespace.ketabdoozak-staging.metadata[0].name
  }

  data = {
    FILE_MANAGER_FILES_BUCKET_URL = var.file_manager_files_bucket_url
  }

  depends_on = [kubernetes_namespace.ketabdoozak-staging]
}

resource "kubernetes_secret" "ketabdoozak-staging-database" {
  metadata {
    name      = "database"
    namespace = kubernetes_namespace.ketabdoozak-staging.metadata[0].name
  }

  data = {
    dsn = var.database_dsn
  }

  depends_on = [kubernetes_namespace.ketabdoozak-staging]
}

resource "kubernetes_secret" "ketabdoozak-staging-redis" {
  metadata {
    name      = "redis"
    namespace = kubernetes_namespace.ketabdoozak-staging.metadata[0].name
  }

  data = {
    url      = var.redis_url
    password = var.redis_password
    db       = var.redis_db
  }

  depends_on = [kubernetes_namespace.ketabdoozak-staging]
}

resource "kubernetes_secret" "ketabdoozak-staging-otel" {
  metadata {
    name      = "otel"
    namespace = kubernetes_namespace.ketabdoozak-staging.metadata[0].name
  }

  data = {
    OTEL_SERVICE_NAME           = var.otel_service_name
    OTEL_TRACES_EXPORTER        = var.otel_traces_exporter
    OTEL_METRICS_EXPORTER       = var.otel_metrics_exporter
    OTEL_LOGS_EXPORTER          = var.otel_logs_exporter
    OTEL_TRACES_SAMPLER         = var.otel_traces_sampler
    OTEL_EXPORTER_OTLP_ENDPOINT = var.otel_exporter_otlp_endpoint
    OTEL_EXPORTER_OTLP_INSECURE = var.otel_exporter_otlp_insecure
  }

  depends_on = [kubernetes_namespace.ketabdoozak-staging]
}


resource "helm_release" "ketabdoozak" {
  chart = "../helm"
  name  = "ketabdoozak"

  namespace = kubernetes_namespace.ketabdoozak-staging.metadata[0].name

  values = [
    yamlencode({
      imagePullSecrets = [
        {
          name = "ghcr-registry-credentials"
        }
      ]
      image = {
        tag        = "latest"
        pullPolicy = "Always"
      }
      fullnameOverride : "ketabdoozak"
      env = [
        {
          name : "ENV"
          value : "staging"
        }
      ]
      ingress = {
        enabled : true
        className : "traefik"
        annotations = {
          "cert-manager.io/cluster-issuer" = "letsencrypt-prod"
        }
        hosts = [
          {
            host : "ketabdoozak.eu-central-1.applicaset.page"
            paths = [
              {
                path : "/"
                pathType : "ImplementationSpecific"
              }
            ]
          }
        ]
        tls = [
          {
            secretName : "ketabdoozak.eu-central-1.applicaset.page-tls"
            hosts : ["ketabdoozak.eu-central-1.applicaset.page"]
          }
        ]
      }
    })
  ]

  depends_on = [
    kubernetes_namespace.ketabdoozak-staging,
    kubernetes_secret.ketabdoozak-staging-ghcr-registry-credentials,
    kubernetes_secret.ketabdoozak-staging-authentication,
    kubernetes_secret.ketabdoozak-staging-www,
    kubernetes_secret.ketabdoozak-staging-minio,
    kubernetes_secret.ketabdoozak-staging-file-manager,
    kubernetes_secret.ketabdoozak-staging-database,
    kubernetes_secret.ketabdoozak-staging-redis,
    kubernetes_secret.ketabdoozak-staging-otel,
  ]
}
