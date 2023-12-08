# Introducing the Drone Docker Image Migration Plugin

The Drone Docker Image Migration Plugin streamlines the transfer of Docker images between registries, enhancing your Continuous Integration (CI) and Continuous Deployment (CD) workflows. This plugin is designed to pull a Docker image from a source registry, tag it, and then push it to a destination registry. The entire process is validated and configured using environment variables.

## Build the Docker Image

Using the plugin is straightforward. You can run the script directly using the following command:

```sh
PLUGIN_SOURCE_DOCKER_REGISTRY=SOURCE_DOCKER_REGISTRY \
PLUGIN_DESTINATION_DOCKER_REGISTRY=DESTINATION_DOCKER_REGISTRY \
PLUGIN_SOURCE_USERNAME=SOURCE_USERNAME \
PLUGIN_SOURCE_PAT=SOURCE_PAT \
PLUGIN_DESTINATION_USERNAME=DESTINATION_USERNAME \
PLUGIN_DESTINATION_PAT=DESTINATION_PAT \
PLUGIN_IMAGE_NAME=IMAGE_NAME \
PLUGIN_IMAGE_TAG=IMAGE_TAG \
sh docker_image_migration_plugin.sh
```

Additionally, you can build the Docker image with these commands:

    docker buildx build -t DOCKER_ORG/drone-docker-image-migration --platform linux/amd64 .

### Usage in Harness CI

Integrating the drone-docker-image-migration plugin into your Harness CI pipeline is seamless. You can use Docker to run the plugin with environment variables. Here's how:

    docker run --rm \
    -e PLUGIN_SOURCE_DOCKER_REGISTRY=${SOURCE_DOCKER_REGISTRY} \
    -e PLUGIN_DESTINATION_DOCKER_REGISTRY=${DESTINATION_DOCKER_REGISTRY} \
    -e PLUGIN_SOURCE_USERNAME=${SOURCE_USERNAME} \
    -e PLUGIN_SOURCE_PAT=${SOURCE_PAT}
    -e PLUGIN_DESTINATION_USERNAME=${DOCKER_PAT}
    -e PLUGIN_DESTINATION_PAT=${DESTINATION_PAT}
    -e PLUGIN_IMAGE_NAME=${IMAGE_NAME}
    -e PLUGIN_IMAGE_TAG=${IMAGE_TAG}
    -v $(pwd):$(pwd) \
    -w $(pwd) \
    harnesscommunity/drone-docker-image-migration

In your Harness CI pipeline, you can define the plugin as a step, like this:

    - step:
        type:  Plugin
        name:  drone-docker-image-migration
        identifier:  drone-docker-image-migration
        spec:
            connectorRef:  docker-registry-connector
            image:  harnesscommunity/drone-docker-image-migration
            settings:
                SOURCE_DOCKER_REGISTRY: registry.hub.docker.com
                SOURCE_USERNAME: <+variable.docker_username>
                SOURCE_PAT: <+secrets.getValue("source_pat")>
                IMAGE_NAME: image_name
                DESTINATION_DOCKER_REGISTRY: registry.hub.docker.com
                DESTINATION_USERNAME: <+variable.docker_username>
                DESTINATION_PAT: <+secrets.getValue("dest_pat")>
                IMAGE_TAG: latest

### Plugin Options

The plugin offers the following customization options:

- **SOURCE_DOCKER_REGISTRY**: The source docker registry

- **SOURCE_USERNAME**: Source docker registry username

- **SOURCE_PAT**: Source docker registry PAT

- **IMAGE_NAME**: The image to be pulled and pushed

- **DESTINATION_DOCKER_REGISTRY**: The destination docker registry

- **DESTINATION_USERNAME**: Destination docker registry username

- **DESTINATION_PAT**: Destination docker registry PAT

- **IMAGE_TAG**: The image tag to be pulled

These environment variables are crucial for configuring and customizing the behavior of the plugin when executed as a Docker container. They allow you to provide specific values and project information required for pulling and tagging your Docker image.
