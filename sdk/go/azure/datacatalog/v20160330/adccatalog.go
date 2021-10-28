


package v20160330

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
			Type: pulumi.String("azure-nextgen:datacatalog/v20160330:ADCCatalog"),
		},
		{
			Type: pulumi.String("azure-native:datacatalog:ADCCatalog"),
		},
		{
			Type: pulumi.String("azure-nextgen:datacatalog:ADCCatalog"),
		},
	})
	opts = append(opts, aliases)
	var resource ADCCatalog
	err := ctx.RegisterResource("azure-native:datacatalog/v20160330:ADCCatalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADCCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADCCatalogState, opts ...pulumi.ResourceOption) (*ADCCatalog, error) {
	var resource ADCCatalog
	err := ctx.ReadResource("azure-native:datacatalog/v20160330:ADCCatalog", name, id, state, &resource, opts...)
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
	Etag                          *string           `pulumi:"etag"`
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
	Etag                          pulumi.StringPtrInput
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
	return reflect.TypeOf((*ADCCatalog)(nil))
}

func (i *ADCCatalog) ToADCCatalogOutput() ADCCatalogOutput {
	return i.ToADCCatalogOutputWithContext(context.Background())
}

func (i *ADCCatalog) ToADCCatalogOutputWithContext(ctx context.Context) ADCCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADCCatalogOutput)
}

type ADCCatalogOutput struct{ *pulumi.OutputState }

func (ADCCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ADCCatalog)(nil))
}

func (o ADCCatalogOutput) ToADCCatalogOutput() ADCCatalogOutput {
	return o
}

func (o ADCCatalogOutput) ToADCCatalogOutputWithContext(ctx context.Context) ADCCatalogOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ADCCatalogOutput{})
}
