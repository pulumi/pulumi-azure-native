


package v20210501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlPoolWorkloadClassifier struct {
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


func NewSqlPoolWorkloadClassifier(ctx *pulumi.Context,
	name string, args *SqlPoolWorkloadClassifierArgs, opts ...pulumi.ResourceOption) (*SqlPoolWorkloadClassifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MemberName == nil {
		return nil, errors.New("invalid value for required argument 'MemberName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlPoolName == nil {
		return nil, errors.New("invalid value for required argument 'SqlPoolName'")
	}
	if args.WorkloadGroupName == nil {
		return nil, errors.New("invalid value for required argument 'WorkloadGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210501:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:synapse:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20190601preview:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20201201:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210301:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210401preview:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210601:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:SqlPoolWorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210601preview:SqlPoolWorkloadClassifier"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlPoolWorkloadClassifier
	err := ctx.RegisterResource("azure-native:synapse/v20210501:SqlPoolWorkloadClassifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlPoolWorkloadClassifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolWorkloadClassifierState, opts ...pulumi.ResourceOption) (*SqlPoolWorkloadClassifier, error) {
	var resource SqlPoolWorkloadClassifier
	err := ctx.ReadResource("azure-native:synapse/v20210501:SqlPoolWorkloadClassifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlPoolWorkloadClassifierState struct {
}

type SqlPoolWorkloadClassifierState struct {
}

func (SqlPoolWorkloadClassifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolWorkloadClassifierState)(nil)).Elem()
}

type sqlPoolWorkloadClassifierArgs struct {
	Context                *string `pulumi:"context"`
	EndTime                *string `pulumi:"endTime"`
	Importance             *string `pulumi:"importance"`
	Label                  *string `pulumi:"label"`
	MemberName             string  `pulumi:"memberName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	SqlPoolName            string  `pulumi:"sqlPoolName"`
	StartTime              *string `pulumi:"startTime"`
	WorkloadClassifierName *string `pulumi:"workloadClassifierName"`
	WorkloadGroupName      string  `pulumi:"workloadGroupName"`
	WorkspaceName          string  `pulumi:"workspaceName"`
}


type SqlPoolWorkloadClassifierArgs struct {
	Context                pulumi.StringPtrInput
	EndTime                pulumi.StringPtrInput
	Importance             pulumi.StringPtrInput
	Label                  pulumi.StringPtrInput
	MemberName             pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	SqlPoolName            pulumi.StringInput
	StartTime              pulumi.StringPtrInput
	WorkloadClassifierName pulumi.StringPtrInput
	WorkloadGroupName      pulumi.StringInput
	WorkspaceName          pulumi.StringInput
}

func (SqlPoolWorkloadClassifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolWorkloadClassifierArgs)(nil)).Elem()
}

type SqlPoolWorkloadClassifierInput interface {
	pulumi.Input

	ToSqlPoolWorkloadClassifierOutput() SqlPoolWorkloadClassifierOutput
	ToSqlPoolWorkloadClassifierOutputWithContext(ctx context.Context) SqlPoolWorkloadClassifierOutput
}

func (*SqlPoolWorkloadClassifier) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolWorkloadClassifier)(nil))
}

func (i *SqlPoolWorkloadClassifier) ToSqlPoolWorkloadClassifierOutput() SqlPoolWorkloadClassifierOutput {
	return i.ToSqlPoolWorkloadClassifierOutputWithContext(context.Background())
}

func (i *SqlPoolWorkloadClassifier) ToSqlPoolWorkloadClassifierOutputWithContext(ctx context.Context) SqlPoolWorkloadClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolWorkloadClassifierOutput)
}

type SqlPoolWorkloadClassifierOutput struct{ *pulumi.OutputState }

func (SqlPoolWorkloadClassifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolWorkloadClassifier)(nil))
}

func (o SqlPoolWorkloadClassifierOutput) ToSqlPoolWorkloadClassifierOutput() SqlPoolWorkloadClassifierOutput {
	return o
}

func (o SqlPoolWorkloadClassifierOutput) ToSqlPoolWorkloadClassifierOutputWithContext(ctx context.Context) SqlPoolWorkloadClassifierOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlPoolWorkloadClassifierInput)(nil)).Elem(), &SqlPoolWorkloadClassifier{})
	pulumi.RegisterOutputType(SqlPoolWorkloadClassifierOutput{})
}
