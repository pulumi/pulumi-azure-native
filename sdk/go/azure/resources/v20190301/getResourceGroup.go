


package v20190301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupResourceGroup(ctx *pulumi.Context, args *LookupResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupResourceGroupResult, error) {
	var rv LookupResourceGroupResult
	err := ctx.Invoke("azure-native:resources/v20190301:getResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGroupArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupResourceGroupResult struct {
	Id         string                          `pulumi:"id"`
	Location   string                          `pulumi:"location"`
	ManagedBy  *string                         `pulumi:"managedBy"`
	Name       string                          `pulumi:"name"`
	Properties ResourceGroupPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string               `pulumi:"tags"`
	Type       string                          `pulumi:"type"`
}

func LookupResourceGroupOutput(ctx *pulumi.Context, args LookupResourceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupResourceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceGroupResult, error) {
			args := v.(LookupResourceGroupArgs)
			r, err := LookupResourceGroup(ctx, &args, opts...)
			var s LookupResourceGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceGroupResultOutput)
}

type LookupResourceGroupOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupResourceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGroupArgs)(nil)).Elem()
}


type LookupResourceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupResourceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGroupResult)(nil)).Elem()
}

func (o LookupResourceGroupResultOutput) ToLookupResourceGroupResultOutput() LookupResourceGroupResultOutput {
	return o
}

func (o LookupResourceGroupResultOutput) ToLookupResourceGroupResultOutputWithContext(ctx context.Context) LookupResourceGroupResultOutput {
	return o
}

func (o LookupResourceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourceGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupResourceGroupResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o LookupResourceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResourceGroupResultOutput) Properties() ResourceGroupPropertiesResponseOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) ResourceGroupPropertiesResponse { return v.Properties }).(ResourceGroupPropertiesResponseOutput)
}

func (o LookupResourceGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupResourceGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceGroupResultOutput{})
}
