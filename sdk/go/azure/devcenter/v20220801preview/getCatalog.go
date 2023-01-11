


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCatalog(ctx *pulumi.Context, args *LookupCatalogArgs, opts ...pulumi.InvokeOption) (*LookupCatalogResult, error) {
	var rv LookupCatalogResult
	err := ctx.Invoke("azure-native:devcenter/v20220801preview:getCatalog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCatalogArgs struct {
	CatalogName       string `pulumi:"catalogName"`
	DevCenterName     string `pulumi:"devCenterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCatalogResult struct {
	AdoGit            *GitCatalogResponse `pulumi:"adoGit"`
	GitHub            *GitCatalogResponse `pulumi:"gitHub"`
	Id                string              `pulumi:"id"`
	LastSyncTime      string              `pulumi:"lastSyncTime"`
	Name              string              `pulumi:"name"`
	ProvisioningState string              `pulumi:"provisioningState"`
	SystemData        SystemDataResponse  `pulumi:"systemData"`
	Type              string              `pulumi:"type"`
}

func LookupCatalogOutput(ctx *pulumi.Context, args LookupCatalogOutputArgs, opts ...pulumi.InvokeOption) LookupCatalogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCatalogResult, error) {
			args := v.(LookupCatalogArgs)
			r, err := LookupCatalog(ctx, &args, opts...)
			var s LookupCatalogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCatalogResultOutput)
}

type LookupCatalogOutputArgs struct {
	CatalogName       pulumi.StringInput `pulumi:"catalogName"`
	DevCenterName     pulumi.StringInput `pulumi:"devCenterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCatalogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogArgs)(nil)).Elem()
}


type LookupCatalogResultOutput struct{ *pulumi.OutputState }

func (LookupCatalogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogResult)(nil)).Elem()
}

func (o LookupCatalogResultOutput) ToLookupCatalogResultOutput() LookupCatalogResultOutput {
	return o
}

func (o LookupCatalogResultOutput) ToLookupCatalogResultOutputWithContext(ctx context.Context) LookupCatalogResultOutput {
	return o
}

func (o LookupCatalogResultOutput) AdoGit() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v LookupCatalogResult) *GitCatalogResponse { return v.AdoGit }).(GitCatalogResponsePtrOutput)
}

func (o LookupCatalogResultOutput) GitHub() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v LookupCatalogResult) *GitCatalogResponse { return v.GitHub }).(GitCatalogResponsePtrOutput)
}

func (o LookupCatalogResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCatalogResultOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

func (o LookupCatalogResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCatalogResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCatalogResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCatalogResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCatalogResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCatalogResultOutput{})
}
