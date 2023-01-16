


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterPrincipalAssignment struct {
	pulumi.CustomResourceState

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


func NewClusterPrincipalAssignment(ctx *pulumi.Context,
	name string, args *ClusterPrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*ClusterPrincipalAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
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
			Type: pulumi.String("azure-native:kusto:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:ClusterPrincipalAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource ClusterPrincipalAssignment
	err := ctx.RegisterResource("azure-native:kusto/v20210101:ClusterPrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetClusterPrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterPrincipalAssignmentState, opts ...pulumi.ResourceOption) (*ClusterPrincipalAssignment, error) {
	var resource ClusterPrincipalAssignment
	err := ctx.ReadResource("azure-native:kusto/v20210101:ClusterPrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type clusterPrincipalAssignmentState struct {
}

type ClusterPrincipalAssignmentState struct {
}

func (ClusterPrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPrincipalAssignmentState)(nil)).Elem()
}

type clusterPrincipalAssignmentArgs struct {
	ClusterName             string  `pulumi:"clusterName"`
	PrincipalAssignmentName *string `pulumi:"principalAssignmentName"`
	PrincipalId             string  `pulumi:"principalId"`
	PrincipalType           string  `pulumi:"principalType"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	Role                    string  `pulumi:"role"`
	TenantId                *string `pulumi:"tenantId"`
}


type ClusterPrincipalAssignmentArgs struct {
	ClusterName             pulumi.StringInput
	PrincipalAssignmentName pulumi.StringPtrInput
	PrincipalId             pulumi.StringInput
	PrincipalType           pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	Role                    pulumi.StringInput
	TenantId                pulumi.StringPtrInput
}

func (ClusterPrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPrincipalAssignmentArgs)(nil)).Elem()
}

type ClusterPrincipalAssignmentInput interface {
	pulumi.Input

	ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput
	ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput
}

func (*ClusterPrincipalAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPrincipalAssignment)(nil)).Elem()
}

func (i *ClusterPrincipalAssignment) ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput {
	return i.ToClusterPrincipalAssignmentOutputWithContext(context.Background())
}

func (i *ClusterPrincipalAssignment) ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPrincipalAssignmentOutput)
}

type ClusterPrincipalAssignmentOutput struct{ *pulumi.OutputState }

func (ClusterPrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPrincipalAssignment)(nil)).Elem()
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput {
	return o
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput {
	return o
}

func (o ClusterPrincipalAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterPrincipalAssignmentOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ClusterPrincipalAssignmentOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o ClusterPrincipalAssignmentOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o ClusterPrincipalAssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ClusterPrincipalAssignmentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o ClusterPrincipalAssignmentOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ClusterPrincipalAssignmentOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.TenantName }).(pulumi.StringOutput)
}

func (o ClusterPrincipalAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterPrincipalAssignmentOutput{})
}
