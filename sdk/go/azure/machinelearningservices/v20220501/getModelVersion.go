


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupModelVersion(ctx *pulumi.Context, args *LookupModelVersionArgs, opts ...pulumi.InvokeOption) (*LookupModelVersionResult, error) {
	var rv LookupModelVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220501:getModelVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupModelVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupModelVersionResult struct {
	Id                     string               `pulumi:"id"`
	ModelVersionProperties ModelVersionResponse `pulumi:"modelVersionProperties"`
	Name                   string               `pulumi:"name"`
	SystemData             SystemDataResponse   `pulumi:"systemData"`
	Type                   string               `pulumi:"type"`
}


func (val *LookupModelVersionResult) Defaults() *LookupModelVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ModelVersionProperties = *tmp.ModelVersionProperties.Defaults()

	return &tmp
}

func LookupModelVersionOutput(ctx *pulumi.Context, args LookupModelVersionOutputArgs, opts ...pulumi.InvokeOption) LookupModelVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelVersionResult, error) {
			args := v.(LookupModelVersionArgs)
			r, err := LookupModelVersion(ctx, &args, opts...)
			var s LookupModelVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelVersionResultOutput)
}

type LookupModelVersionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupModelVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelVersionArgs)(nil)).Elem()
}


type LookupModelVersionResultOutput struct{ *pulumi.OutputState }

func (LookupModelVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelVersionResult)(nil)).Elem()
}

func (o LookupModelVersionResultOutput) ToLookupModelVersionResultOutput() LookupModelVersionResultOutput {
	return o
}

func (o LookupModelVersionResultOutput) ToLookupModelVersionResultOutputWithContext(ctx context.Context) LookupModelVersionResultOutput {
	return o
}

func (o LookupModelVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupModelVersionResultOutput) ModelVersionProperties() ModelVersionResponseOutput {
	return o.ApplyT(func(v LookupModelVersionResult) ModelVersionResponse { return v.ModelVersionProperties }).(ModelVersionResponseOutput)
}

func (o LookupModelVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupModelVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupModelVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupModelVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelVersionResultOutput{})
}
