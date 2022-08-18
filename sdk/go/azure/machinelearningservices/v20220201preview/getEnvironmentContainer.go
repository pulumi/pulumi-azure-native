


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironmentContainer(ctx *pulumi.Context, args *LookupEnvironmentContainerArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentContainerResult, error) {
	var rv LookupEnvironmentContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220201preview:getEnvironmentContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEnvironmentContainerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEnvironmentContainerResult struct {
	EnvironmentContainerDetails EnvironmentContainerResponse `pulumi:"environmentContainerDetails"`
	Id                          string                       `pulumi:"id"`
	Name                        string                       `pulumi:"name"`
	SystemData                  SystemDataResponse           `pulumi:"systemData"`
	Type                        string                       `pulumi:"type"`
}


func (val *LookupEnvironmentContainerResult) Defaults() *LookupEnvironmentContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EnvironmentContainerDetails = *tmp.EnvironmentContainerDetails.Defaults()

	return &tmp
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

func (o LookupEnvironmentContainerResultOutput) EnvironmentContainerDetails() EnvironmentContainerResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) EnvironmentContainerResponse {
		return v.EnvironmentContainerDetails
	}).(EnvironmentContainerResponseOutput)
}

func (o LookupEnvironmentContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) string { return v.Name }).(pulumi.StringOutput)
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
