


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCodeContainer(ctx *pulumi.Context, args *LookupCodeContainerArgs, opts ...pulumi.InvokeOption) (*LookupCodeContainerResult, error) {
	var rv LookupCodeContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getCodeContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodeContainerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupCodeContainerResult struct {
	Id         string                `pulumi:"id"`
	Name       string                `pulumi:"name"`
	Properties CodeContainerResponse `pulumi:"properties"`
	SystemData SystemDataResponse    `pulumi:"systemData"`
	Type       string                `pulumi:"type"`
}

func LookupCodeContainerOutput(ctx *pulumi.Context, args LookupCodeContainerOutputArgs, opts ...pulumi.InvokeOption) LookupCodeContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCodeContainerResult, error) {
			args := v.(LookupCodeContainerArgs)
			r, err := LookupCodeContainer(ctx, &args, opts...)
			var s LookupCodeContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCodeContainerResultOutput)
}

type LookupCodeContainerOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupCodeContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeContainerArgs)(nil)).Elem()
}


type LookupCodeContainerResultOutput struct{ *pulumi.OutputState }

func (LookupCodeContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeContainerResult)(nil)).Elem()
}

func (o LookupCodeContainerResultOutput) ToLookupCodeContainerResultOutput() LookupCodeContainerResultOutput {
	return o
}

func (o LookupCodeContainerResultOutput) ToLookupCodeContainerResultOutputWithContext(ctx context.Context) LookupCodeContainerResultOutput {
	return o
}

func (o LookupCodeContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCodeContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCodeContainerResultOutput) Properties() CodeContainerResponseOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) CodeContainerResponse { return v.Properties }).(CodeContainerResponseOutput)
}

func (o LookupCodeContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCodeContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCodeContainerResultOutput{})
}
