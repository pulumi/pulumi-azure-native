


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KeyValue struct {
	pulumi.CustomResourceState

	ContentType  pulumi.StringPtrOutput `pulumi:"contentType"`
	ETag         pulumi.StringOutput    `pulumi:"eTag"`
	Key          pulumi.StringOutput    `pulumi:"key"`
	Label        pulumi.StringOutput    `pulumi:"label"`
	LastModified pulumi.StringOutput    `pulumi:"lastModified"`
	Locked       pulumi.BoolOutput      `pulumi:"locked"`
	Name         pulumi.StringOutput    `pulumi:"name"`
	Tags         pulumi.StringMapOutput `pulumi:"tags"`
	Type         pulumi.StringOutput    `pulumi:"type"`
	Value        pulumi.StringPtrOutput `pulumi:"value"`
}


func NewKeyValue(ctx *pulumi.Context,
	name string, args *KeyValueArgs, opts ...pulumi.ResourceOption) (*KeyValue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigStoreName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigStoreName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appconfiguration:KeyValue"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20200701preview:KeyValue"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20210301preview:KeyValue"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20211001preview:KeyValue"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20220301preview:KeyValue"),
		},
	})
	opts = append(opts, aliases)
	var resource KeyValue
	err := ctx.RegisterResource("azure-native:appconfiguration/v20220501:KeyValue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKeyValue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyValueState, opts ...pulumi.ResourceOption) (*KeyValue, error) {
	var resource KeyValue
	err := ctx.ReadResource("azure-native:appconfiguration/v20220501:KeyValue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type keyValueState struct {
}

type KeyValueState struct {
}

func (KeyValueState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyValueState)(nil)).Elem()
}

type keyValueArgs struct {
	ConfigStoreName   string            `pulumi:"configStoreName"`
	ContentType       *string           `pulumi:"contentType"`
	KeyValueName      *string           `pulumi:"keyValueName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	Value             *string           `pulumi:"value"`
}


type KeyValueArgs struct {
	ConfigStoreName   pulumi.StringInput
	ContentType       pulumi.StringPtrInput
	KeyValueName      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Value             pulumi.StringPtrInput
}

func (KeyValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyValueArgs)(nil)).Elem()
}

type KeyValueInput interface {
	pulumi.Input

	ToKeyValueOutput() KeyValueOutput
	ToKeyValueOutputWithContext(ctx context.Context) KeyValueOutput
}

func (*KeyValue) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyValue)(nil)).Elem()
}

func (i *KeyValue) ToKeyValueOutput() KeyValueOutput {
	return i.ToKeyValueOutputWithContext(context.Background())
}

func (i *KeyValue) ToKeyValueOutputWithContext(ctx context.Context) KeyValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyValueOutput)
}

type KeyValueOutput struct{ *pulumi.OutputState }

func (KeyValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyValue)(nil)).Elem()
}

func (o KeyValueOutput) ToKeyValueOutput() KeyValueOutput {
	return o
}

func (o KeyValueOutput) ToKeyValueOutputWithContext(ctx context.Context) KeyValueOutput {
	return o
}

func (o KeyValueOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o KeyValueOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

func (o KeyValueOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o KeyValueOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

func (o KeyValueOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

func (o KeyValueOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.BoolOutput { return v.Locked }).(pulumi.BoolOutput)
}

func (o KeyValueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KeyValueOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o KeyValueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o KeyValueOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyValue) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyValueOutput{})
}
