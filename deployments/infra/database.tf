resource "oci_nosql_table" "resource_action_schedule" {
  #Required
  compartment_id = oci_identity_compartment.primary.id
  name           = "resource_action_schedule"
  ddl_statement  = "CREATE TABLE resource_action_schedule ( id string, name string, schedule string, action string, resources string, region string, PRIMARY KEY ( id ) )"

  table_limits {
    capacity_mode      = "ON_DEMAND"
    max_storage_in_gbs = 1

    max_read_units  = 0
    max_write_units = 0
  }
}