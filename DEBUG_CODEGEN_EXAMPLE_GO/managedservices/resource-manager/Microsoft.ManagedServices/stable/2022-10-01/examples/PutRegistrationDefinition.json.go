package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managedservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := managedservices.NewRegistrationDefinition(ctx, "registrationDefinition", &managedservices.RegistrationDefinitionArgs{
Plan: &managedservices.PlanArgs{
Name: pulumi.String("addesai-plan"),
Product: pulumi.String("test"),
Publisher: pulumi.String("marketplace-test"),
Version: pulumi.String("1.0.0"),
},
Properties: managedservices.RegistrationDefinitionPropertiesResponse{
Authorizations: managedservices.AuthorizationArray{
&managedservices.AuthorizationArgs{
PrincipalId: pulumi.String("f98d86a2-4cc4-4e9d-ad47-b3e80a1bcdfc"),
PrincipalIdDisplayName: pulumi.String("Support User"),
RoleDefinitionId: pulumi.String("acdd72a7-3385-48ef-bd42-f606fba81ae7"),
},
&managedservices.AuthorizationArgs{
DelegatedRoleDefinitionIds: pulumi.StringArray{
pulumi.String("b24988ac-6180-42a0-ab88-20f7382dd24c"),
},
PrincipalId: pulumi.String("f98d86a2-4cc4-4e9d-ad47-b3e80a1bcdfc"),
PrincipalIdDisplayName: pulumi.String("User Access Administrator"),
RoleDefinitionId: pulumi.String("18d7d88d-d35e-4fb5-a5c3-7773c20a72d9"),
},
},
Description: pulumi.String("Tes1t"),
EligibleAuthorizations: managedservices.EligibleAuthorizationArray{
interface{}{
JustInTimeAccessPolicy: interface{}{
ManagedByTenantApprovers: managedservices.EligibleApproverArray{
&managedservices.EligibleApproverArgs{
PrincipalId: pulumi.String("d9b22cd6-6407-43cc-8c60-07c56df0b51a"),
PrincipalIdDisplayName: pulumi.String("Approver Group"),
},
},
MaximumActivationDuration: pulumi.String("PT8H"),
MultiFactorAuthProvider: pulumi.String("Azure"),
},
PrincipalId: pulumi.String("3e0ed8c6-e902-4fc5-863c-e3ddbb2ae2a2"),
PrincipalIdDisplayName: pulumi.String("Support User"),
RoleDefinitionId: pulumi.String("ae349356-3a1b-4a5e-921d-050484c6347e"),
},
},
ManagedByTenantId: pulumi.String("83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
RegistrationDefinitionName: pulumi.String("DefinitionName"),
},
RegistrationDefinitionId: pulumi.String("26c128c2-fefa-4340-9bb1-6e081c90ada2"),
Scope: pulumi.String("subscription/0afefe50-734e-4610-8a82-a144ahf49dea"),
})
if err != nil {
return err
}
return nil
})
}
