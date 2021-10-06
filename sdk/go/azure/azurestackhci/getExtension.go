


package azurestackhci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:azurestackhci:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	ArcSettingName    string `pulumi:"arcSettingName"`
	ClusterName       string `pulumi:"clusterName"`
	ExtensionName     string `pulumi:"extensionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExtensionResult struct {
	AggregateState          string                          `pulumi:"aggregateState"`
	AutoUpgradeMinorVersion *bool                           `pulumi:"autoUpgradeMinorVersion"`
	CreatedAt               *string                         `pulumi:"createdAt"`
	CreatedBy               *string                         `pulumi:"createdBy"`
	CreatedByType           *string                         `pulumi:"createdByType"`
	ForceUpdateTag          *string                         `pulumi:"forceUpdateTag"`
	Id                      string                          `pulumi:"id"`
	LastModifiedAt          *string                         `pulumi:"lastModifiedAt"`
	LastModifiedBy          *string                         `pulumi:"lastModifiedBy"`
	LastModifiedByType      *string                         `pulumi:"lastModifiedByType"`
	Name                    string                          `pulumi:"name"`
	PerNodeExtensionDetails []PerNodeExtensionStateResponse `pulumi:"perNodeExtensionDetails"`
	ProtectedSettings       interface{}                     `pulumi:"protectedSettings"`
	ProvisioningState       string                          `pulumi:"provisioningState"`
	Publisher               *string                         `pulumi:"publisher"`
	Settings                interface{}                     `pulumi:"settings"`
	Type                    string                          `pulumi:"type"`
	TypeHandlerVersion      *string                         `pulumi:"typeHandlerVersion"`
}
