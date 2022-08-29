


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataContainer(ctx *pulumi.Context, args *LookupDataContainerArgs, opts ...pulumi.InvokeOption) (*LookupDataContainerResult, error) {
	var rv LookupDataContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220201preview:getDataContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDataContainerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDataContainerResult struct {
	DataContainerDetails DataContainerResponse `pulumi:"dataContainerDetails"`
	Id                   string                `pulumi:"id"`
	Name                 string                `pulumi:"name"`
	SystemData           SystemDataResponse    `pulumi:"systemData"`
	Type                 string                `pulumi:"type"`
}


func (val *LookupDataContainerResult) Defaults() *LookupDataContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DataContainerDetails = *tmp.DataContainerDetails.Defaults()

	return &tmp
}

func LookupDataContainerOutput(ctx *pulumi.Context, args LookupDataContainerOutputArgs, opts ...pulumi.InvokeOption) LookupDataContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataContainerResult, error) {
			args := v.(LookupDataContainerArgs)
			r, err := LookupDataContainer(ctx, &args, opts...)
			var s LookupDataContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataContainerResultOutput)
}

type LookupDataContainerOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDataContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataContainerArgs)(nil)).Elem()
}


type LookupDataContainerResultOutput struct{ *pulumi.OutputState }

func (LookupDataContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataContainerResult)(nil)).Elem()
}

func (o LookupDataContainerResultOutput) ToLookupDataContainerResultOutput() LookupDataContainerResultOutput {
	return o
}

func (o LookupDataContainerResultOutput) ToLookupDataContainerResultOutputWithContext(ctx context.Context) LookupDataContainerResultOutput {
	return o
}

func (o LookupDataContainerResultOutput) DataContainerDetails() DataContainerResponseOutput {
	return o.ApplyT(func(v LookupDataContainerResult) DataContainerResponse { return v.DataContainerDetails }).(DataContainerResponseOutput)
}

func (o LookupDataContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDataContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataContainerResultOutput{})
}
