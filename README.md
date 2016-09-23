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

You can query the application using the following route: `http://<Application Route>/`

It will then return:
```JSON
{
  "opsman-version": "1.7",
  "ert-version":    "1.7.21-build.2"
}
```

## Badges

| Type | Url | Example |
|:----:|:----|:--------|
| ERT  | `http://<Application Route>/badge.svg?type=ert` | ![Ert](https://cdn.rawgit.com/ECSTeam/pcf-status/master/img/ert.svg) |
| Ops Man | `http://<Application Route>/badge.svg` <br> `http://<Application Route>/badge.svg?type=opsman` | ![OpsMan](https://cdn.rawgit.com/ECSTeam/pcf-status/master/img/opsman.svg) |
