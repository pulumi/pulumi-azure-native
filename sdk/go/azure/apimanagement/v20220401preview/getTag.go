


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTag(ctx *pulumi.Context, args *LookupTagArgs, opts ...pulumi.InvokeOption) (*LookupTagResult, error) {
	var rv LookupTagResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getTag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TagId             string `pulumi:"tagId"`
}


type LookupTagResult struct {
	DisplayName string `pulumi:"displayName"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	Type        string `pulumi:"type"`
}

func LookupTagOutput(ctx *pulumi.Context, args LookupTagOutputArgs, opts ...pulumi.InvokeOption) LookupTagResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagResult, error) {
			args := v.(LookupTagArgs)
			r, err := LookupTag(ctx, &args, opts...)
			var s LookupTagResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagResultOutput)
}

type LookupTagOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	TagId             pulumi.StringInput `pulumi:"tagId"`
}

func (LookupTagOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagArgs)(nil)).Elem()
}


type LookupTagResultOutput struct{ *pulumi.OutputState }

func (LookupTagResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagResult)(nil)).Elem()
}

func (o LookupTagResultOutput) ToLookupTagResultOutput() LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) ToLookupTagResultOutputWithContext(ctx context.Context) LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagResultOutput{})
}
