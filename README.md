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
| UAA_ADDRESS          | This is the URL to the UAA and Ops Manager instance. Example: `https://192.168.0.0` |
| OPSMAN_USER          | This is the Ops Man user that you want to use. |
| OPSMAN_PASSWORD      | This is the password to the Ops Man user. |

## Raw Data

You can query the application using the following route: `http://<Application Route>/versions`

It will then return:
```JSON
{
  "versions": {
    "ERT": "1.7.20-build.2",
    "Ops Man": "1.7"
  }
}
```

## Badges
Badges have been removed and another project controls them now: https://github.com/ECSTeam/status-badge
