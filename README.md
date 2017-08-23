# pcf-status

This application will return useful information about a PCF install.

## Installing

Simply push the application to you cloud foundry.
```Bash
> cf push -f manifest.yml
```

for some of the older foundations, you may need to specify the go buildpack:
```Bash
> cf push -f manifest.yml -b https://github.com/cloudfoundry/go-buildpack.git
```

```
> glide install
```


## Configuration

Apply the following configuration to the application instance.

| Environment Variable | Required | Type    | Details |
|:--------------------:|----------|:-------:|:--------|
| `UAA_ADDRESS`        |   Yes    |  URL    | URL to the UAA and Ops Manager instance.          |
| `OPSMAN_USER`        |   Yes    | String  | Ops Man user that you want to use.                |
| `OPSMAN_PASSWORD`    |   Yes    | String  | Password to the Ops Man user.                     |
| `DEBUG`              |    No    | Boolean | Show useful info while it is executing. |

### Sample Manifest

``` YAML
applications:
- name: pcf-status
  host: status
  memory: 32M
  buildpack: go_buildpack
  instances: 1
  env:
    OPSMAN_UAA_ADDRESS: https://<ops-man-ip>/uaa
    OPSMAN_ADDRESS: https://<ops-man-ip>
    OPSMAN_USER: admin
```






## Routes


## Badges
Badges have been removed and another project controls them now: https://github.com/ECSTeam/status-badge


https://apidocs.cloudfoundry.org/

GET /v2/apps
GET /v2/apps/ac6a28be-42d2-49da-9a69-65da93a0e505/stats
GET /v2/apps/cd897c8c-3171-456d-b5d7-3c87feeabbd1/summary
GET /v2/apps/ed5512e3-511f-49b3-b30a-b3630e782d03/instances
GET /v2/apps/684c2eae-28db-45f2-8ad5-fe26e2d9f60c/routes
GET /v2/apps/2a3820bb-febd-4c90-ab66-80faa4362142/service_bindings

GET /v2/app_usage_events?results-per-page=1&after_guid=5a3416b0-cf3c-425a-a14c-45a317c497ed

GET /v2/info

GET /v2/private_domains

GET /v2/routes

GET /v2/routes/81464707-0f48-4ab9-87dc-667ef15489fb/apps

GET /v2/route_mappings

GET /v2/security_groups
GET /v2/security_groups/1452e164-0c3e-4a6c-b3c3-c40ad9fd0159/spaces

GET /v2/service_brokers

GET /v2/service_instances

GET /v2/services

GET /v2/spaces
GET /v2/stacks
GET /v2/users
GET /v2/users/uaa-id-309/spaces
GET /v2/users/uaa-id-299/organizations










// https://opsman-dev-api-docs.cfapps.io/?shell#retrieving-status-of-product-jobs

// https://opsman-dev-api-docs.cfapps.io/?shell#listing-static-ip-assignments-for-product-jobs

// https://opsman-dev-api-docs.cfapps.io/?shell#retrieving-manifest-for-a-deployed-product
// GET /api/v0/deployed/products/:product_guid/manifest

