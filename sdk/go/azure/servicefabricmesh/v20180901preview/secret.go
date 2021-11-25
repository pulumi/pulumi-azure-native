


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Secret struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                    `pulumi:"location"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties SecretResourcePropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                 `pulumi:"tags"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabricmesh:Secret"),
		},
	})
	opts = append(opts, aliases)
	var resource Secret
	err := ctx.RegisterResource("azure-native:servicefabricmesh/v20180901preview:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("azure-native:servicefabricmesh/v20180901preview:Secret", name, id, state, &resource, opts...)
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
	Location           *string                  `pulumi:"location"`
	Properties         SecretResourceProperties `pulumi:"properties"`
	ResourceGroupName  string                   `pulumi:"resourceGroupName"`
	SecretResourceName *string                  `pulumi:"secretResourceName"`
	Tags               map[string]string        `pulumi:"tags"`
}


type SecretArgs struct {
	Location           pulumi.StringPtrInput
	Properties         SecretResourcePropertiesInput
	ResourceGroupName  pulumi.StringInput
	SecretResourceName pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
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
	return reflect.TypeOf((*Secret)(nil))
}

func (i *Secret) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i *Secret) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}

type SecretOutput struct{ *pulumi.OutputState }

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Secret)(nil))
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecretOutput{})
}
