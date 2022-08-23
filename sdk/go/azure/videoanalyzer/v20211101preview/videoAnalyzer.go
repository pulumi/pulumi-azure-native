


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VideoAnalyzer struct {
	pulumi.CustomResourceState

	Encryption                 AccountEncryptionResponsePtrOutput           `pulumi:"encryption"`
	Endpoints                  EndpointResponseArrayOutput                  `pulumi:"endpoints"`
	Identity                   VideoAnalyzerIdentityResponsePtrOutput       `pulumi:"identity"`
	IotHubs                    IotHubResponseArrayOutput                    `pulumi:"iotHubs"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	NetworkAccessControl       NetworkAccessControlResponsePtrOutput        `pulumi:"networkAccessControl"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
	StorageAccounts            StorageAccountResponseArrayOutput            `pulumi:"storageAccounts"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
}


func NewVideoAnalyzer(ctx *pulumi.Context,
	name string, args *VideoAnalyzerArgs, opts ...pulumi.ResourceOption) (*VideoAnalyzer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccounts == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccounts'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer:VideoAnalyzer"),
		},
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20210501preview:VideoAnalyzer"),
		},
	})
	opts = append(opts, aliases)
	var resource VideoAnalyzer
	err := ctx.RegisterResource("azure-native:videoanalyzer/v20211101preview:VideoAnalyzer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVideoAnalyzer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VideoAnalyzerState, opts ...pulumi.ResourceOption) (*VideoAnalyzer, error) {
	var resource VideoAnalyzer
	err := ctx.ReadResource("azure-native:videoanalyzer/v20211101preview:VideoAnalyzer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type videoAnalyzerState struct {
}

type VideoAnalyzerState struct {
}

func (VideoAnalyzerState) ElementType() reflect.Type {
	return reflect.TypeOf((*videoAnalyzerState)(nil)).Elem()
}

type videoAnalyzerArgs struct {
	AccountName          *string                `pulumi:"accountName"`
	Encryption           *AccountEncryption     `pulumi:"encryption"`
	Identity             *VideoAnalyzerIdentity `pulumi:"identity"`
	IotHubs              []IotHub               `pulumi:"iotHubs"`
	Location             *string                `pulumi:"location"`
	NetworkAccessControl *NetworkAccessControl  `pulumi:"networkAccessControl"`
	PublicNetworkAccess  *string                `pulumi:"publicNetworkAccess"`
	ResourceGroupName    string                 `pulumi:"resourceGroupName"`
	StorageAccounts      []StorageAccount       `pulumi:"storageAccounts"`
	Tags                 map[string]string      `pulumi:"tags"`
}


type VideoAnalyzerArgs struct {
	AccountName          pulumi.StringPtrInput
	Encryption           AccountEncryptionPtrInput
	Identity             VideoAnalyzerIdentityPtrInput
	IotHubs              IotHubArrayInput
	Location             pulumi.StringPtrInput
	NetworkAccessControl NetworkAccessControlPtrInput
	PublicNetworkAccess  pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	StorageAccounts      StorageAccountArrayInput
	Tags                 pulumi.StringMapInput
}

func (VideoAnalyzerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*videoAnalyzerArgs)(nil)).Elem()
}

type VideoAnalyzerInput interface {
	pulumi.Input

	ToVideoAnalyzerOutput() VideoAnalyzerOutput
	ToVideoAnalyzerOutputWithContext(ctx context.Context) VideoAnalyzerOutput
}

func (*VideoAnalyzer) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoAnalyzer)(nil)).Elem()
}

func (i *VideoAnalyzer) ToVideoAnalyzerOutput() VideoAnalyzerOutput {
	return i.ToVideoAnalyzerOutputWithContext(context.Background())
}

func (i *VideoAnalyzer) ToVideoAnalyzerOutputWithContext(ctx context.Context) VideoAnalyzerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerOutput)
}

type VideoAnalyzerOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoAnalyzer)(nil)).Elem()
}

func (o VideoAnalyzerOutput) ToVideoAnalyzerOutput() VideoAnalyzerOutput {
	return o
}

func (o VideoAnalyzerOutput) ToVideoAnalyzerOutputWithContext(ctx context.Context) VideoAnalyzerOutput {
	return o
}

func (o VideoAnalyzerOutput) Encryption() AccountEncryptionResponsePtrOutput {
	return o.ApplyT(func(v *VideoAnalyzer) AccountEncryptionResponsePtrOutput { return v.Encryption }).(AccountEncryptionResponsePtrOutput)
}

func (o VideoAnalyzerOutput) Endpoints() EndpointResponseArrayOutput {
	return o.ApplyT(func(v *VideoAnalyzer) EndpointResponseArrayOutput { return v.Endpoints }).(EndpointResponseArrayOutput)
}

func (o VideoAnalyzerOutput) Identity() VideoAnalyzerIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VideoAnalyzer) VideoAnalyzerIdentityResponsePtrOutput { return v.Identity }).(VideoAnalyzerIdentityResponsePtrOutput)
}

func (o VideoAnalyzerOutput) IotHubs() IotHubResponseArrayOutput {
	return o.ApplyT(func(v *VideoAnalyzer) IotHubResponseArrayOutput { return v.IotHubs }).(IotHubResponseArrayOutput)
}

func (o VideoAnalyzerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VideoAnalyzer) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VideoAnalyzerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VideoAnalyzer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VideoAnalyzerOutput) NetworkAccessControl() NetworkAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *VideoAnalyzer) NetworkAccessControlResponsePtrOutput { return v.NetworkAccessControl }).(NetworkAccessControlResponsePtrOutput)
}

func (o VideoAnalyzerOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *VideoAnalyzer) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o VideoAnalyzerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VideoAnalyzer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VideoAnalyzerOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoAnalyzer) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerOutput) StorageAccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v *VideoAnalyzer) StorageAccountResponseArrayOutput { return v.StorageAccounts }).(StorageAccountResponseArrayOutput)
}

func (o VideoAnalyzerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VideoAnalyzer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VideoAnalyzerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VideoAnalyzer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VideoAnalyzerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VideoAnalyzer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VideoAnalyzerOutput{})
}
