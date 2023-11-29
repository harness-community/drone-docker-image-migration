# Drone plugin to pull a docker image from a registry and push it to another

# env vars:
# ! PLUGIN_SOURCE_DOCKER_REGISTRY
# ! PLUGIN_DESTINATION_DOCKER_REGISTRY
# ! PLUGIN_SOURCE_USERNAME
# ! PLUGIN_SOURCE_PAT
# ! PLUGIN_DESTINATION_USERNAME
# ! PLUGIN_DESTINATION_PAT
# ! PLUGIN_IMAGE_NAME
# ? PLUGIN_IMAGE_TAG - optional, defaults to latest

#############################################################################

# env variables validation

if [ -z "$PLUGIN_SOURCE_DOCKER_REGISTRY" ]; then
  echo "Error: PLUGIN_SOURCE_DOCKER_REGISTRY is not set"
  exit 1
fi

if [ -z "$PLUGIN_DESTINATION_DOCKER_REGISTRY" ]; then
  echo "Error: PLUGIN_DESTINATION_DOCKER_REGISTRY is not set"
  exit 1
fi

if [ -z "$PLUGIN_SOURCE_USERNAME" ]; then
  echo "Error: PLUGIN_SOURCE_USERNAME is not set"
  exit 1
fi

if [ -z "$PLUGIN_SOURCE_PAT" ]; then
  echo "Error: PLUGIN_SOURCE_PAT is not set"
  exit 1
fi

if [ -z "$PLUGIN_DESTINATION_USERNAME" ]; then
  echo "Error: PLUGIN_DESTINATION_USERNAME is not set"
  exit 1
fi

if [ -z "$PLUGIN_DESTINATION_PAT" ]; then
  echo "Error: PLUGIN_DESTINATION_PAT is not set"
  exit 1
fi

if [ -z "$PLUGIN_IMAGE_NAME" ]; then
  echo "Error: PLUGIN_IMAGE_NAME is not set"
  exit 1
fi

# optional env var

if [ -z "$PLUGIN_IMAGE_TAG" ]; then
  PLUGIN_IMAGE_TAG="latest"
fi

#############################################################################

# pull image from source registry
docker login $PLUGIN_SOURCE_DOCKER_REGISTRY -u $PLUGIN_SOURCE_USERNAME -p $PLUGIN_SOURCE_PAT || exit 1
docker pull $PLUGIN_SOURCE_DOCKER_REGISTRY/$PLUGIN_SOURCE_USERNAME/$PLUGIN_IMAGE_NAME:$PLUGIN_IMAGE_TAG || exit 1

echo "Image pulled successfully"

# tag image
docker tag $PLUGIN_SOURCE_DOCKER_REGISTRY/$PLUGIN_SOURCE_USERNAME/$PLUGIN_IMAGE_NAME:$PLUGIN_IMAGE_TAG $PLUGIN_DESTINATION_DOCKER_REGISTRY/$PLUGIN_DESTINATION_USERNAME/$PLUGIN_IMAGE_NAME:$PLUGIN_IMAGE_TAG || exit 1

echo "Image tagged successfully"

# push image to destination registry
docker login $PLUGIN_DESTINATION_DOCKER_REGISTRY -u $PLUGIN_DESTINATION_USERNAME -p $PLUGIN_DESTINATION_PAT || exit 1 
docker push $PLUGIN_DESTINATION_DOCKER_REGISTRY/$PLUGIN_DESTINATION_USERNAME/$PLUGIN_IMAGE_NAME:$PLUGIN_IMAGE_TAG || exit 1

echo "Image pushed successfully to $PLUGIN_DESTINATION_DOCKER_REGISTRY/$PLUGIN_DESTINATION_USERNAME/$PLUGIN_IMAGE_NAME:$PLUGIN_IMAGE_TAG"