


package v20150501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Favorite struct {
	pulumi.CustomResourceState

	Category                pulumi.StringPtrOutput   `pulumi:"category"`
	Config                  pulumi.StringPtrOutput   `pulumi:"config"`
	FavoriteId              pulumi.StringOutput      `pulumi:"favoriteId"`
	FavoriteType            pulumi.StringPtrOutput   `pulumi:"favoriteType"`
	IsGeneratedFromTemplate pulumi.BoolPtrOutput     `pulumi:"isGeneratedFromTemplate"`
	Name                    pulumi.StringPtrOutput   `pulumi:"name"`
	SourceType              pulumi.StringPtrOutput   `pulumi:"sourceType"`
	Tags                    pulumi.StringArrayOutput `pulumi:"tags"`
	TimeModified            pulumi.StringOutput      `pulumi:"timeModified"`
	UserId                  pulumi.StringOutput      `pulumi:"userId"`
	Version                 pulumi.StringPtrOutput   `pulumi:"version"`
}


func NewFavorite(ctx *pulumi.Context,
	name string, args *FavoriteArgs, opts ...pulumi.ResourceOption) (*Favorite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:Favorite"),
		},
	})
	opts = append(opts, aliases)
	var resource Favorite
	err := ctx.RegisterResource("azure-native:insights/v20150501:Favorite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFavorite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FavoriteState, opts ...pulumi.ResourceOption) (*Favorite, error) {
	var resource Favorite
	err := ctx.ReadResource("azure-native:insights/v20150501:Favorite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type favoriteState struct {
}

type FavoriteState struct {
}

func (FavoriteState) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteState)(nil)).Elem()
}

type favoriteArgs struct {
	Category                *string       `pulumi:"category"`
	Config                  *string       `pulumi:"config"`
	FavoriteId              *string       `pulumi:"favoriteId"`
	FavoriteType            *FavoriteType `pulumi:"favoriteType"`
	IsGeneratedFromTemplate *bool         `pulumi:"isGeneratedFromTemplate"`
	Name                    *string       `pulumi:"name"`
	ResourceGroupName       string        `pulumi:"resourceGroupName"`
	ResourceName            string        `pulumi:"resourceName"`
	SourceType              *string       `pulumi:"sourceType"`
	Tags                    []string      `pulumi:"tags"`
	Version                 *string       `pulumi:"version"`
}


type FavoriteArgs struct {
	Category                pulumi.StringPtrInput
	Config                  pulumi.StringPtrInput
	FavoriteId              pulumi.StringPtrInput
	FavoriteType            FavoriteTypePtrInput
	IsGeneratedFromTemplate pulumi.BoolPtrInput
	Name                    pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ResourceName            pulumi.StringInput
	SourceType              pulumi.StringPtrInput
	Tags                    pulumi.StringArrayInput
	Version                 pulumi.StringPtrInput
}

func (FavoriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteArgs)(nil)).Elem()
}

type FavoriteInput interface {
	pulumi.Input

	ToFavoriteOutput() FavoriteOutput
	ToFavoriteOutputWithContext(ctx context.Context) FavoriteOutput
}

func (*Favorite) ElementType() reflect.Type {
	return reflect.TypeOf((**Favorite)(nil)).Elem()
}

func (i *Favorite) ToFavoriteOutput() FavoriteOutput {
	return i.ToFavoriteOutputWithContext(context.Background())
}

func (i *Favorite) ToFavoriteOutputWithContext(ctx context.Context) FavoriteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FavoriteOutput)
}

type FavoriteOutput struct{ *pulumi.OutputState }

func (FavoriteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Favorite)(nil)).Elem()
}

func (o FavoriteOutput) ToFavoriteOutput() FavoriteOutput {
	return o
}

func (o FavoriteOutput) ToFavoriteOutputWithContext(ctx context.Context) FavoriteOutput {
	return o
}

func (o FavoriteOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.Category }).(pulumi.StringPtrOutput)
}

func (o FavoriteOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.Config }).(pulumi.StringPtrOutput)
}

func (o FavoriteOutput) FavoriteId() pulumi.StringOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringOutput { return v.FavoriteId }).(pulumi.StringOutput)
}

func (o FavoriteOutput) FavoriteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.FavoriteType }).(pulumi.StringPtrOutput)
}

func (o FavoriteOutput) IsGeneratedFromTemplate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.BoolPtrOutput { return v.IsGeneratedFromTemplate }).(pulumi.BoolPtrOutput)
}

func (o FavoriteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FavoriteOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

func (o FavoriteOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o FavoriteOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

func (o FavoriteOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

func (o FavoriteOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FavoriteOutput{})
}
