


package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityPolicy struct {
	pulumi.CustomResourceState

	DeploymentStatus  pulumi.StringOutput                                             `pulumi:"deploymentStatus"`
	Name              pulumi.StringOutput                                             `pulumi:"name"`
	Parameters        SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput `pulumi:"parameters"`
	ProvisioningState pulumi.StringOutput                                             `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                                        `pulumi:"systemData"`
	Type              pulumi.StringOutput                                             `pulumi:"type"`
}


func NewSecurityPolicy(ctx *pulumi.Context,
	name string, args *SecurityPolicyArgs, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn/v20200901:SecurityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:SecurityPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityPolicy
	err := ctx.RegisterResource("azure-native:cdn:SecurityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityPolicyState, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	var resource SecurityPolicy
	err := ctx.ReadResource("azure-native:cdn:SecurityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityPolicyState struct {
}

type SecurityPolicyState struct {
}

func (SecurityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyState)(nil)).Elem()
}

type securityPolicyArgs struct {
	Parameters         *SecurityPolicyWebApplicationFirewallParameters `pulumi:"parameters"`
	ProfileName        string                                          `pulumi:"profileName"`
	ResourceGroupName  string                                          `pulumi:"resourceGroupName"`
	SecurityPolicyName *string                                         `pulumi:"securityPolicyName"`
}


type SecurityPolicyArgs struct {
	Parameters         SecurityPolicyWebApplicationFirewallParametersPtrInput
	ProfileName        pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	SecurityPolicyName pulumi.StringPtrInput
}

func (SecurityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyArgs)(nil)).Elem()
}

type SecurityPolicyInput interface {
	pulumi.Input

	ToSecurityPolicyOutput() SecurityPolicyOutput
	ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput
}

func (*SecurityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicy)(nil)).Elem()
}

func (i *SecurityPolicy) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return i.ToSecurityPolicyOutputWithContext(context.Background())
}

func (i *SecurityPolicy) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyOutput)
}

type SecurityPolicyOutput struct{ *pulumi.OutputState }

func (SecurityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicy)(nil)).Elem()
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return o
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return o
}

func (o SecurityPolicyOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o SecurityPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecurityPolicyOutput) Parameters() SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput {
	return o.ApplyT(func(v *SecurityPolicy) SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput {
		return v.Parameters
	}).(SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput)
}

func (o SecurityPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SecurityPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SecurityPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SecurityPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityPolicyOutput{})
}
