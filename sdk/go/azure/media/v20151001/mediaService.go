


package v20151001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MediaService struct {
	pulumi.CustomResourceState

	ApiEndpoints    ApiEndpointResponseArrayOutput    `pulumi:"apiEndpoints"`
	Location        pulumi.StringPtrOutput            `pulumi:"location"`
	Name            pulumi.StringOutput               `pulumi:"name"`
	StorageAccounts StorageAccountResponseArrayOutput `pulumi:"storageAccounts"`
	Tags            pulumi.StringMapOutput            `pulumi:"tags"`
	Type            pulumi.StringOutput               `pulumi:"type"`
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
		{
			Type: pulumi.String("azure-native:media/v20210601:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:MediaService"),
		},
	})
	opts = append(opts, aliases)
	var resource MediaService
	err := ctx.RegisterResource("azure-native:media/v20151001:MediaService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMediaService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaServiceState, opts ...pulumi.ResourceOption) (*MediaService, error) {
	var resource MediaService
	err := ctx.ReadResource("azure-native:media/v20151001:MediaService", name, id, state, &resource, opts...)
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
	Location          *string           `pulumi:"location"`
	MediaServiceName  *string           `pulumi:"mediaServiceName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	StorageAccounts   []StorageAccount  `pulumi:"storageAccounts"`
	Tags              map[string]string `pulumi:"tags"`
}


type MediaServiceArgs struct {
	Location          pulumi.StringPtrInput
	MediaServiceName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	StorageAccounts   StorageAccountArrayInput
	Tags              pulumi.StringMapInput
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
	return reflect.TypeOf((**MediaService)(nil)).Elem()
}

func (i *MediaService) ToMediaServiceOutput() MediaServiceOutput {
	return i.ToMediaServiceOutputWithContext(context.Background())
}

func (i *MediaService) ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceOutput)
}

type MediaServiceOutput struct{ *pulumi.OutputState }

func (MediaServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaService)(nil)).Elem()
}

func (o MediaServiceOutput) ToMediaServiceOutput() MediaServiceOutput {
	return o
}

func (o MediaServiceOutput) ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput {
	return o
}

func (o MediaServiceOutput) ApiEndpoints() ApiEndpointResponseArrayOutput {
	return o.ApplyT(func(v *MediaService) ApiEndpointResponseArrayOutput { return v.ApiEndpoints }).(ApiEndpointResponseArrayOutput)
}

func (o MediaServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o MediaServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MediaServiceOutput) StorageAccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v *MediaService) StorageAccountResponseArrayOutput { return v.StorageAccounts }).(StorageAccountResponseArrayOutput)
}

func (o MediaServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MediaServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MediaServiceOutput{})
}
