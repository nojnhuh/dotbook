#!/bin/bash
set -ev

if [ "$TRAVIS_BRANCH" == "master" ]; then
    docker login -u='$DOCKER_USERNAME' -p='$DOCKER_PASSWORD';
    docker-compose push;
fi
