


package v20220131preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FederatedIdentityCredential struct {
	pulumi.CustomResourceState

	Audiences pulumi.StringArrayOutput `pulumi:"audiences"`
	Issuer    pulumi.StringOutput      `pulumi:"issuer"`
	Name      pulumi.StringOutput      `pulumi:"name"`
	Subject   pulumi.StringOutput      `pulumi:"subject"`
	Type      pulumi.StringOutput      `pulumi:"type"`
}


func NewFederatedIdentityCredential(ctx *pulumi.Context,
	name string, args *FederatedIdentityCredentialArgs, opts ...pulumi.ResourceOption) (*FederatedIdentityCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Audiences == nil {
		return nil, errors.New("invalid value for required argument 'Audiences'")
	}
	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.Subject == nil {
		return nil, errors.New("invalid value for required argument 'Subject'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managedidentity:FederatedIdentityCredential"),
		},
	})
	opts = append(opts, aliases)
	var resource FederatedIdentityCredential
	err := ctx.RegisterResource("azure-native:managedidentity/v20220131preview:FederatedIdentityCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFederatedIdentityCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedIdentityCredentialState, opts ...pulumi.ResourceOption) (*FederatedIdentityCredential, error) {
	var resource FederatedIdentityCredential
	err := ctx.ReadResource("azure-native:managedidentity/v20220131preview:FederatedIdentityCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type federatedIdentityCredentialState struct {
}

type FederatedIdentityCredentialState struct {
}

func (FederatedIdentityCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedIdentityCredentialState)(nil)).Elem()
}

type federatedIdentityCredentialArgs struct {
	Audiences                               []string `pulumi:"audiences"`
	FederatedIdentityCredentialResourceName *string  `pulumi:"federatedIdentityCredentialResourceName"`
	Issuer                                  string   `pulumi:"issuer"`
	ResourceGroupName                       string   `pulumi:"resourceGroupName"`
	ResourceName                            string   `pulumi:"resourceName"`
	Subject                                 string   `pulumi:"subject"`
}


type FederatedIdentityCredentialArgs struct {
	Audiences                               pulumi.StringArrayInput
	FederatedIdentityCredentialResourceName pulumi.StringPtrInput
	Issuer                                  pulumi.StringInput
	ResourceGroupName                       pulumi.StringInput
	ResourceName                            pulumi.StringInput
	Subject                                 pulumi.StringInput
}

func (FederatedIdentityCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedIdentityCredentialArgs)(nil)).Elem()
}

type FederatedIdentityCredentialInput interface {
	pulumi.Input

	ToFederatedIdentityCredentialOutput() FederatedIdentityCredentialOutput
	ToFederatedIdentityCredentialOutputWithContext(ctx context.Context) FederatedIdentityCredentialOutput
}

func (*FederatedIdentityCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedIdentityCredential)(nil)).Elem()
}

func (i *FederatedIdentityCredential) ToFederatedIdentityCredentialOutput() FederatedIdentityCredentialOutput {
	return i.ToFederatedIdentityCredentialOutputWithContext(context.Background())
}

func (i *FederatedIdentityCredential) ToFederatedIdentityCredentialOutputWithContext(ctx context.Context) FederatedIdentityCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedIdentityCredentialOutput)
}

type FederatedIdentityCredentialOutput struct{ *pulumi.OutputState }

func (FederatedIdentityCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedIdentityCredential)(nil)).Elem()
}

func (o FederatedIdentityCredentialOutput) ToFederatedIdentityCredentialOutput() FederatedIdentityCredentialOutput {
	return o
}

func (o FederatedIdentityCredentialOutput) ToFederatedIdentityCredentialOutputWithContext(ctx context.Context) FederatedIdentityCredentialOutput {
	return o
}

func (o FederatedIdentityCredentialOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedIdentityCredential) pulumi.StringArrayOutput { return v.Audiences }).(pulumi.StringArrayOutput)
}

func (o FederatedIdentityCredentialOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedIdentityCredential) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

func (o FederatedIdentityCredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedIdentityCredential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FederatedIdentityCredentialOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedIdentityCredential) pulumi.StringOutput { return v.Subject }).(pulumi.StringOutput)
}

func (o FederatedIdentityCredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedIdentityCredential) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FederatedIdentityCredentialOutput{})
}
