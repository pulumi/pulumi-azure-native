


package v20210622

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HybridRunbookWorker struct {
	pulumi.CustomResourceState

	Ip                 pulumi.StringPtrOutput   `pulumi:"ip"`
	LastSeenDateTime   pulumi.StringPtrOutput   `pulumi:"lastSeenDateTime"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	RegisteredDateTime pulumi.StringPtrOutput   `pulumi:"registeredDateTime"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
	VmResourceId       pulumi.StringPtrOutput   `pulumi:"vmResourceId"`
	WorkerName         pulumi.StringPtrOutput   `pulumi:"workerName"`
	WorkerType         pulumi.StringPtrOutput   `pulumi:"workerType"`
}


func NewHybridRunbookWorker(ctx *pulumi.Context,
	name string, args *HybridRunbookWorkerArgs, opts ...pulumi.ResourceOption) (*HybridRunbookWorker, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.HybridRunbookWorkerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'HybridRunbookWorkerGroupName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:HybridRunbookWorker"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridRunbookWorker
	err := ctx.RegisterResource("azure-native:automation/v20210622:HybridRunbookWorker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHybridRunbookWorker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridRunbookWorkerState, opts ...pulumi.ResourceOption) (*HybridRunbookWorker, error) {
	var resource HybridRunbookWorker
	err := ctx.ReadResource("azure-native:automation/v20210622:HybridRunbookWorker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hybridRunbookWorkerState struct {
}

type HybridRunbookWorkerState struct {
}

func (HybridRunbookWorkerState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridRunbookWorkerState)(nil)).Elem()
}

type hybridRunbookWorkerArgs struct {
	AutomationAccountName        string  `pulumi:"automationAccountName"`
	HybridRunbookWorkerGroupName string  `pulumi:"hybridRunbookWorkerGroupName"`
	HybridRunbookWorkerId        *string `pulumi:"hybridRunbookWorkerId"`
	Name                         *string `pulumi:"name"`
	ResourceGroupName            string  `pulumi:"resourceGroupName"`
	VmResourceId                 *string `pulumi:"vmResourceId"`
}


type HybridRunbookWorkerArgs struct {
	AutomationAccountName        pulumi.StringInput
	HybridRunbookWorkerGroupName pulumi.StringInput
	HybridRunbookWorkerId        pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	VmResourceId                 pulumi.StringPtrInput
}

func (HybridRunbookWorkerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridRunbookWorkerArgs)(nil)).Elem()
}

type HybridRunbookWorkerInput interface {
	pulumi.Input

	ToHybridRunbookWorkerOutput() HybridRunbookWorkerOutput
	ToHybridRunbookWorkerOutputWithContext(ctx context.Context) HybridRunbookWorkerOutput
}

func (*HybridRunbookWorker) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridRunbookWorker)(nil)).Elem()
}

func (i *HybridRunbookWorker) ToHybridRunbookWorkerOutput() HybridRunbookWorkerOutput {
	return i.ToHybridRunbookWorkerOutputWithContext(context.Background())
}

func (i *HybridRunbookWorker) ToHybridRunbookWorkerOutputWithContext(ctx context.Context) HybridRunbookWorkerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridRunbookWorkerOutput)
}

type HybridRunbookWorkerOutput struct{ *pulumi.OutputState }

func (HybridRunbookWorkerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridRunbookWorker)(nil)).Elem()
}

func (o HybridRunbookWorkerOutput) ToHybridRunbookWorkerOutput() HybridRunbookWorkerOutput {
	return o
}

func (o HybridRunbookWorkerOutput) ToHybridRunbookWorkerOutputWithContext(ctx context.Context) HybridRunbookWorkerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HybridRunbookWorkerOutput{})
}
