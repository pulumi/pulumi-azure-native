package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearningservices.NewCompute(ctx, "compute", &machinelearningservices.ComputeArgs{
ComputeName: pulumi.String("compute123"),
Location: pulumi.String("eastus"),
Properties: machinelearningservices.Kubernetes{
ComputeType: "Kubernetes",
Description: "some compute",
Properties: machinelearningservices.KubernetesProperties{
DefaultInstanceType: "defaultInstanceType",
InstanceTypes: interface{}{
DefaultInstanceType: machinelearningservices.InstanceTypeSchema{
Resources: machinelearningservices.InstanceTypeSchemaResources{
Limits: interface{}{
Cpu: "1",
Memory: "4Gi",
Nvidia.com/gpu: nil,
},
Requests: interface{}{
Cpu: "1",
Memory: "4Gi",
Nvidia.com/gpu: nil,
},
},
},
},
Namespace: "default",
},
ResourceId: "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2",
},
ResourceGroupName: pulumi.String("testrg123"),
WorkspaceName: pulumi.String("workspaces123"),
})
if err != nil {
return err
}
return nil
})
}
