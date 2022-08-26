


package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADCCatalog struct {
	pulumi.CustomResourceState

	Admins                        PrincipalsResponseArrayOutput `pulumi:"admins"`
	EnableAutomaticUnitAdjustment pulumi.BoolPtrOutput          `pulumi:"enableAutomaticUnitAdjustment"`
	Etag                          pulumi.StringPtrOutput        `pulumi:"etag"`
	Location                      pulumi.StringPtrOutput        `pulumi:"location"`
	Name                          pulumi.StringOutput           `pulumi:"name"`
	Sku                           pulumi.StringPtrOutput        `pulumi:"sku"`
	SuccessfullyProvisioned       pulumi.BoolPtrOutput          `pulumi:"successfullyProvisioned"`
	Tags                          pulumi.StringMapOutput        `pulumi:"tags"`
	Type                          pulumi.StringOutput           `pulumi:"type"`
	Units                         pulumi.IntPtrOutput           `pulumi:"units"`
	Users                         PrincipalsResponseArrayOutput `pulumi:"users"`
}


func NewADCCatalog(ctx *pulumi.Context,
	name string, args *ADCCatalogArgs, opts ...pulumi.ResourceOption) (*ADCCatalog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datacatalog/v20160330:ADCCatalog"),
		},
	})
	opts = append(opts, aliases)
	var resource ADCCatalog
	err := ctx.RegisterResource("azure-native:datacatalog:ADCCatalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADCCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADCCatalogState, opts ...pulumi.ResourceOption) (*ADCCatalog, error) {
	var resource ADCCatalog
	err := ctx.ReadResource("azure-native:datacatalog:ADCCatalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adccatalogState struct {
}

type ADCCatalogState struct {
}

func (ADCCatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*adccatalogState)(nil)).Elem()
}

type adccatalogArgs struct {
	Admins                        []Principals      `pulumi:"admins"`
	CatalogName                   *string           `pulumi:"catalogName"`
	EnableAutomaticUnitAdjustment *bool             `pulumi:"enableAutomaticUnitAdjustment"`
	Location                      *string           `pulumi:"location"`
	ResourceGroupName             string            `pulumi:"resourceGroupName"`
	Sku                           *string           `pulumi:"sku"`
	SuccessfullyProvisioned       *bool             `pulumi:"successfullyProvisioned"`
	Tags                          map[string]string `pulumi:"tags"`
	Units                         *int              `pulumi:"units"`
	Users                         []Principals      `pulumi:"users"`
}


type ADCCatalogArgs struct {
	Admins                        PrincipalsArrayInput
	CatalogName                   pulumi.StringPtrInput
	EnableAutomaticUnitAdjustment pulumi.BoolPtrInput
	Location                      pulumi.StringPtrInput
	ResourceGroupName             pulumi.StringInput
	Sku                           pulumi.StringPtrInput
	SuccessfullyProvisioned       pulumi.BoolPtrInput
	Tags                          pulumi.StringMapInput
	Units                         pulumi.IntPtrInput
	Users                         PrincipalsArrayInput
}

func (ADCCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adccatalogArgs)(nil)).Elem()
}

type ADCCatalogInput interface {
	pulumi.Input

	ToADCCatalogOutput() ADCCatalogOutput
	ToADCCatalogOutputWithContext(ctx context.Context) ADCCatalogOutput
}

func (*ADCCatalog) ElementType() reflect.Type {
	return reflect.TypeOf((**ADCCatalog)(nil)).Elem()
}

func (i *ADCCatalog) ToADCCatalogOutput() ADCCatalogOutput {
	return i.ToADCCatalogOutputWithContext(context.Background())
}

func (i *ADCCatalog) ToADCCatalogOutputWithContext(ctx context.Context) ADCCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADCCatalogOutput)
}

type ADCCatalogOutput struct{ *pulumi.OutputState }

func (ADCCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADCCatalog)(nil)).Elem()
}

func (o ADCCatalogOutput) ToADCCatalogOutput() ADCCatalogOutput {
	return o
}

func (o ADCCatalogOutput) ToADCCatalogOutputWithContext(ctx context.Context) ADCCatalogOutput {
	return o
}

func (o ADCCatalogOutput) Admins() PrincipalsResponseArrayOutput {
	return o.ApplyT(func(v *ADCCatalog) PrincipalsResponseArrayOutput { return v.Admins }).(PrincipalsResponseArrayOutput)
}

func (o ADCCatalogOutput) EnableAutomaticUnitAdjustment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.BoolPtrOutput { return v.EnableAutomaticUnitAdjustment }).(pulumi.BoolPtrOutput)
}

func (o ADCCatalogOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ADCCatalogOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ADCCatalogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ADCCatalogOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringPtrOutput { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ADCCatalogOutput) SuccessfullyProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.BoolPtrOutput { return v.SuccessfullyProvisioned }).(pulumi.BoolPtrOutput)
}

func (o ADCCatalogOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ADCCatalogOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ADCCatalogOutput) Units() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.IntPtrOutput { return v.Units }).(pulumi.IntPtrOutput)
}

func (o ADCCatalogOutput) Users() PrincipalsResponseArrayOutput {
	return o.ApplyT(func(v *ADCCatalog) PrincipalsResponseArrayOutput { return v.Users }).(PrincipalsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ADCCatalogOutput{})
}
