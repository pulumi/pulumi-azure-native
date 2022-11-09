


package v20210404preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContactProfile struct {
	pulumi.CustomResourceState

	AutoTrackingConfiguration    pulumi.StringPtrOutput                `pulumi:"autoTrackingConfiguration"`
	Etag                         pulumi.StringOutput                   `pulumi:"etag"`
	EventHubUri                  pulumi.StringPtrOutput                `pulumi:"eventHubUri"`
	Links                        ContactProfileLinkResponseArrayOutput `pulumi:"links"`
	Location                     pulumi.StringOutput                   `pulumi:"location"`
	MinimumElevationDegrees      pulumi.Float64PtrOutput               `pulumi:"minimumElevationDegrees"`
	MinimumViableContactDuration pulumi.StringPtrOutput                `pulumi:"minimumViableContactDuration"`
	Name                         pulumi.StringOutput                   `pulumi:"name"`
	SystemData                   SystemDataResponseOutput              `pulumi:"systemData"`
	Tags                         pulumi.StringMapOutput                `pulumi:"tags"`
	Type                         pulumi.StringOutput                   `pulumi:"type"`
}


func NewContactProfile(ctx *pulumi.Context,
	name string, args *ContactProfileArgs, opts ...pulumi.ResourceOption) (*ContactProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Links == nil {
		return nil, errors.New("invalid value for required argument 'Links'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:orbital:ContactProfile"),
		},
		{
			Type: pulumi.String("azure-native:orbital/v20220301:ContactProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource ContactProfile
	err := ctx.RegisterResource("azure-native:orbital/v20210404preview:ContactProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContactProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactProfileState, opts ...pulumi.ResourceOption) (*ContactProfile, error) {
	var resource ContactProfile
	err := ctx.ReadResource("azure-native:orbital/v20210404preview:ContactProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type contactProfileState struct {
}

type ContactProfileState struct {
}

func (ContactProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactProfileState)(nil)).Elem()
}

type contactProfileArgs struct {
	AutoTrackingConfiguration    *AutoTrackingConfiguration `pulumi:"autoTrackingConfiguration"`
	ContactProfileName           *string                    `pulumi:"contactProfileName"`
	EventHubUri                  *string                    `pulumi:"eventHubUri"`
	Links                        []ContactProfileLink       `pulumi:"links"`
	Location                     *string                    `pulumi:"location"`
	MinimumElevationDegrees      *float64                   `pulumi:"minimumElevationDegrees"`
	MinimumViableContactDuration *string                    `pulumi:"minimumViableContactDuration"`
	ResourceGroupName            string                     `pulumi:"resourceGroupName"`
	Tags                         map[string]string          `pulumi:"tags"`
}


type ContactProfileArgs struct {
	AutoTrackingConfiguration    AutoTrackingConfigurationPtrInput
	ContactProfileName           pulumi.StringPtrInput
	EventHubUri                  pulumi.StringPtrInput
	Links                        ContactProfileLinkArrayInput
	Location                     pulumi.StringPtrInput
	MinimumElevationDegrees      pulumi.Float64PtrInput
	MinimumViableContactDuration pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
}

func (ContactProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactProfileArgs)(nil)).Elem()
}

type ContactProfileInput interface {
	pulumi.Input

	ToContactProfileOutput() ContactProfileOutput
	ToContactProfileOutputWithContext(ctx context.Context) ContactProfileOutput
}

func (*ContactProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactProfile)(nil)).Elem()
}

func (i *ContactProfile) ToContactProfileOutput() ContactProfileOutput {
	return i.ToContactProfileOutputWithContext(context.Background())
}

func (i *ContactProfile) ToContactProfileOutputWithContext(ctx context.Context) ContactProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileOutput)
}

type ContactProfileOutput struct{ *pulumi.OutputState }

func (ContactProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactProfile)(nil)).Elem()
}

func (o ContactProfileOutput) ToContactProfileOutput() ContactProfileOutput {
	return o
}

func (o ContactProfileOutput) ToContactProfileOutputWithContext(ctx context.Context) ContactProfileOutput {
	return o
}

func (o ContactProfileOutput) AutoTrackingConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactProfile) pulumi.StringPtrOutput { return v.AutoTrackingConfiguration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactProfile) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ContactProfileOutput) EventHubUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactProfile) pulumi.StringPtrOutput { return v.EventHubUri }).(pulumi.StringPtrOutput)
}

func (o ContactProfileOutput) Links() ContactProfileLinkResponseArrayOutput {
	return o.ApplyT(func(v *ContactProfile) ContactProfileLinkResponseArrayOutput { return v.Links }).(ContactProfileLinkResponseArrayOutput)
}

func (o ContactProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactProfile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ContactProfileOutput) MinimumElevationDegrees() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContactProfile) pulumi.Float64PtrOutput { return v.MinimumElevationDegrees }).(pulumi.Float64PtrOutput)
}

func (o ContactProfileOutput) MinimumViableContactDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactProfile) pulumi.StringPtrOutput { return v.MinimumViableContactDuration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContactProfileOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContactProfile) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ContactProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContactProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ContactProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContactProfileOutput{})
}
