package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearningservices.NewJob(ctx, "job", &machinelearningservices.JobArgs{
Id: pulumi.String("string"),
JobBaseProperties: machinelearningservices.PipelineJob{
ComputeId: "string",
Description: "string",
DisplayName: "string",
ExperimentName: "string",
Inputs: interface{}{
String: machinelearningservices.LiteralJobInput{
Description: "string",
JobInputType: "literal",
Value: "string",
},
},
JobType: "Pipeline",
Outputs: interface{}{
String: machinelearningservices.UriFileJobOutput{
Description: "string",
JobOutputType: "uri_file",
Mode: "Upload",
Uri: "string",
},
},
Properties: map[string]interface{}{
"string": "string",
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
Settings: nil,
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
