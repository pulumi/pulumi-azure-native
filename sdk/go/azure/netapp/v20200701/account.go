


package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	ActiveDirectories ActiveDirectoryResponseArrayOutput `pulumi:"activeDirectories"`
	Location          pulumi.StringOutput                `pulumi:"location"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput             `pulumi:"tags"`
	Type              pulumi.StringOutput                `pulumi:"type"`
}


func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200701:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20170815:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190501:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190501:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190601:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190701:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190801:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190801:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191001:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191001:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191101:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200201:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200301:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200301:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200501:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200601:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200801:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200901:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201101:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201201:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210201:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210401:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210401preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210601:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:netapp/v20200701:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:netapp/v20200701:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accountState struct {
}

type AccountState struct {
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	AccountName       *string           `pulumi:"accountName"`
	ActiveDirectories []ActiveDirectory `pulumi:"activeDirectories"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type AccountArgs struct {
	AccountName       pulumi.StringPtrInput
	ActiveDirectories ActiveDirectoryArrayInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
