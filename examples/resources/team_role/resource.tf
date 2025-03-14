resource "cortex_team_role" "engineer" {
  tag                  = "se-1"
  name                 = "Software Engineer 1"
  description          = "A first-level engineer"
  notifcations_enabled = true
}
