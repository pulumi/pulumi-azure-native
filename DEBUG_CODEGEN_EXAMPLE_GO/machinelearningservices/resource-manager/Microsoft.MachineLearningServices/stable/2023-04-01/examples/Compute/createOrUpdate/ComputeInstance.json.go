package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewCompute(ctx, "compute", &machinelearningservices.ComputeArgs{
			ComputeName: pulumi.String("compute123"),
			Location:    pulumi.String("eastus"),
			Properties: machinelearningservices.ComputeInstance{
				ComputeType: "ComputeInstance",
				Properties: machinelearningservices.ComputeInstanceProperties{
					ApplicationSharingPolicy:         "Personal",
					ComputeInstanceAuthorizationType: "personal",
					CustomServices: []machinelearningservices.CustomService{
						{
							Docker: {
								Privileged: true,
							},
							Endpoints: []machinelearningservices.Endpoint{
								{
									Name:      "connect",
									Protocol:  "http",
									Published: 8787,
									Target:    8787,
								},
							},
							EnvironmentVariables: {
								Test_variable: {
									Type:  "local",
									Value: "test_value",
								},
							},
							Image: {
								Reference: "ghcr.io/azure/rocker-rstudio-ml-verse:latest",
								Type:      "docker",
							},
							Name: "rstudio",
							Volumes: []machinelearningservices.VolumeDefinition{
								{
									ReadOnly: false,
									Source:   "/home/azureuser/cloudfiles",
									Target:   "/home/azureuser/cloudfiles",
									Type:     "bind",
								},
							},
						},
					},
					PersonalComputeInstanceSettings: machinelearningservices.PersonalComputeInstanceSettings{
						AssignedUser: machinelearningservices.AssignedUser{
							ObjectId: "00000000-0000-0000-0000-000000000000",
							TenantId: "00000000-0000-0000-0000-000000000000",
						},
					},
					SshSettings: machinelearningservices.ComputeInstanceSshSettings{
						SshPublicAccess: "Disabled",
					},
					Subnet: machinelearningservices.ResourceId{
						Id: "test-subnet-resource-id",
					},
					VmSize: "STANDARD_NC6",
				},
			},
			ResourceGroupName: pulumi.String("testrg123"),
			WorkspaceName:     pulumi.String("workspaces123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
