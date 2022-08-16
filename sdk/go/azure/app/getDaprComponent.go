


package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDaprComponent(ctx *pulumi.Context, args *LookupDaprComponentArgs, opts ...pulumi.InvokeOption) (*LookupDaprComponentResult, error) {
	var rv LookupDaprComponentResult
	err := ctx.Invoke("azure-native:app:getDaprComponent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDaprComponentArgs struct {
	ComponentName     string `pulumi:"componentName"`
	EnvironmentName   string `pulumi:"environmentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDaprComponentResult struct {
	ComponentType *string                `pulumi:"componentType"`
	Id            string                 `pulumi:"id"`
	IgnoreErrors  *bool                  `pulumi:"ignoreErrors"`
	InitTimeout   *string                `pulumi:"initTimeout"`
	Metadata      []DaprMetadataResponse `pulumi:"metadata"`
	Name          string                 `pulumi:"name"`
	Scopes        []string               `pulumi:"scopes"`
	Secrets       []SecretResponse       `pulumi:"secrets"`
	SystemData    SystemDataResponse     `pulumi:"systemData"`
	Type          string                 `pulumi:"type"`
	Version       *string                `pulumi:"version"`
}

func LookupDaprComponentOutput(ctx *pulumi.Context, args LookupDaprComponentOutputArgs, opts ...pulumi.InvokeOption) LookupDaprComponentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDaprComponentResult, error) {
			args := v.(LookupDaprComponentArgs)
			r, err := LookupDaprComponent(ctx, &args, opts...)
			var s LookupDaprComponentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDaprComponentResultOutput)
}

type LookupDaprComponentOutputArgs struct {
	ComponentName     pulumi.StringInput `pulumi:"componentName"`
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDaprComponentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDaprComponentArgs)(nil)).Elem()
}


type LookupDaprComponentResultOutput struct{ *pulumi.OutputState }

func (LookupDaprComponentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDaprComponentResult)(nil)).Elem()
}

func (o LookupDaprComponentResultOutput) ToLookupDaprComponentResultOutput() LookupDaprComponentResultOutput {
	return o
}

func (o LookupDaprComponentResultOutput) ToLookupDaprComponentResultOutputWithContext(ctx context.Context) LookupDaprComponentResultOutput {
	return o
}

func (o LookupDaprComponentResultOutput) ComponentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) *string { return v.ComponentType }).(pulumi.StringPtrOutput)
}

func (o LookupDaprComponentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDaprComponentResultOutput) IgnoreErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) *bool { return v.IgnoreErrors }).(pulumi.BoolPtrOutput)
}

func (o LookupDaprComponentResultOutput) InitTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) *string { return v.InitTimeout }).(pulumi.StringPtrOutput)
}

func (o LookupDaprComponentResultOutput) Metadata() DaprMetadataResponseArrayOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) []DaprMetadataResponse { return v.Metadata }).(DaprMetadataResponseArrayOutput)
}

func (o LookupDaprComponentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDaprComponentResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o LookupDaprComponentResultOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) []SecretResponse { return v.Secrets }).(SecretResponseArrayOutput)
}

func (o LookupDaprComponentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDaprComponentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDaprComponentResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDaprComponentResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDaprComponentResultOutput{})
}
