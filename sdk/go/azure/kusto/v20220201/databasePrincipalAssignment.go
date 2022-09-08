


package v20220201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabasePrincipalAssignment struct {
	pulumi.CustomResourceState

	AadObjectId       pulumi.StringOutput    `pulumi:"aadObjectId"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	PrincipalId       pulumi.StringOutput    `pulumi:"principalId"`
	PrincipalName     pulumi.StringOutput    `pulumi:"principalName"`
	PrincipalType     pulumi.StringOutput    `pulumi:"principalType"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Role              pulumi.StringOutput    `pulumi:"role"`
	TenantId          pulumi.StringPtrOutput `pulumi:"tenantId"`
	TenantName        pulumi.StringOutput    `pulumi:"tenantName"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, args *DatabasePrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*DatabasePrincipalAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:DatabasePrincipalAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabasePrincipalAssignment
	err := ctx.RegisterResource("azure-native:kusto/v20220201:DatabasePrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasePrincipalAssignmentState, opts ...pulumi.ResourceOption) (*DatabasePrincipalAssignment, error) {
	var resource DatabasePrincipalAssignment
	err := ctx.ReadResource("azure-native:kusto/v20220201:DatabasePrincipalAssignment", name, id, state, &resource, opts...)
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
	ClusterName             string  `pulumi:"clusterName"`
	DatabaseName            string  `pulumi:"databaseName"`
	PrincipalAssignmentName *string `pulumi:"principalAssignmentName"`
	PrincipalId             string  `pulumi:"principalId"`
	PrincipalType           string  `pulumi:"principalType"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	Role                    string  `pulumi:"role"`
	TenantId                *string `pulumi:"tenantId"`
}


type DatabasePrincipalAssignmentArgs struct {
	ClusterName             pulumi.StringInput
	DatabaseName            pulumi.StringInput
	PrincipalAssignmentName pulumi.StringPtrInput
	PrincipalId             pulumi.StringInput
	PrincipalType           pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	Role                    pulumi.StringInput
	TenantId                pulumi.StringPtrInput
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
	return reflect.TypeOf((**DatabasePrincipalAssignment)(nil)).Elem()
}

func (i *DatabasePrincipalAssignment) ToDatabasePrincipalAssignmentOutput() DatabasePrincipalAssignmentOutput {
	return i.ToDatabasePrincipalAssignmentOutputWithContext(context.Background())
}

func (i *DatabasePrincipalAssignment) ToDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) DatabasePrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalAssignmentOutput)
}

type DatabasePrincipalAssignmentOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePrincipalAssignment)(nil)).Elem()
}

func (o DatabasePrincipalAssignmentOutput) ToDatabasePrincipalAssignmentOutput() DatabasePrincipalAssignmentOutput {
	return o
}

func (o DatabasePrincipalAssignmentOutput) ToDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) DatabasePrincipalAssignmentOutput {
	return o
}

func (o DatabasePrincipalAssignmentOutput) AadObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.AadObjectId }).(pulumi.StringOutput)
}

func (o DatabasePrincipalAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabasePrincipalAssignmentOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o DatabasePrincipalAssignmentOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o DatabasePrincipalAssignmentOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o DatabasePrincipalAssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatabasePrincipalAssignmentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o DatabasePrincipalAssignmentOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o DatabasePrincipalAssignmentOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.TenantName }).(pulumi.StringOutput)
}

func (o DatabasePrincipalAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabasePrincipalAssignmentOutput{})
}
