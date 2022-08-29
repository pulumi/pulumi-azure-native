


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GalleryApplication struct {
	pulumi.CustomResourceState

	Description         pulumi.StringPtrOutput `pulumi:"description"`
	EndOfLifeDate       pulumi.StringPtrOutput `pulumi:"endOfLifeDate"`
	Eula                pulumi.StringPtrOutput `pulumi:"eula"`
	Location            pulumi.StringOutput    `pulumi:"location"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	PrivacyStatementUri pulumi.StringPtrOutput `pulumi:"privacyStatementUri"`
	ReleaseNoteUri      pulumi.StringPtrOutput `pulumi:"releaseNoteUri"`
	SupportedOSType     pulumi.StringOutput    `pulumi:"supportedOSType"`
	Tags                pulumi.StringMapOutput `pulumi:"tags"`
	Type                pulumi.StringOutput    `pulumi:"type"`
}


func NewGalleryApplication(ctx *pulumi.Context,
	name string, args *GalleryApplicationArgs, opts ...pulumi.ResourceOption) (*GalleryApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SupportedOSType == nil {
		return nil, errors.New("invalid value for required argument 'SupportedOSType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220103:GalleryApplication"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryApplication
	err := ctx.RegisterResource("azure-native:compute/v20211001:GalleryApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGalleryApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryApplicationState, opts ...pulumi.ResourceOption) (*GalleryApplication, error) {
	var resource GalleryApplication
	err := ctx.ReadResource("azure-native:compute/v20211001:GalleryApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type galleryApplicationState struct {
}

type GalleryApplicationState struct {
}

func (GalleryApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryApplicationState)(nil)).Elem()
}

type galleryApplicationArgs struct {
	Description            *string              `pulumi:"description"`
	EndOfLifeDate          *string              `pulumi:"endOfLifeDate"`
	Eula                   *string              `pulumi:"eula"`
	GalleryApplicationName *string              `pulumi:"galleryApplicationName"`
	GalleryName            string               `pulumi:"galleryName"`
	Location               *string              `pulumi:"location"`
	PrivacyStatementUri    *string              `pulumi:"privacyStatementUri"`
	ReleaseNoteUri         *string              `pulumi:"releaseNoteUri"`
	ResourceGroupName      string               `pulumi:"resourceGroupName"`
	SupportedOSType        OperatingSystemTypes `pulumi:"supportedOSType"`
	Tags                   map[string]string    `pulumi:"tags"`
}


type GalleryApplicationArgs struct {
	Description            pulumi.StringPtrInput
	EndOfLifeDate          pulumi.StringPtrInput
	Eula                   pulumi.StringPtrInput
	GalleryApplicationName pulumi.StringPtrInput
	GalleryName            pulumi.StringInput
	Location               pulumi.StringPtrInput
	PrivacyStatementUri    pulumi.StringPtrInput
	ReleaseNoteUri         pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	SupportedOSType        OperatingSystemTypesInput
	Tags                   pulumi.StringMapInput
}

func (GalleryApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryApplicationArgs)(nil)).Elem()
}

type GalleryApplicationInput interface {
	pulumi.Input

	ToGalleryApplicationOutput() GalleryApplicationOutput
	ToGalleryApplicationOutputWithContext(ctx context.Context) GalleryApplicationOutput
}

func (*GalleryApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryApplication)(nil)).Elem()
}

func (i *GalleryApplication) ToGalleryApplicationOutput() GalleryApplicationOutput {
	return i.ToGalleryApplicationOutputWithContext(context.Background())
}

func (i *GalleryApplication) ToGalleryApplicationOutputWithContext(ctx context.Context) GalleryApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryApplicationOutput)
}

type GalleryApplicationOutput struct{ *pulumi.OutputState }

func (GalleryApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryApplication)(nil)).Elem()
}

func (o GalleryApplicationOutput) ToGalleryApplicationOutput() GalleryApplicationOutput {
	return o
}

func (o GalleryApplicationOutput) ToGalleryApplicationOutputWithContext(ctx context.Context) GalleryApplicationOutput {
	return o
}

func (o GalleryApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.Eula }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GalleryApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GalleryApplicationOutput) PrivacyStatementUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.PrivacyStatementUri }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationOutput) ReleaseNoteUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.ReleaseNoteUri }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationOutput) SupportedOSType() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringOutput { return v.SupportedOSType }).(pulumi.StringOutput)
}

func (o GalleryApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GalleryApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryApplicationOutput{})
}
