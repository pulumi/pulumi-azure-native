


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
			Type: pulumi.String("azure-native:servicefabricmesh/v20180901preview:SecretValue"),
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
	return reflect.TypeOf((**SecretValue)(nil)).Elem()
}

func (i *SecretValue) ToSecretValueOutput() SecretValueOutput {
	return i.ToSecretValueOutputWithContext(context.Background())
}

func (i *SecretValue) ToSecretValueOutputWithContext(ctx context.Context) SecretValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretValueOutput)
}

type SecretValueOutput struct{ *pulumi.OutputState }

func (SecretValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretValue)(nil)).Elem()
}

func (o SecretValueOutput) ToSecretValueOutput() SecretValueOutput {
	return o
}

func (o SecretValueOutput) ToSecretValueOutputWithContext(ctx context.Context) SecretValueOutput {
	return o
}

func (o SecretValueOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SecretValueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecretValueOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SecretValueOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SecretValueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SecretValueOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretValueOutput{})
}
