


package managednetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedNetworkPeeringPolicy(ctx *pulumi.Context, args *LookupManagedNetworkPeeringPolicyArgs, opts ...pulumi.InvokeOption) (*LookupManagedNetworkPeeringPolicyResult, error) {
	var rv LookupManagedNetworkPeeringPolicyResult
	err := ctx.Invoke("azure-native:managednetwork:getManagedNetworkPeeringPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedNetworkPeeringPolicyArgs struct {
	ManagedNetworkName              string `pulumi:"managedNetworkName"`
	ManagedNetworkPeeringPolicyName string `pulumi:"managedNetworkPeeringPolicyName"`
	ResourceGroupName               string `pulumi:"resourceGroupName"`
}


type LookupManagedNetworkPeeringPolicyResult struct {
	Id         string                                        `pulumi:"id"`
	Location   *string                                       `pulumi:"location"`
	Name       string                                        `pulumi:"name"`
	Properties ManagedNetworkPeeringPolicyPropertiesResponse `pulumi:"properties"`
	Type       string                                        `pulumi:"type"`
}

func LookupManagedNetworkPeeringPolicyOutput(ctx *pulumi.Context, args LookupManagedNetworkPeeringPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupManagedNetworkPeeringPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedNetworkPeeringPolicyResult, error) {
			args := v.(LookupManagedNetworkPeeringPolicyArgs)
			r, err := LookupManagedNetworkPeeringPolicy(ctx, &args, opts...)
			var s LookupManagedNetworkPeeringPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedNetworkPeeringPolicyResultOutput)
}

type LookupManagedNetworkPeeringPolicyOutputArgs struct {
	ManagedNetworkName              pulumi.StringInput `pulumi:"managedNetworkName"`
	ManagedNetworkPeeringPolicyName pulumi.StringInput `pulumi:"managedNetworkPeeringPolicyName"`
	ResourceGroupName               pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedNetworkPeeringPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkPeeringPolicyArgs)(nil)).Elem()
}


type LookupManagedNetworkPeeringPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupManagedNetworkPeeringPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkPeeringPolicyResult)(nil)).Elem()
}

func (o LookupManagedNetworkPeeringPolicyResultOutput) ToLookupManagedNetworkPeeringPolicyResultOutput() LookupManagedNetworkPeeringPolicyResultOutput {
	return o
}

func (o LookupManagedNetworkPeeringPolicyResultOutput) ToLookupManagedNetworkPeeringPolicyResultOutputWithContext(ctx context.Context) LookupManagedNetworkPeeringPolicyResultOutput {
	return o
}

func (o LookupManagedNetworkPeeringPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkPeeringPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkPeeringPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedNetworkPeeringPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupManagedNetworkPeeringPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkPeeringPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkPeeringPolicyResultOutput) Properties() ManagedNetworkPeeringPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v LookupManagedNetworkPeeringPolicyResult) ManagedNetworkPeeringPolicyPropertiesResponse {
		return v.Properties
	}).(ManagedNetworkPeeringPolicyPropertiesResponseOutput)
}

func (o LookupManagedNetworkPeeringPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkPeeringPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedNetworkPeeringPolicyResultOutput{})
}
