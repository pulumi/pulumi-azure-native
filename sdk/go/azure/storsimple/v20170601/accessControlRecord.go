


package v20170601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessControlRecord struct {
	pulumi.CustomResourceState

	InitiatorName pulumi.StringOutput    `pulumi:"initiatorName"`
	Kind          pulumi.StringPtrOutput `pulumi:"kind"`
	Name          pulumi.StringOutput    `pulumi:"name"`
	Type          pulumi.StringOutput    `pulumi:"type"`
	VolumeCount   pulumi.IntOutput       `pulumi:"volumeCount"`
}


func NewAccessControlRecord(ctx *pulumi.Context,
	name string, args *AccessControlRecordArgs, opts ...pulumi.ResourceOption) (*AccessControlRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InitiatorName == nil {
		return nil, errors.New("invalid value for required argument 'InitiatorName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20170601:AccessControlRecord"),
		},
		{
			Type: pulumi.String("azure-native:storsimple:AccessControlRecord"),
		},
		{
			Type: pulumi.String("azure-nextgen:storsimple:AccessControlRecord"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20161001:AccessControlRecord"),
		},
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20161001:AccessControlRecord"),
		},
	})
	opts = append(opts, aliases)
	var resource AccessControlRecord
	err := ctx.RegisterResource("azure-native:storsimple/v20170601:AccessControlRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccessControlRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessControlRecordState, opts ...pulumi.ResourceOption) (*AccessControlRecord, error) {
	var resource AccessControlRecord
	err := ctx.ReadResource("azure-native:storsimple/v20170601:AccessControlRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accessControlRecordState struct {
}

type AccessControlRecordState struct {
}

func (AccessControlRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessControlRecordState)(nil)).Elem()
}

type accessControlRecordArgs struct {
	AccessControlRecordName *string `pulumi:"accessControlRecordName"`
	InitiatorName           string  `pulumi:"initiatorName"`
	Kind                    *Kind   `pulumi:"kind"`
	ManagerName             string  `pulumi:"managerName"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
}


type AccessControlRecordArgs struct {
	AccessControlRecordName pulumi.StringPtrInput
	InitiatorName           pulumi.StringInput
	Kind                    KindPtrInput
	ManagerName             pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
}

func (AccessControlRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessControlRecordArgs)(nil)).Elem()
}

type AccessControlRecordInput interface {
	pulumi.Input

	ToAccessControlRecordOutput() AccessControlRecordOutput
	ToAccessControlRecordOutputWithContext(ctx context.Context) AccessControlRecordOutput
}

func (*AccessControlRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessControlRecord)(nil))
}

func (i *AccessControlRecord) ToAccessControlRecordOutput() AccessControlRecordOutput {
	return i.ToAccessControlRecordOutputWithContext(context.Background())
}

func (i *AccessControlRecord) ToAccessControlRecordOutputWithContext(ctx context.Context) AccessControlRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessControlRecordOutput)
}

type AccessControlRecordOutput struct{ *pulumi.OutputState }

func (AccessControlRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessControlRecord)(nil))
}

func (o AccessControlRecordOutput) ToAccessControlRecordOutput() AccessControlRecordOutput {
	return o
}

func (o AccessControlRecordOutput) ToAccessControlRecordOutputWithContext(ctx context.Context) AccessControlRecordOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccessControlRecordOutput{})
}
