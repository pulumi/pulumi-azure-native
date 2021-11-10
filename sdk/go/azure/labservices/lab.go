


package labservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Lab struct {
	pulumi.CustomResourceState

	CreatedByObjectId          pulumi.StringOutput                 `pulumi:"createdByObjectId"`
	CreatedByUserPrincipalName pulumi.StringOutput                 `pulumi:"createdByUserPrincipalName"`
	CreatedDate                pulumi.StringOutput                 `pulumi:"createdDate"`
	InvitationCode             pulumi.StringOutput                 `pulumi:"invitationCode"`
	LatestOperationResult      LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	Location                   pulumi.StringPtrOutput              `pulumi:"location"`
	MaxUsersInLab              pulumi.IntPtrOutput                 `pulumi:"maxUsersInLab"`
	Name                       pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState          pulumi.StringPtrOutput              `pulumi:"provisioningState"`
	Tags                       pulumi.StringMapOutput              `pulumi:"tags"`
	Type                       pulumi.StringOutput                 `pulumi:"type"`
	UniqueIdentifier           pulumi.StringPtrOutput              `pulumi:"uniqueIdentifier"`
	UsageQuota                 pulumi.StringPtrOutput              `pulumi:"usageQuota"`
	UserAccessMode             pulumi.StringPtrOutput              `pulumi:"userAccessMode"`
	UserQuota                  pulumi.IntOutput                    `pulumi:"userQuota"`
}


func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LabAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20181015:Lab"),
		},
	})
	opts = append(opts, aliases)
	var resource Lab
	err := ctx.RegisterResource("azure-native:labservices:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure-native:labservices:Lab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labState struct {
}

type LabState struct {
}

func (LabState) ElementType() reflect.Type {
	return reflect.TypeOf((*labState)(nil)).Elem()
}

type labArgs struct {
	LabAccountName    string            `pulumi:"labAccountName"`
	LabName           *string           `pulumi:"labName"`
	Location          *string           `pulumi:"location"`
	MaxUsersInLab     *int              `pulumi:"maxUsersInLab"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	UniqueIdentifier  *string           `pulumi:"uniqueIdentifier"`
	UsageQuota        *string           `pulumi:"usageQuota"`
	UserAccessMode    *string           `pulumi:"userAccessMode"`
}


type LabArgs struct {
	LabAccountName    pulumi.StringInput
	LabName           pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	MaxUsersInLab     pulumi.IntPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	UniqueIdentifier  pulumi.StringPtrInput
	UsageQuota        pulumi.StringPtrInput
	UserAccessMode    pulumi.StringPtrInput
}

func (LabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labArgs)(nil)).Elem()
}

type LabInput interface {
	pulumi.Input

	ToLabOutput() LabOutput
	ToLabOutputWithContext(ctx context.Context) LabOutput
}

func (*Lab) ElementType() reflect.Type {
	return reflect.TypeOf((*Lab)(nil))
}

func (i *Lab) ToLabOutput() LabOutput {
	return i.ToLabOutputWithContext(context.Background())
}

func (i *Lab) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabOutput)
}

type LabOutput struct{ *pulumi.OutputState }

func (LabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Lab)(nil))
}

func (o LabOutput) ToLabOutput() LabOutput {
	return o
}

func (o LabOutput) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
