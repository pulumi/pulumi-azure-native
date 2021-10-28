


package servicefabricmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecretValue struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput    `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
	Value             pulumi.StringPtrOutput `pulumi:"value"`
}


func NewSecretValue(ctx *pulumi.Context,
	name string, args *SecretValueArgs, opts ...pulumi.ResourceOption) (*SecretValue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SecretResourceName == nil {
		return nil, errors.New("invalid value for required argument 'SecretResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicefabricmesh:SecretValue"),
		},
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180901preview:SecretValue"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicefabricmesh/v20180901preview:SecretValue"),
		},
	})
	opts = append(opts, aliases)
	var resource SecretValue
	err := ctx.RegisterResource("azure-native:servicefabricmesh:SecretValue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecretValue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretValueState, opts ...pulumi.ResourceOption) (*SecretValue, error) {
	var resource SecretValue
	err := ctx.ReadResource("azure-native:servicefabricmesh:SecretValue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type secretValueState struct {
}

type SecretValueState struct {
}

func (SecretValueState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretValueState)(nil)).Elem()
}

type secretValueArgs struct {
	Location                *string           `pulumi:"location"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	SecretResourceName      string            `pulumi:"secretResourceName"`
	SecretValueResourceName *string           `pulumi:"secretValueResourceName"`
	Tags                    map[string]string `pulumi:"tags"`
	Value                   *string           `pulumi:"value"`
}


type SecretValueArgs struct {
	Location                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	SecretResourceName      pulumi.StringInput
	SecretValueResourceName pulumi.StringPtrInput
	Tags                    pulumi.StringMapInput
	Value                   pulumi.StringPtrInput
}

func (SecretValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretValueArgs)(nil)).Elem()
}

type SecretValueInput interface {
	pulumi.Input

	ToSecretValueOutput() SecretValueOutput
	ToSecretValueOutputWithContext(ctx context.Context) SecretValueOutput
}

func (*SecretValue) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretValue)(nil))
}

func (i *SecretValue) ToSecretValueOutput() SecretValueOutput {
	return i.ToSecretValueOutputWithContext(context.Background())
}

func (i *SecretValue) ToSecretValueOutputWithContext(ctx context.Context) SecretValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretValueOutput)
}

type SecretValueOutput struct{ *pulumi.OutputState }

func (SecretValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretValue)(nil))
}

func (o SecretValueOutput) ToSecretValueOutput() SecretValueOutput {
	return o
}

func (o SecretValueOutput) ToSecretValueOutputWithContext(ctx context.Context) SecretValueOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretValueInput)(nil)).Elem(), &SecretValue{})
	pulumi.RegisterOutputType(SecretValueOutput{})
}
