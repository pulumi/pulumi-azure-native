


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadClassifier struct {
	pulumi.CustomResourceState

	Context    pulumi.StringPtrOutput `pulumi:"context"`
	EndTime    pulumi.StringPtrOutput `pulumi:"endTime"`
	Importance pulumi.StringPtrOutput `pulumi:"importance"`
	Label      pulumi.StringPtrOutput `pulumi:"label"`
	MemberName pulumi.StringOutput    `pulumi:"memberName"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	StartTime  pulumi.StringPtrOutput `pulumi:"startTime"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewWorkloadClassifier(ctx *pulumi.Context,
	name string, args *WorkloadClassifierArgs, opts ...pulumi.ResourceOption) (*WorkloadClassifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.MemberName == nil {
		return nil, errors.New("invalid value for required argument 'MemberName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.WorkloadGroupName == nil {
		return nil, errors.New("invalid value for required argument 'WorkloadGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:WorkloadClassifier"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadClassifier
	err := ctx.RegisterResource("azure-native:sql/v20210201preview:WorkloadClassifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadClassifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadClassifierState, opts ...pulumi.ResourceOption) (*WorkloadClassifier, error) {
	var resource WorkloadClassifier
	err := ctx.ReadResource("azure-native:sql/v20210201preview:WorkloadClassifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workloadClassifierState struct {
}

type WorkloadClassifierState struct {
}

func (WorkloadClassifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadClassifierState)(nil)).Elem()
}

type workloadClassifierArgs struct {
	Context                *string `pulumi:"context"`
	DatabaseName           string  `pulumi:"databaseName"`
	EndTime                *string `pulumi:"endTime"`
	Importance             *string `pulumi:"importance"`
	Label                  *string `pulumi:"label"`
	MemberName             string  `pulumi:"memberName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ServerName             string  `pulumi:"serverName"`
	StartTime              *string `pulumi:"startTime"`
	WorkloadClassifierName *string `pulumi:"workloadClassifierName"`
	WorkloadGroupName      string  `pulumi:"workloadGroupName"`
}


type WorkloadClassifierArgs struct {
	Context                pulumi.StringPtrInput
	DatabaseName           pulumi.StringInput
	EndTime                pulumi.StringPtrInput
	Importance             pulumi.StringPtrInput
	Label                  pulumi.StringPtrInput
	MemberName             pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	ServerName             pulumi.StringInput
	StartTime              pulumi.StringPtrInput
	WorkloadClassifierName pulumi.StringPtrInput
	WorkloadGroupName      pulumi.StringInput
}

func (WorkloadClassifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadClassifierArgs)(nil)).Elem()
}

type WorkloadClassifierInput interface {
	pulumi.Input

	ToWorkloadClassifierOutput() WorkloadClassifierOutput
	ToWorkloadClassifierOutputWithContext(ctx context.Context) WorkloadClassifierOutput
}

func (*WorkloadClassifier) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadClassifier)(nil)).Elem()
}

func (i *WorkloadClassifier) ToWorkloadClassifierOutput() WorkloadClassifierOutput {
	return i.ToWorkloadClassifierOutputWithContext(context.Background())
}

func (i *WorkloadClassifier) ToWorkloadClassifierOutputWithContext(ctx context.Context) WorkloadClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadClassifierOutput)
}

type WorkloadClassifierOutput struct{ *pulumi.OutputState }

func (WorkloadClassifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadClassifier)(nil)).Elem()
}

func (o WorkloadClassifierOutput) ToWorkloadClassifierOutput() WorkloadClassifierOutput {
	return o
}

func (o WorkloadClassifierOutput) ToWorkloadClassifierOutputWithContext(ctx context.Context) WorkloadClassifierOutput {
	return o
}

func (o WorkloadClassifierOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.Context }).(pulumi.StringPtrOutput)
}

func (o WorkloadClassifierOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o WorkloadClassifierOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.Importance }).(pulumi.StringPtrOutput)
}

func (o WorkloadClassifierOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

func (o WorkloadClassifierOutput) MemberName() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringOutput { return v.MemberName }).(pulumi.StringOutput)
}

func (o WorkloadClassifierOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkloadClassifierOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o WorkloadClassifierOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadClassifierOutput{})
}
