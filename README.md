# dotscience-slack-plugin

[![Build Status](https://drone.app.cloud.dotscience.net/api/badges/dotmesh-io/dotscience-slack-plugin/status.svg)](https://drone.app.cloud.dotscience.net/dotmesh-io/dotscience-slack-plugin)


## Usage

Create a new incoming webhook config via Slack app: https://api.slack.com/messaging/webhooks

Example `.dotscience.yml` configuration:

```yaml
kind: pipeline

after:
- name: circleci
  image: quay.io/dotmesh/dotscience-slack-plugin
  settings:
    slackUrl: https://hooks.slack.com/services/.../...
```
