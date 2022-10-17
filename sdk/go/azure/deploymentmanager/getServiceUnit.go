


package deploymentmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceUnit(ctx *pulumi.Context, args *LookupServiceUnitArgs, opts ...pulumi.InvokeOption) (*LookupServiceUnitResult, error) {
	var rv LookupServiceUnitResult
	err := ctx.Invoke("azure-native:deploymentmanager:getServiceUnit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceUnitArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServiceName         string `pulumi:"serviceName"`
	ServiceTopologyName string `pulumi:"serviceTopologyName"`
	ServiceUnitName     string `pulumi:"serviceUnitName"`
}


type LookupServiceUnitResult struct {
	Artifacts           *ServiceUnitArtifactsResponse `pulumi:"artifacts"`
	DeploymentMode      string                        `pulumi:"deploymentMode"`
	Id                  string                        `pulumi:"id"`
	Location            string                        `pulumi:"location"`
	Name                string                        `pulumi:"name"`
	Tags                map[string]string             `pulumi:"tags"`
	TargetResourceGroup string                        `pulumi:"targetResourceGroup"`
	Type                string                        `pulumi:"type"`
}

func LookupServiceUnitOutput(ctx *pulumi.Context, args LookupServiceUnitOutputArgs, opts ...pulumi.InvokeOption) LookupServiceUnitResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceUnitResult, error) {
			args := v.(LookupServiceUnitArgs)
			r, err := LookupServiceUnit(ctx, &args, opts...)
			var s LookupServiceUnitResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceUnitResultOutput)
}

type LookupServiceUnitOutputArgs struct {
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName         pulumi.StringInput `pulumi:"serviceName"`
	ServiceTopologyName pulumi.StringInput `pulumi:"serviceTopologyName"`
	ServiceUnitName     pulumi.StringInput `pulumi:"serviceUnitName"`
}

func (LookupServiceUnitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceUnitArgs)(nil)).Elem()
}


type LookupServiceUnitResultOutput struct{ *pulumi.OutputState }

func (LookupServiceUnitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceUnitResult)(nil)).Elem()
}

func (o LookupServiceUnitResultOutput) ToLookupServiceUnitResultOutput() LookupServiceUnitResultOutput {
	return o
}

func (o LookupServiceUnitResultOutput) ToLookupServiceUnitResultOutputWithContext(ctx context.Context) LookupServiceUnitResultOutput {
	return o
}

func (o LookupServiceUnitResultOutput) Artifacts() ServiceUnitArtifactsResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceUnitResult) *ServiceUnitArtifactsResponse { return v.Artifacts }).(ServiceUnitArtifactsResponsePtrOutput)
}

func (o LookupServiceUnitResultOutput) DeploymentMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceUnitResult) string { return v.DeploymentMode }).(pulumi.StringOutput)
}

func (o LookupServiceUnitResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceUnitResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceUnitResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceUnitResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServiceUnitResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceUnitResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceUnitResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceUnitResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceUnitResultOutput) TargetResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceUnitResult) string { return v.TargetResourceGroup }).(pulumi.StringOutput)
}

func (o LookupServiceUnitResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceUnitResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceUnitResultOutput{})
}
