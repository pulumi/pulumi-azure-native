


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiskEncryptionSet struct {
	pulumi.CustomResourceState

	ActiveKey                         KeyForDiskEncryptionSetResponsePtrOutput   `pulumi:"activeKey"`
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
			Type: pulumi.String("azure-native:compute/v20210401:DiskEncryptionSet"),
		},
	})
	opts = append(opts, aliases)
	var resource DiskEncryptionSet
	err := ctx.RegisterResource("azure-native:compute/v20201201:DiskEncryptionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiskEncryptionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskEncryptionSetState, opts ...pulumi.ResourceOption) (*DiskEncryptionSet, error) {
	var resource DiskEncryptionSet
	err := ctx.ReadResource("azure-native:compute/v20201201:DiskEncryptionSet", name, id, state, &resource, opts...)
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

func init() {
	pulumi.RegisterOutputType(DiskEncryptionSetOutput{})
}
