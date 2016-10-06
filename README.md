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

## Configuration

Apply the following configuration to the application instance.

| Environment Variable | Details |
|:--------------------:|:--------|
| `UAA_ADDRESS`        | This is the URL to the UAA and Ops Manager instance. Example: `https://192.168.0.0` |
| `OPSMAN_USER`        | This is the Ops Man user that you want to use. |
| `OPSMAN_PASSWORD`    | This is the password to the Ops Man user. |

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
