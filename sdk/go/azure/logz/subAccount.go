


package logz

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubAccount struct {
	pulumi.CustomResourceState

	Identity   IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties MonitorPropertiesResponseOutput     `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewSubAccount(ctx *pulumi.Context,
	name string, args *SubAccountArgs, opts ...pulumi.ResourceOption) (*SubAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logz:SubAccount"),
		},
		{
			Type: pulumi.String("azure-native:logz/v20201001:SubAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:logz/v20201001:SubAccount"),
		},
		{
			Type: pulumi.String("azure-native:logz/v20201001preview:SubAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:logz/v20201001preview:SubAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource SubAccount
	err := ctx.RegisterResource("azure-native:logz:SubAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubAccountState, opts ...pulumi.ResourceOption) (*SubAccount, error) {
	var resource SubAccount
	err := ctx.ReadResource("azure-native:logz:SubAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subAccountState struct {
}

type SubAccountState struct {
}

func (SubAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountState)(nil)).Elem()
}

type subAccountArgs struct {
	Identity          *IdentityProperties `pulumi:"identity"`
	Location          *string             `pulumi:"location"`
	MonitorName       string              `pulumi:"monitorName"`
	Properties        *MonitorProperties  `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	SubAccountName    *string             `pulumi:"subAccountName"`
	Tags              map[string]string   `pulumi:"tags"`
}


type SubAccountArgs struct {
	Identity          IdentityPropertiesPtrInput
	Location          pulumi.StringPtrInput
	MonitorName       pulumi.StringInput
	Properties        MonitorPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	SubAccountName    pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (SubAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountArgs)(nil)).Elem()
}

type SubAccountInput interface {
	pulumi.Input

	ToSubAccountOutput() SubAccountOutput
	ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput
}

func (*SubAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*SubAccount)(nil))
}

func (i *SubAccount) ToSubAccountOutput() SubAccountOutput {
	return i.ToSubAccountOutputWithContext(context.Background())
}

func (i *SubAccount) ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubAccountOutput)
}

type SubAccountOutput struct{ *pulumi.OutputState }

func (SubAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubAccount)(nil))
}

func (o SubAccountOutput) ToSubAccountOutput() SubAccountOutput {
	return o
}

func (o SubAccountOutput) ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubAccountInput)(nil)).Elem(), &SubAccount{})
	pulumi.RegisterOutputType(SubAccountOutput{})
}
