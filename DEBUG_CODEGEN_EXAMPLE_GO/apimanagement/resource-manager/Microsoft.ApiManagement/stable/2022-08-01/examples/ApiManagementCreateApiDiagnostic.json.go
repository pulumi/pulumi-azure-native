package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := apimanagement.NewApiDiagnostic(ctx, "apiDiagnostic", &apimanagement.ApiDiagnosticArgs{
AlwaysLog: pulumi.String("allErrors"),
ApiId: pulumi.String("57d1f7558aa04f15146d9d8a"),
Backend: apimanagement.PipelineDiagnosticSettingsResponse{
Request: interface{}{
Body: &apimanagement.BodyDiagnosticSettingsArgs{
Bytes: pulumi.Int(512),
},
Headers: pulumi.StringArray{
pulumi.String("Content-type"),
},
},
Response: interface{}{
Body: &apimanagement.BodyDiagnosticSettingsArgs{
Bytes: pulumi.Int(512),
},
Headers: pulumi.StringArray{
pulumi.String("Content-type"),
},
},
},
DiagnosticId: pulumi.String("applicationinsights"),
Frontend: apimanagement.PipelineDiagnosticSettingsResponse{
Request: interface{}{
Body: &apimanagement.BodyDiagnosticSettingsArgs{
Bytes: pulumi.Int(512),
},
Headers: pulumi.StringArray{
pulumi.String("Content-type"),
},
},
Response: interface{}{
Body: &apimanagement.BodyDiagnosticSettingsArgs{
Bytes: pulumi.Int(512),
},
Headers: pulumi.StringArray{
pulumi.String("Content-type"),
},
},
},
LoggerId: pulumi.String("/loggers/applicationinsights"),
ResourceGroupName: pulumi.String("rg1"),
Sampling: &apimanagement.SamplingSettingsArgs{
Percentage: pulumi.Float64(50),
SamplingType: pulumi.String("fixed"),
},
ServiceName: pulumi.String("apimService1"),
})
if err != nil {
return err
}
return nil
})
}
