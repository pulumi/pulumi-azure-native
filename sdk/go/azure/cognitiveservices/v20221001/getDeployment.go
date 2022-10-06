


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeployment(ctx *pulumi.Context, args *LookupDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentResult, error) {
	var rv LookupDeploymentResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20221001:getDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentArgs struct {
	AccountName       string `pulumi:"accountName"`
	DeploymentName    string `pulumi:"deploymentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDeploymentResult struct {
	Etag       string                       `pulumi:"etag"`
	Id         string                       `pulumi:"id"`
	Name       string                       `pulumi:"name"`
	Properties DeploymentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse           `pulumi:"systemData"`
	Type       string                       `pulumi:"type"`
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
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DeploymentName    pulumi.StringInput `pulumi:"deploymentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupDeploymentResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeploymentResultOutput) Properties() DeploymentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDeploymentResult) DeploymentPropertiesResponse { return v.Properties }).(DeploymentPropertiesResponseOutput)
}

func (o LookupDeploymentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeploymentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentResultOutput{})
}
