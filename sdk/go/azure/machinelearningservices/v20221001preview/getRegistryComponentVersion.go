


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistryComponentVersion(ctx *pulumi.Context, args *LookupRegistryComponentVersionArgs, opts ...pulumi.InvokeOption) (*LookupRegistryComponentVersionResult, error) {
	var rv LookupRegistryComponentVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getRegistryComponentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryComponentVersionArgs struct {
	ComponentName     string `pulumi:"componentName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
}


type LookupRegistryComponentVersionResult struct {
	ComponentVersionProperties ComponentVersionResponse `pulumi:"componentVersionProperties"`
	Id                         string                   `pulumi:"id"`
	Name                       string                   `pulumi:"name"`
	SystemData                 SystemDataResponse       `pulumi:"systemData"`
	Type                       string                   `pulumi:"type"`
}


func (val *LookupRegistryComponentVersionResult) Defaults() *LookupRegistryComponentVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ComponentVersionProperties = *tmp.ComponentVersionProperties.Defaults()

	return &tmp
}

func LookupRegistryComponentVersionOutput(ctx *pulumi.Context, args LookupRegistryComponentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryComponentVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryComponentVersionResult, error) {
			args := v.(LookupRegistryComponentVersionArgs)
			r, err := LookupRegistryComponentVersion(ctx, &args, opts...)
			var s LookupRegistryComponentVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryComponentVersionResultOutput)
}

type LookupRegistryComponentVersionOutputArgs struct {
	ComponentName     pulumi.StringInput `pulumi:"componentName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
}

func (LookupRegistryComponentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryComponentVersionArgs)(nil)).Elem()
}


type LookupRegistryComponentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryComponentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryComponentVersionResult)(nil)).Elem()
}

func (o LookupRegistryComponentVersionResultOutput) ToLookupRegistryComponentVersionResultOutput() LookupRegistryComponentVersionResultOutput {
	return o
}

func (o LookupRegistryComponentVersionResultOutput) ToLookupRegistryComponentVersionResultOutputWithContext(ctx context.Context) LookupRegistryComponentVersionResultOutput {
	return o
}

func (o LookupRegistryComponentVersionResultOutput) ComponentVersionProperties() ComponentVersionResponseOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) ComponentVersionResponse {
		return v.ComponentVersionProperties
	}).(ComponentVersionResponseOutput)
}

func (o LookupRegistryComponentVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryComponentVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryComponentVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryComponentVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryComponentVersionResultOutput{})
}
