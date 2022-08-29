


package v20200201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCreator(ctx *pulumi.Context, args *LookupCreatorArgs, opts ...pulumi.InvokeOption) (*LookupCreatorResult, error) {
	var rv LookupCreatorResult
	err := ctx.Invoke("azure-native:maps/v20200201preview:getCreator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCreatorArgs struct {
	AccountName       string `pulumi:"accountName"`
	CreatorName       string `pulumi:"creatorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCreatorResult struct {
	Id         string                    `pulumi:"id"`
	Location   string                    `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties CreatorPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string         `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}

func LookupCreatorOutput(ctx *pulumi.Context, args LookupCreatorOutputArgs, opts ...pulumi.InvokeOption) LookupCreatorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCreatorResult, error) {
			args := v.(LookupCreatorArgs)
			r, err := LookupCreator(ctx, &args, opts...)
			var s LookupCreatorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCreatorResultOutput)
}

type LookupCreatorOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	CreatorName       pulumi.StringInput `pulumi:"creatorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCreatorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCreatorArgs)(nil)).Elem()
}


type LookupCreatorResultOutput struct{ *pulumi.OutputState }

func (LookupCreatorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCreatorResult)(nil)).Elem()
}

func (o LookupCreatorResultOutput) ToLookupCreatorResultOutput() LookupCreatorResultOutput {
	return o
}

func (o LookupCreatorResultOutput) ToLookupCreatorResultOutputWithContext(ctx context.Context) LookupCreatorResultOutput {
	return o
}

func (o LookupCreatorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCreatorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCreatorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCreatorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCreatorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCreatorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCreatorResultOutput) Properties() CreatorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCreatorResult) CreatorPropertiesResponse { return v.Properties }).(CreatorPropertiesResponseOutput)
}

func (o LookupCreatorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCreatorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCreatorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCreatorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCreatorResultOutput{})
}
