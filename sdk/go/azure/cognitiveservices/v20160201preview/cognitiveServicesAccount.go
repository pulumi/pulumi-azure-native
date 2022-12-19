


package v20160201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type CognitiveServicesAccount struct {
	pulumi.CustomResourceState

	Endpoint          pulumi.StringPtrOutput `pulumi:"endpoint"`
	Etag              pulumi.StringPtrOutput `pulumi:"etag"`
	Kind              pulumi.StringPtrOutput `pulumi:"kind"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	Name              pulumi.StringPtrOutput `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Sku               SkuResponsePtrOutput   `pulumi:"sku"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringPtrOutput `pulumi:"type"`
}


func NewCognitiveServicesAccount(ctx *pulumi.Context,
	name string, args *CognitiveServicesAccountArgs, opts ...pulumi.ResourceOption) (*CognitiveServicesAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20170418:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20210430:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20211001:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20220301:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20221001:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20221201:CognitiveServicesAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource CognitiveServicesAccount
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20160201preview:CognitiveServicesAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCognitiveServicesAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CognitiveServicesAccountState, opts ...pulumi.ResourceOption) (*CognitiveServicesAccount, error) {
	var resource CognitiveServicesAccount
	err := ctx.ReadResource("azure-native:cognitiveservices/v20160201preview:CognitiveServicesAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cognitiveServicesAccountState struct {
}

type CognitiveServicesAccountState struct {
}

func (CognitiveServicesAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*cognitiveServicesAccountState)(nil)).Elem()
}

type cognitiveServicesAccountArgs struct {
	AccountName       *string           `pulumi:"accountName"`
	Kind              string            `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               Sku               `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type CognitiveServicesAccountArgs struct {
	AccountName       pulumi.StringPtrInput
	Kind              pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
	Tags              pulumi.StringMapInput
}

func (CognitiveServicesAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cognitiveServicesAccountArgs)(nil)).Elem()
}

type CognitiveServicesAccountInput interface {
	pulumi.Input

	ToCognitiveServicesAccountOutput() CognitiveServicesAccountOutput
	ToCognitiveServicesAccountOutputWithContext(ctx context.Context) CognitiveServicesAccountOutput
}

func (*CognitiveServicesAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccount)(nil)).Elem()
}

func (i *CognitiveServicesAccount) ToCognitiveServicesAccountOutput() CognitiveServicesAccountOutput {
	return i.ToCognitiveServicesAccountOutputWithContext(context.Background())
}

func (i *CognitiveServicesAccount) ToCognitiveServicesAccountOutputWithContext(ctx context.Context) CognitiveServicesAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountOutput)
}

type CognitiveServicesAccountOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccount)(nil)).Elem()
}

func (o CognitiveServicesAccountOutput) ToCognitiveServicesAccountOutput() CognitiveServicesAccountOutput {
	return o
}

func (o CognitiveServicesAccountOutput) ToCognitiveServicesAccountOutputWithContext(ctx context.Context) CognitiveServicesAccountOutput {
	return o
}

func (o CognitiveServicesAccountOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CognitiveServicesAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CognitiveServicesAccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o CognitiveServicesAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CognitiveServicesAccountOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CognitiveServicesAccountOutput{})
}
