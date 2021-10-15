


package v20210630preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReferenceDataSet(ctx *pulumi.Context, args *LookupReferenceDataSetArgs, opts ...pulumi.InvokeOption) (*LookupReferenceDataSetResult, error) {
	var rv LookupReferenceDataSetResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20210630preview:getReferenceDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReferenceDataSetArgs struct {
	EnvironmentName      string `pulumi:"environmentName"`
	ReferenceDataSetName string `pulumi:"referenceDataSetName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupReferenceDataSetResult struct {
	CreationTime                 string                                `pulumi:"creationTime"`
	DataStringComparisonBehavior *string                               `pulumi:"dataStringComparisonBehavior"`
	Id                           string                                `pulumi:"id"`
	KeyProperties                []ReferenceDataSetKeyPropertyResponse `pulumi:"keyProperties"`
	Location                     string                                `pulumi:"location"`
	Name                         string                                `pulumi:"name"`
	ProvisioningState            string                                `pulumi:"provisioningState"`
	Tags                         map[string]string                     `pulumi:"tags"`
	Type                         string                                `pulumi:"type"`
}
