


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleAssignmentArtifact struct {
	pulumi.CustomResourceState

	DependsOn        pulumi.StringArrayOutput `pulumi:"dependsOn"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	DisplayName      pulumi.StringPtrOutput   `pulumi:"displayName"`
	Kind             pulumi.StringOutput      `pulumi:"kind"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	PrincipalIds     pulumi.AnyOutput         `pulumi:"principalIds"`
	ResourceGroup    pulumi.StringPtrOutput   `pulumi:"resourceGroup"`
	RoleDefinitionId pulumi.StringOutput      `pulumi:"roleDefinitionId"`
	Type             pulumi.StringOutput      `pulumi:"type"`
}


func NewRoleAssignmentArtifact(ctx *pulumi.Context,
	name string, args *RoleAssignmentArtifactArgs, opts ...pulumi.ResourceOption) (*RoleAssignmentArtifact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintName == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.PrincipalIds == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalIds'")
	}
	if args.ResourceScope == nil {
		return nil, errors.New("invalid value for required argument 'ResourceScope'")
	}
	if args.RoleDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'RoleDefinitionId'")
	}
	args.Kind = pulumi.String("roleAssignment")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:blueprint/v20181101preview:RoleAssignmentArtifact"),
		},
		{
			Type: pulumi.String("azure-native:blueprint:RoleAssignmentArtifact"),
		},
		{
			Type: pulumi.String("azure-nextgen:blueprint:RoleAssignmentArtifact"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleAssignmentArtifact
	err := ctx.RegisterResource("azure-native:blueprint/v20181101preview:RoleAssignmentArtifact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoleAssignmentArtifact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentArtifactState, opts ...pulumi.ResourceOption) (*RoleAssignmentArtifact, error) {
	var resource RoleAssignmentArtifact
	err := ctx.ReadResource("azure-native:blueprint/v20181101preview:RoleAssignmentArtifact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type roleAssignmentArtifactState struct {
}

type RoleAssignmentArtifactState struct {
}

func (RoleAssignmentArtifactState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArtifactState)(nil)).Elem()
}

type roleAssignmentArtifactArgs struct {
	ArtifactName     *string     `pulumi:"artifactName"`
	BlueprintName    string      `pulumi:"blueprintName"`
	DependsOn        []string    `pulumi:"dependsOn"`
	Description      *string     `pulumi:"description"`
	DisplayName      *string     `pulumi:"displayName"`
	Kind             string      `pulumi:"kind"`
	PrincipalIds     interface{} `pulumi:"principalIds"`
	ResourceGroup    *string     `pulumi:"resourceGroup"`
	ResourceScope    string      `pulumi:"resourceScope"`
	RoleDefinitionId string      `pulumi:"roleDefinitionId"`
}


type RoleAssignmentArtifactArgs struct {
	ArtifactName     pulumi.StringPtrInput
	BlueprintName    pulumi.StringInput
	DependsOn        pulumi.StringArrayInput
	Description      pulumi.StringPtrInput
	DisplayName      pulumi.StringPtrInput
	Kind             pulumi.StringInput
	PrincipalIds     pulumi.Input
	ResourceGroup    pulumi.StringPtrInput
	ResourceScope    pulumi.StringInput
	RoleDefinitionId pulumi.StringInput
}

func (RoleAssignmentArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArtifactArgs)(nil)).Elem()
}

type RoleAssignmentArtifactInput interface {
	pulumi.Input

	ToRoleAssignmentArtifactOutput() RoleAssignmentArtifactOutput
	ToRoleAssignmentArtifactOutputWithContext(ctx context.Context) RoleAssignmentArtifactOutput
}

func (*RoleAssignmentArtifact) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignmentArtifact)(nil))
}

func (i *RoleAssignmentArtifact) ToRoleAssignmentArtifactOutput() RoleAssignmentArtifactOutput {
	return i.ToRoleAssignmentArtifactOutputWithContext(context.Background())
}

func (i *RoleAssignmentArtifact) ToRoleAssignmentArtifactOutputWithContext(ctx context.Context) RoleAssignmentArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentArtifactOutput)
}

type RoleAssignmentArtifactOutput struct{ *pulumi.OutputState }

func (RoleAssignmentArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignmentArtifact)(nil))
}

func (o RoleAssignmentArtifactOutput) ToRoleAssignmentArtifactOutput() RoleAssignmentArtifactOutput {
	return o
}

func (o RoleAssignmentArtifactOutput) ToRoleAssignmentArtifactOutputWithContext(ctx context.Context) RoleAssignmentArtifactOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAssignmentArtifactInput)(nil)).Elem(), &RoleAssignmentArtifact{})
	pulumi.RegisterOutputType(RoleAssignmentArtifactOutput{})
}
