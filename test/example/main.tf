terraform {
  required_providers {
    db = {
      source = "hashicorp.com/kasaikou/db"
    }
  }
}

provider "db" {
  driver      = "postgres"
  data_source = "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
}

data "db_current" "current" {}

output "current" {
  value = data.db_current.current
}
