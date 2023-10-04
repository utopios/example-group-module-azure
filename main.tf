resource "azurerm_resource_group" "example_group" {
  name = var.group_name
  location = var.location_group
}