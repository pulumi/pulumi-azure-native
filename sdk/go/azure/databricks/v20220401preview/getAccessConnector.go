


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessConnector(ctx *pulumi.Context, args *LookupAccessConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAccessConnectorResult, error) {
	var rv LookupAccessConnectorResult
	err := ctx.Invoke("azure-native:databricks/v20220401preview:getAccessConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessConnectorArgs struct {
	ConnectorName     string `pulumi:"connectorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccessConnectorResult struct {
	Id         string                            `pulumi:"id"`
	Identity   *IdentityDataResponse             `pulumi:"identity"`
	Location   string                            `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties AccessConnectorPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}

func LookupAccessConnectorOutput(ctx *pulumi.Context, args LookupAccessConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAccessConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessConnectorResult, error) {
			args := v.(LookupAccessConnectorArgs)
			r, err := LookupAccessConnector(ctx, &args, opts...)
			var s LookupAccessConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessConnectorResultOutput)
}

type LookupAccessConnectorOutputArgs struct {
	ConnectorName     pulumi.StringInput `pulumi:"connectorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccessConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessConnectorArgs)(nil)).Elem()
}


type LookupAccessConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAccessConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessConnectorResult)(nil)).Elem()
}

func (o LookupAccessConnectorResultOutput) ToLookupAccessConnectorResultOutput() LookupAccessConnectorResultOutput {
	return o
}

func (o LookupAccessConnectorResultOutput) ToLookupAccessConnectorResultOutputWithContext(ctx context.Context) LookupAccessConnectorResultOutput {
	return o
}

func (o LookupAccessConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessConnectorResultOutput) Identity() IdentityDataResponsePtrOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) *IdentityDataResponse { return v.Identity }).(IdentityDataResponsePtrOutput)
}

func (o LookupAccessConnectorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAccessConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccessConnectorResultOutput) Properties() AccessConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) AccessConnectorPropertiesResponse { return v.Properties }).(AccessConnectorPropertiesResponseOutput)
}

func (o LookupAccessConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAccessConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessConnectorResultOutput{})
}
