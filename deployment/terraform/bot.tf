resource "kubernetes_deployment" "bot" {
  metadata {
    name = var.gke_deployment_name
    labels = {
      app = var.gke_deployment_name
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        app = var.gke_deployment_name
      }
    }
    template {
      metadata {
        labels = {
          app = var.gke_deployment_name
        }
      }
      spec {
        container {
          name    = var.gke_deployment_name
          image   = var.gcr_image_name
          command = ["./nxd"]
          args    = ["bot"]
          env {
            name  = "NAKA_DISCORD_TOKEN"
            value = var.naka_discord_token
          }
          env {
            name  = "NXD_DISCORD_PREFIX"
            value = var.nxd_discord_prefix
          }
          env {
            name  = "NXD_DISCORD_TOKEN"
            value = var.nxd_discord_token
          }
          env {
            name  = "NXD_DB_URI"
            value = var.nxd_db_uri
          }
          env {
            name  = "NXD_DB_NAME"
            value = var.nxd_db_name
          }
          env {
            name  = "NXD_DB_USER"
            value = var.nxd_db_user
          }
          env {
            name  = "NXD_DB_PASSWORD"
            value = var.nxd_db_password
          }
          env {
            name  = "NXD_CRON_INTERVAL"
            value = var.nxd_cron_interval
          }
          env {
            name  = "NXD_LOG_JSON"
            value = var.nxd_log_json
          }
          env {
            name  = "NXD_LOG_LEVEL"
            value = var.nxd_log_level
          }
          env {
            name  = "NXD_NEWRELIC_NAME"
            value = var.nxd_newrelic_name
          }
          env {
            name  = "NXD_NEWRELIC_LICENSE_KEY"
            value = var.nxd_newrelic_license_key
          }
        }
      }
    }
  }
}
