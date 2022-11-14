


package v20200113preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Credential struct {
	pulumi.CustomResourceState

	CreationTime     pulumi.StringOutput    `pulumi:"creationTime"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	LastModifiedTime pulumi.StringOutput    `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Type             pulumi.StringOutput    `pulumi:"type"`
	UserName         pulumi.StringOutput    `pulumi:"userName"`
}


func NewCredential(ctx *pulumi.Context,
	name string, args *CredentialArgs, opts ...pulumi.ResourceOption) (*Credential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Credential"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Credential"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Credential"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:Credential"),
		},
	})
	opts = append(opts, aliases)
	var resource Credential
	err := ctx.RegisterResource("azure-native:automation/v20200113preview:Credential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialState, opts ...pulumi.ResourceOption) (*Credential, error) {
	var resource Credential
	err := ctx.ReadResource("azure-native:automation/v20200113preview:Credential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type credentialState struct {
}

type CredentialState struct {
}

func (CredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialState)(nil)).Elem()
}

type credentialArgs struct {
	AutomationAccountName string  `pulumi:"automationAccountName"`
	CredentialName        *string `pulumi:"credentialName"`
	Description           *string `pulumi:"description"`
	Name                  string  `pulumi:"name"`
	Password              string  `pulumi:"password"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	UserName              string  `pulumi:"userName"`
}


type CredentialArgs struct {
	AutomationAccountName pulumi.StringInput
	CredentialName        pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	Name                  pulumi.StringInput
	Password              pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	UserName              pulumi.StringInput
}

func (CredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialArgs)(nil)).Elem()
}

type CredentialInput interface {
	pulumi.Input

	ToCredentialOutput() CredentialOutput
	ToCredentialOutputWithContext(ctx context.Context) CredentialOutput
}

func (*Credential) ElementType() reflect.Type {
	return reflect.TypeOf((**Credential)(nil)).Elem()
}

func (i *Credential) ToCredentialOutput() CredentialOutput {
	return i.ToCredentialOutputWithContext(context.Background())
}

func (i *Credential) ToCredentialOutputWithContext(ctx context.Context) CredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialOutput)
}

type CredentialOutput struct{ *pulumi.OutputState }

func (CredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Credential)(nil)).Elem()
}

func (o CredentialOutput) ToCredentialOutput() CredentialOutput {
	return o
}

func (o CredentialOutput) ToCredentialOutputWithContext(ctx context.Context) CredentialOutput {
	return o
}

func (o CredentialOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o CredentialOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CredentialOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o CredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CredentialOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CredentialOutput{})
}
