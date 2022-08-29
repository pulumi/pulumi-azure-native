


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteConnectionStrings struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput                   `pulumi:"kind"`
	Location   pulumi.StringOutput                      `pulumi:"location"`
	Name       pulumi.StringPtrOutput                   `pulumi:"name"`
	Properties ConnStringValueTypePairResponseMapOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                   `pulumi:"tags"`
	Type       pulumi.StringPtrOutput                   `pulumi:"type"`
}


func NewSiteConnectionStrings(ctx *pulumi.Context,
	name string, args *SiteConnectionStringsArgs, opts ...pulumi.ResourceOption) (*SiteConnectionStrings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteConnectionStrings"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteConnectionStrings
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteConnectionStrings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteConnectionStrings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteConnectionStringsState, opts ...pulumi.ResourceOption) (*SiteConnectionStrings, error) {
	var resource SiteConnectionStrings
	err := ctx.ReadResource("azure-native:web/v20150801:SiteConnectionStrings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteConnectionStringsState struct {
}

type SiteConnectionStringsState struct {
}

func (SiteConnectionStringsState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteConnectionStringsState)(nil)).Elem()
}

type siteConnectionStringsArgs struct {
	Id                *string                            `pulumi:"id"`
	Kind              *string                            `pulumi:"kind"`
	Location          *string                            `pulumi:"location"`
	Name              string                             `pulumi:"name"`
	Properties        map[string]ConnStringValueTypePair `pulumi:"properties"`
	ResourceGroupName string                             `pulumi:"resourceGroupName"`
	Tags              map[string]string                  `pulumi:"tags"`
	Type              *string                            `pulumi:"type"`
}


type SiteConnectionStringsArgs struct {
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        ConnStringValueTypePairMapInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (SiteConnectionStringsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteConnectionStringsArgs)(nil)).Elem()
}

type SiteConnectionStringsInput interface {
	pulumi.Input

	ToSiteConnectionStringsOutput() SiteConnectionStringsOutput
	ToSiteConnectionStringsOutputWithContext(ctx context.Context) SiteConnectionStringsOutput
}

func (*SiteConnectionStrings) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConnectionStrings)(nil)).Elem()
}

func (i *SiteConnectionStrings) ToSiteConnectionStringsOutput() SiteConnectionStringsOutput {
	return i.ToSiteConnectionStringsOutputWithContext(context.Background())
}

func (i *SiteConnectionStrings) ToSiteConnectionStringsOutputWithContext(ctx context.Context) SiteConnectionStringsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConnectionStringsOutput)
}

type SiteConnectionStringsOutput struct{ *pulumi.OutputState }

func (SiteConnectionStringsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConnectionStrings)(nil)).Elem()
}

func (o SiteConnectionStringsOutput) ToSiteConnectionStringsOutput() SiteConnectionStringsOutput {
	return o
}

func (o SiteConnectionStringsOutput) ToSiteConnectionStringsOutputWithContext(ctx context.Context) SiteConnectionStringsOutput {
	return o
}

func (o SiteConnectionStringsOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConnectionStrings) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteConnectionStringsOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteConnectionStrings) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteConnectionStringsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConnectionStrings) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteConnectionStringsOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v *SiteConnectionStrings) ConnStringValueTypePairResponseMapOutput { return v.Properties }).(ConnStringValueTypePairResponseMapOutput)
}

func (o SiteConnectionStringsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteConnectionStrings) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteConnectionStringsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConnectionStrings) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteConnectionStringsOutput{})
}
