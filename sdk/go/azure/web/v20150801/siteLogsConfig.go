


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteLogsConfig struct {
	pulumi.CustomResourceState

	ApplicationLogs       ApplicationLogsConfigResponsePtrOutput `pulumi:"applicationLogs"`
	DetailedErrorMessages EnabledConfigResponsePtrOutput         `pulumi:"detailedErrorMessages"`
	FailedRequestsTracing EnabledConfigResponsePtrOutput         `pulumi:"failedRequestsTracing"`
	HttpLogs              HttpLogsConfigResponsePtrOutput        `pulumi:"httpLogs"`
	Kind                  pulumi.StringPtrOutput                 `pulumi:"kind"`
	Location              pulumi.StringOutput                    `pulumi:"location"`
	Name                  pulumi.StringPtrOutput                 `pulumi:"name"`
	Tags                  pulumi.StringMapOutput                 `pulumi:"tags"`
	Type                  pulumi.StringPtrOutput                 `pulumi:"type"`
}


func NewSiteLogsConfig(ctx *pulumi.Context,
	name string, args *SiteLogsConfigArgs, opts ...pulumi.ResourceOption) (*SiteLogsConfig, error) {
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
			Type: pulumi.String("azure-nextgen:web/v20150801:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteLogsConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:SiteLogsConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteLogsConfig
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteLogsConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteLogsConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteLogsConfigState, opts ...pulumi.ResourceOption) (*SiteLogsConfig, error) {
	var resource SiteLogsConfig
	err := ctx.ReadResource("azure-native:web/v20150801:SiteLogsConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteLogsConfigState struct {
}

type SiteLogsConfigState struct {
}

func (SiteLogsConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteLogsConfigState)(nil)).Elem()
}

type siteLogsConfigArgs struct {
	ApplicationLogs       *ApplicationLogsConfig `pulumi:"applicationLogs"`
	DetailedErrorMessages *EnabledConfig         `pulumi:"detailedErrorMessages"`
	FailedRequestsTracing *EnabledConfig         `pulumi:"failedRequestsTracing"`
	HttpLogs              *HttpLogsConfig        `pulumi:"httpLogs"`
	Id                    *string                `pulumi:"id"`
	Kind                  *string                `pulumi:"kind"`
	Location              *string                `pulumi:"location"`
	Name                  string                 `pulumi:"name"`
	ResourceGroupName     string                 `pulumi:"resourceGroupName"`
	Tags                  map[string]string      `pulumi:"tags"`
	Type                  *string                `pulumi:"type"`
}


type SiteLogsConfigArgs struct {
	ApplicationLogs       ApplicationLogsConfigPtrInput
	DetailedErrorMessages EnabledConfigPtrInput
	FailedRequestsTracing EnabledConfigPtrInput
	HttpLogs              HttpLogsConfigPtrInput
	Id                    pulumi.StringPtrInput
	Kind                  pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
	Type                  pulumi.StringPtrInput
}

func (SiteLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteLogsConfigArgs)(nil)).Elem()
}

type SiteLogsConfigInput interface {
	pulumi.Input

	ToSiteLogsConfigOutput() SiteLogsConfigOutput
	ToSiteLogsConfigOutputWithContext(ctx context.Context) SiteLogsConfigOutput
}

func (*SiteLogsConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLogsConfig)(nil))
}

func (i *SiteLogsConfig) ToSiteLogsConfigOutput() SiteLogsConfigOutput {
	return i.ToSiteLogsConfigOutputWithContext(context.Background())
}

func (i *SiteLogsConfig) ToSiteLogsConfigOutputWithContext(ctx context.Context) SiteLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLogsConfigOutput)
}

type SiteLogsConfigOutput struct{ *pulumi.OutputState }

func (SiteLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLogsConfig)(nil))
}

func (o SiteLogsConfigOutput) ToSiteLogsConfigOutput() SiteLogsConfigOutput {
	return o
}

func (o SiteLogsConfigOutput) ToSiteLogsConfigOutputWithContext(ctx context.Context) SiteLogsConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteLogsConfigOutput{})
}
