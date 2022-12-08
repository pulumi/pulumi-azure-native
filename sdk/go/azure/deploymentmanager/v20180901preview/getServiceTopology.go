


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupServiceTopology(ctx *pulumi.Context, args *LookupServiceTopologyArgs, opts ...pulumi.InvokeOption) (*LookupServiceTopologyResult, error) {
	var rv LookupServiceTopologyResult
	err := ctx.Invoke("azure-native:deploymentmanager/v20180901preview:getServiceTopology", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceTopologyArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServiceTopologyName string `pulumi:"serviceTopologyName"`
}


type LookupServiceTopologyResult struct {
	ArtifactSourceId *string           `pulumi:"artifactSourceId"`
	Id               string            `pulumi:"id"`
	Location         string            `pulumi:"location"`
	Name             string            `pulumi:"name"`
	Tags             map[string]string `pulumi:"tags"`
	Type             string            `pulumi:"type"`
}

func LookupServiceTopologyOutput(ctx *pulumi.Context, args LookupServiceTopologyOutputArgs, opts ...pulumi.InvokeOption) LookupServiceTopologyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceTopologyResult, error) {
			args := v.(LookupServiceTopologyArgs)
			r, err := LookupServiceTopology(ctx, &args, opts...)
			var s LookupServiceTopologyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceTopologyResultOutput)
}

type LookupServiceTopologyOutputArgs struct {
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceTopologyName pulumi.StringInput `pulumi:"serviceTopologyName"`
}

func (LookupServiceTopologyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceTopologyArgs)(nil)).Elem()
}


type LookupServiceTopologyResultOutput struct{ *pulumi.OutputState }

func (LookupServiceTopologyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceTopologyResult)(nil)).Elem()
}

func (o LookupServiceTopologyResultOutput) ToLookupServiceTopologyResultOutput() LookupServiceTopologyResultOutput {
	return o
}

func (o LookupServiceTopologyResultOutput) ToLookupServiceTopologyResultOutputWithContext(ctx context.Context) LookupServiceTopologyResultOutput {
	return o
}

func (o LookupServiceTopologyResultOutput) ArtifactSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) *string { return v.ArtifactSourceId }).(pulumi.StringPtrOutput)
}

func (o LookupServiceTopologyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceTopologyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServiceTopologyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceTopologyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceTopologyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceTopologyResultOutput{})
}
