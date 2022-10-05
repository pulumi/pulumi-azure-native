


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudService(ctx *pulumi.Context, args *LookupCloudServiceArgs, opts ...pulumi.InvokeOption) (*LookupCloudServiceResult, error) {
	var rv LookupCloudServiceResult
	err := ctx.Invoke("azure-native:compute/v20210301:getCloudService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudServiceArgs struct {
	CloudServiceName  string `pulumi:"cloudServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCloudServiceResult struct {
	Id         string                         `pulumi:"id"`
	Location   string                         `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Properties CloudServicePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}

func LookupCloudServiceOutput(ctx *pulumi.Context, args LookupCloudServiceOutputArgs, opts ...pulumi.InvokeOption) LookupCloudServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudServiceResult, error) {
			args := v.(LookupCloudServiceArgs)
			r, err := LookupCloudService(ctx, &args, opts...)
			var s LookupCloudServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudServiceResultOutput)
}

type LookupCloudServiceOutputArgs struct {
	CloudServiceName  pulumi.StringInput `pulumi:"cloudServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCloudServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudServiceArgs)(nil)).Elem()
}


type LookupCloudServiceResultOutput struct{ *pulumi.OutputState }

func (LookupCloudServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudServiceResult)(nil)).Elem()
}

func (o LookupCloudServiceResultOutput) ToLookupCloudServiceResultOutput() LookupCloudServiceResultOutput {
	return o
}

func (o LookupCloudServiceResultOutput) ToLookupCloudServiceResultOutputWithContext(ctx context.Context) LookupCloudServiceResultOutput {
	return o
}

func (o LookupCloudServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCloudServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudServiceResultOutput) Properties() CloudServicePropertiesResponseOutput {
	return o.ApplyT(func(v LookupCloudServiceResult) CloudServicePropertiesResponse { return v.Properties }).(CloudServicePropertiesResponseOutput)
}

func (o LookupCloudServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCloudServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCloudServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudServiceResultOutput{})
}
