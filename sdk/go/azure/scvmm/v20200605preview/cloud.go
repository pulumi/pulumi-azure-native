


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cloud struct {
	pulumi.CustomResourceState

	CloudCapacity      CloudCapacityResponseOutput         `pulumi:"cloudCapacity"`
	CloudName          pulumi.StringOutput                 `pulumi:"cloudName"`
	ExtendedLocation   ExtendedLocationResponseOutput      `pulumi:"extendedLocation"`
	InventoryItemId    pulumi.StringPtrOutput              `pulumi:"inventoryItemId"`
	Location           pulumi.StringOutput                 `pulumi:"location"`
	Name               pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput                 `pulumi:"provisioningState"`
	StorageQoSPolicies StorageQoSPolicyResponseArrayOutput `pulumi:"storageQoSPolicies"`
	SystemData         SystemDataResponseOutput            `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput              `pulumi:"tags"`
	Type               pulumi.StringOutput                 `pulumi:"type"`
	Uuid               pulumi.StringPtrOutput              `pulumi:"uuid"`
	VmmServerId        pulumi.StringPtrOutput              `pulumi:"vmmServerId"`
}


func NewCloud(ctx *pulumi.Context,
	name string, args *CloudArgs, opts ...pulumi.ResourceOption) (*Cloud, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm:Cloud"),
		},
	})
	opts = append(opts, aliases)
	var resource Cloud
	err := ctx.RegisterResource("azure-native:scvmm/v20200605preview:Cloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudState, opts ...pulumi.ResourceOption) (*Cloud, error) {
	var resource Cloud
	err := ctx.ReadResource("azure-native:scvmm/v20200605preview:Cloud", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cloudState struct {
}

type CloudState struct {
}

func (CloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudState)(nil)).Elem()
}

type cloudArgs struct {
	CloudName         *string           `pulumi:"cloudName"`
	ExtendedLocation  ExtendedLocation  `pulumi:"extendedLocation"`
	InventoryItemId   *string           `pulumi:"inventoryItemId"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	Uuid              *string           `pulumi:"uuid"`
	VmmServerId       *string           `pulumi:"vmmServerId"`
}


type CloudArgs struct {
	CloudName         pulumi.StringPtrInput
	ExtendedLocation  ExtendedLocationInput
	InventoryItemId   pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Uuid              pulumi.StringPtrInput
	VmmServerId       pulumi.StringPtrInput
}

func (CloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudArgs)(nil)).Elem()
}

type CloudInput interface {
	pulumi.Input

	ToCloudOutput() CloudOutput
	ToCloudOutputWithContext(ctx context.Context) CloudOutput
}

func (*Cloud) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloud)(nil)).Elem()
}

func (i *Cloud) ToCloudOutput() CloudOutput {
	return i.ToCloudOutputWithContext(context.Background())
}

func (i *Cloud) ToCloudOutputWithContext(ctx context.Context) CloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudOutput)
}

type CloudOutput struct{ *pulumi.OutputState }

func (CloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloud)(nil)).Elem()
}

func (o CloudOutput) ToCloudOutput() CloudOutput {
	return o
}

func (o CloudOutput) ToCloudOutputWithContext(ctx context.Context) CloudOutput {
	return o
}

func (o CloudOutput) CloudCapacity() CloudCapacityResponseOutput {
	return o.ApplyT(func(v *Cloud) CloudCapacityResponseOutput { return v.CloudCapacity }).(CloudCapacityResponseOutput)
}

func (o CloudOutput) CloudName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.CloudName }).(pulumi.StringOutput)
}

func (o CloudOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Cloud) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o CloudOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o CloudOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CloudOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CloudOutput) StorageQoSPolicies() StorageQoSPolicyResponseArrayOutput {
	return o.ApplyT(func(v *Cloud) StorageQoSPolicyResponseArrayOutput { return v.StorageQoSPolicies }).(StorageQoSPolicyResponseArrayOutput)
}

func (o CloudOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cloud) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CloudOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CloudOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CloudOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringPtrOutput { return v.Uuid }).(pulumi.StringPtrOutput)
}

func (o CloudOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudOutput{})
}
