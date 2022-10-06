


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:devcenter/v20220901preview:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPoolArgs struct {
	PoolName          string `pulumi:"poolName"`
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPoolResult struct {
	DevBoxDefinitionName  string             `pulumi:"devBoxDefinitionName"`
	Id                    string             `pulumi:"id"`
	LicenseType           string             `pulumi:"licenseType"`
	LocalAdministrator    string             `pulumi:"localAdministrator"`
	Location              string             `pulumi:"location"`
	Name                  string             `pulumi:"name"`
	NetworkConnectionName string             `pulumi:"networkConnectionName"`
	ProvisioningState     string             `pulumi:"provisioningState"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	Tags                  map[string]string  `pulumi:"tags"`
	Type                  string             `pulumi:"type"`
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPoolResult, error) {
			args := v.(LookupPoolArgs)
			r, err := LookupPool(ctx, &args, opts...)
			var s LookupPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPoolResultOutput)
}

type LookupPoolOutputArgs struct {
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ProjectName       pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}


type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) DevBoxDefinitionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.DevBoxDefinitionName }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.LicenseType }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) LocalAdministrator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.LocalAdministrator }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) NetworkConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.NetworkConnectionName }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
