


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Datastore struct {
	pulumi.CustomResourceState

	CustomResourceName pulumi.StringOutput               `pulumi:"customResourceName"`
	ExtendedLocation   ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	InventoryItemId    pulumi.StringPtrOutput            `pulumi:"inventoryItemId"`
	Kind               pulumi.StringPtrOutput            `pulumi:"kind"`
	Location           pulumi.StringOutput               `pulumi:"location"`
	MoName             pulumi.StringOutput               `pulumi:"moName"`
	MoRefId            pulumi.StringPtrOutput            `pulumi:"moRefId"`
	Name               pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput               `pulumi:"provisioningState"`
	Statuses           ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	SystemData         SystemDataResponseOutput          `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput            `pulumi:"tags"`
	Type               pulumi.StringOutput               `pulumi:"type"`
	Uuid               pulumi.StringOutput               `pulumi:"uuid"`
	VCenterId          pulumi.StringPtrOutput            `pulumi:"vCenterId"`
}


func NewDatastore(ctx *pulumi.Context,
	name string, args *DatastoreArgs, opts ...pulumi.ResourceOption) (*Datastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:Datastore"),
		},
	})
	opts = append(opts, aliases)
	var resource Datastore
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:Datastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatastoreState, opts ...pulumi.ResourceOption) (*Datastore, error) {
	var resource Datastore
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:Datastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type datastoreState struct {
}

type DatastoreState struct {
}

func (DatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreState)(nil)).Elem()
}

type datastoreArgs struct {
	DatastoreName     *string           `pulumi:"datastoreName"`
	ExtendedLocation  *ExtendedLocation `pulumi:"extendedLocation"`
	InventoryItemId   *string           `pulumi:"inventoryItemId"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	MoRefId           *string           `pulumi:"moRefId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VCenterId         *string           `pulumi:"vCenterId"`
}


type DatastoreArgs struct {
	DatastoreName     pulumi.StringPtrInput
	ExtendedLocation  ExtendedLocationPtrInput
	InventoryItemId   pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	MoRefId           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VCenterId         pulumi.StringPtrInput
}

func (DatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreArgs)(nil)).Elem()
}

type DatastoreInput interface {
	pulumi.Input

	ToDatastoreOutput() DatastoreOutput
	ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput
}

func (*Datastore) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (i *Datastore) ToDatastoreOutput() DatastoreOutput {
	return i.ToDatastoreOutputWithContext(context.Background())
}

func (i *Datastore) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreOutput)
}

type DatastoreOutput struct{ *pulumi.OutputState }

func (DatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (o DatastoreOutput) ToDatastoreOutput() DatastoreOutput {
	return o
}

func (o DatastoreOutput) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return o
}

func (o DatastoreOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o DatastoreOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Datastore) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o DatastoreOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o DatastoreOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o DatastoreOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DatastoreOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

func (o DatastoreOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o DatastoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatastoreOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatastoreOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *Datastore) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o DatastoreOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Datastore) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DatastoreOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatastoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DatastoreOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o DatastoreOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatastoreOutput{})
}
