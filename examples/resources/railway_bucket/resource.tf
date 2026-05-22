resource "railway_bucket" "example" {
  name       = "assets"
  project_id = railway_project.example.id
}
