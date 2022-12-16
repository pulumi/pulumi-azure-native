


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IncidentTask struct {
	pulumi.CustomResourceState

	CreatedBy           ClientInfoResponsePtrOutput `pulumi:"createdBy"`
	CreatedTimeUtc      pulumi.StringOutput         `pulumi:"createdTimeUtc"`
	Description         pulumi.StringPtrOutput      `pulumi:"description"`
	Etag                pulumi.StringPtrOutput      `pulumi:"etag"`
	LastModifiedBy      ClientInfoResponsePtrOutput `pulumi:"lastModifiedBy"`
	LastModifiedTimeUtc pulumi.StringOutput         `pulumi:"lastModifiedTimeUtc"`
	Name                pulumi.StringOutput         `pulumi:"name"`
	Status              pulumi.StringOutput         `pulumi:"status"`
	SystemData          SystemDataResponseOutput    `pulumi:"systemData"`
	Title               pulumi.StringOutput         `pulumi:"title"`
	Type                pulumi.StringOutput         `pulumi:"type"`
}


func NewIncidentTask(ctx *pulumi.Context,
	name string, args *IncidentTaskArgs, opts ...pulumi.ResourceOption) (*IncidentTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IncidentId == nil {
		return nil, errors.New("invalid value for required argument 'IncidentId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	var resource IncidentTask
	err := ctx.RegisterResource("azure-native:securityinsights/v20221201preview:IncidentTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIncidentTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentTaskState, opts ...pulumi.ResourceOption) (*IncidentTask, error) {
	var resource IncidentTask
	err := ctx.ReadResource("azure-native:securityinsights/v20221201preview:IncidentTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type incidentTaskState struct {
}

type IncidentTaskState struct {
}

func (IncidentTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentTaskState)(nil)).Elem()
}

type incidentTaskArgs struct {
	CreatedBy         *ClientInfo `pulumi:"createdBy"`
	Description       *string     `pulumi:"description"`
	IncidentId        string      `pulumi:"incidentId"`
	IncidentTaskId    *string     `pulumi:"incidentTaskId"`
	LastModifiedBy    *ClientInfo `pulumi:"lastModifiedBy"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	Status            string      `pulumi:"status"`
	Title             string      `pulumi:"title"`
	WorkspaceName     string      `pulumi:"workspaceName"`
}


type IncidentTaskArgs struct {
	CreatedBy         ClientInfoPtrInput
	Description       pulumi.StringPtrInput
	IncidentId        pulumi.StringInput
	IncidentTaskId    pulumi.StringPtrInput
	LastModifiedBy    ClientInfoPtrInput
	ResourceGroupName pulumi.StringInput
	Status            pulumi.StringInput
	Title             pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (IncidentTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentTaskArgs)(nil)).Elem()
}

type IncidentTaskInput interface {
	pulumi.Input

	ToIncidentTaskOutput() IncidentTaskOutput
	ToIncidentTaskOutputWithContext(ctx context.Context) IncidentTaskOutput
}

func (*IncidentTask) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentTask)(nil)).Elem()
}

func (i *IncidentTask) ToIncidentTaskOutput() IncidentTaskOutput {
	return i.ToIncidentTaskOutputWithContext(context.Background())
}

func (i *IncidentTask) ToIncidentTaskOutputWithContext(ctx context.Context) IncidentTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentTaskOutput)
}

type IncidentTaskOutput struct{ *pulumi.OutputState }

func (IncidentTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentTask)(nil)).Elem()
}

func (o IncidentTaskOutput) ToIncidentTaskOutput() IncidentTaskOutput {
	return o
}

func (o IncidentTaskOutput) ToIncidentTaskOutputWithContext(ctx context.Context) IncidentTaskOutput {
	return o
}

func (o IncidentTaskOutput) CreatedBy() ClientInfoResponsePtrOutput {
	return o.ApplyT(func(v *IncidentTask) ClientInfoResponsePtrOutput { return v.CreatedBy }).(ClientInfoResponsePtrOutput)
}

func (o IncidentTaskOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o IncidentTaskOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IncidentTaskOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IncidentTaskOutput) LastModifiedBy() ClientInfoResponsePtrOutput {
	return o.ApplyT(func(v *IncidentTask) ClientInfoResponsePtrOutput { return v.LastModifiedBy }).(ClientInfoResponsePtrOutput)
}

func (o IncidentTaskOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o IncidentTaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IncidentTaskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o IncidentTaskOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IncidentTask) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o IncidentTaskOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

func (o IncidentTaskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentTaskOutput{})
}
