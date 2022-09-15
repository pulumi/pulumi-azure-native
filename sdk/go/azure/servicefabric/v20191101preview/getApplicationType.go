


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplicationType(ctx *pulumi.Context, args *LookupApplicationTypeArgs, opts ...pulumi.InvokeOption) (*LookupApplicationTypeResult, error) {
	var rv LookupApplicationTypeResult
	err := ctx.Invoke("azure-native:servicefabric/v20191101preview:getApplicationType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationTypeArgs struct {
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	ClusterName         string `pulumi:"clusterName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupApplicationTypeResult struct {
	Etag              string            `pulumi:"etag"`
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}

func LookupApplicationTypeOutput(ctx *pulumi.Context, args LookupApplicationTypeOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationTypeResult, error) {
			args := v.(LookupApplicationTypeArgs)
			r, err := LookupApplicationType(ctx, &args, opts...)
			var s LookupApplicationTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationTypeResultOutput)
}

type LookupApplicationTypeOutputArgs struct {
	ApplicationTypeName pulumi.StringInput `pulumi:"applicationTypeName"`
	ClusterName         pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationTypeArgs)(nil)).Elem()
}


type LookupApplicationTypeResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationTypeResult)(nil)).Elem()
}

func (o LookupApplicationTypeResultOutput) ToLookupApplicationTypeResultOutput() LookupApplicationTypeResultOutput {
	return o
}

func (o LookupApplicationTypeResultOutput) ToLookupApplicationTypeResultOutputWithContext(ctx context.Context) LookupApplicationTypeResultOutput {
	return o
}

func (o LookupApplicationTypeResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApplicationTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationTypeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationTypeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplicationTypeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationTypeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationTypeResultOutput{})
}
