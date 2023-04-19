import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { Function, Code, Runtime } from 'aws-cdk-lib/aws-lambda';

export class PipelineDemoStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

	   const fn = new Function(this, "AddHandler", {
		  runtime: Runtime.GO_1_X,
		  code: Code.fromAsset("lambda/bin"),
		  handler: "app.bin",
	  });
  }
}
