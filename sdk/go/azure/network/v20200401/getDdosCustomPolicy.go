


package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDdosCustomPolicy(ctx *pulumi.Context, args *LookupDdosCustomPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDdosCustomPolicyResult, error) {
	var rv LookupDdosCustomPolicyResult
	err := ctx.Invoke("azure-native:network/v20200401:getDdosCustomPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDdosCustomPolicyArgs struct {
	DdosCustomPolicyName string `pulumi:"ddosCustomPolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupDdosCustomPolicyResult struct {
	Etag                   string                                 `pulumi:"etag"`
	Id                     *string                                `pulumi:"id"`
	Location               *string                                `pulumi:"location"`
	Name                   string                                 `pulumi:"name"`
	ProtocolCustomSettings []ProtocolCustomSettingsFormatResponse `pulumi:"protocolCustomSettings"`
	ProvisioningState      string                                 `pulumi:"provisioningState"`
	PublicIPAddresses      []SubResourceResponse                  `pulumi:"publicIPAddresses"`
	ResourceGuid           string                                 `pulumi:"resourceGuid"`
	Tags                   map[string]string                      `pulumi:"tags"`
	Type                   string                                 `pulumi:"type"`
}
