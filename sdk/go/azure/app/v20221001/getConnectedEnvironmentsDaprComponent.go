


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectedEnvironmentsDaprComponent(ctx *pulumi.Context, args *LookupConnectedEnvironmentsDaprComponentArgs, opts ...pulumi.InvokeOption) (*LookupConnectedEnvironmentsDaprComponentResult, error) {
	var rv LookupConnectedEnvironmentsDaprComponentResult
	err := ctx.Invoke("azure-native:app/v20221001:getConnectedEnvironmentsDaprComponent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectedEnvironmentsDaprComponentArgs struct {
	ComponentName            string `pulumi:"componentName"`
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type LookupConnectedEnvironmentsDaprComponentResult struct {
	ComponentType        *string                `pulumi:"componentType"`
	Id                   string                 `pulumi:"id"`
	IgnoreErrors         *bool                  `pulumi:"ignoreErrors"`
	InitTimeout          *string                `pulumi:"initTimeout"`
	Metadata             []DaprMetadataResponse `pulumi:"metadata"`
	Name                 string                 `pulumi:"name"`
	Scopes               []string               `pulumi:"scopes"`
	SecretStoreComponent *string                `pulumi:"secretStoreComponent"`
	Secrets              []SecretResponse       `pulumi:"secrets"`
	SystemData           SystemDataResponse     `pulumi:"systemData"`
	Type                 string                 `pulumi:"type"`
	Version              *string                `pulumi:"version"`
}


func (val *LookupConnectedEnvironmentsDaprComponentResult) Defaults() *LookupConnectedEnvironmentsDaprComponentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IgnoreErrors) {
		ignoreErrors_ := false
		tmp.IgnoreErrors = &ignoreErrors_
	}
	return &tmp
}

func LookupConnectedEnvironmentsDaprComponentOutput(ctx *pulumi.Context, args LookupConnectedEnvironmentsDaprComponentOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedEnvironmentsDaprComponentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedEnvironmentsDaprComponentResult, error) {
			args := v.(LookupConnectedEnvironmentsDaprComponentArgs)
			r, err := LookupConnectedEnvironmentsDaprComponent(ctx, &args, opts...)
			var s LookupConnectedEnvironmentsDaprComponentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedEnvironmentsDaprComponentResultOutput)
}

type LookupConnectedEnvironmentsDaprComponentOutputArgs struct {
	ComponentName            pulumi.StringInput `pulumi:"componentName"`
	ConnectedEnvironmentName pulumi.StringInput `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectedEnvironmentsDaprComponentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsDaprComponentArgs)(nil)).Elem()
}


type LookupConnectedEnvironmentsDaprComponentResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedEnvironmentsDaprComponentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsDaprComponentResult)(nil)).Elem()
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) ToLookupConnectedEnvironmentsDaprComponentResultOutput() LookupConnectedEnvironmentsDaprComponentResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) ToLookupConnectedEnvironmentsDaprComponentResultOutputWithContext(ctx context.Context) LookupConnectedEnvironmentsDaprComponentResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) ComponentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *string { return v.ComponentType }).(pulumi.StringPtrOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) IgnoreErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *bool { return v.IgnoreErrors }).(pulumi.BoolPtrOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) InitTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *string { return v.InitTimeout }).(pulumi.StringPtrOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Metadata() DaprMetadataResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) []DaprMetadataResponse { return v.Metadata }).(DaprMetadataResponseArrayOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) SecretStoreComponent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *string { return v.SecretStoreComponent }).(pulumi.StringPtrOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) []SecretResponse { return v.Secrets }).(SecretResponseArrayOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedEnvironmentsDaprComponentResultOutput{})
}
