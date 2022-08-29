


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkExperimentProfile(ctx *pulumi.Context, args *LookupNetworkExperimentProfileArgs, opts ...pulumi.InvokeOption) (*LookupNetworkExperimentProfileResult, error) {
	var rv LookupNetworkExperimentProfileResult
	err := ctx.Invoke("azure-native:network/v20191101:getNetworkExperimentProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkExperimentProfileArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNetworkExperimentProfileResult struct {
	EnabledState  *string           `pulumi:"enabledState"`
	Etag          *string           `pulumi:"etag"`
	Id            string            `pulumi:"id"`
	Location      *string           `pulumi:"location"`
	Name          string            `pulumi:"name"`
	ResourceState string            `pulumi:"resourceState"`
	Tags          map[string]string `pulumi:"tags"`
	Type          string            `pulumi:"type"`
}

func LookupNetworkExperimentProfileOutput(ctx *pulumi.Context, args LookupNetworkExperimentProfileOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkExperimentProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkExperimentProfileResult, error) {
			args := v.(LookupNetworkExperimentProfileArgs)
			r, err := LookupNetworkExperimentProfile(ctx, &args, opts...)
			var s LookupNetworkExperimentProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkExperimentProfileResultOutput)
}

type LookupNetworkExperimentProfileOutputArgs struct {
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkExperimentProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkExperimentProfileArgs)(nil)).Elem()
}


type LookupNetworkExperimentProfileResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkExperimentProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkExperimentProfileResult)(nil)).Elem()
}

func (o LookupNetworkExperimentProfileResultOutput) ToLookupNetworkExperimentProfileResultOutput() LookupNetworkExperimentProfileResultOutput {
	return o
}

func (o LookupNetworkExperimentProfileResultOutput) ToLookupNetworkExperimentProfileResultOutputWithContext(ctx context.Context) LookupNetworkExperimentProfileResultOutput {
	return o
}

func (o LookupNetworkExperimentProfileResultOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkExperimentProfileResult) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkExperimentProfileResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkExperimentProfileResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkExperimentProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkExperimentProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkExperimentProfileResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkExperimentProfileResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkExperimentProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkExperimentProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkExperimentProfileResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkExperimentProfileResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupNetworkExperimentProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkExperimentProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkExperimentProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkExperimentProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkExperimentProfileResultOutput{})
}
