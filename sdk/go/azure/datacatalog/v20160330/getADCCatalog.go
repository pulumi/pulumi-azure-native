


package v20160330

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADCCatalog(ctx *pulumi.Context, args *LookupADCCatalogArgs, opts ...pulumi.InvokeOption) (*LookupADCCatalogResult, error) {
	var rv LookupADCCatalogResult
	err := ctx.Invoke("azure-native:datacatalog/v20160330:getADCCatalog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADCCatalogArgs struct {
	CatalogName       string `pulumi:"catalogName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupADCCatalogResult struct {
	Admins                        []PrincipalsResponse `pulumi:"admins"`
	EnableAutomaticUnitAdjustment *bool                `pulumi:"enableAutomaticUnitAdjustment"`
	Etag                          *string              `pulumi:"etag"`
	Id                            string               `pulumi:"id"`
	Location                      *string              `pulumi:"location"`
	Name                          string               `pulumi:"name"`
	Sku                           *string              `pulumi:"sku"`
	SuccessfullyProvisioned       *bool                `pulumi:"successfullyProvisioned"`
	Tags                          map[string]string    `pulumi:"tags"`
	Type                          string               `pulumi:"type"`
	Units                         *int                 `pulumi:"units"`
	Users                         []PrincipalsResponse `pulumi:"users"`
}

func LookupADCCatalogOutput(ctx *pulumi.Context, args LookupADCCatalogOutputArgs, opts ...pulumi.InvokeOption) LookupADCCatalogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADCCatalogResult, error) {
			args := v.(LookupADCCatalogArgs)
			r, err := LookupADCCatalog(ctx, &args, opts...)
			var s LookupADCCatalogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADCCatalogResultOutput)
}

type LookupADCCatalogOutputArgs struct {
	CatalogName       pulumi.StringInput `pulumi:"catalogName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupADCCatalogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADCCatalogArgs)(nil)).Elem()
}


type LookupADCCatalogResultOutput struct{ *pulumi.OutputState }

func (LookupADCCatalogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADCCatalogResult)(nil)).Elem()
}

func (o LookupADCCatalogResultOutput) ToLookupADCCatalogResultOutput() LookupADCCatalogResultOutput {
	return o
}

func (o LookupADCCatalogResultOutput) ToLookupADCCatalogResultOutputWithContext(ctx context.Context) LookupADCCatalogResultOutput {
	return o
}

func (o LookupADCCatalogResultOutput) Admins() PrincipalsResponseArrayOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) []PrincipalsResponse { return v.Admins }).(PrincipalsResponseArrayOutput)
}

func (o LookupADCCatalogResultOutput) EnableAutomaticUnitAdjustment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) *bool { return v.EnableAutomaticUnitAdjustment }).(pulumi.BoolPtrOutput)
}

func (o LookupADCCatalogResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupADCCatalogResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADCCatalogResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupADCCatalogResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADCCatalogResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LookupADCCatalogResultOutput) SuccessfullyProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) *bool { return v.SuccessfullyProvisioned }).(pulumi.BoolPtrOutput)
}

func (o LookupADCCatalogResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupADCCatalogResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupADCCatalogResultOutput) Units() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) *int { return v.Units }).(pulumi.IntPtrOutput)
}

func (o LookupADCCatalogResultOutput) Users() PrincipalsResponseArrayOutput {
	return o.ApplyT(func(v LookupADCCatalogResult) []PrincipalsResponse { return v.Users }).(PrincipalsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADCCatalogResultOutput{})
}
