


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MediaService struct {
	pulumi.CustomResourceState

	Encryption            AccountEncryptionResponsePtrOutput    `pulumi:"encryption"`
	Identity              MediaServiceIdentityResponsePtrOutput `pulumi:"identity"`
	KeyDelivery           KeyDeliveryResponsePtrOutput          `pulumi:"keyDelivery"`
	Location              pulumi.StringOutput                   `pulumi:"location"`
	MediaServiceId        pulumi.StringOutput                   `pulumi:"mediaServiceId"`
	Name                  pulumi.StringOutput                   `pulumi:"name"`
	PublicNetworkAccess   pulumi.StringPtrOutput                `pulumi:"publicNetworkAccess"`
	StorageAccounts       StorageAccountResponseArrayOutput     `pulumi:"storageAccounts"`
	StorageAuthentication pulumi.StringPtrOutput                `pulumi:"storageAuthentication"`
	SystemData            SystemDataResponseOutput              `pulumi:"systemData"`
	Tags                  pulumi.StringMapOutput                `pulumi:"tags"`
	Type                  pulumi.StringOutput                   `pulumi:"type"`
}


func NewMediaService(ctx *pulumi.Context,
	name string, args *MediaServiceArgs, opts ...pulumi.ResourceOption) (*MediaService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20151001:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210501:MediaService"),
		},
	})
	opts = append(opts, aliases)
	var resource MediaService
	err := ctx.RegisterResource("azure-native:media/v20210601:MediaService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMediaService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaServiceState, opts ...pulumi.ResourceOption) (*MediaService, error) {
	var resource MediaService
	err := ctx.ReadResource("azure-native:media/v20210601:MediaService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mediaServiceState struct {
}

type MediaServiceState struct {
}

func (MediaServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceState)(nil)).Elem()
}

type mediaServiceArgs struct {
	AccountName           *string               `pulumi:"accountName"`
	Encryption            *AccountEncryption    `pulumi:"encryption"`
	Identity              *MediaServiceIdentity `pulumi:"identity"`
	KeyDelivery           *KeyDelivery          `pulumi:"keyDelivery"`
	Location              *string               `pulumi:"location"`
	PublicNetworkAccess   *string               `pulumi:"publicNetworkAccess"`
	ResourceGroupName     string                `pulumi:"resourceGroupName"`
	StorageAccounts       []StorageAccount      `pulumi:"storageAccounts"`
	StorageAuthentication *string               `pulumi:"storageAuthentication"`
	Tags                  map[string]string     `pulumi:"tags"`
}


type MediaServiceArgs struct {
	AccountName           pulumi.StringPtrInput
	Encryption            AccountEncryptionPtrInput
	Identity              MediaServiceIdentityPtrInput
	KeyDelivery           KeyDeliveryPtrInput
	Location              pulumi.StringPtrInput
	PublicNetworkAccess   pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	StorageAccounts       StorageAccountArrayInput
	StorageAuthentication pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
}

func (MediaServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceArgs)(nil)).Elem()
}

type MediaServiceInput interface {
	pulumi.Input

	ToMediaServiceOutput() MediaServiceOutput
	ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput
}

func (*MediaService) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaService)(nil))
}

func (i *MediaService) ToMediaServiceOutput() MediaServiceOutput {
	return i.ToMediaServiceOutputWithContext(context.Background())
}

func (i *MediaService) ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceOutput)
}

type MediaServiceOutput struct{ *pulumi.OutputState }

func (MediaServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaService)(nil))
}

func (o MediaServiceOutput) ToMediaServiceOutput() MediaServiceOutput {
	return o
}

func (o MediaServiceOutput) ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MediaServiceOutput{})
}
