resource "oci_identity_compartment" "test_compartment" {
  compartment_id = var.root_compartment_id
  description    = "Resources for the ${var.service_name} service, supporting the ${var.environment} environment"
  name           = "${var.service_name}-${var.environment}"
  
  freeform_tags = {"service"= var.service_name, "environment"=var.environment}
}