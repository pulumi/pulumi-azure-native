


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScopeAssignment struct {
	pulumi.CustomResourceState

	AssignedManagedNetwork pulumi.StringPtrOutput `pulumi:"assignedManagedNetwork"`
	Etag                   pulumi.StringOutput    `pulumi:"etag"`
	Location               pulumi.StringPtrOutput `pulumi:"location"`
	Name                   pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput    `pulumi:"provisioningState"`
	Type                   pulumi.StringOutput    `pulumi:"type"`
}


func NewScopeAssignment(ctx *pulumi.Context,
	name string, args *ScopeAssignmentArgs, opts ...pulumi.ResourceOption) (*ScopeAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetwork:ScopeAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource ScopeAssignment
	err := ctx.RegisterResource("azure-native:managednetwork/v20190601preview:ScopeAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScopeAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScopeAssignmentState, opts ...pulumi.ResourceOption) (*ScopeAssignment, error) {
	var resource ScopeAssignment
	err := ctx.ReadResource("azure-native:managednetwork/v20190601preview:ScopeAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scopeAssignmentState struct {
}

type ScopeAssignmentState struct {
}

func (ScopeAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeAssignmentState)(nil)).Elem()
}

type scopeAssignmentArgs struct {
	AssignedManagedNetwork *string `pulumi:"assignedManagedNetwork"`
	Location               *string `pulumi:"location"`
	Scope                  string  `pulumi:"scope"`
	ScopeAssignmentName    *string `pulumi:"scopeAssignmentName"`
}


type ScopeAssignmentArgs struct {
	AssignedManagedNetwork pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	Scope                  pulumi.StringInput
	ScopeAssignmentName    pulumi.StringPtrInput
}

func (ScopeAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeAssignmentArgs)(nil)).Elem()
}

type ScopeAssignmentInput interface {
	pulumi.Input

	ToScopeAssignmentOutput() ScopeAssignmentOutput
	ToScopeAssignmentOutputWithContext(ctx context.Context) ScopeAssignmentOutput
}

func (*ScopeAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeAssignment)(nil)).Elem()
}

func (i *ScopeAssignment) ToScopeAssignmentOutput() ScopeAssignmentOutput {
	return i.ToScopeAssignmentOutputWithContext(context.Background())
}

func (i *ScopeAssignment) ToScopeAssignmentOutputWithContext(ctx context.Context) ScopeAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeAssignmentOutput)
}

type ScopeAssignmentOutput struct{ *pulumi.OutputState }

func (ScopeAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeAssignment)(nil)).Elem()
}

func (o ScopeAssignmentOutput) ToScopeAssignmentOutput() ScopeAssignmentOutput {
	return o
}

func (o ScopeAssignmentOutput) ToScopeAssignmentOutputWithContext(ctx context.Context) ScopeAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScopeAssignmentOutput{})
}
