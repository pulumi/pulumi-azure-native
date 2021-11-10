


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppAzureStorageAccounts struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput                 `pulumi:"kind"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties AzureStorageInfoValueResponseMapOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput               `pulumi:"systemData"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewWebAppAzureStorageAccounts(ctx *pulumi.Context,
	name string, args *WebAppAzureStorageAccountsArgs, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccounts, error) {
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
			Type: pulumi.String("azure-native:web:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppAzureStorageAccounts"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppAzureStorageAccounts
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppAzureStorageAccounts", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppAzureStorageAccounts(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAzureStorageAccountsState, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccounts, error) {
	var resource WebAppAzureStorageAccounts
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppAzureStorageAccounts", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppAzureStorageAccountsState struct {
}

type WebAppAzureStorageAccountsState struct {
}

func (WebAppAzureStorageAccountsState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsState)(nil)).Elem()
}

type webAppAzureStorageAccountsArgs struct {
	Kind              *string                          `pulumi:"kind"`
	Name              string                           `pulumi:"name"`
	Properties        map[string]AzureStorageInfoValue `pulumi:"properties"`
	ResourceGroupName string                           `pulumi:"resourceGroupName"`
}


type WebAppAzureStorageAccountsArgs struct {
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        AzureStorageInfoValueMapInput
	ResourceGroupName pulumi.StringInput
}

func (WebAppAzureStorageAccountsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsArgs)(nil)).Elem()
}

type WebAppAzureStorageAccountsInput interface {
	pulumi.Input

	ToWebAppAzureStorageAccountsOutput() WebAppAzureStorageAccountsOutput
	ToWebAppAzureStorageAccountsOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsOutput
}

func (*WebAppAzureStorageAccounts) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppAzureStorageAccounts)(nil))
}

func (i *WebAppAzureStorageAccounts) ToWebAppAzureStorageAccountsOutput() WebAppAzureStorageAccountsOutput {
	return i.ToWebAppAzureStorageAccountsOutputWithContext(context.Background())
}

func (i *WebAppAzureStorageAccounts) ToWebAppAzureStorageAccountsOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAzureStorageAccountsOutput)
}

type WebAppAzureStorageAccountsOutput struct{ *pulumi.OutputState }

func (WebAppAzureStorageAccountsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppAzureStorageAccounts)(nil))
}

func (o WebAppAzureStorageAccountsOutput) ToWebAppAzureStorageAccountsOutput() WebAppAzureStorageAccountsOutput {
	return o
}

func (o WebAppAzureStorageAccountsOutput) ToWebAppAzureStorageAccountsOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppAzureStorageAccountsOutput{})
}
