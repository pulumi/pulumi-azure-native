


package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudConnector(ctx *pulumi.Context, args *LookupCloudConnectorArgs, opts ...pulumi.InvokeOption) (*LookupCloudConnectorResult, error) {
	var rv LookupCloudConnectorResult
	err := ctx.Invoke("azure-native:costmanagement:getCloudConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudConnectorArgs struct {
	ConnectorName string  `pulumi:"connectorName"`
	Expand        *string `pulumi:"expand"`
}


type LookupCloudConnectorResult struct {
	BillingModel                      *string                         `pulumi:"billingModel"`
	CollectionInfo                    ConnectorCollectionInfoResponse `pulumi:"collectionInfo"`
	CreatedOn                         string                          `pulumi:"createdOn"`
	CredentialsKey                    *string                         `pulumi:"credentialsKey"`
	DaysTrialRemaining                int                             `pulumi:"daysTrialRemaining"`
	DefaultManagementGroupId          *string                         `pulumi:"defaultManagementGroupId"`
	DisplayName                       *string                         `pulumi:"displayName"`
	ExternalBillingAccountId          string                          `pulumi:"externalBillingAccountId"`
	Id                                string                          `pulumi:"id"`
	Kind                              *string                         `pulumi:"kind"`
	ModifiedOn                        string                          `pulumi:"modifiedOn"`
	Name                              string                          `pulumi:"name"`
	ProviderBillingAccountDisplayName string                          `pulumi:"providerBillingAccountDisplayName"`
	ProviderBillingAccountId          string                          `pulumi:"providerBillingAccountId"`
	ReportId                          *string                         `pulumi:"reportId"`
	Status                            string                          `pulumi:"status"`
	SubscriptionId                    *string                         `pulumi:"subscriptionId"`
	Type                              string                          `pulumi:"type"`
}

func LookupCloudConnectorOutput(ctx *pulumi.Context, args LookupCloudConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupCloudConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudConnectorResult, error) {
			args := v.(LookupCloudConnectorArgs)
			r, err := LookupCloudConnector(ctx, &args, opts...)
			var s LookupCloudConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudConnectorResultOutput)
}

type LookupCloudConnectorOutputArgs struct {
	ConnectorName pulumi.StringInput    `pulumi:"connectorName"`
	Expand        pulumi.StringPtrInput `pulumi:"expand"`
}

func (LookupCloudConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudConnectorArgs)(nil)).Elem()
}


type LookupCloudConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupCloudConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudConnectorResult)(nil)).Elem()
}

func (o LookupCloudConnectorResultOutput) ToLookupCloudConnectorResultOutput() LookupCloudConnectorResultOutput {
	return o
}

func (o LookupCloudConnectorResultOutput) ToLookupCloudConnectorResultOutputWithContext(ctx context.Context) LookupCloudConnectorResultOutput {
	return o
}

func (o LookupCloudConnectorResultOutput) BillingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.BillingModel }).(pulumi.StringPtrOutput)
}

func (o LookupCloudConnectorResultOutput) CollectionInfo() ConnectorCollectionInfoResponseOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) ConnectorCollectionInfoResponse { return v.CollectionInfo }).(ConnectorCollectionInfoResponseOutput)
}

func (o LookupCloudConnectorResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupCloudConnectorResultOutput) CredentialsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.CredentialsKey }).(pulumi.StringPtrOutput)
}

func (o LookupCloudConnectorResultOutput) DaysTrialRemaining() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) int { return v.DaysTrialRemaining }).(pulumi.IntOutput)
}

func (o LookupCloudConnectorResultOutput) DefaultManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.DefaultManagementGroupId }).(pulumi.StringPtrOutput)
}

func (o LookupCloudConnectorResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupCloudConnectorResultOutput) ExternalBillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.ExternalBillingAccountId }).(pulumi.StringOutput)
}

func (o LookupCloudConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudConnectorResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupCloudConnectorResultOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o LookupCloudConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudConnectorResultOutput) ProviderBillingAccountDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.ProviderBillingAccountDisplayName }).(pulumi.StringOutput)
}

func (o LookupCloudConnectorResultOutput) ProviderBillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.ProviderBillingAccountId }).(pulumi.StringOutput)
}

func (o LookupCloudConnectorResultOutput) ReportId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.ReportId }).(pulumi.StringPtrOutput)
}

func (o LookupCloudConnectorResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupCloudConnectorResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupCloudConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudConnectorResultOutput{})
}
