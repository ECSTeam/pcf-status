# pcf-status

This application will return useful information about a PCF install.

## Installing

Simply push the application to you cloud foundry.
```Bash
> cf push -f manifest.yml
```

For some of the older foundations, you may need to specify the go buildpack:
```Bash
> cf push -f manifest.yml -b https://github.com/cloudfoundry/go-buildpack.git
```

Add dependancies using glide.
```
> glide install
```

## Configuration

Apply the following configuration to the application instance.

| Environment Variable | Required | Details                                  |
|:--------------------:|----------|:-----------------------------------------|
| `OPSMAN`             |   Yes    | A JSON object with an `address` of the opsman server, the `user` and `passwrd` of the user to use. |

### Sample Manifest

``` YAML
applications:
- name: pcf-status
  host: status
  memory: 32M
  buildpack: go_buildpack
  instances: 1
  env:
    OPSMAN: |
      {
        "address": "<opsman-address>",
        "user": "<user>",
        "password": "<password>"
      }
```

## API Routes

### [GET] `/models/buildpacks`

Get the collection of buildpacks

#### Sample Response
```JSON
[
    {
        "name": "staticfile_buildpack",
        "filename": "staticfile_buildpack-cached-v1.4.11.zip",
        "guid": "516ba0d6-1e27-4c9b-bbe2-fb9c6d00a940",
        "created_at": "2017-08-04T13:38:24Z",
        "updated_at": "2017-08-04T13:38:24Z"
    },
    {
        "name": "java_buildpack_offline",
        "filename": "java-buildpack-offline-v3.18.zip",
        "guid": "3cdef2b7-1dbf-4f9f-a6fb-a7a57e294a2c",
        "created_at": "2017-08-04T13:38:24Z",
        "updated_at": "2017-08-04T13:38:28Z"
    },
    {
        "name": "ruby_buildpack",
        "filename": "ruby_buildpack-cached-v1.6.44.zip",
        "guid": "567c219c-be1f-49ec-b852-3794e5e4e4a7",
        "created_at": "2017-08-04T13:38:25Z",
        "updated_at": "2017-08-04T13:38:29Z"
    },
    {
        "name": "nodejs_buildpack",
        "filename": "nodejs_buildpack-cached-v1.6.3.zip",
        "guid": "2ab6b2a6-595c-40a4-bb3c-a57926c333f9",
        "created_at": "2017-08-04T13:38:28Z",
        "updated_at": "2017-08-04T13:38:31Z"
    },
    {
        "name": "go_buildpack",
        "filename": "go_buildpack-cached-v1.8.5.zip",
        "guid": "fadfa871-f4bb-4ea1-8fdd-e7a6415ea5bc",
        "created_at": "2017-08-04T13:38:29Z",
        "updated_at": "2017-08-04T13:38:33Z"
    },
    {
        "name": "python_buildpack",
        "filename": "python_buildpack-cached-v1.5.20.zip",
        "guid": "5c8b2f39-0ccb-488f-b56f-733d36bb67ef",
        "created_at": "2017-08-04T13:38:31Z",
        "updated_at": "2017-08-04T13:38:37Z"
    },
    {
        "name": "php_buildpack",
        "filename": "php_buildpack-cached-v4.3.38.zip",
        "guid": "1d8189ea-4953-4cd7-9ca8-b03694b4d984",
        "created_at": "2017-08-04T13:38:33Z",
        "updated_at": "2017-08-04T13:38:38Z"
    },
    {
        "name": "dotnet_core_buildpack",
        "filename": "dotnet-core_buildpack-cached-v1.0.22.zip",
        "guid": "483a1093-673e-4dc8-9e3a-d938430f659a",
        "created_at": "2017-08-04T13:38:37Z",
        "updated_at": "2017-08-04T13:39:02Z"
    },
    {
        "name": "binary_buildpack",
        "filename": "binary_buildpack-cached-v1.0.13.zip",
        "guid": "9d8f48fe-01ad-4592-8324-19585cc62e92",
        "created_at": "2017-08-04T13:38:38Z",
        "updated_at": "2017-08-04T13:38:38Z"
    }
]
```

### [GET] `/models/info`

Get general information about the platform.

#### Sample Response
```JSON
{
    "iaas-type": "vsphere",
    "version": "1.12.0.0"
}
```

### [GET] `/models/products`

Get the collection of installed products (tiles).

#### Sample Response
```JSON
[
    {
        "installation_name": "p-bosh",
        "guid": "p-bosh-43bdb35efc3c504823e4",
        "type": "p-bosh"
    },
    {
        "installation_name": "cf-712c1d330ebea47e9e1e",
        "guid": "cf-712c1d330ebea47e9e1e",
        "type": "cf"
    },
    {
        "installation_name": "pivotal-mysql-a0606c50a3894ed40af3",
        "guid": "pivotal-mysql-a0606c50a3894ed40af3",
        "type": "pivotal-mysql"
    }
]
```

### [GET] `/models/products/{guid}`

Get the details of a product.

#### Sample Response
```JSON
{
    "status": [
        {
            "az_guid": "az1",
            "az_name": null,
            "cid": "vm-de71b4de-3ae8-434b-828e-5677318fa208",
            "cpu": {
                "sys": "1.5",
                "user": "0.7",
                "wait": "0.0"
            },
            "ephemeral_disk": {
                "inode_percent": "1",
                "percent": "6"
            },
            "index": 0,
            "ips": [
                "172.28.61.72"
            ],
            "job-name": "diego_cell",
            "load_avg": [
                "0.02",
                "0.10",
                "0.09"
            ],
            "memory": {
                "kb": "2328936",
                "percent": "14"
            },
            "persistent_disk": null,
            "swap": {
                "kb": "13020",
                "percent": "0"
            },
            "system_disk": {
                "inode_percent": "30",
                "percent": "36"
            }
        },
				...
    ]
}
```

