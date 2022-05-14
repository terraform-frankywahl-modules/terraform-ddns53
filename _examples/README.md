Example folder for using the module

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_ddns"></a> [ddns](#module\_ddns) | ../. | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_domain"></a> [domain](#input\_domain) | The TLD of the domain we want to create a DDNS for | `string` | `"example.com"` | no |
| <a name="input_subdomain"></a> [subdomain](#input\_subdomain) | The subdomain that needs to be updated with DDNS | `string` | `"test"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_password"></a> [password](#output\_password) | n/a |
| <a name="output_url"></a> [url](#output\_url) | n/a |
| <a name="output_username"></a> [username](#output\_username) | n/a |
