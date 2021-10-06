


package servicefabricmesh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:servicefabricmesh:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VolumeResourceName string `pulumi:"volumeResourceName"`
}


type LookupVolumeResult struct {
	AzureFileParameters *VolumeProviderParametersAzureFileResponse `pulumi:"azureFileParameters"`
	Description         *string                                    `pulumi:"description"`
	Id                  string                                     `pulumi:"id"`
	Location            string                                     `pulumi:"location"`
	Name                string                                     `pulumi:"name"`
	Provider            string                                     `pulumi:"provider"`
	ProvisioningState   string                                     `pulumi:"provisioningState"`
	Status              string                                     `pulumi:"status"`
	StatusDetails       string                                     `pulumi:"statusDetails"`
	Tags                map[string]string                          `pulumi:"tags"`
	Type                string                                     `pulumi:"type"`
}
