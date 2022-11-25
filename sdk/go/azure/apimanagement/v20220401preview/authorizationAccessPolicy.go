


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthorizationAccessPolicy struct {
	pulumi.CustomResourceState

	Name     pulumi.StringOutput    `pulumi:"name"`
	ObjectId pulumi.StringPtrOutput `pulumi:"objectId"`
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewAuthorizationAccessPolicy(ctx *pulumi.Context,
	name string, args *AuthorizationAccessPolicyArgs, opts ...pulumi.ResourceOption) (*AuthorizationAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizationId == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationId'")
	}
	if args.AuthorizationProviderId == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationProviderId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource AuthorizationAccessPolicy
	err := ctx.RegisterResource("azure-native:apimanagement/v20220401preview:AuthorizationAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAuthorizationAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationAccessPolicyState, opts ...pulumi.ResourceOption) (*AuthorizationAccessPolicy, error) {
	var resource AuthorizationAccessPolicy
	err := ctx.ReadResource("azure-native:apimanagement/v20220401preview:AuthorizationAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type authorizationAccessPolicyState struct {
}

type AuthorizationAccessPolicyState struct {
}

func (AuthorizationAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationAccessPolicyState)(nil)).Elem()
}

type authorizationAccessPolicyArgs struct {
	AuthorizationAccessPolicyId *string `pulumi:"authorizationAccessPolicyId"`
	AuthorizationId             string  `pulumi:"authorizationId"`
	AuthorizationProviderId     string  `pulumi:"authorizationProviderId"`
	ObjectId                    *string `pulumi:"objectId"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
	ServiceName                 string  `pulumi:"serviceName"`
	TenantId                    *string `pulumi:"tenantId"`
}


type AuthorizationAccessPolicyArgs struct {
	AuthorizationAccessPolicyId pulumi.StringPtrInput
	AuthorizationId             pulumi.StringInput
	AuthorizationProviderId     pulumi.StringInput
	ObjectId                    pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	ServiceName                 pulumi.StringInput
	TenantId                    pulumi.StringPtrInput
}

func (AuthorizationAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationAccessPolicyArgs)(nil)).Elem()
}

type AuthorizationAccessPolicyInput interface {
	pulumi.Input

	ToAuthorizationAccessPolicyOutput() AuthorizationAccessPolicyOutput
	ToAuthorizationAccessPolicyOutputWithContext(ctx context.Context) AuthorizationAccessPolicyOutput
}

func (*AuthorizationAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationAccessPolicy)(nil)).Elem()
}

func (i *AuthorizationAccessPolicy) ToAuthorizationAccessPolicyOutput() AuthorizationAccessPolicyOutput {
	return i.ToAuthorizationAccessPolicyOutputWithContext(context.Background())
}

func (i *AuthorizationAccessPolicy) ToAuthorizationAccessPolicyOutputWithContext(ctx context.Context) AuthorizationAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationAccessPolicyOutput)
}

type AuthorizationAccessPolicyOutput struct{ *pulumi.OutputState }

func (AuthorizationAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationAccessPolicy)(nil)).Elem()
}

func (o AuthorizationAccessPolicyOutput) ToAuthorizationAccessPolicyOutput() AuthorizationAccessPolicyOutput {
	return o
}

func (o AuthorizationAccessPolicyOutput) ToAuthorizationAccessPolicyOutputWithContext(ctx context.Context) AuthorizationAccessPolicyOutput {
	return o
}

func (o AuthorizationAccessPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationAccessPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AuthorizationAccessPolicyOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationAccessPolicy) pulumi.StringPtrOutput { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o AuthorizationAccessPolicyOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationAccessPolicy) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AuthorizationAccessPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationAccessPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationAccessPolicyOutput{})
}
