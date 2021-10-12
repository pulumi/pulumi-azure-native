


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppDeployment struct {
	pulumi.CustomResourceState

	Active      pulumi.BoolPtrOutput   `pulumi:"active"`
	Author      pulumi.StringPtrOutput `pulumi:"author"`
	AuthorEmail pulumi.StringPtrOutput `pulumi:"authorEmail"`
	Deployer    pulumi.StringPtrOutput `pulumi:"deployer"`
	Details     pulumi.StringPtrOutput `pulumi:"details"`
	EndTime     pulumi.StringPtrOutput `pulumi:"endTime"`
	Kind        pulumi.StringPtrOutput `pulumi:"kind"`
	Message     pulumi.StringPtrOutput `pulumi:"message"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	StartTime   pulumi.StringPtrOutput `pulumi:"startTime"`
	Status      pulumi.IntPtrOutput    `pulumi:"status"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppDeployment(ctx *pulumi.Context,
	name string, args *WebAppDeploymentArgs, opts ...pulumi.ResourceOption) (*WebAppDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:WebAppDeployment"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppDeployment
	err := ctx.RegisterResource("azure-native:web/v20210101:WebAppDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDeploymentState, opts ...pulumi.ResourceOption) (*WebAppDeployment, error) {
	var resource WebAppDeployment
	err := ctx.ReadResource("azure-native:web/v20210101:WebAppDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppDeploymentState struct {
}

type WebAppDeploymentState struct {
}

func (WebAppDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDeploymentState)(nil)).Elem()
}

type webAppDeploymentArgs struct {
	Active            *bool   `pulumi:"active"`
	Author            *string `pulumi:"author"`
	AuthorEmail       *string `pulumi:"authorEmail"`
	Deployer          *string `pulumi:"deployer"`
	Details           *string `pulumi:"details"`
	EndTime           *string `pulumi:"endTime"`
	Id                *string `pulumi:"id"`
	Kind              *string `pulumi:"kind"`
	Message           *string `pulumi:"message"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	StartTime         *string `pulumi:"startTime"`
	Status            *int    `pulumi:"status"`
}


type WebAppDeploymentArgs struct {
	Active            pulumi.BoolPtrInput
	Author            pulumi.StringPtrInput
	AuthorEmail       pulumi.StringPtrInput
	Deployer          pulumi.StringPtrInput
	Details           pulumi.StringPtrInput
	EndTime           pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Message           pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	StartTime         pulumi.StringPtrInput
	Status            pulumi.IntPtrInput
}

func (WebAppDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDeploymentArgs)(nil)).Elem()
}

type WebAppDeploymentInput interface {
	pulumi.Input

	ToWebAppDeploymentOutput() WebAppDeploymentOutput
	ToWebAppDeploymentOutputWithContext(ctx context.Context) WebAppDeploymentOutput
}

func (*WebAppDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppDeployment)(nil))
}

func (i *WebAppDeployment) ToWebAppDeploymentOutput() WebAppDeploymentOutput {
	return i.ToWebAppDeploymentOutputWithContext(context.Background())
}

func (i *WebAppDeployment) ToWebAppDeploymentOutputWithContext(ctx context.Context) WebAppDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDeploymentOutput)
}

type WebAppDeploymentOutput struct{ *pulumi.OutputState }

func (WebAppDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppDeployment)(nil))
}

func (o WebAppDeploymentOutput) ToWebAppDeploymentOutput() WebAppDeploymentOutput {
	return o
}

func (o WebAppDeploymentOutput) ToWebAppDeploymentOutputWithContext(ctx context.Context) WebAppDeploymentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppDeploymentOutput{})
}
