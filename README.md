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
> go run main.go -opsmanuaa https://172.28.61.5/uaa -opsman https://172.28.61.5 -opsmanuser admin -opsmanpassword welcome1 -appsmanuaa https://uaa.system.lab06.den.ecsteam.io -appsmanpassword JPTvnTxOR_6-5YPfhUA3EWruvg1a26rg -appsmanuser admin -appsman https://api.system.lab06.den.ecsteam.io
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
    OPSMAN_PASSWORD:  <ops-man-password>
    APPSMAN_UAA_ADDRESS: https://<apps-man-uaa>
    APPSMAN_ADDRESS: https://<apps-man-api>
    APPSMAN_USER: admin
    APPSMAN_PASSWORD:  <apps-man-password>
```





    OPSMAN_UAA_ADDRESS: https://172.28.61.5/uaa
    OPSMAN_ADDRESS: https://172.28.61.5
    OPSMAN_USER: admin
    OPSMAN_PASSWORD:  welcome1
    APPSMAN_UAA_ADDRESS: https://uaa.system.lab06.den.ecsteam.io
    APPSMAN_ADDRESS: https://api.system.lab06.den.ecsteam.io

    # TODO: We should get this from UAA
    APPSMAN_USER: admin
    APPSMAN_PASSWORD:  JPTvnTxOR_6-5YPfhUA3EWruvg1a26rg

    BOSH_ADDRESS:

    # TODO: We should get this from UAA
    BOSH_USER: director
    BOSH_PASSWORD: PBa0FoQtB0iFfLjWnDfTniAJtj_rFtM0








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








/*
/////////
	"/api/v0/deployed/products/:product-guid/variables"

	{
	  "variables": ["first-variable", "second-variable", "third-variable"]
	}

/////////
	"/api/v0/deployed/products/:product-guid/variables?name=:variable_name"

	{
  	"credhub-password": "example-password"
	}

/////////
	"/api/v0/deployed/products/:product_guid/status"

	{
	  "status": [
	    {
	      "job-name": "web_server-7f841fc2af9c2b357cc4",
	      "index": 0,
	      "az_guid": "ee61aa1e420ed3fdf276",
	      "az_name": "first-az",
	      "ips": [
	        "10.85.42.58"
	      ],
	      "cid": "vm-448ef313-86ee-4049-87cf-764ca2fa97e7",
	      "load_avg": [
	        "0.00",
	        "0.01",
	        "0.03"
	      ],
	      "cpu": {
	        "sys": "0.1",
	        "user": "0.2",
	        "wait": "0.3"
	      },
	      "memory": {
	        "kb": "60632",
	        "percent": "6"
	      },
	      "swap": {
	        "kb": "0",
	        "percent": "0"
	      },
	      "system_disk": {
	        "inode_percent": "31",
	        "percent": "42"
	      },
	      "ephemeral_disk": {
	        "inode_percent": "0",
	        "percent": "1"
	      },
	      "persistent_disk": {
	        "inode_percent": "0",
	        "percent": "0"
	      }
	    }
	  ]
	}

/////////
"/api/v0/deployed/products/:product_guid/static_ips"
[
{
	"name": "job-type1-guid-partition-default-az-guid",
	"ips": [
		"192.168.163.4"
	]
},
{
	"name": "credentials-job-guid-partition-default-az-guid",
	"ips": [
		"192.168.163.7"
	]
}
]

//////////
/api/v0/disk_types

{
  "disk_types": [
    {
      "name": "1024",
      "builtin": true,
      "size_mb": 1024
    },
    {
      "name": "2048",
      "builtin": true,
      "size_mb": 2048
    },
    {
      "name": "5120",
      "builtin": true,
      "size_mb": 5120
    }
  ]
}



//////////////
/api/v0/vm_types

{
  "vm_types": [
    {
      "name": "nano",
      "ram": 512,
      "cpu": 1,
      "ephemeral_disk": 1024,
      "builtin": true
    },
    {
      "name": "micro",
      "ram": 1024,
      "cpu": 1,
      "ephemeral_disk": 2048,
      "builtin": true
    },
    {
      "name": "small.disk",
      "ram": 2048,
      "cpu": 1,
      "ephemeral_disk": 16384,
      "builtin": true
      }
  ]
}


///////////////////
/api/v0/deployed/products/:product_guid/manifest

/api/v0/diagnostic_report
*/
