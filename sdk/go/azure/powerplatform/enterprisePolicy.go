


package powerplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnterprisePolicy struct {
	pulumi.CustomResourceState

	Encryption       PropertiesResponseEncryptionPtrOutput       `pulumi:"encryption"`
	Identity         EnterprisePolicyIdentityResponsePtrOutput   `pulumi:"identity"`
	Kind             pulumi.StringOutput                         `pulumi:"kind"`
	Location         pulumi.StringOutput                         `pulumi:"location"`
	Lockbox          PropertiesResponseLockboxPtrOutput          `pulumi:"lockbox"`
	Name             pulumi.StringOutput                         `pulumi:"name"`
	NetworkInjection PropertiesResponseNetworkInjectionPtrOutput `pulumi:"networkInjection"`
	SystemData       SystemDataResponseOutput                    `pulumi:"systemData"`
	Tags             pulumi.StringMapOutput                      `pulumi:"tags"`
	Type             pulumi.StringOutput                         `pulumi:"type"`
}


func NewEnterprisePolicy(ctx *pulumi.Context,
	name string, args *EnterprisePolicyArgs, opts ...pulumi.ResourceOption) (*EnterprisePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:powerplatform:EnterprisePolicy"),
		},
		{
			Type: pulumi.String("azure-native:powerplatform/v20201030preview:EnterprisePolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:powerplatform/v20201030preview:EnterprisePolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource EnterprisePolicy
	err := ctx.RegisterResource("azure-native:powerplatform:EnterprisePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnterprisePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterprisePolicyState, opts ...pulumi.ResourceOption) (*EnterprisePolicy, error) {
	var resource EnterprisePolicy
	err := ctx.ReadResource("azure-native:powerplatform:EnterprisePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type enterprisePolicyState struct {
}

type EnterprisePolicyState struct {
}

func (EnterprisePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterprisePolicyState)(nil)).Elem()
}

type enterprisePolicyArgs struct {
	Encryption           *PropertiesEncryption       `pulumi:"encryption"`
	EnterprisePolicyName *string                     `pulumi:"enterprisePolicyName"`
	Identity             *EnterprisePolicyIdentity   `pulumi:"identity"`
	Kind                 string                      `pulumi:"kind"`
	Location             *string                     `pulumi:"location"`
	Lockbox              *PropertiesLockbox          `pulumi:"lockbox"`
	NetworkInjection     *PropertiesNetworkInjection `pulumi:"networkInjection"`
	ResourceGroupName    string                      `pulumi:"resourceGroupName"`
	Tags                 map[string]string           `pulumi:"tags"`
}


type EnterprisePolicyArgs struct {
	Encryption           PropertiesEncryptionPtrInput
	EnterprisePolicyName pulumi.StringPtrInput
	Identity             EnterprisePolicyIdentityPtrInput
	Kind                 pulumi.StringInput
	Location             pulumi.StringPtrInput
	Lockbox              PropertiesLockboxPtrInput
	NetworkInjection     PropertiesNetworkInjectionPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
}

func (EnterprisePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterprisePolicyArgs)(nil)).Elem()
}

type EnterprisePolicyInput interface {
	pulumi.Input

	ToEnterprisePolicyOutput() EnterprisePolicyOutput
	ToEnterprisePolicyOutputWithContext(ctx context.Context) EnterprisePolicyOutput
}

func (*EnterprisePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterprisePolicy)(nil))
}

func (i *EnterprisePolicy) ToEnterprisePolicyOutput() EnterprisePolicyOutput {
	return i.ToEnterprisePolicyOutputWithContext(context.Background())
}

func (i *EnterprisePolicy) ToEnterprisePolicyOutputWithContext(ctx context.Context) EnterprisePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePolicyOutput)
}

type EnterprisePolicyOutput struct{ *pulumi.OutputState }

func (EnterprisePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterprisePolicy)(nil))
}

func (o EnterprisePolicyOutput) ToEnterprisePolicyOutput() EnterprisePolicyOutput {
	return o
}

func (o EnterprisePolicyOutput) ToEnterprisePolicyOutputWithContext(ctx context.Context) EnterprisePolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EnterprisePolicyOutput{})
}
