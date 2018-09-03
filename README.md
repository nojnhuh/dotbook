# dotbook

[![Build Status](https://travis-ci.org/nojnhuh/dotbook.svg?branch=development)](https://travis-ci.org/nojnhuh/dotbook)
[![Go Report Card](https://goreportcard.com/badge/github.com/nojnhuh/dotbook)](https://goreportcard.com/report/github.com/nojnhuh/dotbook)

A REST API to help marching band and drum corps members better navigate their drill.

## Features
- Midsets
- Yard line crossing counts
- Body-center to foot dot conversion

## Installation

### Kubernetes
`$ helm install ./charts/dotbook-api --name dotbook-api`

### Docker Compose

A sample [docker-compose.yml](docker-compose.yml) can also be used to set up a local Docker environment outside of kubernetes with `$ docker-compose up --build`.
