


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSiteLinkedBackendForBuild(ctx *pulumi.Context, args *LookupStaticSiteLinkedBackendForBuildArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteLinkedBackendForBuildResult, error) {
	var rv LookupStaticSiteLinkedBackendForBuildResult
	err := ctx.Invoke("azure-native:web/v20220301:getStaticSiteLinkedBackendForBuild", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteLinkedBackendForBuildArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	LinkedBackendName string `pulumi:"linkedBackendName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteLinkedBackendForBuildResult struct {
	BackendResourceId *string `pulumi:"backendResourceId"`
	CreatedOn         string  `pulumi:"createdOn"`
	Id                string  `pulumi:"id"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Region            *string `pulumi:"region"`
	Type              string  `pulumi:"type"`
}

func LookupStaticSiteLinkedBackendForBuildOutput(ctx *pulumi.Context, args LookupStaticSiteLinkedBackendForBuildOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteLinkedBackendForBuildResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteLinkedBackendForBuildResult, error) {
			args := v.(LookupStaticSiteLinkedBackendForBuildArgs)
			r, err := LookupStaticSiteLinkedBackendForBuild(ctx, &args, opts...)
			var s LookupStaticSiteLinkedBackendForBuildResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteLinkedBackendForBuildResultOutput)
}

type LookupStaticSiteLinkedBackendForBuildOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	LinkedBackendName pulumi.StringInput `pulumi:"linkedBackendName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteLinkedBackendForBuildOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteLinkedBackendForBuildArgs)(nil)).Elem()
}


type LookupStaticSiteLinkedBackendForBuildResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteLinkedBackendForBuildResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteLinkedBackendForBuildResult)(nil)).Elem()
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) ToLookupStaticSiteLinkedBackendForBuildResultOutput() LookupStaticSiteLinkedBackendForBuildResultOutput {
	return o
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) ToLookupStaticSiteLinkedBackendForBuildResultOutputWithContext(ctx context.Context) LookupStaticSiteLinkedBackendForBuildResultOutput {
	return o
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) BackendResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendForBuildResult) *string { return v.BackendResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendForBuildResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendForBuildResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendForBuildResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendForBuildResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendForBuildResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendForBuildResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteLinkedBackendForBuildResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteLinkedBackendForBuildResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteLinkedBackendForBuildResultOutput{})
}
