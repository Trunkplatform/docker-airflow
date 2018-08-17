#!/usr/bin/env bash

export TAG=${GO_PIPELINE_LABEL}
export DOCKERFILE=DockerfileSpark
export ARTIFACT=/ailohq/airflow-pyspark

docker build --no-cache -t ${DOCKER_REPOSITORY_HOST}${ARTIFACT}:${TAG} -f ${DOCKERFILE} .
docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD} ${DOCKER_REPOSITORY_HOST}
docker push ${DOCKER_REPOSITORY_HOST}${ARTIFACT}:${TAG}
