


package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataMaskingPolicy struct {
	pulumi.CustomResourceState

	ApplicationPrincipals pulumi.StringOutput    `pulumi:"applicationPrincipals"`
	DataMaskingState      pulumi.StringOutput    `pulumi:"dataMaskingState"`
	ExemptPrincipals      pulumi.StringPtrOutput `pulumi:"exemptPrincipals"`
	Kind                  pulumi.StringOutput    `pulumi:"kind"`
	Location              pulumi.StringOutput    `pulumi:"location"`
	MaskingLevel          pulumi.StringOutput    `pulumi:"maskingLevel"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
}


func NewDataMaskingPolicy(ctx *pulumi.Context,
	name string, args *DataMaskingPolicyArgs, opts ...pulumi.ResourceOption) (*DataMaskingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataMaskingState == nil {
		return nil, errors.New("invalid value for required argument 'DataMaskingState'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20140401:DataMaskingPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource DataMaskingPolicy
	err := ctx.RegisterResource("azure-native:sql:DataMaskingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataMaskingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataMaskingPolicyState, opts ...pulumi.ResourceOption) (*DataMaskingPolicy, error) {
	var resource DataMaskingPolicy
	err := ctx.ReadResource("azure-native:sql:DataMaskingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataMaskingPolicyState struct {
}

type DataMaskingPolicyState struct {
}

func (DataMaskingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataMaskingPolicyState)(nil)).Elem()
}

type dataMaskingPolicyArgs struct {
	DataMaskingPolicyName *string          `pulumi:"dataMaskingPolicyName"`
	DataMaskingState      DataMaskingState `pulumi:"dataMaskingState"`
	DatabaseName          string           `pulumi:"databaseName"`
	ExemptPrincipals      *string          `pulumi:"exemptPrincipals"`
	ResourceGroupName     string           `pulumi:"resourceGroupName"`
	ServerName            string           `pulumi:"serverName"`
}


type DataMaskingPolicyArgs struct {
	DataMaskingPolicyName pulumi.StringPtrInput
	DataMaskingState      DataMaskingStateInput
	DatabaseName          pulumi.StringInput
	ExemptPrincipals      pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	ServerName            pulumi.StringInput
}

func (DataMaskingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataMaskingPolicyArgs)(nil)).Elem()
}

type DataMaskingPolicyInput interface {
	pulumi.Input

	ToDataMaskingPolicyOutput() DataMaskingPolicyOutput
	ToDataMaskingPolicyOutputWithContext(ctx context.Context) DataMaskingPolicyOutput
}

func (*DataMaskingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingPolicy)(nil))
}

func (i *DataMaskingPolicy) ToDataMaskingPolicyOutput() DataMaskingPolicyOutput {
	return i.ToDataMaskingPolicyOutputWithContext(context.Background())
}

func (i *DataMaskingPolicy) ToDataMaskingPolicyOutputWithContext(ctx context.Context) DataMaskingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataMaskingPolicyOutput)
}

type DataMaskingPolicyOutput struct{ *pulumi.OutputState }

func (DataMaskingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingPolicy)(nil))
}

func (o DataMaskingPolicyOutput) ToDataMaskingPolicyOutput() DataMaskingPolicyOutput {
	return o
}

func (o DataMaskingPolicyOutput) ToDataMaskingPolicyOutputWithContext(ctx context.Context) DataMaskingPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataMaskingPolicyOutput{})
}
