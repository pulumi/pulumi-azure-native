package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearningservices.NewJob(ctx, "job", &machinelearningservices.JobArgs{
Id: pulumi.String("string"),
JobBaseProperties: machinelearningservices.CommandJob{
CodeId: "string",
Command: "string",
ComputeId: "string",
Description: "string",
DisplayName: "string",
Distribution: machinelearningservices.TensorFlow{
DistributionType: "TensorFlow",
ParameterServerCount: 1,
WorkerCount: 1,
},
EnvironmentId: "string",
EnvironmentVariables: map[string]interface{}{
"string": "string",
},
ExperimentName: "string",
Identity: machinelearningservices.AmlToken{
IdentityType: "AMLToken",
},
Inputs: interface{}{
String: machinelearningservices.LiteralJobInput{
Description: "string",
JobInputType: "literal",
Value: "string",
},
},
JobType: "Command",
Limits: machinelearningservices.CommandJobLimits{
JobLimitsType: "Command",
Timeout: "PT5M",
},
Outputs: interface{}{
String: machinelearningservices.UriFileJobOutput{
Description: "string",
JobOutputType: "uri_file",
Mode: "ReadWriteMount",
Uri: "string",
},
},
Properties: map[string]interface{}{
"string": "string",
},
Resources: machinelearningservices.JobResourceConfiguration{
InstanceCount: 1,
InstanceType: "string",
Properties: map[string]interface{}{
"string": map[string]interface{}{
"e6b6493e-7d5e-4db3-be1e-306ec641327e": nil,
},
},
},
Services: interface{}{
String: machinelearningservices.JobService{
Endpoint: "string",
JobServiceType: "string",
Port: 1,
Properties: map[string]interface{}{
"string": "string",
},
},
},
Tags: map[string]interface{}{
"string": "string",
},
},
ResourceGroupName: pulumi.String("test-rg"),
WorkspaceName: pulumi.String("my-aml-workspace"),
})
if err != nil {
return err
}
return nil
})
}
