# drone-image-migration

- [Synopsis](#Synopsis)
- [Parameters](#Paramaters)
- [Notes](#Notes)
- [Plugin Image](#Plugin-Image)
- [Examples](#Examples)

## Synopsis

This plugin is designed to migrate a Docker image from one registry to another.

To learn how to utilize Drone plugins in Harness CI, please consult the provided [documentation](https://developer.harness.io/docs/continuous-integration/use-ci/use-drone-plugins/run-a-drone-plugin-in-ci).

## Parameters

| Parameter                                                                                                                 | Choices/<span style="color:blue;">Defaults</span> | Comments                                                         |     |
| :------------------------------------------------------------------------------------------------------------------------ | :------------------------------------------------ | :--------------------------------------------------------------- | --- |
| source <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>      |                                                   | The source image to be copied                                    |     |
| username <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>    |                                                   | Username to login to the destination registry                    |     |
| password <span style="font-size: 10px"><br/>`string`</span>                                                               |                                                   | PAT / access token to authenticate with the destination registry |     |
| destination <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span> |                                                   | The destination where image will be copied                       |     |
| source_username <span style="font-size: 10px"><br/>`string`</span>                                                        |                                                   | Username to login to the source registry                         |     |
| source_password <span style="font-size: 10px"><br/>`string`</span>                                                        |                                                   | PAT / access token to authenticate with the source registry      |     |
| aws_access_key_id <span style="font-size: 10px"><br/>`string`</span>                                                      |                                                   | AWS access key ID for generating access token                    |     |
| aws_secret_access_key <span style="font-size: 10px"><br/>`string`</span>                                                  |                                                   | AWS secret access key for generating access token                |     |
| aws_region <span style="font-size: 10px"><br/>`string`</span>                                                             |                                                   | AWS region containing the ECR registry                           |     |
| overwrite <span style="font-size: 10px"><br/>`boolean`</span>                                                             | Default: `false`                                  | Overwrite the existing image at destination, if present          |     |
| insecure <span style="font-size: 10px"><br/>`boolean`</span>                                                              | Default: `false`                                  | Disable TLS                                                      |     |

## Notes

While using AWS ECR as destination registy, set `destination_username` as `AWS`, and either provide the AWS access token as `destination_password`, or provide the `aws_access_key_id`, `aws_secret_access_key` and `aws_region`.

While using Google Artifact Registry, use `oauth2accesstoken` as the relevant username and access-token as the password.

## Plugin Image

The plugin `plugins/image-migration` is available for the following architectures:

| OS          | Tag                          |
| ----------- | ---------------------------- |
| latest      | `linux-amd64/arm64, windows` |
| linux/amd64 | `linux-amd64`                |
| linux/arm64 | `linux-arm64`                |
| windows     | `windows-amd64`              |

## Examples

```
# Plugin YAML
- step:
    type: Plugin
    name: Migration Plugin
    identifier: Migration_Plugin
    spec:
        connectorRef: my-docker-connector
        image: plugins/image-migration
        settings:
                source: footloose/gitness:1.2.3
                destination: tremors/gitness:1.2.3
                username: kevinbacon
                password: <+secrets.getValue("docker_pat")>

- step:
    type: Plugin
    name: Migration Plugin
    identifier: Migration_Plugin
    spec:
        connectorRef: my-docker-connector
        image: plugins/image-migration
        settings:
                source: aws_account_id.dkr.ecr.us-west-2.amazonaws.com/gitness-dev:1.2.3
                destination: aws_account_id.dkr.ecr.us-west-2.amazonaws.com/gitness-prod:1.2.3
                username: AWS
                aws_access_key_id: "012345678901"
                aws_secret_access_key: <+secrets.getValue("aws_secret_access_key")>
                aws_region: us-west-2

- step:
    type: Plugin
    name: Migration Plugin
    identifier: Migration_Plugin
    spec:
        connectorRef: my-docker-connector
        image: plugins/image-migration
        settings:
                source: registry-1.example.com/gitness:1.2.3
                destination: registry-2.example.com/gitness:1.2.3
                source_username: finncarter
                source_password: <+secrets.getValue("source_docker_pat")>
                username: kevinbacon
                password: <+secrets.getValue("docker_pat")>
                overwrite: true

- step:
    type: Plugin
    name: Migration Plugin
    identifier: Migration_Plugin
    spec:
        connectorRef: my-docker-connector
        image: plugins/image-migration
        settings:
                source: registry-1.example.com/gitness:1.2.3
                destination: LOCATION-docker.pkg.dev/PROJECT-ID/REPO-NAME/IMAGE-NAME
                source_username: finncarter
                source_password: <+secrets.getValue("source_docker_pat")>
                username: oauth2accesstoken
                password: <+secrets.getValue("gcr_pat")>
                overwrite: true

```

> <span style="font-size: 14px; margin-left:5px; background-color: #d3d3d3; padding: 4px; border-radius: 4px;">ℹ️ If you notice any issues in this documentation, you can [edit this document](https://github.com/harness-community/drone-docker-image-migration/blob/main/README.md) to improve it.</span>
