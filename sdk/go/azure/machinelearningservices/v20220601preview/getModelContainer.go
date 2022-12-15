


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupModelContainer(ctx *pulumi.Context, args *LookupModelContainerArgs, opts ...pulumi.InvokeOption) (*LookupModelContainerResult, error) {
	var rv LookupModelContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220601preview:getModelContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupModelContainerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupModelContainerResult struct {
	Id                       string                 `pulumi:"id"`
	ModelContainerProperties ModelContainerResponse `pulumi:"modelContainerProperties"`
	Name                     string                 `pulumi:"name"`
	SystemData               SystemDataResponse     `pulumi:"systemData"`
	Type                     string                 `pulumi:"type"`
}


func (val *LookupModelContainerResult) Defaults() *LookupModelContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ModelContainerProperties = *tmp.ModelContainerProperties.Defaults()

	return &tmp
}

func LookupModelContainerOutput(ctx *pulumi.Context, args LookupModelContainerOutputArgs, opts ...pulumi.InvokeOption) LookupModelContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelContainerResult, error) {
			args := v.(LookupModelContainerArgs)
			r, err := LookupModelContainer(ctx, &args, opts...)
			var s LookupModelContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelContainerResultOutput)
}

type LookupModelContainerOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupModelContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelContainerArgs)(nil)).Elem()
}


type LookupModelContainerResultOutput struct{ *pulumi.OutputState }

func (LookupModelContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelContainerResult)(nil)).Elem()
}

func (o LookupModelContainerResultOutput) ToLookupModelContainerResultOutput() LookupModelContainerResultOutput {
	return o
}

func (o LookupModelContainerResultOutput) ToLookupModelContainerResultOutputWithContext(ctx context.Context) LookupModelContainerResultOutput {
	return o
}

func (o LookupModelContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupModelContainerResultOutput) ModelContainerProperties() ModelContainerResponseOutput {
	return o.ApplyT(func(v LookupModelContainerResult) ModelContainerResponse { return v.ModelContainerProperties }).(ModelContainerResponseOutput)
}

func (o LookupModelContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupModelContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupModelContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupModelContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelContainerResultOutput{})
}
