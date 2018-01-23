#!/usr/bin/env bash

export TAG=1.${GO_PIPELINE_COUNTER}
export DOCKERFILE=DockerfileSpark
export ARTIFACT=/ailohq/airflow-pyspark

if [ -d "libs" ]; then
    docker build --no-cache -t ${DOCKER_REPOSITORY_HOST}${ARTIFACT}:${TAG} -f ${DOCKERFILE} .
    docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD} ${DOCKER_REPOSITORY_HOST}
    docker push ${DOCKER_REPOSITORY_HOST}${ARTIFACT}:${TAG}
else
    echo "/libs folder does not exist, aborting build"
    exit 1
fi
