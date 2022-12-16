


package v20220904

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SyncSet struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput      `pulumi:"name"`
	Resources  pulumi.StringPtrOutput   `pulumi:"resources"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewSyncSet(ctx *pulumi.Context,
	name string, args *SyncSetArgs, opts ...pulumi.ResourceOption) (*SyncSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	var resource SyncSet
	err := ctx.RegisterResource("azure-native:redhatopenshift/v20220904:SyncSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSyncSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncSetState, opts ...pulumi.ResourceOption) (*SyncSet, error) {
	var resource SyncSet
	err := ctx.ReadResource("azure-native:redhatopenshift/v20220904:SyncSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type syncSetState struct {
}

type SyncSetState struct {
}

func (SyncSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncSetState)(nil)).Elem()
}

type syncSetArgs struct {
	ChildResourceName *string `pulumi:"childResourceName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	Resources         *string `pulumi:"resources"`
}


type SyncSetArgs struct {
	ChildResourceName pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
	Resources         pulumi.StringPtrInput
}

func (SyncSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncSetArgs)(nil)).Elem()
}

type SyncSetInput interface {
	pulumi.Input

	ToSyncSetOutput() SyncSetOutput
	ToSyncSetOutputWithContext(ctx context.Context) SyncSetOutput
}

func (*SyncSet) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncSet)(nil)).Elem()
}

func (i *SyncSet) ToSyncSetOutput() SyncSetOutput {
	return i.ToSyncSetOutputWithContext(context.Background())
}

func (i *SyncSet) ToSyncSetOutputWithContext(ctx context.Context) SyncSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncSetOutput)
}

type SyncSetOutput struct{ *pulumi.OutputState }

func (SyncSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncSet)(nil)).Elem()
}

func (o SyncSetOutput) ToSyncSetOutput() SyncSetOutput {
	return o
}

func (o SyncSetOutput) ToSyncSetOutputWithContext(ctx context.Context) SyncSetOutput {
	return o
}

func (o SyncSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SyncSetOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncSet) pulumi.StringPtrOutput { return v.Resources }).(pulumi.StringPtrOutput)
}

func (o SyncSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SyncSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SyncSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SyncSetOutput{})
}
