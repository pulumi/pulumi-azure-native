


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBatchDeployment(ctx *pulumi.Context, args *LookupBatchDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupBatchDeploymentResult, error) {
	var rv LookupBatchDeploymentResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220501:getBatchDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBatchDeploymentArgs struct {
	DeploymentName    string `pulumi:"deploymentName"`
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type LookupBatchDeploymentResult struct {
	BatchDeploymentProperties BatchDeploymentResponse         `pulumi:"batchDeploymentProperties"`
	Id                        string                          `pulumi:"id"`
	Identity                  *ManagedServiceIdentityResponse `pulumi:"identity"`
	Kind                      *string                         `pulumi:"kind"`
	Location                  string                          `pulumi:"location"`
	Name                      string                          `pulumi:"name"`
	Sku                       *SkuResponse                    `pulumi:"sku"`
	SystemData                SystemDataResponse              `pulumi:"systemData"`
	Tags                      map[string]string               `pulumi:"tags"`
	Type                      string                          `pulumi:"type"`
}


func (val *LookupBatchDeploymentResult) Defaults() *LookupBatchDeploymentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BatchDeploymentProperties = *tmp.BatchDeploymentProperties.Defaults()

	return &tmp
}

func LookupBatchDeploymentOutput(ctx *pulumi.Context, args LookupBatchDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupBatchDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBatchDeploymentResult, error) {
			args := v.(LookupBatchDeploymentArgs)
			r, err := LookupBatchDeployment(ctx, &args, opts...)
			var s LookupBatchDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBatchDeploymentResultOutput)
}

type LookupBatchDeploymentOutputArgs struct {
	DeploymentName    pulumi.StringInput `pulumi:"deploymentName"`
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBatchDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchDeploymentArgs)(nil)).Elem()
}

type LookupBatchDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupBatchDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchDeploymentResult)(nil)).Elem()
}

func (o LookupBatchDeploymentResultOutput) ToLookupBatchDeploymentResultOutput() LookupBatchDeploymentResultOutput {
	return o
}

func (o LookupBatchDeploymentResultOutput) ToLookupBatchDeploymentResultOutputWithContext(ctx context.Context) LookupBatchDeploymentResultOutput {
	return o
}

func (o LookupBatchDeploymentResultOutput) BatchDeploymentProperties() BatchDeploymentResponseOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) BatchDeploymentResponse { return v.BatchDeploymentProperties }).(BatchDeploymentResponseOutput)
}

func (o LookupBatchDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBatchDeploymentResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupBatchDeploymentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupBatchDeploymentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBatchDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBatchDeploymentResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupBatchDeploymentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBatchDeploymentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBatchDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBatchDeploymentResultOutput{})
}
