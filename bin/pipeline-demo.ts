#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { PipelineDemoStack } from '../lib/pipeline-demo-stack';

const app = new cdk.App();
new PipelineDemoStack(app, 'PipelineDemoStack', {
  env: {
	  account: '525038205906',
	  region: 'us-east-1'
  }

  /* For more information, see https://docs.aws.amazon.com/cdk/latest/guide/environments.html */
});
