


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDscpConfiguration(ctx *pulumi.Context, args *LookupDscpConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDscpConfigurationResult, error) {
	var rv LookupDscpConfigurationResult
	err := ctx.Invoke("azure-native:network/v20210501:getDscpConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDscpConfigurationArgs struct {
	DscpConfigurationName string `pulumi:"dscpConfigurationName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupDscpConfigurationResult struct {
	AssociatedNetworkInterfaces []NetworkInterfaceResponse `pulumi:"associatedNetworkInterfaces"`
	DestinationIpRanges         []QosIpRangeResponse       `pulumi:"destinationIpRanges"`
	DestinationPortRanges       []QosPortRangeResponse     `pulumi:"destinationPortRanges"`
	Etag                        string                     `pulumi:"etag"`
	Id                          *string                    `pulumi:"id"`
	Location                    *string                    `pulumi:"location"`
	Markings                    []int                      `pulumi:"markings"`
	Name                        string                     `pulumi:"name"`
	Protocol                    *string                    `pulumi:"protocol"`
	ProvisioningState           string                     `pulumi:"provisioningState"`
	QosCollectionId             string                     `pulumi:"qosCollectionId"`
	QosDefinitionCollection     []QosDefinitionResponse    `pulumi:"qosDefinitionCollection"`
	ResourceGuid                string                     `pulumi:"resourceGuid"`
	SourceIpRanges              []QosIpRangeResponse       `pulumi:"sourceIpRanges"`
	SourcePortRanges            []QosPortRangeResponse     `pulumi:"sourcePortRanges"`
	Tags                        map[string]string          `pulumi:"tags"`
	Type                        string                     `pulumi:"type"`
}
