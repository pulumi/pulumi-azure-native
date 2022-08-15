


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoPoolDatabasePrincipalAssignment struct {
	pulumi.CustomResourceState

	AadObjectId       pulumi.StringOutput      `pulumi:"aadObjectId"`
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


func NewKustoPoolDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, args *KustoPoolDatabasePrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*KustoPoolDatabasePrincipalAssignment, error) {
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
			Type: pulumi.String("azure-native:synapse:KustoPoolDatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:KustoPoolDatabasePrincipalAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPoolDatabasePrincipalAssignment
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:KustoPoolDatabasePrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoPoolDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolDatabasePrincipalAssignmentState, opts ...pulumi.ResourceOption) (*KustoPoolDatabasePrincipalAssignment, error) {
	var resource KustoPoolDatabasePrincipalAssignment
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:KustoPoolDatabasePrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoPoolDatabasePrincipalAssignmentState struct {
}

type KustoPoolDatabasePrincipalAssignmentState struct {
}

func (KustoPoolDatabasePrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolDatabasePrincipalAssignmentState)(nil)).Elem()
}

type kustoPoolDatabasePrincipalAssignmentArgs struct {
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


type KustoPoolDatabasePrincipalAssignmentArgs struct {
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

func (KustoPoolDatabasePrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolDatabasePrincipalAssignmentArgs)(nil)).Elem()
}

type KustoPoolDatabasePrincipalAssignmentInput interface {
	pulumi.Input

	ToKustoPoolDatabasePrincipalAssignmentOutput() KustoPoolDatabasePrincipalAssignmentOutput
	ToKustoPoolDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) KustoPoolDatabasePrincipalAssignmentOutput
}

func (*KustoPoolDatabasePrincipalAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolDatabasePrincipalAssignment)(nil)).Elem()
}

func (i *KustoPoolDatabasePrincipalAssignment) ToKustoPoolDatabasePrincipalAssignmentOutput() KustoPoolDatabasePrincipalAssignmentOutput {
	return i.ToKustoPoolDatabasePrincipalAssignmentOutputWithContext(context.Background())
}

func (i *KustoPoolDatabasePrincipalAssignment) ToKustoPoolDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) KustoPoolDatabasePrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolDatabasePrincipalAssignmentOutput)
}

type KustoPoolDatabasePrincipalAssignmentOutput struct{ *pulumi.OutputState }

func (KustoPoolDatabasePrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolDatabasePrincipalAssignment)(nil)).Elem()
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) ToKustoPoolDatabasePrincipalAssignmentOutput() KustoPoolDatabasePrincipalAssignmentOutput {
	return o
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) ToKustoPoolDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) KustoPoolDatabasePrincipalAssignmentOutput {
	return o
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) AadObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringOutput { return v.AadObjectId }).(pulumi.StringOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringOutput { return v.TenantName }).(pulumi.StringOutput)
}

func (o KustoPoolDatabasePrincipalAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabasePrincipalAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoPoolDatabasePrincipalAssignmentOutput{})
}
