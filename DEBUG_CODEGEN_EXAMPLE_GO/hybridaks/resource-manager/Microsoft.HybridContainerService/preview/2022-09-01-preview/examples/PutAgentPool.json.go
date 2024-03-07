package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcontainerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcontainerservice.NewAgentPool(ctx, "agentPool", &hybridcontainerservice.AgentPoolArgs{
			AgentPoolName:     pulumi.String("test-hybridaksnodepool"),
			Count:             pulumi.Int(1),
			Location:          pulumi.String("westus"),
			OsType:            pulumi.String("Linux"),
			ResourceGroupName: pulumi.String("test-arcappliance-resgrp"),
			ResourceName:      pulumi.String("test-hybridakscluster"),
			VmSize:            pulumi.String("Standard_A4_v2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
