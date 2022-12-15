


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Secret struct {
	pulumi.CustomResourceState

	DeploymentStatus  pulumi.StringOutput      `pulumi:"deploymentStatus"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	Parameters        pulumi.AnyOutput         `pulumi:"parameters"`
	ProfileName       pulumi.StringOutput      `pulumi:"profileName"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
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
			Type: pulumi.String("azure-native:cdn:Secret"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Secret"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Secret"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:Secret"),
		},
	})
	opts = append(opts, aliases)
	var resource Secret
	err := ctx.RegisterResource("azure-native:cdn/v20220501preview:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("azure-native:cdn/v20220501preview:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type secretState struct {
}

type SecretState struct {
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	Parameters        interface{} `pulumi:"parameters"`
	ProfileName       string      `pulumi:"profileName"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	SecretName        *string     `pulumi:"secretName"`
}


type SecretArgs struct {
	Parameters        pulumi.Input
	ProfileName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SecretName        pulumi.StringPtrInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}

type SecretInput interface {
	pulumi.Input

	ToSecretOutput() SecretOutput
	ToSecretOutputWithContext(ctx context.Context) SecretOutput
}

func (*Secret) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (i *Secret) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i *Secret) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}

type SecretOutput struct{ *pulumi.OutputState }

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

func (o SecretOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o SecretOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecretOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *Secret) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

func (o SecretOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.ProfileName }).(pulumi.StringOutput)
}

func (o SecretOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SecretOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Secret) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SecretOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretOutput{})
}
