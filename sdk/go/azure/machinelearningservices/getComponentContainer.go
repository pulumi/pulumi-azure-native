


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComponentContainer(ctx *pulumi.Context, args *LookupComponentContainerArgs, opts ...pulumi.InvokeOption) (*LookupComponentContainerResult, error) {
	var rv LookupComponentContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices:getComponentContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupComponentContainerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupComponentContainerResult struct {
	ComponentContainerDetails ComponentContainerResponse `pulumi:"componentContainerDetails"`
	Id                        string                     `pulumi:"id"`
	Name                      string                     `pulumi:"name"`
	SystemData                SystemDataResponse         `pulumi:"systemData"`
	Type                      string                     `pulumi:"type"`
}


func (val *LookupComponentContainerResult) Defaults() *LookupComponentContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ComponentContainerDetails = *tmp.ComponentContainerDetails.Defaults()

	return &tmp
}

func LookupComponentContainerOutput(ctx *pulumi.Context, args LookupComponentContainerOutputArgs, opts ...pulumi.InvokeOption) LookupComponentContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComponentContainerResult, error) {
			args := v.(LookupComponentContainerArgs)
			r, err := LookupComponentContainer(ctx, &args, opts...)
			var s LookupComponentContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComponentContainerResultOutput)
}

type LookupComponentContainerOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupComponentContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentContainerArgs)(nil)).Elem()
}


type LookupComponentContainerResultOutput struct{ *pulumi.OutputState }

func (LookupComponentContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentContainerResult)(nil)).Elem()
}

func (o LookupComponentContainerResultOutput) ToLookupComponentContainerResultOutput() LookupComponentContainerResultOutput {
	return o
}

func (o LookupComponentContainerResultOutput) ToLookupComponentContainerResultOutputWithContext(ctx context.Context) LookupComponentContainerResultOutput {
	return o
}

func (o LookupComponentContainerResultOutput) ComponentContainerDetails() ComponentContainerResponseOutput {
	return o.ApplyT(func(v LookupComponentContainerResult) ComponentContainerResponse { return v.ComponentContainerDetails }).(ComponentContainerResponseOutput)
}

func (o LookupComponentContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComponentContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComponentContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupComponentContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupComponentContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentContainerResultOutput{})
}
