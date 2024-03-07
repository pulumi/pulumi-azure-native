package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearningservices.NewJob(ctx, "job", &machinelearningservices.JobArgs{
Id: pulumi.String("string"),
JobBaseProperties: machinelearningservices.AutoMLJob{
ComputeId: "string",
Description: "string",
DisplayName: "string",
EnvironmentId: "string",
EnvironmentVariables: map[string]interface{}{
"string": "string",
},
ExperimentName: "string",
Identity: machinelearningservices.AmlToken{
IdentityType: "AMLToken",
},
IsArchived: false,
JobType: "AutoML",
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
"9bec0ab0-c62f-4fa9-a97c-7b24bbcc90ad": nil,
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
TaskDetails: machinelearningservices.ImageClassification{
LimitSettings: machinelearningservices.ImageLimitSettings{
MaxTrials: 2,
},
ModelSettings: machinelearningservices.ImageModelSettingsClassification{
ValidationCropSize: 2,
},
SearchSpace: []machinelearningservices.ImageModelDistributionSettingsClassification{
{
ValidationCropSize: "choice(2, 360)",
},
},
TargetColumnName: "string",
TaskType: "ImageClassification",
TrainingData: machinelearningservices.MLTableJobInput{
JobInputType: "mltable",
Uri: "string",
},
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
