


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:security:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	ConnectorName string `pulumi:"connectorName"`
}


type LookupConnectorResult struct {
	AuthenticationDetails interface{}                              `pulumi:"authenticationDetails"`
	HybridComputeSettings *HybridComputeSettingsPropertiesResponse `pulumi:"hybridComputeSettings"`
	Id                    string                                   `pulumi:"id"`
	Name                  string                                   `pulumi:"name"`
	Type                  string                                   `pulumi:"type"`
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
	ConnectorName pulumi.StringInput `pulumi:"connectorName"`
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

func (o LookupConnectorResultOutput) AuthenticationDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupConnectorResult) interface{} { return v.AuthenticationDetails }).(pulumi.AnyOutput)
}

func (o LookupConnectorResultOutput) HybridComputeSettings() HybridComputeSettingsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *HybridComputeSettingsPropertiesResponse { return v.HybridComputeSettings }).(HybridComputeSettingsPropertiesResponsePtrOutput)
}

func (o LookupConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorResultOutput{})
}
