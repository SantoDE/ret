## Overview

If you're using Rancher in combination with e.g. [AWS Auto Scaling Groups](https://aws.amazon.com/en/autoscaling/) you may have already encountered the issue, that the Rancher Registration Tokens expire after a fairly short amount of time.

Therfore seamless autoscaling with new Agent-Hosts will become tricky as they can't join the environment.

RET will help you solve this issue.

By only providing:

- the Rancher API URL (by either *--api-url* or as Environment Variable *RANCHER_API_URL*)
- the Access Key (by either *--access-key* or as Environment Variable *RANCHER_ACCESS_KEY*)
- the Secret Key (by either *--secret-key* or as Environment Variable *RANCHER_SECRET_KEY*)
- the Rancher Environment Id (by either *--environment-id* or as Environment Variable *RANCHER_ENVIRONMENT_ID*)

it will give you the current token, so your agents can join easily

## Example

```
echo $tmp `./ret --api-url=http://rancherurl --access-key=XXX --secret-key=XXX --environment-id=1a7`
F30581DB7EC2EE52AF3F:1485460800000:203kD7zc4gwwrfRMWvmsB30fhI
```

## Getting Started

You can either download the latest release [here](https://github.com/SantoDE/ret/releases/tag/v0.1-alpha) or just the latest docker image [here](https://hub.docker.com/r/santode/ret/) and set the required variables mentioned above. All good from there!