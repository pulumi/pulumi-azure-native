


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationServiceEnvironment(ctx *pulumi.Context, args *LookupIntegrationServiceEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationServiceEnvironmentResult, error) {
	var rv LookupIntegrationServiceEnvironmentResult
	err := ctx.Invoke("azure-native:logic/v20190501:getIntegrationServiceEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationServiceEnvironmentArgs struct {
	IntegrationServiceEnvironmentName string `pulumi:"integrationServiceEnvironmentName"`
	ResourceGroup                     string `pulumi:"resourceGroup"`
}


type LookupIntegrationServiceEnvironmentResult struct {
	Id         string                                          `pulumi:"id"`
	Identity   *ManagedServiceIdentityResponse                 `pulumi:"identity"`
	Location   *string                                         `pulumi:"location"`
	Name       string                                          `pulumi:"name"`
	Properties IntegrationServiceEnvironmentPropertiesResponse `pulumi:"properties"`
	Sku        *IntegrationServiceEnvironmentSkuResponse       `pulumi:"sku"`
	Tags       map[string]string                               `pulumi:"tags"`
	Type       string                                          `pulumi:"type"`
}

func LookupIntegrationServiceEnvironmentOutput(ctx *pulumi.Context, args LookupIntegrationServiceEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationServiceEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationServiceEnvironmentResult, error) {
			args := v.(LookupIntegrationServiceEnvironmentArgs)
			r, err := LookupIntegrationServiceEnvironment(ctx, &args, opts...)
			var s LookupIntegrationServiceEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationServiceEnvironmentResultOutput)
}

type LookupIntegrationServiceEnvironmentOutputArgs struct {
	IntegrationServiceEnvironmentName pulumi.StringInput `pulumi:"integrationServiceEnvironmentName"`
	ResourceGroup                     pulumi.StringInput `pulumi:"resourceGroup"`
}

func (LookupIntegrationServiceEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationServiceEnvironmentArgs)(nil)).Elem()
}


type LookupIntegrationServiceEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationServiceEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationServiceEnvironmentResult)(nil)).Elem()
}

func (o LookupIntegrationServiceEnvironmentResultOutput) ToLookupIntegrationServiceEnvironmentResultOutput() LookupIntegrationServiceEnvironmentResultOutput {
	return o
}

func (o LookupIntegrationServiceEnvironmentResultOutput) ToLookupIntegrationServiceEnvironmentResultOutputWithContext(ctx context.Context) LookupIntegrationServiceEnvironmentResultOutput {
	return o
}

func (o LookupIntegrationServiceEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationServiceEnvironmentResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupIntegrationServiceEnvironmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationServiceEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationServiceEnvironmentResultOutput) Properties() IntegrationServiceEnvironmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentResult) IntegrationServiceEnvironmentPropertiesResponse {
		return v.Properties
	}).(IntegrationServiceEnvironmentPropertiesResponseOutput)
}

func (o LookupIntegrationServiceEnvironmentResultOutput) Sku() IntegrationServiceEnvironmentSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentResult) *IntegrationServiceEnvironmentSkuResponse {
		return v.Sku
	}).(IntegrationServiceEnvironmentSkuResponsePtrOutput)
}

func (o LookupIntegrationServiceEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationServiceEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationServiceEnvironmentResultOutput{})
}
