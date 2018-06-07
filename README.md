# dotbook

[![Build Status](https://travis-ci.org/nojnhuh/dotbook.svg?branch=development)](https://travis-ci.org/nojnhuh/dotbook)
[![Go Report Card](https://goreportcard.com/badge/github.com/nojnhuh/dotbook)](https://goreportcard.com/report/github.com/nojnhuh/dotbook)

A REST API to help marching band and drum corps members better navigate their drill.

## Features
- Midsets
- Yard line crossing counts
- Body-center to foot dot conversion

## Pre-requisites
1. `stable/mongodb-replicaset` helm chart installed
	- `$ helm install stable/mongodb-replicaset --name mongodb -f mongodb-values.yaml`
1. `stable/nginx-ingress` helm chart installed
	- `$ helm install stable/nginx-ingress --name nginx -f nginx-ingress-values.yaml`
1. `dotbook-tls` kubernetes secret
	- `$ openssl req -x509 -nodes -newkey rsa:2048 -keyout server.key -out server.crt -subj "/CN=dotbook.local"`
	- `$ kubectl create secret tls dotbook-tls --key server.key --cert server.crt`

## Installation
	- `$ helm install charts/dotbook-api --name dotbook-api`
