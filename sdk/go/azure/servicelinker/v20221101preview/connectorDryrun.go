


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectorDryrun struct {
	pulumi.CustomResourceState

	Name                pulumi.StringOutput                             `pulumi:"name"`
	OperationPreviews   DryrunOperationPreviewResponseArrayOutput       `pulumi:"operationPreviews"`
	Parameters          CreateOrUpdateDryrunParametersResponsePtrOutput `pulumi:"parameters"`
	PrerequisiteResults pulumi.ArrayOutput                              `pulumi:"prerequisiteResults"`
	ProvisioningState   pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData          SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                pulumi.StringOutput                             `pulumi:"type"`
}


func NewConnectorDryrun(ctx *pulumi.Context,
	name string, args *ConnectorDryrunArgs, opts ...pulumi.ResourceOption) (*ConnectorDryrun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ConnectorDryrun
	err := ctx.RegisterResource("azure-native:servicelinker/v20221101preview:ConnectorDryrun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectorDryrun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorDryrunState, opts ...pulumi.ResourceOption) (*ConnectorDryrun, error) {
	var resource ConnectorDryrun
	err := ctx.ReadResource("azure-native:servicelinker/v20221101preview:ConnectorDryrun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectorDryrunState struct {
}

type ConnectorDryrunState struct {
}

func (ConnectorDryrunState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorDryrunState)(nil)).Elem()
}

type connectorDryrunArgs struct {
	DryrunName        *string                         `pulumi:"dryrunName"`
	Location          string                          `pulumi:"location"`
	Parameters        *CreateOrUpdateDryrunParameters `pulumi:"parameters"`
	ResourceGroupName string                          `pulumi:"resourceGroupName"`
	SubscriptionId    *string                         `pulumi:"subscriptionId"`
}


type ConnectorDryrunArgs struct {
	DryrunName        pulumi.StringPtrInput
	Location          pulumi.StringInput
	Parameters        CreateOrUpdateDryrunParametersPtrInput
	ResourceGroupName pulumi.StringInput
	SubscriptionId    pulumi.StringPtrInput
}

func (ConnectorDryrunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorDryrunArgs)(nil)).Elem()
}

type ConnectorDryrunInput interface {
	pulumi.Input

	ToConnectorDryrunOutput() ConnectorDryrunOutput
	ToConnectorDryrunOutputWithContext(ctx context.Context) ConnectorDryrunOutput
}

func (*ConnectorDryrun) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorDryrun)(nil)).Elem()
}

func (i *ConnectorDryrun) ToConnectorDryrunOutput() ConnectorDryrunOutput {
	return i.ToConnectorDryrunOutputWithContext(context.Background())
}

func (i *ConnectorDryrun) ToConnectorDryrunOutputWithContext(ctx context.Context) ConnectorDryrunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorDryrunOutput)
}

type ConnectorDryrunOutput struct{ *pulumi.OutputState }

func (ConnectorDryrunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorDryrun)(nil)).Elem()
}

func (o ConnectorDryrunOutput) ToConnectorDryrunOutput() ConnectorDryrunOutput {
	return o
}

func (o ConnectorDryrunOutput) ToConnectorDryrunOutputWithContext(ctx context.Context) ConnectorDryrunOutput {
	return o
}

func (o ConnectorDryrunOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorDryrun) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectorDryrunOutput) OperationPreviews() DryrunOperationPreviewResponseArrayOutput {
	return o.ApplyT(func(v *ConnectorDryrun) DryrunOperationPreviewResponseArrayOutput { return v.OperationPreviews }).(DryrunOperationPreviewResponseArrayOutput)
}

func (o ConnectorDryrunOutput) Parameters() CreateOrUpdateDryrunParametersResponsePtrOutput {
	return o.ApplyT(func(v *ConnectorDryrun) CreateOrUpdateDryrunParametersResponsePtrOutput { return v.Parameters }).(CreateOrUpdateDryrunParametersResponsePtrOutput)
}

func (o ConnectorDryrunOutput) PrerequisiteResults() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ConnectorDryrun) pulumi.ArrayOutput { return v.PrerequisiteResults }).(pulumi.ArrayOutput)
}

func (o ConnectorDryrunOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorDryrun) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConnectorDryrunOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectorDryrun) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConnectorDryrunOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorDryrun) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectorDryrunOutput{})
}
