variable "gcp_project_id" {
  type        = string
  description = "GCP project id"
}

variable "gcp_region" {
  type        = string
  description = "GCP project region"
}

variable "gke_cluster_name" {
  type        = string
  description = "GKE cluster name"
}

variable "gke_location" {
  type        = string
  description = "GKE location"
}

variable "gke_pool_name" {
  type        = string
  description = "GKE node pool name"
}

variable "gke_node_preemptible" {
  type        = bool
  description = "GKE node preemptible"
}

variable "gke_node_machine_type" {
  type        = string
  description = "GKE node machine type"
}

variable "gke_node_disk_size_gb" {
  type        = number
  description = "GKE node disk size in gb"
}

variable "gcr_image_name" {
  type        = string
  description = "GCR image name"
}

variable "gke_deployment_name" {
  type        = string
  description = "GKE deployment bot name"
}

variable "gke_cron_name" {
  type        = string
  description = "GKE cron name"
}

variable "gke_cron_schedule" {
  type        = string
  description = "GKE cron schedule"
}

variable "nxd_discord_prefix" {
  type        = string
  description = "Discord command prefix"
}

variable "nxd_discord_token" {
  type        = string
  description = "Discord bot token"
}

variable "nxd_db_uri" {
  type        = string
  description = "Database URI"
}

variable "nxd_db_name" {
  type        = string
  description = "Database name"
}

variable "nxd_db_user" {
  type        = string
  description = "Database name"
}

variable "nxd_db_password" {
  type        = string
  description = "Database password"
}

variable "nxd_cron_interval" {
  type        = string
  description = "Cron interval"
}

variable "nxd_log_json" {
  type        = bool
  description = "Log json"
}

variable "nxd_log_level" {
  type        = number
  description = "Log level"
}

variable "nxd_newrelic_name" {
  type        = string
  description = "Newrelic name"
}

variable "nxd_newrelic_license_key" {
  type        = string
  description = "Newrelic license key"
}
