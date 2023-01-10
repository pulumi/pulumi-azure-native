


package security

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityConnectorApplication struct {
	pulumi.CustomResourceState

	Description        pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName        pulumi.StringPtrOutput `pulumi:"displayName"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	SourceResourceType pulumi.StringOutput    `pulumi:"sourceResourceType"`
	Type               pulumi.StringOutput    `pulumi:"type"`
}


func NewSecurityConnectorApplication(ctx *pulumi.Context,
	name string, args *SecurityConnectorApplicationArgs, opts ...pulumi.ResourceOption) (*SecurityConnectorApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SecurityConnectorName == nil {
		return nil, errors.New("invalid value for required argument 'SecurityConnectorName'")
	}
	if args.SourceResourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20220701preview:SecurityConnectorApplication"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityConnectorApplication
	err := ctx.RegisterResource("azure-native:security:SecurityConnectorApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityConnectorApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityConnectorApplicationState, opts ...pulumi.ResourceOption) (*SecurityConnectorApplication, error) {
	var resource SecurityConnectorApplication
	err := ctx.ReadResource("azure-native:security:SecurityConnectorApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityConnectorApplicationState struct {
}

type SecurityConnectorApplicationState struct {
}

func (SecurityConnectorApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorApplicationState)(nil)).Elem()
}

type securityConnectorApplicationArgs struct {
	ApplicationId         *string `pulumi:"applicationId"`
	Description           *string `pulumi:"description"`
	DisplayName           *string `pulumi:"displayName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	SecurityConnectorName string  `pulumi:"securityConnectorName"`
	SourceResourceType    string  `pulumi:"sourceResourceType"`
}


type SecurityConnectorApplicationArgs struct {
	ApplicationId         pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	DisplayName           pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	SecurityConnectorName pulumi.StringInput
	SourceResourceType    pulumi.StringInput
}

func (SecurityConnectorApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorApplicationArgs)(nil)).Elem()
}

type SecurityConnectorApplicationInput interface {
	pulumi.Input

	ToSecurityConnectorApplicationOutput() SecurityConnectorApplicationOutput
	ToSecurityConnectorApplicationOutputWithContext(ctx context.Context) SecurityConnectorApplicationOutput
}

func (*SecurityConnectorApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorApplication)(nil)).Elem()
}

func (i *SecurityConnectorApplication) ToSecurityConnectorApplicationOutput() SecurityConnectorApplicationOutput {
	return i.ToSecurityConnectorApplicationOutputWithContext(context.Background())
}

func (i *SecurityConnectorApplication) ToSecurityConnectorApplicationOutputWithContext(ctx context.Context) SecurityConnectorApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorApplicationOutput)
}

type SecurityConnectorApplicationOutput struct{ *pulumi.OutputState }

func (SecurityConnectorApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorApplication)(nil)).Elem()
}

func (o SecurityConnectorApplicationOutput) ToSecurityConnectorApplicationOutput() SecurityConnectorApplicationOutput {
	return o
}

func (o SecurityConnectorApplicationOutput) ToSecurityConnectorApplicationOutputWithContext(ctx context.Context) SecurityConnectorApplicationOutput {
	return o
}

func (o SecurityConnectorApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorApplication) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorApplicationOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorApplication) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecurityConnectorApplicationOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorApplication) pulumi.StringOutput { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o SecurityConnectorApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorApplication) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityConnectorApplicationOutput{})
}
