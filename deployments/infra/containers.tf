resource "oci_artifacts_container_repository" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  display_name   = "${var.service_name}-${var.environment}"
  is_public      = false
}