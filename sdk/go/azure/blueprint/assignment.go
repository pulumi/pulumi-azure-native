


package blueprint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Assignment struct {
	pulumi.CustomResourceState

	BlueprintId       pulumi.StringPtrOutput                  `pulumi:"blueprintId"`
	Description       pulumi.StringPtrOutput                  `pulumi:"description"`
	DisplayName       pulumi.StringPtrOutput                  `pulumi:"displayName"`
	Identity          ManagedServiceIdentityResponseOutput    `pulumi:"identity"`
	Location          pulumi.StringOutput                     `pulumi:"location"`
	Locks             AssignmentLockSettingsResponsePtrOutput `pulumi:"locks"`
	Name              pulumi.StringOutput                     `pulumi:"name"`
	Parameters        ParameterValueResponseMapOutput         `pulumi:"parameters"`
	ProvisioningState pulumi.StringOutput                     `pulumi:"provisioningState"`
	ResourceGroups    ResourceGroupValueResponseMapOutput     `pulumi:"resourceGroups"`
	Scope             pulumi.StringPtrOutput                  `pulumi:"scope"`
	Status            AssignmentStatusResponseOutput          `pulumi:"status"`
	Type              pulumi.StringOutput                     `pulumi:"type"`
}


func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.ResourceGroups == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroups'")
	}
	if args.ResourceScope == nil {
		return nil, errors.New("invalid value for required argument 'ResourceScope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:blueprint/v20181101preview:Assignment"),
		},
	})
	opts = append(opts, aliases)
	var resource Assignment
	err := ctx.RegisterResource("azure-native:blueprint:Assignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentState, opts ...pulumi.ResourceOption) (*Assignment, error) {
	var resource Assignment
	err := ctx.ReadResource("azure-native:blueprint:Assignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type assignmentState struct {
}

type AssignmentState struct {
}

func (AssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentState)(nil)).Elem()
}

type assignmentArgs struct {
	AssignmentName *string                       `pulumi:"assignmentName"`
	BlueprintId    *string                       `pulumi:"blueprintId"`
	Description    *string                       `pulumi:"description"`
	DisplayName    *string                       `pulumi:"displayName"`
	Identity       ManagedServiceIdentity        `pulumi:"identity"`
	Location       *string                       `pulumi:"location"`
	Locks          *AssignmentLockSettings       `pulumi:"locks"`
	Parameters     map[string]ParameterValue     `pulumi:"parameters"`
	ResourceGroups map[string]ResourceGroupValue `pulumi:"resourceGroups"`
	ResourceScope  string                        `pulumi:"resourceScope"`
	Scope          *string                       `pulumi:"scope"`
}


type AssignmentArgs struct {
	AssignmentName pulumi.StringPtrInput
	BlueprintId    pulumi.StringPtrInput
	Description    pulumi.StringPtrInput
	DisplayName    pulumi.StringPtrInput
	Identity       ManagedServiceIdentityInput
	Location       pulumi.StringPtrInput
	Locks          AssignmentLockSettingsPtrInput
	Parameters     ParameterValueMapInput
	ResourceGroups ResourceGroupValueMapInput
	ResourceScope  pulumi.StringInput
	Scope          pulumi.StringPtrInput
}

func (AssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentArgs)(nil)).Elem()
}

type AssignmentInput interface {
	pulumi.Input

	ToAssignmentOutput() AssignmentOutput
	ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput
}

func (*Assignment) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (i *Assignment) ToAssignmentOutput() AssignmentOutput {
	return i.ToAssignmentOutputWithContext(context.Background())
}

func (i *Assignment) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentOutput)
}

type AssignmentOutput struct{ *pulumi.OutputState }

func (AssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (o AssignmentOutput) ToAssignmentOutput() AssignmentOutput {
	return o
}

func (o AssignmentOutput) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AssignmentOutput{})
}
