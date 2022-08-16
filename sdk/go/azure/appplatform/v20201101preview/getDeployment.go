


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeployment(ctx *pulumi.Context, args *LookupDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentResult, error) {
	var rv LookupDeploymentResult
	err := ctx.Invoke("azure-native:appplatform/v20201101preview:getDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDeploymentArgs struct {
	AppName           string `pulumi:"appName"`
	DeploymentName    string `pulumi:"deploymentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupDeploymentResult struct {
	Id         string                               `pulumi:"id"`
	Name       string                               `pulumi:"name"`
	Properties DeploymentResourcePropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                         `pulumi:"sku"`
	Type       string                               `pulumi:"type"`
}


func (val *LookupDeploymentResult) Defaults() *LookupDeploymentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupDeploymentOutput(ctx *pulumi.Context, args LookupDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentResult, error) {
			args := v.(LookupDeploymentArgs)
			r, err := LookupDeployment(ctx, &args, opts...)
			var s LookupDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentResultOutput)
}

type LookupDeploymentOutputArgs struct {
	AppName           pulumi.StringInput `pulumi:"appName"`
	DeploymentName    pulumi.StringInput `pulumi:"deploymentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentArgs)(nil)).Elem()
}


type LookupDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentResult)(nil)).Elem()
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutput() LookupDeploymentResultOutput {
	return o
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutputWithContext(ctx context.Context) LookupDeploymentResultOutput {
	return o
}

func (o LookupDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeploymentResultOutput) Properties() DeploymentResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupDeploymentResult) DeploymentResourcePropertiesResponse { return v.Properties }).(DeploymentResourcePropertiesResponseOutput)
}

func (o LookupDeploymentResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentResultOutput{})
}
