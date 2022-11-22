


package v20221120preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAPICollection(ctx *pulumi.Context, args *LookupAPICollectionArgs, opts ...pulumi.InvokeOption) (*LookupAPICollectionResult, error) {
	var rv LookupAPICollectionResult
	err := ctx.Invoke("azure-native:security/v20221120preview:getAPICollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAPICollectionArgs struct {
	ApiCollectionId   string `pulumi:"apiCollectionId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupAPICollectionResult struct {
	AdditionalData map[string]string `pulumi:"additionalData"`
	DisplayName    *string           `pulumi:"displayName"`
	Id             string            `pulumi:"id"`
	Name           string            `pulumi:"name"`
	Type           string            `pulumi:"type"`
}

func LookupAPICollectionOutput(ctx *pulumi.Context, args LookupAPICollectionOutputArgs, opts ...pulumi.InvokeOption) LookupAPICollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAPICollectionResult, error) {
			args := v.(LookupAPICollectionArgs)
			r, err := LookupAPICollection(ctx, &args, opts...)
			var s LookupAPICollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAPICollectionResultOutput)
}

type LookupAPICollectionOutputArgs struct {
	ApiCollectionId   pulumi.StringInput `pulumi:"apiCollectionId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAPICollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAPICollectionArgs)(nil)).Elem()
}


type LookupAPICollectionResultOutput struct{ *pulumi.OutputState }

func (LookupAPICollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAPICollectionResult)(nil)).Elem()
}

func (o LookupAPICollectionResultOutput) ToLookupAPICollectionResultOutput() LookupAPICollectionResultOutput {
	return o
}

func (o LookupAPICollectionResultOutput) ToLookupAPICollectionResultOutputWithContext(ctx context.Context) LookupAPICollectionResultOutput {
	return o
}

func (o LookupAPICollectionResultOutput) AdditionalData() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) map[string]string { return v.AdditionalData }).(pulumi.StringMapOutput)
}

func (o LookupAPICollectionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupAPICollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAPICollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAPICollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAPICollectionResultOutput{})
}
