


package documentdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCassandraResourceCassandraView(ctx *pulumi.Context, args *LookupCassandraResourceCassandraViewArgs, opts ...pulumi.InvokeOption) (*LookupCassandraResourceCassandraViewResult, error) {
	var rv LookupCassandraResourceCassandraViewResult
	err := ctx.Invoke("azure-native:documentdb:getCassandraResourceCassandraView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraResourceCassandraViewArgs struct {
	AccountName       string `pulumi:"accountName"`
	KeyspaceName      string `pulumi:"keyspaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ViewName          string `pulumi:"viewName"`
}


type LookupCassandraResourceCassandraViewResult struct {
	Id       string                                      `pulumi:"id"`
	Identity *ManagedServiceIdentityResponse             `pulumi:"identity"`
	Location *string                                     `pulumi:"location"`
	Name     string                                      `pulumi:"name"`
	Options  *CassandraViewGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *CassandraViewGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                           `pulumi:"tags"`
	Type     string                                      `pulumi:"type"`
}

func LookupCassandraResourceCassandraViewOutput(ctx *pulumi.Context, args LookupCassandraResourceCassandraViewOutputArgs, opts ...pulumi.InvokeOption) LookupCassandraResourceCassandraViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCassandraResourceCassandraViewResult, error) {
			args := v.(LookupCassandraResourceCassandraViewArgs)
			r, err := LookupCassandraResourceCassandraView(ctx, &args, opts...)
			var s LookupCassandraResourceCassandraViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCassandraResourceCassandraViewResultOutput)
}

type LookupCassandraResourceCassandraViewOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	KeyspaceName      pulumi.StringInput `pulumi:"keyspaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ViewName          pulumi.StringInput `pulumi:"viewName"`
}

func (LookupCassandraResourceCassandraViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraResourceCassandraViewArgs)(nil)).Elem()
}


type LookupCassandraResourceCassandraViewResultOutput struct{ *pulumi.OutputState }

func (LookupCassandraResourceCassandraViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraResourceCassandraViewResult)(nil)).Elem()
}

func (o LookupCassandraResourceCassandraViewResultOutput) ToLookupCassandraResourceCassandraViewResultOutput() LookupCassandraResourceCassandraViewResultOutput {
	return o
}

func (o LookupCassandraResourceCassandraViewResultOutput) ToLookupCassandraResourceCassandraViewResultOutputWithContext(ctx context.Context) LookupCassandraResourceCassandraViewResultOutput {
	return o
}

func (o LookupCassandraResourceCassandraViewResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraViewResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCassandraResourceCassandraViewResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraViewResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupCassandraResourceCassandraViewResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraViewResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCassandraResourceCassandraViewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraViewResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCassandraResourceCassandraViewResultOutput) Options() CassandraViewGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraViewResult) *CassandraViewGetPropertiesResponseOptions {
		return v.Options
	}).(CassandraViewGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupCassandraResourceCassandraViewResultOutput) Resource() CassandraViewGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraViewResult) *CassandraViewGetPropertiesResponseResource {
		return v.Resource
	}).(CassandraViewGetPropertiesResponseResourcePtrOutput)
}

func (o LookupCassandraResourceCassandraViewResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraViewResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCassandraResourceCassandraViewResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraViewResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCassandraResourceCassandraViewResultOutput{})
}
