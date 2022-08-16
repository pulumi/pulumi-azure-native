


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedDatabase(ctx *pulumi.Context, args *LookupManagedDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupManagedDatabaseResult, error) {
	var rv LookupManagedDatabaseResult
	err := ctx.Invoke("azure-native:sql/v20220201preview:getManagedDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedDatabaseArgs struct {
	DatabaseName        string `pulumi:"databaseName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedDatabaseResult struct {
	CatalogCollation         *string           `pulumi:"catalogCollation"`
	Collation                *string           `pulumi:"collation"`
	CreationDate             string            `pulumi:"creationDate"`
	DefaultSecondaryLocation string            `pulumi:"defaultSecondaryLocation"`
	EarliestRestorePoint     string            `pulumi:"earliestRestorePoint"`
	FailoverGroupId          string            `pulumi:"failoverGroupId"`
	Id                       string            `pulumi:"id"`
	Location                 string            `pulumi:"location"`
	Name                     string            `pulumi:"name"`
	Status                   string            `pulumi:"status"`
	Tags                     map[string]string `pulumi:"tags"`
	Type                     string            `pulumi:"type"`
}

func LookupManagedDatabaseOutput(ctx *pulumi.Context, args LookupManagedDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupManagedDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedDatabaseResult, error) {
			args := v.(LookupManagedDatabaseArgs)
			r, err := LookupManagedDatabase(ctx, &args, opts...)
			var s LookupManagedDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedDatabaseResultOutput)
}

type LookupManagedDatabaseOutputArgs struct {
	DatabaseName        pulumi.StringInput `pulumi:"databaseName"`
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedDatabaseArgs)(nil)).Elem()
}


type LookupManagedDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupManagedDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedDatabaseResult)(nil)).Elem()
}

func (o LookupManagedDatabaseResultOutput) ToLookupManagedDatabaseResultOutput() LookupManagedDatabaseResultOutput {
	return o
}

func (o LookupManagedDatabaseResultOutput) ToLookupManagedDatabaseResultOutputWithContext(ctx context.Context) LookupManagedDatabaseResultOutput {
	return o
}

func (o LookupManagedDatabaseResultOutput) CatalogCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) *string { return v.CatalogCollation }).(pulumi.StringPtrOutput)
}

func (o LookupManagedDatabaseResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o LookupManagedDatabaseResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseResultOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseResultOutput) EarliestRestorePoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.EarliestRestorePoint }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseResultOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.FailoverGroupId }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedDatabaseResultOutput{})
}
