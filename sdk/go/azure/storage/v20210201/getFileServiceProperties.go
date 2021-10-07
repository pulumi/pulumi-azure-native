


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFileServiceProperties(ctx *pulumi.Context, args *LookupFileServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupFileServicePropertiesResult, error) {
	var rv LookupFileServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20210201:getFileServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileServicePropertiesArgs struct {
	AccountName       string `pulumi:"accountName"`
	FileServicesName  string `pulumi:"fileServicesName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFileServicePropertiesResult struct {
	Cors                       *CorsRulesResponse             `pulumi:"cors"`
	Id                         string                         `pulumi:"id"`
	Name                       string                         `pulumi:"name"`
	ProtocolSettings           *ProtocolSettingsResponse      `pulumi:"protocolSettings"`
	ShareDeleteRetentionPolicy *DeleteRetentionPolicyResponse `pulumi:"shareDeleteRetentionPolicy"`
	Sku                        SkuResponse                    `pulumi:"sku"`
	Type                       string                         `pulumi:"type"`
}
