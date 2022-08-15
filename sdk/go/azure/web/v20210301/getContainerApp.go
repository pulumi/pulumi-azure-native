


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainerApp(ctx *pulumi.Context, args *LookupContainerAppArgs, opts ...pulumi.InvokeOption) (*LookupContainerAppResult, error) {
	var rv LookupContainerAppResult
	err := ctx.Invoke("azure-native:web/v20210301:getContainerApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupContainerAppArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupContainerAppResult struct {
	Configuration      *ConfigurationResponse `pulumi:"configuration"`
	Id                 string                 `pulumi:"id"`
	Kind               *string                `pulumi:"kind"`
	KubeEnvironmentId  *string                `pulumi:"kubeEnvironmentId"`
	LatestRevisionFqdn string                 `pulumi:"latestRevisionFqdn"`
	LatestRevisionName string                 `pulumi:"latestRevisionName"`
	Location           string                 `pulumi:"location"`
	Name               string                 `pulumi:"name"`
	ProvisioningState  string                 `pulumi:"provisioningState"`
	Tags               map[string]string      `pulumi:"tags"`
	Template           *TemplateResponse      `pulumi:"template"`
	Type               string                 `pulumi:"type"`
}


func (val *LookupContainerAppResult) Defaults() *LookupContainerAppResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Configuration = tmp.Configuration.Defaults()

	return &tmp
}

func LookupContainerAppOutput(ctx *pulumi.Context, args LookupContainerAppOutputArgs, opts ...pulumi.InvokeOption) LookupContainerAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerAppResult, error) {
			args := v.(LookupContainerAppArgs)
			r, err := LookupContainerApp(ctx, &args, opts...)
			var s LookupContainerAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerAppResultOutput)
}

type LookupContainerAppOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppArgs)(nil)).Elem()
}


type LookupContainerAppResultOutput struct{ *pulumi.OutputState }

func (LookupContainerAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppResult)(nil)).Elem()
}

func (o LookupContainerAppResultOutput) ToLookupContainerAppResultOutput() LookupContainerAppResultOutput {
	return o
}

func (o LookupContainerAppResultOutput) ToLookupContainerAppResultOutputWithContext(ctx context.Context) LookupContainerAppResultOutput {
	return o
}

func (o LookupContainerAppResultOutput) Configuration() ConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppResult) *ConfigurationResponse { return v.Configuration }).(ConfigurationResponsePtrOutput)
}

func (o LookupContainerAppResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupContainerAppResultOutput) KubeEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppResult) *string { return v.KubeEnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LookupContainerAppResultOutput) LatestRevisionFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.LatestRevisionFqdn }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) LatestRevisionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.LatestRevisionName }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerAppResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupContainerAppResultOutput) Template() TemplateResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppResult) *TemplateResponse { return v.Template }).(TemplateResponsePtrOutput)
}

func (o LookupContainerAppResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerAppResultOutput{})
}
