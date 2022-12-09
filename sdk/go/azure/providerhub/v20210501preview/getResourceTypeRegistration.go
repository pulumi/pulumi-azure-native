


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceTypeRegistration(ctx *pulumi.Context, args *LookupResourceTypeRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupResourceTypeRegistrationResult, error) {
	var rv LookupResourceTypeRegistrationResult
	err := ctx.Invoke("azure-native:providerhub/v20210501preview:getResourceTypeRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceTypeRegistrationArgs struct {
	ProviderNamespace string `pulumi:"providerNamespace"`
	ResourceType      string `pulumi:"resourceType"`
}

type LookupResourceTypeRegistrationResult struct {
	Id         string                                     `pulumi:"id"`
	Name       string                                     `pulumi:"name"`
	Properties ResourceTypeRegistrationResponseProperties `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}

func LookupResourceTypeRegistrationOutput(ctx *pulumi.Context, args LookupResourceTypeRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupResourceTypeRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceTypeRegistrationResult, error) {
			args := v.(LookupResourceTypeRegistrationArgs)
			r, err := LookupResourceTypeRegistration(ctx, &args, opts...)
			var s LookupResourceTypeRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceTypeRegistrationResultOutput)
}

type LookupResourceTypeRegistrationOutputArgs struct {
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
	ResourceType      pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupResourceTypeRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceTypeRegistrationArgs)(nil)).Elem()
}

type LookupResourceTypeRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupResourceTypeRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceTypeRegistrationResult)(nil)).Elem()
}

func (o LookupResourceTypeRegistrationResultOutput) ToLookupResourceTypeRegistrationResultOutput() LookupResourceTypeRegistrationResultOutput {
	return o
}

func (o LookupResourceTypeRegistrationResultOutput) ToLookupResourceTypeRegistrationResultOutputWithContext(ctx context.Context) LookupResourceTypeRegistrationResultOutput {
	return o
}

func (o LookupResourceTypeRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceTypeRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourceTypeRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceTypeRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResourceTypeRegistrationResultOutput) Properties() ResourceTypeRegistrationResponsePropertiesOutput {
	return o.ApplyT(func(v LookupResourceTypeRegistrationResult) ResourceTypeRegistrationResponseProperties {
		return v.Properties
	}).(ResourceTypeRegistrationResponsePropertiesOutput)
}

func (o LookupResourceTypeRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceTypeRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceTypeRegistrationResultOutput{})
}
