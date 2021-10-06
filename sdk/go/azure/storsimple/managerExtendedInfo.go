


package storsimple

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagerExtendedInfo struct {
	pulumi.CustomResourceState

	Algorithm                   pulumi.StringOutput    `pulumi:"algorithm"`
	EncryptionKey               pulumi.StringPtrOutput `pulumi:"encryptionKey"`
	EncryptionKeyThumbprint     pulumi.StringPtrOutput `pulumi:"encryptionKeyThumbprint"`
	Etag                        pulumi.StringPtrOutput `pulumi:"etag"`
	IntegrityKey                pulumi.StringOutput    `pulumi:"integrityKey"`
	Kind                        pulumi.StringPtrOutput `pulumi:"kind"`
	Name                        pulumi.StringOutput    `pulumi:"name"`
	PortalCertificateThumbprint pulumi.StringPtrOutput `pulumi:"portalCertificateThumbprint"`
	Type                        pulumi.StringOutput    `pulumi:"type"`
	Version                     pulumi.StringPtrOutput `pulumi:"version"`
}


func NewManagerExtendedInfo(ctx *pulumi.Context,
	name string, args *ManagerExtendedInfoArgs, opts ...pulumi.ResourceOption) (*ManagerExtendedInfo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Algorithm == nil {
		return nil, errors.New("invalid value for required argument 'Algorithm'")
	}
	if args.IntegrityKey == nil {
		return nil, errors.New("invalid value for required argument 'IntegrityKey'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple:ManagerExtendedInfo"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20161001:ManagerExtendedInfo"),
		},
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20161001:ManagerExtendedInfo"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20170601:ManagerExtendedInfo"),
		},
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20170601:ManagerExtendedInfo"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagerExtendedInfo
	err := ctx.RegisterResource("azure-native:storsimple:ManagerExtendedInfo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagerExtendedInfo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagerExtendedInfoState, opts ...pulumi.ResourceOption) (*ManagerExtendedInfo, error) {
	var resource ManagerExtendedInfo
	err := ctx.ReadResource("azure-native:storsimple:ManagerExtendedInfo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managerExtendedInfoState struct {
}

type ManagerExtendedInfoState struct {
}

func (ManagerExtendedInfoState) ElementType() reflect.Type {
	return reflect.TypeOf((*managerExtendedInfoState)(nil)).Elem()
}

type managerExtendedInfoArgs struct {
	Algorithm                   string  `pulumi:"algorithm"`
	EncryptionKey               *string `pulumi:"encryptionKey"`
	EncryptionKeyThumbprint     *string `pulumi:"encryptionKeyThumbprint"`
	Etag                        *string `pulumi:"etag"`
	IntegrityKey                string  `pulumi:"integrityKey"`
	Kind                        *Kind   `pulumi:"kind"`
	ManagerName                 string  `pulumi:"managerName"`
	PortalCertificateThumbprint *string `pulumi:"portalCertificateThumbprint"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
	Version                     *string `pulumi:"version"`
}


type ManagerExtendedInfoArgs struct {
	Algorithm                   pulumi.StringInput
	EncryptionKey               pulumi.StringPtrInput
	EncryptionKeyThumbprint     pulumi.StringPtrInput
	Etag                        pulumi.StringPtrInput
	IntegrityKey                pulumi.StringInput
	Kind                        KindPtrInput
	ManagerName                 pulumi.StringInput
	PortalCertificateThumbprint pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	Version                     pulumi.StringPtrInput
}

func (ManagerExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managerExtendedInfoArgs)(nil)).Elem()
}

type ManagerExtendedInfoInput interface {
	pulumi.Input

	ToManagerExtendedInfoOutput() ManagerExtendedInfoOutput
	ToManagerExtendedInfoOutputWithContext(ctx context.Context) ManagerExtendedInfoOutput
}

func (*ManagerExtendedInfo) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerExtendedInfo)(nil))
}

func (i *ManagerExtendedInfo) ToManagerExtendedInfoOutput() ManagerExtendedInfoOutput {
	return i.ToManagerExtendedInfoOutputWithContext(context.Background())
}

func (i *ManagerExtendedInfo) ToManagerExtendedInfoOutputWithContext(ctx context.Context) ManagerExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerExtendedInfoOutput)
}

type ManagerExtendedInfoOutput struct{ *pulumi.OutputState }

func (ManagerExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerExtendedInfo)(nil))
}

func (o ManagerExtendedInfoOutput) ToManagerExtendedInfoOutput() ManagerExtendedInfoOutput {
	return o
}

func (o ManagerExtendedInfoOutput) ToManagerExtendedInfoOutputWithContext(ctx context.Context) ManagerExtendedInfoOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagerExtendedInfoOutput{})
}
