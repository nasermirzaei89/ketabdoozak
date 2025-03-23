variable "kubeconfig" {
  type        = string
  description = "Path to the kubeconfig file"
  default     = "~/.kube/config"
}

variable "kubeconfig_context" {
  type        = string
  description = "Kubeconfig context to use"
}

variable "registry_server" {
  type        = string
  description = "Registry server"
}

variable "registry_username" {
  type        = string
  description = "Registry username"
  sensitive   = true
}

variable "registry_password" {
  type        = string
  description = "Registry password"
  sensitive   = true
}

variable "authentication_oidc_issuer_url" {
  type = string
}

variable "authentication_oidc_client_id" {
  type    = string
  default = "api"
}

variable "authentication_oidc_username_claim" {
  type    = string
  default = "preferred_username"
}

variable "authentication_oidc_groups_claim" {
  type    = string
  default = "groups"
}

variable "www_cookie_store_key" {
  type      = string
  sensitive = true
}

variable "www_csrf_auth_key" {
  type      = string
  sensitive = true
}

variable "www_oidc_issuer_url" {
  type = string
}

variable "www_oidc_client_id" {
  type = string
}

variable "www_oidc_client_secret" {
  type      = string
  sensitive = true
}

variable "www_oidc_redirect_url" {
  type = string
}

variable "www_oidc_logout_redirect_url" {
  type = string
}

variable "aws_access_key_id" {
  type      = string
  sensitive = true
}

variable "aws_secret_access_key" {
  type      = string
  sensitive = true
}

variable "file_manager_files_bucket_url" {
  type = string
}

variable "database_dsn" {
  type      = string
  sensitive = true
}

variable "redis_url" {
  type = string
}

variable "redis_password" {
  type      = string
  sensitive = true
}

variable "redis_db" {
  type = string
}

variable "otel_service_name" {
  type = string
}

variable "otel_traces_exporter" {
  type = string
}

variable "otel_metrics_exporter" {
  type = string
}

variable "otel_logs_exporter" {
  type = string
}

variable "otel_traces_sampler" {
  type = string
}

variable "otel_exporter_otlp_endpoint" {
  type = string
}

variable "otel_exporter_otlp_insecure" {
  type = string
}
