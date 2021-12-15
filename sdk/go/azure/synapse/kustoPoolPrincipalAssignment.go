


package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoPoolPrincipalAssignment struct {
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


func NewKustoPoolPrincipalAssignment(ctx *pulumi.Context,
	name string, args *KustoPoolPrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*KustoPoolPrincipalAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:synapse/v20210401preview:KustoPoolPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:KustoPoolPrincipalAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPoolPrincipalAssignment
	err := ctx.RegisterResource("azure-native:synapse:KustoPoolPrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoPoolPrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolPrincipalAssignmentState, opts ...pulumi.ResourceOption) (*KustoPoolPrincipalAssignment, error) {
	var resource KustoPoolPrincipalAssignment
	err := ctx.ReadResource("azure-native:synapse:KustoPoolPrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoPoolPrincipalAssignmentState struct {
}

type KustoPoolPrincipalAssignmentState struct {
}

func (KustoPoolPrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolPrincipalAssignmentState)(nil)).Elem()
}

type kustoPoolPrincipalAssignmentArgs struct {
	KustoPoolName           string  `pulumi:"kustoPoolName"`
	PrincipalAssignmentName *string `pulumi:"principalAssignmentName"`
	PrincipalId             string  `pulumi:"principalId"`
	PrincipalType           string  `pulumi:"principalType"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	Role                    string  `pulumi:"role"`
	TenantId                *string `pulumi:"tenantId"`
	WorkspaceName           string  `pulumi:"workspaceName"`
}


type KustoPoolPrincipalAssignmentArgs struct {
	KustoPoolName           pulumi.StringInput
	PrincipalAssignmentName pulumi.StringPtrInput
	PrincipalId             pulumi.StringInput
	PrincipalType           pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	Role                    pulumi.StringInput
	TenantId                pulumi.StringPtrInput
	WorkspaceName           pulumi.StringInput
}

func (KustoPoolPrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolPrincipalAssignmentArgs)(nil)).Elem()
}

type KustoPoolPrincipalAssignmentInput interface {
	pulumi.Input

	ToKustoPoolPrincipalAssignmentOutput() KustoPoolPrincipalAssignmentOutput
	ToKustoPoolPrincipalAssignmentOutputWithContext(ctx context.Context) KustoPoolPrincipalAssignmentOutput
}

func (*KustoPoolPrincipalAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoPoolPrincipalAssignment)(nil))
}

func (i *KustoPoolPrincipalAssignment) ToKustoPoolPrincipalAssignmentOutput() KustoPoolPrincipalAssignmentOutput {
	return i.ToKustoPoolPrincipalAssignmentOutputWithContext(context.Background())
}

func (i *KustoPoolPrincipalAssignment) ToKustoPoolPrincipalAssignmentOutputWithContext(ctx context.Context) KustoPoolPrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolPrincipalAssignmentOutput)
}

type KustoPoolPrincipalAssignmentOutput struct{ *pulumi.OutputState }

func (KustoPoolPrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoPoolPrincipalAssignment)(nil))
}

func (o KustoPoolPrincipalAssignmentOutput) ToKustoPoolPrincipalAssignmentOutput() KustoPoolPrincipalAssignmentOutput {
	return o
}

func (o KustoPoolPrincipalAssignmentOutput) ToKustoPoolPrincipalAssignmentOutputWithContext(ctx context.Context) KustoPoolPrincipalAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KustoPoolPrincipalAssignmentOutput{})
}
