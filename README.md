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
    UAA_ADDRESS: https://uaa.system.example.com/
    OPSMAN_USER: admin
    OPSMAN_PASSWORD: superpasswordthatkillsusall
```

## Raw Data

Note that for sorting reasons, we inject the `\u0001` and `\u0002`. This may not be a great
implementation, but it helped keep the Ops Man and ERT in the order we wanted.

You can query the application using the following route: `http://<Application Route>/versions`

It will then return:
```JSON
{
  "versions":{
    "\u0001Ops Man":{
      "ver":"1.7.13.0"
    },
    "\u0002ERT":{
      "ver":"1.7.21-build.2"
    },
    "MySql Tile":{
      "ver":"1.7.13"
    }
  }
}
```

To include the stemcell versions add a query parameter named `scv` and set it to `true`.
For example `http://<Application Route>/versions?scv=true` would return:

```JSON
{
  "versions":{
    "\u0001Ops Man":{
      "ver":"1.7.13.0"
    },
    "\u0002ERT":{
      "ver":"1.7.21-build.2",
      "sc":"3232.19"
    },
    "MySql Tile":{
      "ver":"1.7.13",
      "sc":"3232.19"
    }
  }
}
```

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
