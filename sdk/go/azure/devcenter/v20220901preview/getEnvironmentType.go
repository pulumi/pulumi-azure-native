


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironmentType(ctx *pulumi.Context, args *LookupEnvironmentTypeArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentTypeResult, error) {
	var rv LookupEnvironmentTypeResult
	err := ctx.Invoke("azure-native:devcenter/v20220901preview:getEnvironmentType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentTypeArgs struct {
	DevCenterName       string `pulumi:"devCenterName"`
	EnvironmentTypeName string `pulumi:"environmentTypeName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupEnvironmentTypeResult struct {
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}

func LookupEnvironmentTypeOutput(ctx *pulumi.Context, args LookupEnvironmentTypeOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentTypeResult, error) {
			args := v.(LookupEnvironmentTypeArgs)
			r, err := LookupEnvironmentType(ctx, &args, opts...)
			var s LookupEnvironmentTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentTypeResultOutput)
}

type LookupEnvironmentTypeOutputArgs struct {
	DevCenterName       pulumi.StringInput `pulumi:"devCenterName"`
	EnvironmentTypeName pulumi.StringInput `pulumi:"environmentTypeName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEnvironmentTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentTypeArgs)(nil)).Elem()
}


type LookupEnvironmentTypeResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentTypeResult)(nil)).Elem()
}

func (o LookupEnvironmentTypeResultOutput) ToLookupEnvironmentTypeResultOutput() LookupEnvironmentTypeResultOutput {
	return o
}

func (o LookupEnvironmentTypeResultOutput) ToLookupEnvironmentTypeResultOutputWithContext(ctx context.Context) LookupEnvironmentTypeResultOutput {
	return o
}

func (o LookupEnvironmentTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupEnvironmentTypeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEnvironmentTypeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEnvironmentTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentTypeResultOutput{})
}
