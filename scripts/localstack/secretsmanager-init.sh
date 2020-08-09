#!/bin/sh
# add client secrets to secretsmanager
awslocal secretsmanager create-secret --name "authorization/secret-key" \
--description "Secret key for authorization token" \
--secret-string 'secret'