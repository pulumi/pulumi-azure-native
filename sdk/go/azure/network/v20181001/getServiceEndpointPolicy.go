


package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceEndpointPolicy(ctx *pulumi.Context, args *LookupServiceEndpointPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServiceEndpointPolicyResult, error) {
	var rv LookupServiceEndpointPolicyResult
	err := ctx.Invoke("azure-native:network/v20181001:getServiceEndpointPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceEndpointPolicyArgs struct {
	Expand                    *string `pulumi:"expand"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	ServiceEndpointPolicyName string  `pulumi:"serviceEndpointPolicyName"`
}


type LookupServiceEndpointPolicyResult struct {
	Etag                             *string                                   `pulumi:"etag"`
	Id                               *string                                   `pulumi:"id"`
	Location                         *string                                   `pulumi:"location"`
	Name                             string                                    `pulumi:"name"`
	ProvisioningState                string                                    `pulumi:"provisioningState"`
	ResourceGuid                     string                                    `pulumi:"resourceGuid"`
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinitionResponse `pulumi:"serviceEndpointPolicyDefinitions"`
	Subnets                          []SubnetResponse                          `pulumi:"subnets"`
	Tags                             map[string]string                         `pulumi:"tags"`
	Type                             string                                    `pulumi:"type"`
}

func LookupServiceEndpointPolicyOutput(ctx *pulumi.Context, args LookupServiceEndpointPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupServiceEndpointPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceEndpointPolicyResult, error) {
			args := v.(LookupServiceEndpointPolicyArgs)
			r, err := LookupServiceEndpointPolicy(ctx, &args, opts...)
			var s LookupServiceEndpointPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceEndpointPolicyResultOutput)
}

type LookupServiceEndpointPolicyOutputArgs struct {
	Expand                    pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName         pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServiceEndpointPolicyName pulumi.StringInput    `pulumi:"serviceEndpointPolicyName"`
}

func (LookupServiceEndpointPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointPolicyArgs)(nil)).Elem()
}


type LookupServiceEndpointPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupServiceEndpointPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointPolicyResult)(nil)).Elem()
}

func (o LookupServiceEndpointPolicyResultOutput) ToLookupServiceEndpointPolicyResultOutput() LookupServiceEndpointPolicyResultOutput {
	return o
}

func (o LookupServiceEndpointPolicyResultOutput) ToLookupServiceEndpointPolicyResultOutputWithContext(ctx context.Context) LookupServiceEndpointPolicyResultOutput {
	return o
}

func (o LookupServiceEndpointPolicyResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupServiceEndpointPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupServiceEndpointPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupServiceEndpointPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointPolicyResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointPolicyResultOutput) ServiceEndpointPolicyDefinitions() ServiceEndpointPolicyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) []ServiceEndpointPolicyDefinitionResponse {
		return v.ServiceEndpointPolicyDefinitions
	}).(ServiceEndpointPolicyDefinitionResponseArrayOutput)
}

func (o LookupServiceEndpointPolicyResultOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o LookupServiceEndpointPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceEndpointPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceEndpointPolicyResultOutput{})
}
