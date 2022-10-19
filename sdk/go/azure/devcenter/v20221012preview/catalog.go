


package v20221012preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Catalog struct {
	pulumi.CustomResourceState

	AdoGit            GitCatalogResponsePtrOutput `pulumi:"adoGit"`
	GitHub            GitCatalogResponsePtrOutput `pulumi:"gitHub"`
	LastSyncTime      pulumi.StringOutput         `pulumi:"lastSyncTime"`
	Name              pulumi.StringOutput         `pulumi:"name"`
	ProvisioningState pulumi.StringOutput         `pulumi:"provisioningState"`
	SyncState         pulumi.StringOutput         `pulumi:"syncState"`
	SystemData        SystemDataResponseOutput    `pulumi:"systemData"`
	Type              pulumi.StringOutput         `pulumi:"type"`
}


func NewCatalog(ctx *pulumi.Context,
	name string, args *CatalogArgs, opts ...pulumi.ResourceOption) (*Catalog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:Catalog"),
		},
	})
	opts = append(opts, aliases)
	var resource Catalog
	err := ctx.RegisterResource("azure-native:devcenter/v20221012preview:Catalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatalogState, opts ...pulumi.ResourceOption) (*Catalog, error) {
	var resource Catalog
	err := ctx.ReadResource("azure-native:devcenter/v20221012preview:Catalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type catalogState struct {
}

type CatalogState struct {
}

func (CatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogState)(nil)).Elem()
}

type catalogArgs struct {
	AdoGit            *GitCatalog `pulumi:"adoGit"`
	CatalogName       *string     `pulumi:"catalogName"`
	DevCenterName     string      `pulumi:"devCenterName"`
	GitHub            *GitCatalog `pulumi:"gitHub"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
}


type CatalogArgs struct {
	AdoGit            GitCatalogPtrInput
	CatalogName       pulumi.StringPtrInput
	DevCenterName     pulumi.StringInput
	GitHub            GitCatalogPtrInput
	ResourceGroupName pulumi.StringInput
}

func (CatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogArgs)(nil)).Elem()
}

type CatalogInput interface {
	pulumi.Input

	ToCatalogOutput() CatalogOutput
	ToCatalogOutputWithContext(ctx context.Context) CatalogOutput
}

func (*Catalog) ElementType() reflect.Type {
	return reflect.TypeOf((**Catalog)(nil)).Elem()
}

func (i *Catalog) ToCatalogOutput() CatalogOutput {
	return i.ToCatalogOutputWithContext(context.Background())
}

func (i *Catalog) ToCatalogOutputWithContext(ctx context.Context) CatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogOutput)
}

type CatalogOutput struct{ *pulumi.OutputState }

func (CatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Catalog)(nil)).Elem()
}

func (o CatalogOutput) ToCatalogOutput() CatalogOutput {
	return o
}

func (o CatalogOutput) ToCatalogOutputWithContext(ctx context.Context) CatalogOutput {
	return o
}

func (o CatalogOutput) AdoGit() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v *Catalog) GitCatalogResponsePtrOutput { return v.AdoGit }).(GitCatalogResponsePtrOutput)
}

func (o CatalogOutput) GitHub() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v *Catalog) GitCatalogResponsePtrOutput { return v.GitHub }).(GitCatalogResponsePtrOutput)
}

func (o CatalogOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.LastSyncTime }).(pulumi.StringOutput)
}

func (o CatalogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CatalogOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CatalogOutput) SyncState() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.SyncState }).(pulumi.StringOutput)
}

func (o CatalogOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Catalog) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CatalogOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CatalogOutput{})
}
