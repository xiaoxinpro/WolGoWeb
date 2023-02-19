#!/bin/bash
BUILD_NAME="buildx-wol-go-web"
BUILD_PLAT=linux/386,linux/amd64,linux/arm/v7,linux/arm64/v8
BUILD_PUSH=chishin/wol-go-web
BUILD_VERS=version-1.5.51

echo "=========== ${BUILD_NAME} ==========="

docker buildx create --name ${BUILD_NAME}
docker buildx use ${BUILD_NAME}
docker buildx build -t ${BUILD_PUSH}:${BUILD_VERS} -t ${BUILD_PUSH} --progress=plain --platform=${BUILD_PLAT} . --push
docker buildx rm ${BUILD_NAME}

echo "Multiarch build Complete ${BUILD_NAME} !!!"