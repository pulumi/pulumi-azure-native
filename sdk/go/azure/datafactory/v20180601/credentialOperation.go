


package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CredentialOperation struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput                     `pulumi:"etag"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties ManagedIdentityCredentialResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewCredentialOperation(ctx *pulumi.Context,
	name string, args *CredentialOperationArgs, opts ...pulumi.ResourceOption) (*CredentialOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory:CredentialOperation"),
		},
	})
	opts = append(opts, aliases)
	var resource CredentialOperation
	err := ctx.RegisterResource("azure-native:datafactory/v20180601:CredentialOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCredentialOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialOperationState, opts ...pulumi.ResourceOption) (*CredentialOperation, error) {
	var resource CredentialOperation
	err := ctx.ReadResource("azure-native:datafactory/v20180601:CredentialOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type credentialOperationState struct {
}

type CredentialOperationState struct {
}

func (CredentialOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialOperationState)(nil)).Elem()
}

type credentialOperationArgs struct {
	CredentialName    *string                   `pulumi:"credentialName"`
	FactoryName       string                    `pulumi:"factoryName"`
	Properties        ManagedIdentityCredential `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
}


type CredentialOperationArgs struct {
	CredentialName    pulumi.StringPtrInput
	FactoryName       pulumi.StringInput
	Properties        ManagedIdentityCredentialInput
	ResourceGroupName pulumi.StringInput
}

func (CredentialOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialOperationArgs)(nil)).Elem()
}

type CredentialOperationInput interface {
	pulumi.Input

	ToCredentialOperationOutput() CredentialOperationOutput
	ToCredentialOperationOutputWithContext(ctx context.Context) CredentialOperationOutput
}

func (*CredentialOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialOperation)(nil)).Elem()
}

func (i *CredentialOperation) ToCredentialOperationOutput() CredentialOperationOutput {
	return i.ToCredentialOperationOutputWithContext(context.Background())
}

func (i *CredentialOperation) ToCredentialOperationOutputWithContext(ctx context.Context) CredentialOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialOperationOutput)
}

type CredentialOperationOutput struct{ *pulumi.OutputState }

func (CredentialOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialOperation)(nil)).Elem()
}

func (o CredentialOperationOutput) ToCredentialOperationOutput() CredentialOperationOutput {
	return o
}

func (o CredentialOperationOutput) ToCredentialOperationOutputWithContext(ctx context.Context) CredentialOperationOutput {
	return o
}

func (o CredentialOperationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialOperation) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o CredentialOperationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialOperation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CredentialOperationOutput) Properties() ManagedIdentityCredentialResponseOutput {
	return o.ApplyT(func(v *CredentialOperation) ManagedIdentityCredentialResponseOutput { return v.Properties }).(ManagedIdentityCredentialResponseOutput)
}

func (o CredentialOperationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialOperation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CredentialOperationOutput{})
}
