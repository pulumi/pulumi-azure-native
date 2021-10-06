


package v20210101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupArcSetting(ctx *pulumi.Context, args *LookupArcSettingArgs, opts ...pulumi.InvokeOption) (*LookupArcSettingResult, error) {
	var rv LookupArcSettingResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210101preview:getArcSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArcSettingArgs struct {
	ArcSettingName    string `pulumi:"arcSettingName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupArcSettingResult struct {
	AggregateState           string                 `pulumi:"aggregateState"`
	ArcInstanceResourceGroup string                 `pulumi:"arcInstanceResourceGroup"`
	CreatedAt                *string                `pulumi:"createdAt"`
	CreatedBy                *string                `pulumi:"createdBy"`
	CreatedByType            *string                `pulumi:"createdByType"`
	Id                       string                 `pulumi:"id"`
	LastModifiedAt           *string                `pulumi:"lastModifiedAt"`
	LastModifiedBy           *string                `pulumi:"lastModifiedBy"`
	LastModifiedByType       *string                `pulumi:"lastModifiedByType"`
	Name                     string                 `pulumi:"name"`
	PerNodeDetails           []PerNodeStateResponse `pulumi:"perNodeDetails"`
	ProvisioningState        string                 `pulumi:"provisioningState"`
	Type                     string                 `pulumi:"type"`
}
