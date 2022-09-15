


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteSlotConfigNames struct {
	pulumi.CustomResourceState

	AppSettingNames       pulumi.StringArrayOutput `pulumi:"appSettingNames"`
	ConnectionStringNames pulumi.StringArrayOutput `pulumi:"connectionStringNames"`
	Kind                  pulumi.StringPtrOutput   `pulumi:"kind"`
	Location              pulumi.StringOutput      `pulumi:"location"`
	Name                  pulumi.StringPtrOutput   `pulumi:"name"`
	Tags                  pulumi.StringMapOutput   `pulumi:"tags"`
	Type                  pulumi.StringPtrOutput   `pulumi:"type"`
}


func NewSiteSlotConfigNames(ctx *pulumi.Context,
	name string, args *SiteSlotConfigNamesArgs, opts ...pulumi.ResourceOption) (*SiteSlotConfigNames, error) {
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
			Type: pulumi.String("azure-native:web:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteSlotConfigNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteSlotConfigNames"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteSlotConfigNames
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteSlotConfigNames", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteSlotConfigNames(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteSlotConfigNamesState, opts ...pulumi.ResourceOption) (*SiteSlotConfigNames, error) {
	var resource SiteSlotConfigNames
	err := ctx.ReadResource("azure-native:web/v20150801:SiteSlotConfigNames", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteSlotConfigNamesState struct {
}

type SiteSlotConfigNamesState struct {
}

func (SiteSlotConfigNamesState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSlotConfigNamesState)(nil)).Elem()
}

type siteSlotConfigNamesArgs struct {
	AppSettingNames       []string          `pulumi:"appSettingNames"`
	ConnectionStringNames []string          `pulumi:"connectionStringNames"`
	Id                    *string           `pulumi:"id"`
	Kind                  *string           `pulumi:"kind"`
	Location              *string           `pulumi:"location"`
	Name                  string            `pulumi:"name"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  *string           `pulumi:"type"`
}


type SiteSlotConfigNamesArgs struct {
	AppSettingNames       pulumi.StringArrayInput
	ConnectionStringNames pulumi.StringArrayInput
	Id                    pulumi.StringPtrInput
	Kind                  pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
	Type                  pulumi.StringPtrInput
}

func (SiteSlotConfigNamesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSlotConfigNamesArgs)(nil)).Elem()
}

type SiteSlotConfigNamesInput interface {
	pulumi.Input

	ToSiteSlotConfigNamesOutput() SiteSlotConfigNamesOutput
	ToSiteSlotConfigNamesOutputWithContext(ctx context.Context) SiteSlotConfigNamesOutput
}

func (*SiteSlotConfigNames) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteSlotConfigNames)(nil)).Elem()
}

func (i *SiteSlotConfigNames) ToSiteSlotConfigNamesOutput() SiteSlotConfigNamesOutput {
	return i.ToSiteSlotConfigNamesOutputWithContext(context.Background())
}

func (i *SiteSlotConfigNames) ToSiteSlotConfigNamesOutputWithContext(ctx context.Context) SiteSlotConfigNamesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteSlotConfigNamesOutput)
}

type SiteSlotConfigNamesOutput struct{ *pulumi.OutputState }

func (SiteSlotConfigNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteSlotConfigNames)(nil)).Elem()
}

func (o SiteSlotConfigNamesOutput) ToSiteSlotConfigNamesOutput() SiteSlotConfigNamesOutput {
	return o
}

func (o SiteSlotConfigNamesOutput) ToSiteSlotConfigNamesOutputWithContext(ctx context.Context) SiteSlotConfigNamesOutput {
	return o
}

func (o SiteSlotConfigNamesOutput) AppSettingNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteSlotConfigNames) pulumi.StringArrayOutput { return v.AppSettingNames }).(pulumi.StringArrayOutput)
}

func (o SiteSlotConfigNamesOutput) ConnectionStringNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteSlotConfigNames) pulumi.StringArrayOutput { return v.ConnectionStringNames }).(pulumi.StringArrayOutput)
}

func (o SiteSlotConfigNamesOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSlotConfigNames) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteSlotConfigNamesOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteSlotConfigNames) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteSlotConfigNamesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSlotConfigNames) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteSlotConfigNamesOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteSlotConfigNames) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteSlotConfigNamesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSlotConfigNames) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteSlotConfigNamesOutput{})
}
