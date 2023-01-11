


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRestorePointCollection(ctx *pulumi.Context, args *LookupRestorePointCollectionArgs, opts ...pulumi.InvokeOption) (*LookupRestorePointCollectionResult, error) {
	var rv LookupRestorePointCollectionResult
	err := ctx.Invoke("azure-native:compute/v20210401:getRestorePointCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRestorePointCollectionArgs struct {
	Expand                     *string `pulumi:"expand"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	RestorePointCollectionName string  `pulumi:"restorePointCollectionName"`
}


type LookupRestorePointCollectionResult struct {
	Id                       string                                          `pulumi:"id"`
	Location                 string                                          `pulumi:"location"`
	Name                     string                                          `pulumi:"name"`
	ProvisioningState        string                                          `pulumi:"provisioningState"`
	RestorePointCollectionId string                                          `pulumi:"restorePointCollectionId"`
	RestorePoints            []RestorePointResponse                          `pulumi:"restorePoints"`
	Source                   *RestorePointCollectionSourcePropertiesResponse `pulumi:"source"`
	Tags                     map[string]string                               `pulumi:"tags"`
	Type                     string                                          `pulumi:"type"`
}

func LookupRestorePointCollectionOutput(ctx *pulumi.Context, args LookupRestorePointCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupRestorePointCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRestorePointCollectionResult, error) {
			args := v.(LookupRestorePointCollectionArgs)
			r, err := LookupRestorePointCollection(ctx, &args, opts...)
			var s LookupRestorePointCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRestorePointCollectionResultOutput)
}

type LookupRestorePointCollectionOutputArgs struct {
	Expand                     pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName          pulumi.StringInput    `pulumi:"resourceGroupName"`
	RestorePointCollectionName pulumi.StringInput    `pulumi:"restorePointCollectionName"`
}

func (LookupRestorePointCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointCollectionArgs)(nil)).Elem()
}


type LookupRestorePointCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupRestorePointCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointCollectionResult)(nil)).Elem()
}

func (o LookupRestorePointCollectionResultOutput) ToLookupRestorePointCollectionResultOutput() LookupRestorePointCollectionResultOutput {
	return o
}

func (o LookupRestorePointCollectionResultOutput) ToLookupRestorePointCollectionResultOutputWithContext(ctx context.Context) LookupRestorePointCollectionResultOutput {
	return o
}

func (o LookupRestorePointCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRestorePointCollectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRestorePointCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRestorePointCollectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRestorePointCollectionResultOutput) RestorePointCollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.RestorePointCollectionId }).(pulumi.StringOutput)
}

func (o LookupRestorePointCollectionResultOutput) RestorePoints() RestorePointResponseArrayOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) []RestorePointResponse { return v.RestorePoints }).(RestorePointResponseArrayOutput)
}

func (o LookupRestorePointCollectionResultOutput) Source() RestorePointCollectionSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) *RestorePointCollectionSourcePropertiesResponse {
		return v.Source
	}).(RestorePointCollectionSourcePropertiesResponsePtrOutput)
}

func (o LookupRestorePointCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRestorePointCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRestorePointCollectionResultOutput{})
}
