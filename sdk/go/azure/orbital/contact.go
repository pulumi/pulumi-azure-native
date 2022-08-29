


package orbital

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Contact struct {
	pulumi.CustomResourceState

	ContactProfile          ResourceReferenceResponseOutput `pulumi:"contactProfile"`
	EndAzimuthDegrees       pulumi.Float64Output            `pulumi:"endAzimuthDegrees"`
	EndElevationDegrees     pulumi.Float64Output            `pulumi:"endElevationDegrees"`
	ErrorMessage            pulumi.StringOutput             `pulumi:"errorMessage"`
	Etag                    pulumi.StringOutput             `pulumi:"etag"`
	GroundStationName       pulumi.StringOutput             `pulumi:"groundStationName"`
	MaximumElevationDegrees pulumi.Float64Output            `pulumi:"maximumElevationDegrees"`
	Name                    pulumi.StringOutput             `pulumi:"name"`
	ReservationEndTime      pulumi.StringOutput             `pulumi:"reservationEndTime"`
	ReservationStartTime    pulumi.StringOutput             `pulumi:"reservationStartTime"`
	RxEndTime               pulumi.StringOutput             `pulumi:"rxEndTime"`
	RxStartTime             pulumi.StringOutput             `pulumi:"rxStartTime"`
	StartAzimuthDegrees     pulumi.Float64Output            `pulumi:"startAzimuthDegrees"`
	StartElevationDegrees   pulumi.Float64Output            `pulumi:"startElevationDegrees"`
	Status                  pulumi.StringOutput             `pulumi:"status"`
	SystemData              SystemDataResponseOutput        `pulumi:"systemData"`
	TxEndTime               pulumi.StringOutput             `pulumi:"txEndTime"`
	TxStartTime             pulumi.StringOutput             `pulumi:"txStartTime"`
	Type                    pulumi.StringOutput             `pulumi:"type"`
}


func NewContact(ctx *pulumi.Context,
	name string, args *ContactArgs, opts ...pulumi.ResourceOption) (*Contact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactProfile == nil {
		return nil, errors.New("invalid value for required argument 'ContactProfile'")
	}
	if args.GroundStationName == nil {
		return nil, errors.New("invalid value for required argument 'GroundStationName'")
	}
	if args.ReservationEndTime == nil {
		return nil, errors.New("invalid value for required argument 'ReservationEndTime'")
	}
	if args.ReservationStartTime == nil {
		return nil, errors.New("invalid value for required argument 'ReservationStartTime'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SpacecraftName == nil {
		return nil, errors.New("invalid value for required argument 'SpacecraftName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:orbital/v20210404preview:Contact"),
		},
		{
			Type: pulumi.String("azure-native:orbital/v20220301:Contact"),
		},
	})
	opts = append(opts, aliases)
	var resource Contact
	err := ctx.RegisterResource("azure-native:orbital:Contact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactState, opts ...pulumi.ResourceOption) (*Contact, error) {
	var resource Contact
	err := ctx.ReadResource("azure-native:orbital:Contact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type contactState struct {
}

type ContactState struct {
}

func (ContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactState)(nil)).Elem()
}

type contactArgs struct {
	ContactName          *string           `pulumi:"contactName"`
	ContactProfile       ResourceReference `pulumi:"contactProfile"`
	GroundStationName    string            `pulumi:"groundStationName"`
	ReservationEndTime   string            `pulumi:"reservationEndTime"`
	ReservationStartTime string            `pulumi:"reservationStartTime"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	SpacecraftName       string            `pulumi:"spacecraftName"`
}


type ContactArgs struct {
	ContactName          pulumi.StringPtrInput
	ContactProfile       ResourceReferenceInput
	GroundStationName    pulumi.StringInput
	ReservationEndTime   pulumi.StringInput
	ReservationStartTime pulumi.StringInput
	ResourceGroupName    pulumi.StringInput
	SpacecraftName       pulumi.StringInput
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactArgs)(nil)).Elem()
}

type ContactInput interface {
	pulumi.Input

	ToContactOutput() ContactOutput
	ToContactOutputWithContext(ctx context.Context) ContactOutput
}

func (*Contact) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (i *Contact) ToContactOutput() ContactOutput {
	return i.ToContactOutputWithContext(context.Background())
}

func (i *Contact) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput)
}

type ContactOutput struct{ *pulumi.OutputState }

func (ContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (o ContactOutput) ToContactOutput() ContactOutput {
	return o
}

func (o ContactOutput) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return o
}

func (o ContactOutput) ContactProfile() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *Contact) ResourceReferenceResponseOutput { return v.ContactProfile }).(ResourceReferenceResponseOutput)
}

func (o ContactOutput) EndAzimuthDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v *Contact) pulumi.Float64Output { return v.EndAzimuthDegrees }).(pulumi.Float64Output)
}

func (o ContactOutput) EndElevationDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v *Contact) pulumi.Float64Output { return v.EndElevationDegrees }).(pulumi.Float64Output)
}

func (o ContactOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o ContactOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ContactOutput) GroundStationName() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.GroundStationName }).(pulumi.StringOutput)
}

func (o ContactOutput) MaximumElevationDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v *Contact) pulumi.Float64Output { return v.MaximumElevationDegrees }).(pulumi.Float64Output)
}

func (o ContactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContactOutput) ReservationEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.ReservationEndTime }).(pulumi.StringOutput)
}

func (o ContactOutput) ReservationStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.ReservationStartTime }).(pulumi.StringOutput)
}

func (o ContactOutput) RxEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.RxEndTime }).(pulumi.StringOutput)
}

func (o ContactOutput) RxStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.RxStartTime }).(pulumi.StringOutput)
}

func (o ContactOutput) StartAzimuthDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v *Contact) pulumi.Float64Output { return v.StartAzimuthDegrees }).(pulumi.Float64Output)
}

func (o ContactOutput) StartElevationDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v *Contact) pulumi.Float64Output { return v.StartElevationDegrees }).(pulumi.Float64Output)
}

func (o ContactOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ContactOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Contact) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ContactOutput) TxEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.TxEndTime }).(pulumi.StringOutput)
}

func (o ContactOutput) TxStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.TxStartTime }).(pulumi.StringOutput)
}

func (o ContactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContactOutput{})
}
