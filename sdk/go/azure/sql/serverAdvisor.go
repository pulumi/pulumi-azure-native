


package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerAdvisor struct {
	pulumi.CustomResourceState

	AdvisorStatus                  pulumi.StringOutput                  `pulumi:"advisorStatus"`
	AutoExecuteStatus              pulumi.StringOutput                  `pulumi:"autoExecuteStatus"`
	AutoExecuteStatusInheritedFrom pulumi.StringOutput                  `pulumi:"autoExecuteStatusInheritedFrom"`
	Kind                           pulumi.StringOutput                  `pulumi:"kind"`
	LastChecked                    pulumi.StringOutput                  `pulumi:"lastChecked"`
	Location                       pulumi.StringOutput                  `pulumi:"location"`
	Name                           pulumi.StringOutput                  `pulumi:"name"`
	RecommendationsStatus          pulumi.StringOutput                  `pulumi:"recommendationsStatus"`
	RecommendedActions             RecommendedActionResponseArrayOutput `pulumi:"recommendedActions"`
	Type                           pulumi.StringOutput                  `pulumi:"type"`
}


func NewServerAdvisor(ctx *pulumi.Context,
	name string, args *ServerAdvisorArgs, opts ...pulumi.ResourceOption) (*ServerAdvisor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoExecuteStatus == nil {
		return nil, errors.New("invalid value for required argument 'AutoExecuteStatus'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20140401:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerAdvisor"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerAdvisor
	err := ctx.RegisterResource("azure-native:sql:ServerAdvisor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerAdvisor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerAdvisorState, opts ...pulumi.ResourceOption) (*ServerAdvisor, error) {
	var resource ServerAdvisor
	err := ctx.ReadResource("azure-native:sql:ServerAdvisor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverAdvisorState struct {
}

type ServerAdvisorState struct {
}

func (ServerAdvisorState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAdvisorState)(nil)).Elem()
}

type serverAdvisorArgs struct {
	AdvisorName       *string           `pulumi:"advisorName"`
	AutoExecuteStatus AutoExecuteStatus `pulumi:"autoExecuteStatus"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ServerName        string            `pulumi:"serverName"`
}


type ServerAdvisorArgs struct {
	AdvisorName       pulumi.StringPtrInput
	AutoExecuteStatus AutoExecuteStatusInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
}

func (ServerAdvisorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAdvisorArgs)(nil)).Elem()
}

type ServerAdvisorInput interface {
	pulumi.Input

	ToServerAdvisorOutput() ServerAdvisorOutput
	ToServerAdvisorOutputWithContext(ctx context.Context) ServerAdvisorOutput
}

func (*ServerAdvisor) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdvisor)(nil))
}

func (i *ServerAdvisor) ToServerAdvisorOutput() ServerAdvisorOutput {
	return i.ToServerAdvisorOutputWithContext(context.Background())
}

func (i *ServerAdvisor) ToServerAdvisorOutputWithContext(ctx context.Context) ServerAdvisorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdvisorOutput)
}

type ServerAdvisorOutput struct{ *pulumi.OutputState }

func (ServerAdvisorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdvisor)(nil))
}

func (o ServerAdvisorOutput) ToServerAdvisorOutput() ServerAdvisorOutput {
	return o
}

func (o ServerAdvisorOutput) ToServerAdvisorOutputWithContext(ctx context.Context) ServerAdvisorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerAdvisorOutput{})
}
