#!/bin/sh

echo "Setting up S3"
/scripts/s3-init.sh

echo "Setting up SecretsManager"
/scripts/secretsmanager-init.sh

echo "done!"
