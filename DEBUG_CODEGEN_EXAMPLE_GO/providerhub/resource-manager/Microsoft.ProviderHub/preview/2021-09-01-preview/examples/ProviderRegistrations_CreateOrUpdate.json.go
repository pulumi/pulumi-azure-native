package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/providerhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := providerhub.NewProviderRegistration(ctx, "providerRegistration", &providerhub.ProviderRegistrationArgs{
Properties: providerhub.ProviderRegistrationResponseProperties{
Capabilities: providerhub.ResourceProviderCapabilitiesArray{
&providerhub.ResourceProviderCapabilitiesArgs{
Effect: pulumi.String("Allow"),
QuotaId: pulumi.String("CSP_2015-05-01"),
},
&providerhub.ResourceProviderCapabilitiesArgs{
Effect: pulumi.String("Allow"),
QuotaId: pulumi.String("CSP_MG_2017-12-01"),
},
},
Management: interface{}{
IncidentContactEmail: pulumi.String("helpme@contoso.com"),
IncidentRoutingService: pulumi.String("Contoso Resource Provider"),
IncidentRoutingTeam: pulumi.String("Contoso Triage"),
ServiceTreeInfos: providerhub.ServiceTreeInfoArray{
&providerhub.ServiceTreeInfoArgs{
ComponentId: pulumi.String("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
Readiness: pulumi.String("InDevelopment"),
ServiceId: pulumi.String("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
},
},
},
ProviderType: pulumi.String("Internal"),
ProviderVersion: pulumi.String("2.0"),
},
ProviderNamespace: pulumi.String("Microsoft.Contoso"),
})
if err != nil {
return err
}
return nil
})
}
