


package v20190401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupController(ctx *pulumi.Context, args *LookupControllerArgs, opts ...pulumi.InvokeOption) (*LookupControllerResult, error) {
	var rv LookupControllerResult
	err := ctx.Invoke("azure-native:devspaces/v20190401:getController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupControllerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupControllerResult struct {
	DataPlaneFqdn                    string            `pulumi:"dataPlaneFqdn"`
	HostSuffix                       string            `pulumi:"hostSuffix"`
	Id                               string            `pulumi:"id"`
	Location                         string            `pulumi:"location"`
	Name                             string            `pulumi:"name"`
	ProvisioningState                string            `pulumi:"provisioningState"`
	Sku                              SkuResponse       `pulumi:"sku"`
	Tags                             map[string]string `pulumi:"tags"`
	TargetContainerHostApiServerFqdn string            `pulumi:"targetContainerHostApiServerFqdn"`
	TargetContainerHostResourceId    string            `pulumi:"targetContainerHostResourceId"`
	Type                             string            `pulumi:"type"`
}

func LookupControllerOutput(ctx *pulumi.Context, args LookupControllerOutputArgs, opts ...pulumi.InvokeOption) LookupControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupControllerResult, error) {
			args := v.(LookupControllerArgs)
			r, err := LookupController(ctx, &args, opts...)
			var s LookupControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupControllerResultOutput)
}

type LookupControllerOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControllerArgs)(nil)).Elem()
}

type LookupControllerResultOutput struct{ *pulumi.OutputState }

func (LookupControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControllerResult)(nil)).Elem()
}

func (o LookupControllerResultOutput) ToLookupControllerResultOutput() LookupControllerResultOutput {
	return o
}

func (o LookupControllerResultOutput) ToLookupControllerResultOutputWithContext(ctx context.Context) LookupControllerResultOutput {
	return o
}

func (o LookupControllerResultOutput) DataPlaneFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.DataPlaneFqdn }).(pulumi.StringOutput)
}

func (o LookupControllerResultOutput) HostSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.HostSuffix }).(pulumi.StringOutput)
}

func (o LookupControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupControllerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupControllerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupControllerResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupControllerResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupControllerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupControllerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupControllerResultOutput) TargetContainerHostApiServerFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.TargetContainerHostApiServerFqdn }).(pulumi.StringOutput)
}

func (o LookupControllerResultOutput) TargetContainerHostResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.TargetContainerHostResourceId }).(pulumi.StringOutput)
}

func (o LookupControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupControllerResultOutput{})
}
