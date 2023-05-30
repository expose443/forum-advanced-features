#!/bin/bash
openssl genrsa -out '../certificats/private.key' 2048

openssl req -new -x509 -sha256 -key '../certificats/private.key' -out '../certificats/certificate.crt' -days 365

