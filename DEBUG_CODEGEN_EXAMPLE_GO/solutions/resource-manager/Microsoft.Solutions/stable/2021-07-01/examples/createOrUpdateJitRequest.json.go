package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/solutions/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := solutions.NewJitRequest(ctx, "jitRequest", &solutions.JitRequestArgs{
			ApplicationResourceId: pulumi.String("/subscriptions/00c76877-e316-48a7-af60-4a09fec9d43f/resourceGroups/52F30DB2/providers/Microsoft.Solutions/applications/7E193158"),
			JitAuthorizationPolicies: []solutions.JitAuthorizationPoliciesArgs{
				{
					PrincipalId:      pulumi.String("1db8e132e2934dbcb8e1178a61319491"),
					RoleDefinitionId: pulumi.String("ecd05a23-931a-4c38-a52b-ac7c4c583334"),
				},
			},
			JitRequestName: pulumi.String("myJitRequest"),
			JitSchedulingPolicy: &solutions.JitSchedulingPolicyArgs{
				Duration:  pulumi.String("PT8H"),
				StartTime: pulumi.String("2021-04-22T05:48:30.6661804Z"),
				Type:      pulumi.String("Once"),
			},
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
