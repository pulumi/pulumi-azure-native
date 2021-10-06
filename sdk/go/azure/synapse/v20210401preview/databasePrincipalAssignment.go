


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabasePrincipalAssignment struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput      `pulumi:"name"`
	PrincipalId       pulumi.StringOutput      `pulumi:"principalId"`
	PrincipalName     pulumi.StringOutput      `pulumi:"principalName"`
	PrincipalType     pulumi.StringOutput      `pulumi:"principalType"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Role              pulumi.StringOutput      `pulumi:"role"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	TenantId          pulumi.StringPtrOutput   `pulumi:"tenantId"`
	TenantName        pulumi.StringOutput      `pulumi:"tenantName"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, args *DatabasePrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*DatabasePrincipalAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.PrincipalType == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210401preview:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:synapse:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210601preview:DatabasePrincipalAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabasePrincipalAssignment
	err := ctx.RegisterResource("azure-native:synapse/v20210401preview:DatabasePrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasePrincipalAssignmentState, opts ...pulumi.ResourceOption) (*DatabasePrincipalAssignment, error) {
	var resource DatabasePrincipalAssignment
	err := ctx.ReadResource("azure-native:synapse/v20210401preview:DatabasePrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databasePrincipalAssignmentState struct {
}

type DatabasePrincipalAssignmentState struct {
}

func (DatabasePrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrincipalAssignmentState)(nil)).Elem()
}

type databasePrincipalAssignmentArgs struct {
	DatabaseName            string  `pulumi:"databaseName"`
	KustoPoolName           string  `pulumi:"kustoPoolName"`
	PrincipalAssignmentName *string `pulumi:"principalAssignmentName"`
	PrincipalId             string  `pulumi:"principalId"`
	PrincipalType           string  `pulumi:"principalType"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	Role                    string  `pulumi:"role"`
	TenantId                *string `pulumi:"tenantId"`
	WorkspaceName           string  `pulumi:"workspaceName"`
}


type DatabasePrincipalAssignmentArgs struct {
	DatabaseName            pulumi.StringInput
	KustoPoolName           pulumi.StringInput
	PrincipalAssignmentName pulumi.StringPtrInput
	PrincipalId             pulumi.StringInput
	PrincipalType           pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	Role                    pulumi.StringInput
	TenantId                pulumi.StringPtrInput
	WorkspaceName           pulumi.StringInput
}

func (DatabasePrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrincipalAssignmentArgs)(nil)).Elem()
}

type DatabasePrincipalAssignmentInput interface {
	pulumi.Input

	ToDatabasePrincipalAssignmentOutput() DatabasePrincipalAssignmentOutput
	ToDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) DatabasePrincipalAssignmentOutput
}

func (*DatabasePrincipalAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipalAssignment)(nil))
}

func (i *DatabasePrincipalAssignment) ToDatabasePrincipalAssignmentOutput() DatabasePrincipalAssignmentOutput {
	return i.ToDatabasePrincipalAssignmentOutputWithContext(context.Background())
}

func (i *DatabasePrincipalAssignment) ToDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) DatabasePrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalAssignmentOutput)
}

type DatabasePrincipalAssignmentOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipalAssignment)(nil))
}

func (o DatabasePrincipalAssignmentOutput) ToDatabasePrincipalAssignmentOutput() DatabasePrincipalAssignmentOutput {
	return o
}

func (o DatabasePrincipalAssignmentOutput) ToDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) DatabasePrincipalAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabasePrincipalAssignmentOutput{})
}
