


package v20220904

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachinePool(ctx *pulumi.Context, args *LookupMachinePoolArgs, opts ...pulumi.InvokeOption) (*LookupMachinePoolResult, error) {
	var rv LookupMachinePoolResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20220904:getMachinePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachinePoolArgs struct {
	ChildResourceName string `pulumi:"childResourceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupMachinePoolResult struct {
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	Resources  *string            `pulumi:"resources"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupMachinePoolOutput(ctx *pulumi.Context, args LookupMachinePoolOutputArgs, opts ...pulumi.InvokeOption) LookupMachinePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachinePoolResult, error) {
			args := v.(LookupMachinePoolArgs)
			r, err := LookupMachinePool(ctx, &args, opts...)
			var s LookupMachinePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachinePoolResultOutput)
}

type LookupMachinePoolOutputArgs struct {
	ChildResourceName pulumi.StringInput `pulumi:"childResourceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupMachinePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachinePoolArgs)(nil)).Elem()
}


type LookupMachinePoolResultOutput struct{ *pulumi.OutputState }

func (LookupMachinePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachinePoolResult)(nil)).Elem()
}

func (o LookupMachinePoolResultOutput) ToLookupMachinePoolResultOutput() LookupMachinePoolResultOutput {
	return o
}

func (o LookupMachinePoolResultOutput) ToLookupMachinePoolResultOutputWithContext(ctx context.Context) LookupMachinePoolResultOutput {
	return o
}

func (o LookupMachinePoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachinePoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachinePoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachinePoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachinePoolResultOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachinePoolResult) *string { return v.Resources }).(pulumi.StringPtrOutput)
}

func (o LookupMachinePoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMachinePoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMachinePoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachinePoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachinePoolResultOutput{})
}
