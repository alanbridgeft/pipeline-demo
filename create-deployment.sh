#!/bin/bash

filename=lambda-binary-${RANDOM}.zip
cp lambda/bin/app.bin .
echo "Filename is $filename"
zip $filename app.bin
rm app.bin

# upload the zipfile to s3
aws s3 cp $filename s3://com.bridgeft.codepipeline

# remove the zip locally
rm $filename

# create the deployment
aws deploy create-deployment --application-name=demo-pipeline-app \
	--deployment-group-name=code-pipeline-deployment-group \
	--s3-location bucket=com.bridgeft.codepipeline,bundleType=YAML,key=appspec.yml \
	--region=us-east-1 \
	--revision=revisionType=S3,s3Location={bucket=com.bridgeft.codepipeline,key=$filename,bundleType=zip}