### [GET] `/models/releases`

Get the release names.

```JSON
[
    "loggregator-91.0.0-3421.9.0.tgz",
    "nodejs-offline-buildpack-1.6.3-3421.9.0.tgz",
    "php-offline-buildpack-4.3.38-3421.9.0.tgz",
    "nfs-volume-1.0.6-3421.9.0.tgz",
    "syslog-migration-7.0.0-3421.9.0.tgz",
    "release-service-metrics-1.5.7-on-ubuntu-trusty-stemcell-3363.26.tgz",
    "haproxy-8.3.0-3421.9.0.tgz",
    "release-syslog-migration-5.0.0-on-ubuntu-trusty-stemcell-3363.25.tgz",
    "cf-backup-and-restore-0.0.9-3421.9.0.tgz",
    "nats-22.0.0-3421.9.0.tgz",
    "notifications-36.0.0-3421.9.0.tgz",
    "python-offline-buildpack-1.5.20-3421.9.0.tgz",
    "java-offline-buildpack-3.18.0-3421.9.0.tgz",
    "release-dedicated-mysql-0.17.3-on-ubuntu-trusty-stemcell-3363.26.tgz",
    "release-loggregator-87.0.0-on-ubuntu-trusty-stemcell-3363.26.tgz",
    "release-consul-165.0.0-on-ubuntu-trusty-stemcell-3363.26.tgz",
    "uaa-41.0.0-3421.9.0.tgz",
    "cf-autoscaling-94.0.0-3421.9.0.tgz",
    "capi-1.35.0-3421.9.0.tgz",
    "scalablesyslog-8.0.0-3421.9.0.tgz",
    "diego-1.22.0-3421.9.0.tgz",
    "pivotal-account-1.6.0-3421.9.0.tgz",
    "service-backup-18.1.2-3421.9.0.tgz",
    "cf-mysql-36.0.0-3421.9.0.tgz",
    "garden-runc-1.9.0-3421.9.0.tgz",
    "push-apps-manager-release-662.0.3-3421.9.0.tgz",
    "mysql-monitoring-8.5.0-3421.9.0.tgz",
    "routing-0.159.0-3421.9.0.tgz",
    "release-dedicated-mysql-adapter-0.25.4-on-ubuntu-trusty-stemcell-3363.26.tgz",
    "staticfile-offline-buildpack-1.4.11-3421.9.0.tgz",
    "statsd-injector-1.0.28-3421.9.0.tgz",
    "ruby-offline-buildpack-1.6.44-3421.9.0.tgz",
    "consul-171.0.0-3421.9.0.tgz",
    "release-on-demand-service-broker-0.16.1-on-ubuntu-trusty-stemcell-3363.26.tgz",
    "cflinuxfs2-1.138.0-3421.9.0.tgz",
    "release-mysql-backup-1.34.0-on-ubuntu-trusty-stemcell-3363.26.tgz",
    "go-offline-buildpack-1.8.5-3421.9.0.tgz",
    "release-mysql-monitoring-8.2.0-on-ubuntu-trusty-stemcell-3363.25.tgz",
    "dotnet-core-offline-buildpack-1.0.22-3421.9.0.tgz",
    "release-service-backup-18.1.2-on-ubuntu-trusty-stemcell-3363.26.tgz",
    "notifications-ui-28.0.0-3421.9.0.tgz",
    "cf-smoke-tests-36.0.0-3421.9.0.tgz",
    "mysql-backup-1.35.0-3421.9.0.tgz",
    "cf-networking-1.2.0-3421.9.0.tgz",
    "binary-offline-buildpack-1.0.13-3421.9.0.tgz"
]
```

### [GET] `/models/stemcells`

Get information about the stemcells.

```JSON
[
    {
        "source": "cf-712c1d330ebea47e9e1e",
        "alias": "bosh-vsphere-esxi-ubuntu-trusty-go_agent",
        "os": "ubuntu-trusty",
        "version": "3421.9"
    },
    {
        "source": "pivotal-mysql-a0606c50a3894ed40af3",
        "alias": "bosh-vsphere-esxi-ubuntu-trusty-go_agent",
        "os": "ubuntu-trusty",
        "version": "3363.26"
    }
]
```

### [GET] `/models/vm_types`

Get information about the types of virtual machines available.

```JSON
[
    {
        "name": "nano",
        "ram": 512,
        "cpu": 1,
        "ephemeral_disk": 8192,
        "builtin": true
    },
    {
        "name": "micro",
        "ram": 1024,
        "cpu": 1,
        "ephemeral_disk": 8192,
        "builtin": true
    },
		....
    {
        "name": "2xlarge.cpu",
        "ram": 16384,
        "cpu": 16,
        "ephemeral_disk": 65536,
        "builtin": true
    }
]
```

### [GET] `/models/vms`

Get information about the allocated virtual machines.

```JSON
[
    {
        "product": "cf-712c1d330ebea47e9e1e",
        "instances": {
            "bosh-vsphere-esxi-ubuntu-trusty-go_agent": {
                "large.disk": 1,
                "medium.disk": 4,
                "medium.mem": 1,
                "micro": 32,
                "nano": 1,
                "small": 7,
                "xlarge.disk": 3
            }
        }
    },
    {
        "product": "pivotal-mysql-a0606c50a3894ed40af3",
        "instances": {
            "bosh-vsphere-esxi-ubuntu-trusty-go_agent": {
                "micro": 7
            }
        }
    }
]
```
