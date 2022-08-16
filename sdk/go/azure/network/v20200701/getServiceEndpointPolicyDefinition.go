


package v20200701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceEndpointPolicyDefinition(ctx *pulumi.Context, args *LookupServiceEndpointPolicyDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupServiceEndpointPolicyDefinitionResult, error) {
	var rv LookupServiceEndpointPolicyDefinitionResult
	err := ctx.Invoke("azure-native:network/v20200701:getServiceEndpointPolicyDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceEndpointPolicyDefinitionArgs struct {
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	ServiceEndpointPolicyDefinitionName string `pulumi:"serviceEndpointPolicyDefinitionName"`
	ServiceEndpointPolicyName           string `pulumi:"serviceEndpointPolicyName"`
}


type LookupServiceEndpointPolicyDefinitionResult struct {
	Description       *string  `pulumi:"description"`
	Etag              string   `pulumi:"etag"`
	Id                *string  `pulumi:"id"`
	Name              *string  `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Service           *string  `pulumi:"service"`
	ServiceResources  []string `pulumi:"serviceResources"`
}

func LookupServiceEndpointPolicyDefinitionOutput(ctx *pulumi.Context, args LookupServiceEndpointPolicyDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupServiceEndpointPolicyDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceEndpointPolicyDefinitionResult, error) {
			args := v.(LookupServiceEndpointPolicyDefinitionArgs)
			r, err := LookupServiceEndpointPolicyDefinition(ctx, &args, opts...)
			var s LookupServiceEndpointPolicyDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceEndpointPolicyDefinitionResultOutput)
}

type LookupServiceEndpointPolicyDefinitionOutputArgs struct {
	ResourceGroupName                   pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceEndpointPolicyDefinitionName pulumi.StringInput `pulumi:"serviceEndpointPolicyDefinitionName"`
	ServiceEndpointPolicyName           pulumi.StringInput `pulumi:"serviceEndpointPolicyName"`
}

func (LookupServiceEndpointPolicyDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointPolicyDefinitionArgs)(nil)).Elem()
}


type LookupServiceEndpointPolicyDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupServiceEndpointPolicyDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointPolicyDefinitionResult)(nil)).Elem()
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) ToLookupServiceEndpointPolicyDefinitionResultOutput() LookupServiceEndpointPolicyDefinitionResultOutput {
	return o
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) ToLookupServiceEndpointPolicyDefinitionResultOutputWithContext(ctx context.Context) LookupServiceEndpointPolicyDefinitionResultOutput {
	return o
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) *string { return v.Service }).(pulumi.StringPtrOutput)
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) ServiceResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) []string { return v.ServiceResources }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceEndpointPolicyDefinitionResultOutput{})
}
