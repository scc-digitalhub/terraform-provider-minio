output "minio_sts_access_key" {
  value = minio_sts_key.minio_key.access_key_id
}

output "minio_sts_secret_key" {
  value = minio_sts_key.minio_key.secret_access_key
  sensitive = true
}
output "minio_sts_session_token" {
  value = minio_sts_key.minio_key.session_token
  sensitive = true
}
