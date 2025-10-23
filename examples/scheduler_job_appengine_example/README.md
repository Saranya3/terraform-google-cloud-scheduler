# Cloud Scheduler Job with App Engine Target Example

This example demonstrates how to create a Google Cloud Scheduler job using the `terraform-google-cloud-scheduler` module.

The job is configured to target a specific App Engine HTTP endpoint every 4 minutes.

## Prerequisites

*   A Google Cloud Project.
*   A Service Account named `ci-account` in your project. The example in `main.tf` is hardcoded to use `ci-account@${var.project_id}.iam.gserviceaccount.com`. This specific name is often created by the test setup process (e.g., `make docker_test_prepare`) in this repository.
*   If you are not using the `make` based setup, ensure a service account with the name `ci-account` exists and has the necessary permissions (e.g. permissions to invoke the target App Engine service). Alternatively, modify the `service_account_email` in `examples/simple_example/main.tf` to point to an existing service account.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| name | The name of the job | `string` | n/a | yes |
| project\_id | The project ID to deploy to | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| id | An identifier for the resource |
| state | State of the job |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

To provision this example, run the following from within this directory:
- `terraform init` to get the plugins
- `terraform plan` to see the infrastructure plan
- `terraform apply` to apply the infrastructure build
- `terraform destroy` to destroy the built infrastructure
