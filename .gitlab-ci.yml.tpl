image: $DOCKER_IMAGE
services:
- docker:dind

variables:
  DOCKER_HOST: tcp://localhost:2375
  APP_NAME: versions
  CHART_NAME: versions
  
stages:
- build_latest
- deploy_dev

build_latest:
  stage: build_latest
  only:
  - dev
  script:
  - docker login -u ${REGISTRY_USER} -p ${REGISTRY_PASSWORD} ${REGISTRY_URL}
  - docker build --network host --build-arg DOCKER_IMAGE_NODEJS=${DOCKER_IMAGE_NODEJS} --build-arg APP_ENV=devserver -t ${REGISTRY_URL}/${CI_PROJECT_NAME} .
  # - docker tag ${REGISTRY_URL}/${CI_PROJECT_NAME} ${REGISTRY_URL}/${CI_PROJECT_NAME}:${CI_COMMIT_SHA}
  # - docker push ${REGISTRY_URL}/${CI_PROJECT_NAME}:${CI_COMMIT_SHA}
  # - docker pull ${REGISTRY_URL}/${CI_PROJECT_NAME}:${CI_COMMIT_SHA}
  - docker tag ${REGISTRY_URL}/${CI_PROJECT_NAME} ${REGISTRY_URL}/${CI_PROJECT_NAME}:latest
  - docker push ${REGISTRY_URL}/${CI_PROJECT_NAME}:latest

deploy_dev:
  stage: deploy_dev
  only:
  - dev
  image: ${DOCKER_IMAGE_HELM}
  environment:
    name: development
  script:
  - rm -rf .git
  - apk update && apk add --no-cache git ca-certificates
  - git clone https://${CI_REGISTRY_USER}:${CI_JOB_TOKEN}@${GITLAB_DOMAIN}/${HELM_VALUES_REPOSITORY}
  - helm init --client-only
  - helm upgrade -i ${APP_NAME} ${VERSIONS_HELM_REPO_NAME}/${CHART_NAME} --namespace ${VERSIONS_NAMESPACE} -f ${VERSIONS_HELM_VALUES_PATH}/${APP_NAME}/devserver.yaml --wait --force