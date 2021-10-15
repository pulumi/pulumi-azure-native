


package v20210622

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HybridRunbookWorkerGroup struct {
	pulumi.CustomResourceState

	Credential           RunAsCredentialAssociationPropertyResponsePtrOutput `pulumi:"credential"`
	GroupType            pulumi.StringPtrOutput                              `pulumi:"groupType"`
	HybridRunbookWorkers HybridRunbookWorkerLegacyResponseArrayOutput        `pulumi:"hybridRunbookWorkers"`
	Name                 pulumi.StringPtrOutput                              `pulumi:"name"`
	SystemData           SystemDataResponseOutput                            `pulumi:"systemData"`
	Type                 pulumi.StringOutput                                 `pulumi:"type"`
}


func NewHybridRunbookWorkerGroup(ctx *pulumi.Context,
	name string, args *HybridRunbookWorkerGroupArgs, opts ...pulumi.ResourceOption) (*HybridRunbookWorkerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation/v20210622:HybridRunbookWorkerGroup"),
		},
		{
			Type: pulumi.String("azure-native:automation:HybridRunbookWorkerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation:HybridRunbookWorkerGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridRunbookWorkerGroup
	err := ctx.RegisterResource("azure-native:automation/v20210622:HybridRunbookWorkerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHybridRunbookWorkerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridRunbookWorkerGroupState, opts ...pulumi.ResourceOption) (*HybridRunbookWorkerGroup, error) {
	var resource HybridRunbookWorkerGroup
	err := ctx.ReadResource("azure-native:automation/v20210622:HybridRunbookWorkerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hybridRunbookWorkerGroupState struct {
}

type HybridRunbookWorkerGroupState struct {
}

func (HybridRunbookWorkerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridRunbookWorkerGroupState)(nil)).Elem()
}

type hybridRunbookWorkerGroupArgs struct {
	AutomationAccountName        string                              `pulumi:"automationAccountName"`
	Credential                   *RunAsCredentialAssociationProperty `pulumi:"credential"`
	HybridRunbookWorkerGroupName *string                             `pulumi:"hybridRunbookWorkerGroupName"`
	ResourceGroupName            string                              `pulumi:"resourceGroupName"`
}


type HybridRunbookWorkerGroupArgs struct {
	AutomationAccountName        pulumi.StringInput
	Credential                   RunAsCredentialAssociationPropertyPtrInput
	HybridRunbookWorkerGroupName pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
}

func (HybridRunbookWorkerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridRunbookWorkerGroupArgs)(nil)).Elem()
}

type HybridRunbookWorkerGroupInput interface {
	pulumi.Input

	ToHybridRunbookWorkerGroupOutput() HybridRunbookWorkerGroupOutput
	ToHybridRunbookWorkerGroupOutputWithContext(ctx context.Context) HybridRunbookWorkerGroupOutput
}

func (*HybridRunbookWorkerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridRunbookWorkerGroup)(nil))
}

func (i *HybridRunbookWorkerGroup) ToHybridRunbookWorkerGroupOutput() HybridRunbookWorkerGroupOutput {
	return i.ToHybridRunbookWorkerGroupOutputWithContext(context.Background())
}

func (i *HybridRunbookWorkerGroup) ToHybridRunbookWorkerGroupOutputWithContext(ctx context.Context) HybridRunbookWorkerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridRunbookWorkerGroupOutput)
}

type HybridRunbookWorkerGroupOutput struct{ *pulumi.OutputState }

func (HybridRunbookWorkerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridRunbookWorkerGroup)(nil))
}

func (o HybridRunbookWorkerGroupOutput) ToHybridRunbookWorkerGroupOutput() HybridRunbookWorkerGroupOutput {
	return o
}

func (o HybridRunbookWorkerGroupOutput) ToHybridRunbookWorkerGroupOutputWithContext(ctx context.Context) HybridRunbookWorkerGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HybridRunbookWorkerGroupOutput{})
}
