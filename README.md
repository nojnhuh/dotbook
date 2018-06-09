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

#### Pre-requisites

1. PostgreSQL helm chart installed
	- `$ helm install stable/postgresql --name postgresql -f postgresql-values.yaml`
1. NGINX Ingress helm chart installed
	- `$ helm install stable/nginx-ingress --name nginx -f nginx-ingress-values.yaml`
1. `dotbook-tls` kubernetes secret
	- `$ openssl req -x509 -nodes -newkey rsa:2048 -keyout server.key -out server.crt -subj "/CN=dotbook.local"`
	- `$ kubectl create secret tls dotbook-tls --key server.key --cert server.crt`

#### Deploy
- `$ helm install charts/dotbook-api --name dotbook-api`

### Docker Compose

A sample [docker-compose.yml](docker-compose.yml) can also be used to set up a local Docker environment outside of kubernetes with `$ docker-compose up --build`.
