


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationFabric(ctx *pulumi.Context, args *LookupReplicationFabricArgs, opts ...pulumi.InvokeOption) (*LookupReplicationFabricResult, error) {
	var rv LookupReplicationFabricResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210301:getReplicationFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationFabricArgs struct {
	FabricName        string `pulumi:"fabricName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupReplicationFabricResult struct {
	Id         string                   `pulumi:"id"`
	Location   *string                  `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Properties FabricPropertiesResponse `pulumi:"properties"`
	Type       string                   `pulumi:"type"`
}

func LookupReplicationFabricOutput(ctx *pulumi.Context, args LookupReplicationFabricOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationFabricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationFabricResult, error) {
			args := v.(LookupReplicationFabricArgs)
			r, err := LookupReplicationFabric(ctx, &args, opts...)
			var s LookupReplicationFabricResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationFabricResultOutput)
}

type LookupReplicationFabricOutputArgs struct {
	FabricName        pulumi.StringInput `pulumi:"fabricName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationFabricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationFabricArgs)(nil)).Elem()
}


type LookupReplicationFabricResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationFabricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationFabricResult)(nil)).Elem()
}

func (o LookupReplicationFabricResultOutput) ToLookupReplicationFabricResultOutput() LookupReplicationFabricResultOutput {
	return o
}

func (o LookupReplicationFabricResultOutput) ToLookupReplicationFabricResultOutputWithContext(ctx context.Context) LookupReplicationFabricResultOutput {
	return o
}

func (o LookupReplicationFabricResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationFabricResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationFabricResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationFabricResultOutput) Properties() FabricPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) FabricPropertiesResponse { return v.Properties }).(FabricPropertiesResponseOutput)
}

func (o LookupReplicationFabricResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationFabricResultOutput{})
}
