package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearningservices.NewJob(ctx, "job", &machinelearningservices.JobArgs{
Id: pulumi.String("string"),
JobBaseProperties: machinelearningservices.SweepJob{
ComputeId: "string",
Description: "string",
DisplayName: "string",
EarlyTermination: machinelearningservices.MedianStoppingPolicy{
DelayEvaluation: 1,
EvaluationInterval: 1,
PolicyType: "MedianStopping",
},
ExperimentName: "string",
JobType: "Sweep",
Limits: machinelearningservices.SweepJobLimits{
JobLimitsType: "Sweep",
MaxConcurrentTrials: 1,
MaxTotalTrials: 1,
TrialTimeout: "PT1S",
},
Objective: machinelearningservices.Objective{
Goal: "Minimize",
PrimaryMetric: "string",
},
Properties: map[string]interface{}{
"string": "string",
},
SamplingAlgorithm: machinelearningservices.GridSamplingAlgorithm{
SamplingAlgorithmType: "Grid",
},
SearchSpace: map[string]interface{}{
"string": nil,
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
Trial: machinelearningservices.TrialComponent{
CodeId: "string",
Command: "string",
Distribution: machinelearningservices.Mpi{
DistributionType: "Mpi",
ProcessCountPerInstance: 1,
},
EnvironmentId: "string",
EnvironmentVariables: map[string]interface{}{
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
