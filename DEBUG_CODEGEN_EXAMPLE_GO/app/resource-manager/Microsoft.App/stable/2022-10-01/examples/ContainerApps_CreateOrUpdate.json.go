package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := app.NewContainerApp(ctx, "containerApp", &app.ContainerAppArgs{
Configuration: app.ConfigurationResponse{
Dapr: &app.DaprArgs{
AppPort: pulumi.Int(3000),
AppProtocol: pulumi.String("http"),
EnableApiLogging: pulumi.Bool(true),
Enabled: pulumi.Bool(true),
HttpMaxRequestSize: pulumi.Int(10),
HttpReadBufferSize: pulumi.Int(30),
LogLevel: pulumi.String("debug"),
},
Ingress: interface{}{
ClientCertificateMode: pulumi.String("accept"),
CorsPolicy: &app.CorsPolicyArgs{
AllowCredentials: pulumi.Bool(true),
AllowedHeaders: pulumi.StringArray{
pulumi.String("HEADER1"),
pulumi.String("HEADER2"),
},
AllowedMethods: pulumi.StringArray{
pulumi.String("GET"),
pulumi.String("POST"),
},
AllowedOrigins: pulumi.StringArray{
pulumi.String("https://a.test.com"),
pulumi.String("https://b.test.com"),
},
ExposeHeaders: pulumi.StringArray{
pulumi.String("HEADER3"),
pulumi.String("HEADER4"),
},
MaxAge: pulumi.Int(1234),
},
CustomDomains: app.CustomDomainArray{
&app.CustomDomainArgs{
BindingType: pulumi.String("SniEnabled"),
CertificateId: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
Name: pulumi.String("www.my-name.com"),
},
&app.CustomDomainArgs{
BindingType: pulumi.String("SniEnabled"),
CertificateId: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"),
Name: pulumi.String("www.my-other-name.com"),
},
},
External: pulumi.Bool(true),
IpSecurityRestrictions: app.IpSecurityRestrictionRuleArray{
&app.IpSecurityRestrictionRuleArgs{
Action: pulumi.String("Allow"),
Description: pulumi.String("Allowing all IP's within the subnet below to access containerapp"),
IpAddressRange: pulumi.String("192.168.1.1/32"),
Name: pulumi.String("Allow work IP A subnet"),
},
&app.IpSecurityRestrictionRuleArgs{
Action: pulumi.String("Allow"),
Description: pulumi.String("Allowing all IP's within the subnet below to access containerapp"),
IpAddressRange: pulumi.String("192.168.1.1/8"),
Name: pulumi.String("Allow work IP B subnet"),
},
},
TargetPort: pulumi.Int(3000),
Traffic: app.TrafficWeightArray{
&app.TrafficWeightArgs{
Label: pulumi.String("production"),
RevisionName: pulumi.String("testcontainerapp0-ab1234"),
Weight: pulumi.Int(100),
},
},
},
MaxInactiveRevisions: pulumi.Int(10),
},
ContainerAppName: pulumi.String("testcontainerapp0"),
EnvironmentId: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
Location: pulumi.String("East US"),
ResourceGroupName: pulumi.String("rg"),
Template: app.TemplateResponse{
Containers: app.ContainerArray{
interface{}{
Image: pulumi.String("repo/testcontainerapp0:v1"),
Name: pulumi.String("testcontainerapp0"),
Probes: app.ContainerAppProbeArray{
interface{}{
HttpGet: interface{}{
HttpHeaders: app.ContainerAppProbeHttpHeadersArray{
&app.ContainerAppProbeHttpHeadersArgs{
Name: pulumi.String("Custom-Header"),
Value: pulumi.String("Awesome"),
},
},
Path: pulumi.String("/health"),
Port: pulumi.Int(8080),
},
InitialDelaySeconds: pulumi.Int(3),
PeriodSeconds: pulumi.Int(3),
Type: pulumi.String("Liveness"),
},
},
},
},
InitContainers: app.InitContainerArray{
interface{}{
Args: pulumi.StringArray{
pulumi.String("-c"),
pulumi.String("while true; do echo hello; sleep 10;done"),
},
Command: pulumi.StringArray{
pulumi.String("/bin/sh"),
},
Image: pulumi.String("repo/testcontainerapp0:v4"),
Name: pulumi.String("testinitcontainerApp0"),
Resources: &app.ContainerResourcesArgs{
Cpu: pulumi.Float64(0.5),
Memory: pulumi.String("1Gi"),
},
},
},
Scale: interface{}{
MaxReplicas: pulumi.Int(5),
MinReplicas: pulumi.Int(1),
Rules: app.ScaleRuleArray{
interface{}{
Custom: &app.CustomScaleRuleArgs{
Metadata: pulumi.StringMap{
"concurrentRequests": pulumi.String("50"),
},
Type: pulumi.String("http"),
},
Name: pulumi.String("httpscalingrule"),
},
},
},
},
WorkloadProfileType: pulumi.String("GeneralPurpose"),
})
if err != nil {
return err
}
return nil
})
}
