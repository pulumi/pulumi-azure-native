


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Replica struct {
	pulumi.CustomResourceState

	Endpoint          pulumi.StringOutput      `pulumi:"endpoint"`
	Location          pulumi.StringPtrOutput   `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewReplica(ctx *pulumi.Context,
	name string, args *ReplicaArgs, opts ...pulumi.ResourceOption) (*Replica, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigStoreName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigStoreName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Replica
	err := ctx.RegisterResource("azure-native:appconfiguration/v20220301preview:Replica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicaState, opts ...pulumi.ResourceOption) (*Replica, error) {
	var resource Replica
	err := ctx.ReadResource("azure-native:appconfiguration/v20220301preview:Replica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicaState struct {
}

type ReplicaState struct {
}

func (ReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaState)(nil)).Elem()
}

type replicaArgs struct {
	ConfigStoreName   string  `pulumi:"configStoreName"`
	Location          *string `pulumi:"location"`
	ReplicaName       *string `pulumi:"replicaName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ReplicaArgs struct {
	ConfigStoreName   pulumi.StringInput
	Location          pulumi.StringPtrInput
	ReplicaName       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (ReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaArgs)(nil)).Elem()
}

type ReplicaInput interface {
	pulumi.Input

	ToReplicaOutput() ReplicaOutput
	ToReplicaOutputWithContext(ctx context.Context) ReplicaOutput
}

func (*Replica) ElementType() reflect.Type {
	return reflect.TypeOf((**Replica)(nil)).Elem()
}

func (i *Replica) ToReplicaOutput() ReplicaOutput {
	return i.ToReplicaOutputWithContext(context.Background())
}

func (i *Replica) ToReplicaOutputWithContext(ctx context.Context) ReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaOutput)
}

type ReplicaOutput struct{ *pulumi.OutputState }

func (ReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Replica)(nil)).Elem()
}

func (o ReplicaOutput) ToReplicaOutput() ReplicaOutput {
	return o
}

func (o ReplicaOutput) ToReplicaOutputWithContext(ctx context.Context) ReplicaOutput {
	return o
}

func (o ReplicaOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o ReplicaOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ReplicaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReplicaOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ReplicaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Replica) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ReplicaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicaOutput{})
}
