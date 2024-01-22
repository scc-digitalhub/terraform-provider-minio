resource "minio_sts_key" "minio_key" {
  oidc_access_token = var.oidc_access_token
  oidc_id_token = var.oidc_access_token
}


