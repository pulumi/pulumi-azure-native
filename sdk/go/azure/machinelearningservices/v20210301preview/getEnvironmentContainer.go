


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironmentContainer(ctx *pulumi.Context, args *LookupEnvironmentContainerArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentContainerResult, error) {
	var rv LookupEnvironmentContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getEnvironmentContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentContainerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEnvironmentContainerResult struct {
	Id         string                       `pulumi:"id"`
	Name       string                       `pulumi:"name"`
	Properties EnvironmentContainerResponse `pulumi:"properties"`
	SystemData SystemDataResponse           `pulumi:"systemData"`
	Type       string                       `pulumi:"type"`
}

func LookupEnvironmentContainerOutput(ctx *pulumi.Context, args LookupEnvironmentContainerOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentContainerResult, error) {
			args := v.(LookupEnvironmentContainerArgs)
			r, err := LookupEnvironmentContainer(ctx, &args, opts...)
			var s LookupEnvironmentContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentContainerResultOutput)
}

type LookupEnvironmentContainerOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEnvironmentContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentContainerArgs)(nil)).Elem()
}


type LookupEnvironmentContainerResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentContainerResult)(nil)).Elem()
}

func (o LookupEnvironmentContainerResultOutput) ToLookupEnvironmentContainerResultOutput() LookupEnvironmentContainerResultOutput {
	return o
}

func (o LookupEnvironmentContainerResultOutput) ToLookupEnvironmentContainerResultOutputWithContext(ctx context.Context) LookupEnvironmentContainerResultOutput {
	return o
}

func (o LookupEnvironmentContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentContainerResultOutput) Properties() EnvironmentContainerResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) EnvironmentContainerResponse { return v.Properties }).(EnvironmentContainerResponseOutput)
}

func (o LookupEnvironmentContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEnvironmentContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentContainerResultOutput{})
}
