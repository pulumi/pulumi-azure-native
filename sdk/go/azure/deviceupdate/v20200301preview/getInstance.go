


package v20200301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("azure-native:deviceupdate/v20200301preview:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	AccountName       string `pulumi:"accountName"`
	InstanceName      string `pulumi:"instanceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInstanceResult struct {
	AccountName                 string                               `pulumi:"accountName"`
	DiagnosticStorageProperties *DiagnosticStoragePropertiesResponse `pulumi:"diagnosticStorageProperties"`
	EnableDiagnostics           *bool                                `pulumi:"enableDiagnostics"`
	Id                          string                               `pulumi:"id"`
	IotHubs                     []IotHubSettingsResponse             `pulumi:"iotHubs"`
	Location                    string                               `pulumi:"location"`
	Name                        string                               `pulumi:"name"`
	ProvisioningState           string                               `pulumi:"provisioningState"`
	SystemData                  SystemDataResponse                   `pulumi:"systemData"`
	Tags                        map[string]string                    `pulumi:"tags"`
	Type                        string                               `pulumi:"type"`
}
