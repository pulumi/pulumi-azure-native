package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPolicyAssignment(ctx, "policyAssignment", &authorization.PolicyAssignmentArgs{
			DisplayName: pulumi.String("Enforce security policies"),
			NonComplianceMessages: authorization.NonComplianceMessageArray{
				&authorization.NonComplianceMessageArgs{
					Message: pulumi.String("Resources must comply with all internal security policies. See <internal site URL> for more info."),
				},
				&authorization.NonComplianceMessageArgs{
					Message:                     pulumi.String("Resource names must start with 'DeptA' and end with '-LC'."),
					PolicyDefinitionReferenceId: pulumi.String("10420126870854049575"),
				},
				&authorization.NonComplianceMessageArgs{
					Message:                     pulumi.String("Storage accounts must have firewall rules configured."),
					PolicyDefinitionReferenceId: pulumi.String("8572513655450389710"),
				},
			},
			PolicyAssignmentName: pulumi.String("securityInitAssignment"),
			PolicyDefinitionId:   pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/securityInitiative"),
			Scope:                pulumi.String("subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
