# dotbook-go

## Pre-requisites
1. `stable/mongodb-replicaset` helm chart installed (`-f mongodb-values.yaml`)
1. `stable/nginx-ingress` helm chart installed (`-f nginx-ingress-values.yaml`)
1. `dotbook-tls` kubernetes secret
	- `openssl req -x509 -nodes -newkey rsa:2048 -keyout server.key -out server.crt -subj "/CN=dotbook.local"`
	- `kubectl create secret tls dotbook-tls --key server.key --cert server.crt`
