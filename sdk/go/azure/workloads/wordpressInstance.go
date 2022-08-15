


package workloads

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WordpressInstance struct {
	pulumi.CustomResourceState

	DatabaseName      pulumi.StringPtrOutput   `pulumi:"databaseName"`
	DatabaseUser      pulumi.StringPtrOutput   `pulumi:"databaseUser"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SiteUrl           pulumi.StringOutput      `pulumi:"siteUrl"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
	Version           pulumi.StringOutput      `pulumi:"version"`
}


func NewWordpressInstance(ctx *pulumi.Context,
	name string, args *WordpressInstanceArgs, opts ...pulumi.ResourceOption) (*WordpressInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PhpWorkloadName == nil {
		return nil, errors.New("invalid value for required argument 'PhpWorkloadName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:WordpressInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource WordpressInstance
	err := ctx.RegisterResource("azure-native:workloads:WordpressInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWordpressInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WordpressInstanceState, opts ...pulumi.ResourceOption) (*WordpressInstance, error) {
	var resource WordpressInstance
	err := ctx.ReadResource("azure-native:workloads:WordpressInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type wordpressInstanceState struct {
}

type WordpressInstanceState struct {
}

func (WordpressInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*wordpressInstanceState)(nil)).Elem()
}

type wordpressInstanceArgs struct {
	DatabaseName      *string `pulumi:"databaseName"`
	DatabaseUser      *string `pulumi:"databaseUser"`
	PhpWorkloadName   string  `pulumi:"phpWorkloadName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Version           string  `pulumi:"version"`
}


type WordpressInstanceArgs struct {
	DatabaseName      pulumi.StringPtrInput
	DatabaseUser      pulumi.StringPtrInput
	PhpWorkloadName   pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Version           pulumi.StringInput
}

func (WordpressInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wordpressInstanceArgs)(nil)).Elem()
}

type WordpressInstanceInput interface {
	pulumi.Input

	ToWordpressInstanceOutput() WordpressInstanceOutput
	ToWordpressInstanceOutputWithContext(ctx context.Context) WordpressInstanceOutput
}

func (*WordpressInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**WordpressInstance)(nil)).Elem()
}

func (i *WordpressInstance) ToWordpressInstanceOutput() WordpressInstanceOutput {
	return i.ToWordpressInstanceOutputWithContext(context.Background())
}

func (i *WordpressInstance) ToWordpressInstanceOutputWithContext(ctx context.Context) WordpressInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WordpressInstanceOutput)
}

type WordpressInstanceOutput struct{ *pulumi.OutputState }

func (WordpressInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WordpressInstance)(nil)).Elem()
}

func (o WordpressInstanceOutput) ToWordpressInstanceOutput() WordpressInstanceOutput {
	return o
}

func (o WordpressInstanceOutput) ToWordpressInstanceOutputWithContext(ctx context.Context) WordpressInstanceOutput {
	return o
}

func (o WordpressInstanceOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WordpressInstance) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o WordpressInstanceOutput) DatabaseUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WordpressInstance) pulumi.StringPtrOutput { return v.DatabaseUser }).(pulumi.StringPtrOutput)
}

func (o WordpressInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WordpressInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WordpressInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WordpressInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WordpressInstanceOutput) SiteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *WordpressInstance) pulumi.StringOutput { return v.SiteUrl }).(pulumi.StringOutput)
}

func (o WordpressInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WordpressInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WordpressInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WordpressInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WordpressInstanceOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *WordpressInstance) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WordpressInstanceOutput{})
}
