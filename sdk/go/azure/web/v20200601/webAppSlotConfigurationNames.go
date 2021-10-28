


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSlotConfigurationNames struct {
	pulumi.CustomResourceState

	AppSettingNames         pulumi.StringArrayOutput `pulumi:"appSettingNames"`
	AzureStorageConfigNames pulumi.StringArrayOutput `pulumi:"azureStorageConfigNames"`
	ConnectionStringNames   pulumi.StringArrayOutput `pulumi:"connectionStringNames"`
	Kind                    pulumi.StringPtrOutput   `pulumi:"kind"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
}


func NewWebAppSlotConfigurationNames(ctx *pulumi.Context,
	name string, args *WebAppSlotConfigurationNamesArgs, opts ...pulumi.ResourceOption) (*WebAppSlotConfigurationNames, error) {
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
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:WebAppSlotConfigurationNames"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSlotConfigurationNames
	err := ctx.RegisterResource("azure-native:web/v20200601:WebAppSlotConfigurationNames", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSlotConfigurationNames(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSlotConfigurationNamesState, opts ...pulumi.ResourceOption) (*WebAppSlotConfigurationNames, error) {
	var resource WebAppSlotConfigurationNames
	err := ctx.ReadResource("azure-native:web/v20200601:WebAppSlotConfigurationNames", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppSlotConfigurationNamesState struct {
}

type WebAppSlotConfigurationNamesState struct {
}

func (WebAppSlotConfigurationNamesState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotConfigurationNamesState)(nil)).Elem()
}

type webAppSlotConfigurationNamesArgs struct {
	AppSettingNames         []string `pulumi:"appSettingNames"`
	AzureStorageConfigNames []string `pulumi:"azureStorageConfigNames"`
	ConnectionStringNames   []string `pulumi:"connectionStringNames"`
	Kind                    *string  `pulumi:"kind"`
	Name                    string   `pulumi:"name"`
	ResourceGroupName       string   `pulumi:"resourceGroupName"`
}


type WebAppSlotConfigurationNamesArgs struct {
	AppSettingNames         pulumi.StringArrayInput
	AzureStorageConfigNames pulumi.StringArrayInput
	ConnectionStringNames   pulumi.StringArrayInput
	Kind                    pulumi.StringPtrInput
	Name                    pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
}

func (WebAppSlotConfigurationNamesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotConfigurationNamesArgs)(nil)).Elem()
}

type WebAppSlotConfigurationNamesInput interface {
	pulumi.Input

	ToWebAppSlotConfigurationNamesOutput() WebAppSlotConfigurationNamesOutput
	ToWebAppSlotConfigurationNamesOutputWithContext(ctx context.Context) WebAppSlotConfigurationNamesOutput
}

func (*WebAppSlotConfigurationNames) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSlotConfigurationNames)(nil))
}

func (i *WebAppSlotConfigurationNames) ToWebAppSlotConfigurationNamesOutput() WebAppSlotConfigurationNamesOutput {
	return i.ToWebAppSlotConfigurationNamesOutputWithContext(context.Background())
}

func (i *WebAppSlotConfigurationNames) ToWebAppSlotConfigurationNamesOutputWithContext(ctx context.Context) WebAppSlotConfigurationNamesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSlotConfigurationNamesOutput)
}

type WebAppSlotConfigurationNamesOutput struct{ *pulumi.OutputState }

func (WebAppSlotConfigurationNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSlotConfigurationNames)(nil))
}

func (o WebAppSlotConfigurationNamesOutput) ToWebAppSlotConfigurationNamesOutput() WebAppSlotConfigurationNamesOutput {
	return o
}

func (o WebAppSlotConfigurationNamesOutput) ToWebAppSlotConfigurationNamesOutputWithContext(ctx context.Context) WebAppSlotConfigurationNamesOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebAppSlotConfigurationNamesInput)(nil)).Elem(), &WebAppSlotConfigurationNames{})
	pulumi.RegisterOutputType(WebAppSlotConfigurationNamesOutput{})
}
