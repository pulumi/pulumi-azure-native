


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistryModelVersion(ctx *pulumi.Context, args *LookupRegistryModelVersionArgs, opts ...pulumi.InvokeOption) (*LookupRegistryModelVersionResult, error) {
	var rv LookupRegistryModelVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getRegistryModelVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryModelVersionArgs struct {
	ModelName         string `pulumi:"modelName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
}


type LookupRegistryModelVersionResult struct {
	Id                     string               `pulumi:"id"`
	ModelVersionProperties ModelVersionResponse `pulumi:"modelVersionProperties"`
	Name                   string               `pulumi:"name"`
	SystemData             SystemDataResponse   `pulumi:"systemData"`
	Type                   string               `pulumi:"type"`
}


func (val *LookupRegistryModelVersionResult) Defaults() *LookupRegistryModelVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ModelVersionProperties = *tmp.ModelVersionProperties.Defaults()

	return &tmp
}

func LookupRegistryModelVersionOutput(ctx *pulumi.Context, args LookupRegistryModelVersionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryModelVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryModelVersionResult, error) {
			args := v.(LookupRegistryModelVersionArgs)
			r, err := LookupRegistryModelVersion(ctx, &args, opts...)
			var s LookupRegistryModelVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryModelVersionResultOutput)
}

type LookupRegistryModelVersionOutputArgs struct {
	ModelName         pulumi.StringInput `pulumi:"modelName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
}

func (LookupRegistryModelVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryModelVersionArgs)(nil)).Elem()
}


type LookupRegistryModelVersionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryModelVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryModelVersionResult)(nil)).Elem()
}

func (o LookupRegistryModelVersionResultOutput) ToLookupRegistryModelVersionResultOutput() LookupRegistryModelVersionResultOutput {
	return o
}

func (o LookupRegistryModelVersionResultOutput) ToLookupRegistryModelVersionResultOutputWithContext(ctx context.Context) LookupRegistryModelVersionResultOutput {
	return o
}

func (o LookupRegistryModelVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryModelVersionResultOutput) ModelVersionProperties() ModelVersionResponseOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) ModelVersionResponse { return v.ModelVersionProperties }).(ModelVersionResponseOutput)
}

func (o LookupRegistryModelVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryModelVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryModelVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryModelVersionResultOutput{})
}
