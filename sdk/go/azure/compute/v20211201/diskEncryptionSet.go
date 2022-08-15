


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiskEncryptionSet struct {
	pulumi.CustomResourceState

	ActiveKey                         KeyForDiskEncryptionSetResponsePtrOutput   `pulumi:"activeKey"`
	AutoKeyRotationError              ApiErrorResponseOutput                     `pulumi:"autoKeyRotationError"`
	EncryptionType                    pulumi.StringPtrOutput                     `pulumi:"encryptionType"`
	Identity                          EncryptionSetIdentityResponsePtrOutput     `pulumi:"identity"`
	LastKeyRotationTimestamp          pulumi.StringOutput                        `pulumi:"lastKeyRotationTimestamp"`
	Location                          pulumi.StringOutput                        `pulumi:"location"`
	Name                              pulumi.StringOutput                        `pulumi:"name"`
	PreviousKeys                      KeyForDiskEncryptionSetResponseArrayOutput `pulumi:"previousKeys"`
	ProvisioningState                 pulumi.StringOutput                        `pulumi:"provisioningState"`
	RotationToLatestKeyVersionEnabled pulumi.BoolPtrOutput                       `pulumi:"rotationToLatestKeyVersionEnabled"`
	Tags                              pulumi.StringMapOutput                     `pulumi:"tags"`
	Type                              pulumi.StringOutput                        `pulumi:"type"`
}


func NewDiskEncryptionSet(ctx *pulumi.Context,
	name string, args *DiskEncryptionSetArgs, opts ...pulumi.ResourceOption) (*DiskEncryptionSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191101:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210801:DiskEncryptionSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220302:DiskEncryptionSet"),
		},
	})
	opts = append(opts, aliases)
	var resource DiskEncryptionSet
	err := ctx.RegisterResource("azure-native:compute/v20211201:DiskEncryptionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiskEncryptionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskEncryptionSetState, opts ...pulumi.ResourceOption) (*DiskEncryptionSet, error) {
	var resource DiskEncryptionSet
	err := ctx.ReadResource("azure-native:compute/v20211201:DiskEncryptionSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diskEncryptionSetState struct {
}

type DiskEncryptionSetState struct {
}

func (DiskEncryptionSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskEncryptionSetState)(nil)).Elem()
}

type diskEncryptionSetArgs struct {
	ActiveKey                         *KeyForDiskEncryptionSet `pulumi:"activeKey"`
	DiskEncryptionSetName             *string                  `pulumi:"diskEncryptionSetName"`
	EncryptionType                    *string                  `pulumi:"encryptionType"`
	Identity                          *EncryptionSetIdentity   `pulumi:"identity"`
	Location                          *string                  `pulumi:"location"`
	ResourceGroupName                 string                   `pulumi:"resourceGroupName"`
	RotationToLatestKeyVersionEnabled *bool                    `pulumi:"rotationToLatestKeyVersionEnabled"`
	Tags                              map[string]string        `pulumi:"tags"`
}


type DiskEncryptionSetArgs struct {
	ActiveKey                         KeyForDiskEncryptionSetPtrInput
	DiskEncryptionSetName             pulumi.StringPtrInput
	EncryptionType                    pulumi.StringPtrInput
	Identity                          EncryptionSetIdentityPtrInput
	Location                          pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	RotationToLatestKeyVersionEnabled pulumi.BoolPtrInput
	Tags                              pulumi.StringMapInput
}

func (DiskEncryptionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskEncryptionSetArgs)(nil)).Elem()
}

type DiskEncryptionSetInput interface {
	pulumi.Input

	ToDiskEncryptionSetOutput() DiskEncryptionSetOutput
	ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput
}

func (*DiskEncryptionSet) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSet)(nil)).Elem()
}

func (i *DiskEncryptionSet) ToDiskEncryptionSetOutput() DiskEncryptionSetOutput {
	return i.ToDiskEncryptionSetOutputWithContext(context.Background())
}

func (i *DiskEncryptionSet) ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetOutput)
}

type DiskEncryptionSetOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSet)(nil)).Elem()
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetOutput() DiskEncryptionSetOutput {
	return o
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput {
	return o
}

func (o DiskEncryptionSetOutput) ActiveKey() KeyForDiskEncryptionSetResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) KeyForDiskEncryptionSetResponsePtrOutput { return v.ActiveKey }).(KeyForDiskEncryptionSetResponsePtrOutput)
}

func (o DiskEncryptionSetOutput) AutoKeyRotationError() ApiErrorResponseOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) ApiErrorResponseOutput { return v.AutoKeyRotationError }).(ApiErrorResponseOutput)
}

func (o DiskEncryptionSetOutput) EncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringPtrOutput { return v.EncryptionType }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionSetOutput) Identity() EncryptionSetIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) EncryptionSetIdentityResponsePtrOutput { return v.Identity }).(EncryptionSetIdentityResponsePtrOutput)
}

func (o DiskEncryptionSetOutput) LastKeyRotationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.LastKeyRotationTimestamp }).(pulumi.StringOutput)
}

func (o DiskEncryptionSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DiskEncryptionSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiskEncryptionSetOutput) PreviousKeys() KeyForDiskEncryptionSetResponseArrayOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) KeyForDiskEncryptionSetResponseArrayOutput { return v.PreviousKeys }).(KeyForDiskEncryptionSetResponseArrayOutput)
}

func (o DiskEncryptionSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DiskEncryptionSetOutput) RotationToLatestKeyVersionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.BoolPtrOutput { return v.RotationToLatestKeyVersionEnabled }).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DiskEncryptionSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskEncryptionSetOutput{})
}
