


package servicefabricmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	var rv LookupNetworkResult
	err := ctx.Invoke("azure-native:servicefabricmesh:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkArgs struct {
	NetworkResourceName string `pulumi:"networkResourceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupNetworkResult struct {
	Id         string                            `pulumi:"id"`
	Location   string                            `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties NetworkResourcePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}

func LookupNetworkOutput(ctx *pulumi.Context, args LookupNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkResult, error) {
			args := v.(LookupNetworkArgs)
			r, err := LookupNetwork(ctx, &args, opts...)
			var s LookupNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkResultOutput)
}

type LookupNetworkOutputArgs struct {
	NetworkResourceName pulumi.StringInput `pulumi:"networkResourceName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkArgs)(nil)).Elem()
}


type LookupNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkResult)(nil)).Elem()
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutput() LookupNetworkResultOutput {
	return o
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutputWithContext(ctx context.Context) LookupNetworkResultOutput {
	return o
}

func (o LookupNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Properties() NetworkResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupNetworkResult) NetworkResourcePropertiesResponse { return v.Properties }).(NetworkResourcePropertiesResponseOutput)
}

func (o LookupNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkResultOutput{})
}
