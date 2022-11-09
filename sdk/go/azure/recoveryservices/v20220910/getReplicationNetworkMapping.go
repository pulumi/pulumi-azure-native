


package v20220910

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationNetworkMapping(ctx *pulumi.Context, args *LookupReplicationNetworkMappingArgs, opts ...pulumi.InvokeOption) (*LookupReplicationNetworkMappingResult, error) {
	var rv LookupReplicationNetworkMappingResult
	err := ctx.Invoke("azure-native:recoveryservices/v20220910:getReplicationNetworkMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationNetworkMappingArgs struct {
	FabricName         string `pulumi:"fabricName"`
	NetworkMappingName string `pulumi:"networkMappingName"`
	NetworkName        string `pulumi:"networkName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	ResourceName       string `pulumi:"resourceName"`
}


type LookupReplicationNetworkMappingResult struct {
	Id         string                           `pulumi:"id"`
	Location   *string                          `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties NetworkMappingPropertiesResponse `pulumi:"properties"`
	Type       string                           `pulumi:"type"`
}

func LookupReplicationNetworkMappingOutput(ctx *pulumi.Context, args LookupReplicationNetworkMappingOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationNetworkMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationNetworkMappingResult, error) {
			args := v.(LookupReplicationNetworkMappingArgs)
			r, err := LookupReplicationNetworkMapping(ctx, &args, opts...)
			var s LookupReplicationNetworkMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationNetworkMappingResultOutput)
}

type LookupReplicationNetworkMappingOutputArgs struct {
	FabricName         pulumi.StringInput `pulumi:"fabricName"`
	NetworkMappingName pulumi.StringInput `pulumi:"networkMappingName"`
	NetworkName        pulumi.StringInput `pulumi:"networkName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName       pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationNetworkMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationNetworkMappingArgs)(nil)).Elem()
}


type LookupReplicationNetworkMappingResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationNetworkMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationNetworkMappingResult)(nil)).Elem()
}

func (o LookupReplicationNetworkMappingResultOutput) ToLookupReplicationNetworkMappingResultOutput() LookupReplicationNetworkMappingResultOutput {
	return o
}

func (o LookupReplicationNetworkMappingResultOutput) ToLookupReplicationNetworkMappingResultOutputWithContext(ctx context.Context) LookupReplicationNetworkMappingResultOutput {
	return o
}

func (o LookupReplicationNetworkMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationNetworkMappingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationNetworkMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationNetworkMappingResultOutput) Properties() NetworkMappingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) NetworkMappingPropertiesResponse { return v.Properties }).(NetworkMappingPropertiesResponseOutput)
}

func (o LookupReplicationNetworkMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationNetworkMappingResultOutput{})
}
