


package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:costmanagement/v20180801preview:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	ConnectorName     string `pulumi:"connectorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectorResult struct {
	Collection        ConnectorCollectionInfoResponse `pulumi:"collection"`
	CreatedOn         string                          `pulumi:"createdOn"`
	CredentialsKey    *string                         `pulumi:"credentialsKey"`
	DisplayName       *string                         `pulumi:"displayName"`
	Id                string                          `pulumi:"id"`
	Kind              *string                         `pulumi:"kind"`
	Location          *string                         `pulumi:"location"`
	ModifiedOn        string                          `pulumi:"modifiedOn"`
	Name              string                          `pulumi:"name"`
	ProviderAccountId string                          `pulumi:"providerAccountId"`
	ReportId          *string                         `pulumi:"reportId"`
	Status            *string                         `pulumi:"status"`
	Tags              map[string]string               `pulumi:"tags"`
	Type              string                          `pulumi:"type"`
}

func LookupConnectorOutput(ctx *pulumi.Context, args LookupConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorResult, error) {
			args := v.(LookupConnectorArgs)
			r, err := LookupConnector(ctx, &args, opts...)
			var s LookupConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorResultOutput)
}

type LookupConnectorOutputArgs struct {
	ConnectorName     pulumi.StringInput `pulumi:"connectorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorArgs)(nil)).Elem()
}


type LookupConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorResult)(nil)).Elem()
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutput() LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutputWithContext(ctx context.Context) LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) Collection() ConnectorCollectionInfoResponseOutput {
	return o.ApplyT(func(v LookupConnectorResult) ConnectorCollectionInfoResponse { return v.Collection }).(ConnectorCollectionInfoResponseOutput)
}

func (o LookupConnectorResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) CredentialsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.CredentialsKey }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) ProviderAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.ProviderAccountId }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) ReportId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.ReportId }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorResultOutput{})
}
