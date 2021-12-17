


package elastic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CompanyInfo struct {
	Business        *string `pulumi:"business"`
	Country         *string `pulumi:"country"`
	Domain          *string `pulumi:"domain"`
	EmployeesNumber *string `pulumi:"employeesNumber"`
	State           *string `pulumi:"state"`
}





type CompanyInfoInput interface {
	pulumi.Input

	ToCompanyInfoOutput() CompanyInfoOutput
	ToCompanyInfoOutputWithContext(context.Context) CompanyInfoOutput
}

type CompanyInfoArgs struct {
	Business        pulumi.StringPtrInput `pulumi:"business"`
	Country         pulumi.StringPtrInput `pulumi:"country"`
	Domain          pulumi.StringPtrInput `pulumi:"domain"`
	EmployeesNumber pulumi.StringPtrInput `pulumi:"employeesNumber"`
	State           pulumi.StringPtrInput `pulumi:"state"`
}

func (CompanyInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CompanyInfo)(nil)).Elem()
}

func (i CompanyInfoArgs) ToCompanyInfoOutput() CompanyInfoOutput {
	return i.ToCompanyInfoOutputWithContext(context.Background())
}

func (i CompanyInfoArgs) ToCompanyInfoOutputWithContext(ctx context.Context) CompanyInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompanyInfoOutput)
}

func (i CompanyInfoArgs) ToCompanyInfoPtrOutput() CompanyInfoPtrOutput {
	return i.ToCompanyInfoPtrOutputWithContext(context.Background())
}

func (i CompanyInfoArgs) ToCompanyInfoPtrOutputWithContext(ctx context.Context) CompanyInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompanyInfoOutput).ToCompanyInfoPtrOutputWithContext(ctx)
}









type CompanyInfoPtrInput interface {
	pulumi.Input

	ToCompanyInfoPtrOutput() CompanyInfoPtrOutput
	ToCompanyInfoPtrOutputWithContext(context.Context) CompanyInfoPtrOutput
}

type companyInfoPtrType CompanyInfoArgs

func CompanyInfoPtr(v *CompanyInfoArgs) CompanyInfoPtrInput {
	return (*companyInfoPtrType)(v)
}

func (*companyInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CompanyInfo)(nil)).Elem()
}

func (i *companyInfoPtrType) ToCompanyInfoPtrOutput() CompanyInfoPtrOutput {
	return i.ToCompanyInfoPtrOutputWithContext(context.Background())
}

func (i *companyInfoPtrType) ToCompanyInfoPtrOutputWithContext(ctx context.Context) CompanyInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompanyInfoPtrOutput)
}

type CompanyInfoOutput struct{ *pulumi.OutputState }

func (CompanyInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompanyInfo)(nil)).Elem()
}

func (o CompanyInfoOutput) ToCompanyInfoOutput() CompanyInfoOutput {
	return o
}

func (o CompanyInfoOutput) ToCompanyInfoOutputWithContext(ctx context.Context) CompanyInfoOutput {
	return o
}

func (o CompanyInfoOutput) ToCompanyInfoPtrOutput() CompanyInfoPtrOutput {
	return o.ToCompanyInfoPtrOutputWithContext(context.Background())
}

func (o CompanyInfoOutput) ToCompanyInfoPtrOutputWithContext(ctx context.Context) CompanyInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CompanyInfo) *CompanyInfo {
		return &v
	}).(CompanyInfoPtrOutput)
}

func (o CompanyInfoOutput) Business() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompanyInfo) *string { return v.Business }).(pulumi.StringPtrOutput)
}

func (o CompanyInfoOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompanyInfo) *string { return v.Country }).(pulumi.StringPtrOutput)
}

func (o CompanyInfoOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompanyInfo) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o CompanyInfoOutput) EmployeesNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompanyInfo) *string { return v.EmployeesNumber }).(pulumi.StringPtrOutput)
}

func (o CompanyInfoOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompanyInfo) *string { return v.State }).(pulumi.StringPtrOutput)
}

type CompanyInfoPtrOutput struct{ *pulumi.OutputState }

func (CompanyInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompanyInfo)(nil)).Elem()
}

func (o CompanyInfoPtrOutput) ToCompanyInfoPtrOutput() CompanyInfoPtrOutput {
	return o
}

func (o CompanyInfoPtrOutput) ToCompanyInfoPtrOutputWithContext(ctx context.Context) CompanyInfoPtrOutput {
	return o
}

func (o CompanyInfoPtrOutput) Elem() CompanyInfoOutput {
	return o.ApplyT(func(v *CompanyInfo) CompanyInfo {
		if v != nil {
			return *v
		}
		var ret CompanyInfo
		return ret
	}).(CompanyInfoOutput)
}

func (o CompanyInfoPtrOutput) Business() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompanyInfo) *string {
		if v == nil {
			return nil
		}
		return v.Business
	}).(pulumi.StringPtrOutput)
}

func (o CompanyInfoPtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompanyInfo) *string {
		if v == nil {
			return nil
		}
		return v.Country
	}).(pulumi.StringPtrOutput)
}

func (o CompanyInfoPtrOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompanyInfo) *string {
		if v == nil {
			return nil
		}
		return v.Domain
	}).(pulumi.StringPtrOutput)
}

func (o CompanyInfoPtrOutput) EmployeesNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompanyInfo) *string {
		if v == nil {
			return nil
		}
		return v.EmployeesNumber
	}).(pulumi.StringPtrOutput)
}

func (o CompanyInfoPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompanyInfo) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ElasticCloudDeploymentResponse struct {
	AzureSubscriptionId     string `pulumi:"azureSubscriptionId"`
	DeploymentId            string `pulumi:"deploymentId"`
	ElasticsearchRegion     string `pulumi:"elasticsearchRegion"`
	ElasticsearchServiceUrl string `pulumi:"elasticsearchServiceUrl"`
	KibanaServiceUrl        string `pulumi:"kibanaServiceUrl"`
	KibanaSsoUrl            string `pulumi:"kibanaSsoUrl"`
	Name                    string `pulumi:"name"`
}

type ElasticCloudDeploymentResponseOutput struct{ *pulumi.OutputState }

func (ElasticCloudDeploymentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticCloudDeploymentResponse)(nil)).Elem()
}

func (o ElasticCloudDeploymentResponseOutput) ToElasticCloudDeploymentResponseOutput() ElasticCloudDeploymentResponseOutput {
	return o
}

func (o ElasticCloudDeploymentResponseOutput) ToElasticCloudDeploymentResponseOutputWithContext(ctx context.Context) ElasticCloudDeploymentResponseOutput {
	return o
}

func (o ElasticCloudDeploymentResponseOutput) AzureSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudDeploymentResponse) string { return v.AzureSubscriptionId }).(pulumi.StringOutput)
}

func (o ElasticCloudDeploymentResponseOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudDeploymentResponse) string { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o ElasticCloudDeploymentResponseOutput) ElasticsearchRegion() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudDeploymentResponse) string { return v.ElasticsearchRegion }).(pulumi.StringOutput)
}

func (o ElasticCloudDeploymentResponseOutput) ElasticsearchServiceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudDeploymentResponse) string { return v.ElasticsearchServiceUrl }).(pulumi.StringOutput)
}

func (o ElasticCloudDeploymentResponseOutput) KibanaServiceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudDeploymentResponse) string { return v.KibanaServiceUrl }).(pulumi.StringOutput)
}

func (o ElasticCloudDeploymentResponseOutput) KibanaSsoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudDeploymentResponse) string { return v.KibanaSsoUrl }).(pulumi.StringOutput)
}

func (o ElasticCloudDeploymentResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudDeploymentResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ElasticCloudDeploymentResponsePtrOutput struct{ *pulumi.OutputState }

func (ElasticCloudDeploymentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticCloudDeploymentResponse)(nil)).Elem()
}

func (o ElasticCloudDeploymentResponsePtrOutput) ToElasticCloudDeploymentResponsePtrOutput() ElasticCloudDeploymentResponsePtrOutput {
	return o
}

func (o ElasticCloudDeploymentResponsePtrOutput) ToElasticCloudDeploymentResponsePtrOutputWithContext(ctx context.Context) ElasticCloudDeploymentResponsePtrOutput {
	return o
}

func (o ElasticCloudDeploymentResponsePtrOutput) Elem() ElasticCloudDeploymentResponseOutput {
	return o.ApplyT(func(v *ElasticCloudDeploymentResponse) ElasticCloudDeploymentResponse {
		if v != nil {
			return *v
		}
		var ret ElasticCloudDeploymentResponse
		return ret
	}).(ElasticCloudDeploymentResponseOutput)
}

func (o ElasticCloudDeploymentResponsePtrOutput) AzureSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudDeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzureSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o ElasticCloudDeploymentResponsePtrOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudDeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DeploymentId
	}).(pulumi.StringPtrOutput)
}

func (o ElasticCloudDeploymentResponsePtrOutput) ElasticsearchRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudDeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ElasticsearchRegion
	}).(pulumi.StringPtrOutput)
}

func (o ElasticCloudDeploymentResponsePtrOutput) ElasticsearchServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudDeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ElasticsearchServiceUrl
	}).(pulumi.StringPtrOutput)
}

func (o ElasticCloudDeploymentResponsePtrOutput) KibanaServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudDeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KibanaServiceUrl
	}).(pulumi.StringPtrOutput)
}

func (o ElasticCloudDeploymentResponsePtrOutput) KibanaSsoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudDeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KibanaSsoUrl
	}).(pulumi.StringPtrOutput)
}

func (o ElasticCloudDeploymentResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudDeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ElasticCloudUserResponse struct {
	ElasticCloudSsoDefaultUrl string `pulumi:"elasticCloudSsoDefaultUrl"`
	EmailAddress              string `pulumi:"emailAddress"`
	Id                        string `pulumi:"id"`
}

type ElasticCloudUserResponseOutput struct{ *pulumi.OutputState }

func (ElasticCloudUserResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticCloudUserResponse)(nil)).Elem()
}

func (o ElasticCloudUserResponseOutput) ToElasticCloudUserResponseOutput() ElasticCloudUserResponseOutput {
	return o
}

func (o ElasticCloudUserResponseOutput) ToElasticCloudUserResponseOutputWithContext(ctx context.Context) ElasticCloudUserResponseOutput {
	return o
}

func (o ElasticCloudUserResponseOutput) ElasticCloudSsoDefaultUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudUserResponse) string { return v.ElasticCloudSsoDefaultUrl }).(pulumi.StringOutput)
}

func (o ElasticCloudUserResponseOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudUserResponse) string { return v.EmailAddress }).(pulumi.StringOutput)
}

func (o ElasticCloudUserResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticCloudUserResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ElasticCloudUserResponsePtrOutput struct{ *pulumi.OutputState }

func (ElasticCloudUserResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticCloudUserResponse)(nil)).Elem()
}

func (o ElasticCloudUserResponsePtrOutput) ToElasticCloudUserResponsePtrOutput() ElasticCloudUserResponsePtrOutput {
	return o
}

func (o ElasticCloudUserResponsePtrOutput) ToElasticCloudUserResponsePtrOutputWithContext(ctx context.Context) ElasticCloudUserResponsePtrOutput {
	return o
}

func (o ElasticCloudUserResponsePtrOutput) Elem() ElasticCloudUserResponseOutput {
	return o.ApplyT(func(v *ElasticCloudUserResponse) ElasticCloudUserResponse {
		if v != nil {
			return *v
		}
		var ret ElasticCloudUserResponse
		return ret
	}).(ElasticCloudUserResponseOutput)
}

func (o ElasticCloudUserResponsePtrOutput) ElasticCloudSsoDefaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudUserResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ElasticCloudSsoDefaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o ElasticCloudUserResponsePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudUserResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o ElasticCloudUserResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticCloudUserResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ElasticPropertiesResponse struct {
	ElasticCloudDeployment *ElasticCloudDeploymentResponse `pulumi:"elasticCloudDeployment"`
	ElasticCloudUser       *ElasticCloudUserResponse       `pulumi:"elasticCloudUser"`
}

type ElasticPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ElasticPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPropertiesResponse)(nil)).Elem()
}

func (o ElasticPropertiesResponseOutput) ToElasticPropertiesResponseOutput() ElasticPropertiesResponseOutput {
	return o
}

func (o ElasticPropertiesResponseOutput) ToElasticPropertiesResponseOutputWithContext(ctx context.Context) ElasticPropertiesResponseOutput {
	return o
}

func (o ElasticPropertiesResponseOutput) ElasticCloudDeployment() ElasticCloudDeploymentResponsePtrOutput {
	return o.ApplyT(func(v ElasticPropertiesResponse) *ElasticCloudDeploymentResponse { return v.ElasticCloudDeployment }).(ElasticCloudDeploymentResponsePtrOutput)
}

func (o ElasticPropertiesResponseOutput) ElasticCloudUser() ElasticCloudUserResponsePtrOutput {
	return o.ApplyT(func(v ElasticPropertiesResponse) *ElasticCloudUserResponse { return v.ElasticCloudUser }).(ElasticCloudUserResponsePtrOutput)
}

type ElasticPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ElasticPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPropertiesResponse)(nil)).Elem()
}

func (o ElasticPropertiesResponsePtrOutput) ToElasticPropertiesResponsePtrOutput() ElasticPropertiesResponsePtrOutput {
	return o
}

func (o ElasticPropertiesResponsePtrOutput) ToElasticPropertiesResponsePtrOutputWithContext(ctx context.Context) ElasticPropertiesResponsePtrOutput {
	return o
}

func (o ElasticPropertiesResponsePtrOutput) Elem() ElasticPropertiesResponseOutput {
	return o.ApplyT(func(v *ElasticPropertiesResponse) ElasticPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ElasticPropertiesResponse
		return ret
	}).(ElasticPropertiesResponseOutput)
}

func (o ElasticPropertiesResponsePtrOutput) ElasticCloudDeployment() ElasticCloudDeploymentResponsePtrOutput {
	return o.ApplyT(func(v *ElasticPropertiesResponse) *ElasticCloudDeploymentResponse {
		if v == nil {
			return nil
		}
		return v.ElasticCloudDeployment
	}).(ElasticCloudDeploymentResponsePtrOutput)
}

func (o ElasticPropertiesResponsePtrOutput) ElasticCloudUser() ElasticCloudUserResponsePtrOutput {
	return o.ApplyT(func(v *ElasticPropertiesResponse) *ElasticCloudUserResponse {
		if v == nil {
			return nil
		}
		return v.ElasticCloudUser
	}).(ElasticCloudUserResponsePtrOutput)
}

type FilteringTag struct {
	Action *string `pulumi:"action"`
	Name   *string `pulumi:"name"`
	Value  *string `pulumi:"value"`
}





type FilteringTagInput interface {
	pulumi.Input

	ToFilteringTagOutput() FilteringTagOutput
	ToFilteringTagOutputWithContext(context.Context) FilteringTagOutput
}

type FilteringTagArgs struct {
	Action pulumi.StringPtrInput `pulumi:"action"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Value  pulumi.StringPtrInput `pulumi:"value"`
}

func (FilteringTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilteringTag)(nil)).Elem()
}

func (i FilteringTagArgs) ToFilteringTagOutput() FilteringTagOutput {
	return i.ToFilteringTagOutputWithContext(context.Background())
}

func (i FilteringTagArgs) ToFilteringTagOutputWithContext(ctx context.Context) FilteringTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilteringTagOutput)
}





type FilteringTagArrayInput interface {
	pulumi.Input

	ToFilteringTagArrayOutput() FilteringTagArrayOutput
	ToFilteringTagArrayOutputWithContext(context.Context) FilteringTagArrayOutput
}

type FilteringTagArray []FilteringTagInput

func (FilteringTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilteringTag)(nil)).Elem()
}

func (i FilteringTagArray) ToFilteringTagArrayOutput() FilteringTagArrayOutput {
	return i.ToFilteringTagArrayOutputWithContext(context.Background())
}

func (i FilteringTagArray) ToFilteringTagArrayOutputWithContext(ctx context.Context) FilteringTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilteringTagArrayOutput)
}

type FilteringTagOutput struct{ *pulumi.OutputState }

func (FilteringTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilteringTag)(nil)).Elem()
}

func (o FilteringTagOutput) ToFilteringTagOutput() FilteringTagOutput {
	return o
}

func (o FilteringTagOutput) ToFilteringTagOutputWithContext(ctx context.Context) FilteringTagOutput {
	return o
}

func (o FilteringTagOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTag) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o FilteringTagOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTag) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FilteringTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type FilteringTagArrayOutput struct{ *pulumi.OutputState }

func (FilteringTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilteringTag)(nil)).Elem()
}

func (o FilteringTagArrayOutput) ToFilteringTagArrayOutput() FilteringTagArrayOutput {
	return o
}

func (o FilteringTagArrayOutput) ToFilteringTagArrayOutputWithContext(ctx context.Context) FilteringTagArrayOutput {
	return o
}

func (o FilteringTagArrayOutput) Index(i pulumi.IntInput) FilteringTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilteringTag {
		return vs[0].([]FilteringTag)[vs[1].(int)]
	}).(FilteringTagOutput)
}

type FilteringTagResponse struct {
	Action *string `pulumi:"action"`
	Name   *string `pulumi:"name"`
	Value  *string `pulumi:"value"`
}

type FilteringTagResponseOutput struct{ *pulumi.OutputState }

func (FilteringTagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilteringTagResponse)(nil)).Elem()
}

func (o FilteringTagResponseOutput) ToFilteringTagResponseOutput() FilteringTagResponseOutput {
	return o
}

func (o FilteringTagResponseOutput) ToFilteringTagResponseOutputWithContext(ctx context.Context) FilteringTagResponseOutput {
	return o
}

func (o FilteringTagResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTagResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o FilteringTagResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTagResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FilteringTagResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTagResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type FilteringTagResponseArrayOutput struct{ *pulumi.OutputState }

func (FilteringTagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilteringTagResponse)(nil)).Elem()
}

func (o FilteringTagResponseArrayOutput) ToFilteringTagResponseArrayOutput() FilteringTagResponseArrayOutput {
	return o
}

func (o FilteringTagResponseArrayOutput) ToFilteringTagResponseArrayOutputWithContext(ctx context.Context) FilteringTagResponseArrayOutput {
	return o
}

func (o FilteringTagResponseArrayOutput) Index(i pulumi.IntInput) FilteringTagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilteringTagResponse {
		return vs[0].([]FilteringTagResponse)[vs[1].(int)]
	}).(FilteringTagResponseOutput)
}

type IdentityProperties struct {
	Type *string `pulumi:"type"`
}





type IdentityPropertiesInput interface {
	pulumi.Input

	ToIdentityPropertiesOutput() IdentityPropertiesOutput
	ToIdentityPropertiesOutputWithContext(context.Context) IdentityPropertiesOutput
}

type IdentityPropertiesArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return i.ToIdentityPropertiesOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput)
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput).ToIdentityPropertiesPtrOutputWithContext(ctx)
}









type IdentityPropertiesPtrInput interface {
	pulumi.Input

	ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput
	ToIdentityPropertiesPtrOutputWithContext(context.Context) IdentityPropertiesPtrOutput
}

type identityPropertiesPtrType IdentityPropertiesArgs

func IdentityPropertiesPtr(v *IdentityPropertiesArgs) IdentityPropertiesPtrInput {
	return (*identityPropertiesPtrType)(v)
}

func (*identityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesPtrOutput)
}

type IdentityPropertiesOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProperties) *IdentityProperties {
		return &v
	}).(IdentityPropertiesPtrOutput)
}

func (o IdentityPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) Elem() IdentityPropertiesOutput {
	return o.ApplyT(func(v *IdentityProperties) IdentityProperties {
		if v != nil {
			return *v
		}
		var ret IdentityProperties
		return ret
	}).(IdentityPropertiesOutput)
}

func (o IdentityPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) Elem() IdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) IdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IdentityPropertiesResponse
		return ret
	}).(IdentityPropertiesResponseOutput)
}

func (o IdentityPropertiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type LogRules struct {
	FilteringTags        []FilteringTag `pulumi:"filteringTags"`
	SendAadLogs          *bool          `pulumi:"sendAadLogs"`
	SendActivityLogs     *bool          `pulumi:"sendActivityLogs"`
	SendSubscriptionLogs *bool          `pulumi:"sendSubscriptionLogs"`
}





type LogRulesInput interface {
	pulumi.Input

	ToLogRulesOutput() LogRulesOutput
	ToLogRulesOutputWithContext(context.Context) LogRulesOutput
}

type LogRulesArgs struct {
	FilteringTags        FilteringTagArrayInput `pulumi:"filteringTags"`
	SendAadLogs          pulumi.BoolPtrInput    `pulumi:"sendAadLogs"`
	SendActivityLogs     pulumi.BoolPtrInput    `pulumi:"sendActivityLogs"`
	SendSubscriptionLogs pulumi.BoolPtrInput    `pulumi:"sendSubscriptionLogs"`
}

func (LogRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogRules)(nil)).Elem()
}

func (i LogRulesArgs) ToLogRulesOutput() LogRulesOutput {
	return i.ToLogRulesOutputWithContext(context.Background())
}

func (i LogRulesArgs) ToLogRulesOutputWithContext(ctx context.Context) LogRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogRulesOutput)
}

func (i LogRulesArgs) ToLogRulesPtrOutput() LogRulesPtrOutput {
	return i.ToLogRulesPtrOutputWithContext(context.Background())
}

func (i LogRulesArgs) ToLogRulesPtrOutputWithContext(ctx context.Context) LogRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogRulesOutput).ToLogRulesPtrOutputWithContext(ctx)
}









type LogRulesPtrInput interface {
	pulumi.Input

	ToLogRulesPtrOutput() LogRulesPtrOutput
	ToLogRulesPtrOutputWithContext(context.Context) LogRulesPtrOutput
}

type logRulesPtrType LogRulesArgs

func LogRulesPtr(v *LogRulesArgs) LogRulesPtrInput {
	return (*logRulesPtrType)(v)
}

func (*logRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogRules)(nil)).Elem()
}

func (i *logRulesPtrType) ToLogRulesPtrOutput() LogRulesPtrOutput {
	return i.ToLogRulesPtrOutputWithContext(context.Background())
}

func (i *logRulesPtrType) ToLogRulesPtrOutputWithContext(ctx context.Context) LogRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogRulesPtrOutput)
}

type LogRulesOutput struct{ *pulumi.OutputState }

func (LogRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogRules)(nil)).Elem()
}

func (o LogRulesOutput) ToLogRulesOutput() LogRulesOutput {
	return o
}

func (o LogRulesOutput) ToLogRulesOutputWithContext(ctx context.Context) LogRulesOutput {
	return o
}

func (o LogRulesOutput) ToLogRulesPtrOutput() LogRulesPtrOutput {
	return o.ToLogRulesPtrOutputWithContext(context.Background())
}

func (o LogRulesOutput) ToLogRulesPtrOutputWithContext(ctx context.Context) LogRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogRules) *LogRules {
		return &v
	}).(LogRulesPtrOutput)
}

func (o LogRulesOutput) FilteringTags() FilteringTagArrayOutput {
	return o.ApplyT(func(v LogRules) []FilteringTag { return v.FilteringTags }).(FilteringTagArrayOutput)
}

func (o LogRulesOutput) SendAadLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRules) *bool { return v.SendAadLogs }).(pulumi.BoolPtrOutput)
}

func (o LogRulesOutput) SendActivityLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRules) *bool { return v.SendActivityLogs }).(pulumi.BoolPtrOutput)
}

func (o LogRulesOutput) SendSubscriptionLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRules) *bool { return v.SendSubscriptionLogs }).(pulumi.BoolPtrOutput)
}

type LogRulesPtrOutput struct{ *pulumi.OutputState }

func (LogRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogRules)(nil)).Elem()
}

func (o LogRulesPtrOutput) ToLogRulesPtrOutput() LogRulesPtrOutput {
	return o
}

func (o LogRulesPtrOutput) ToLogRulesPtrOutputWithContext(ctx context.Context) LogRulesPtrOutput {
	return o
}

func (o LogRulesPtrOutput) Elem() LogRulesOutput {
	return o.ApplyT(func(v *LogRules) LogRules {
		if v != nil {
			return *v
		}
		var ret LogRules
		return ret
	}).(LogRulesOutput)
}

func (o LogRulesPtrOutput) FilteringTags() FilteringTagArrayOutput {
	return o.ApplyT(func(v *LogRules) []FilteringTag {
		if v == nil {
			return nil
		}
		return v.FilteringTags
	}).(FilteringTagArrayOutput)
}

func (o LogRulesPtrOutput) SendAadLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRules) *bool {
		if v == nil {
			return nil
		}
		return v.SendAadLogs
	}).(pulumi.BoolPtrOutput)
}

func (o LogRulesPtrOutput) SendActivityLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRules) *bool {
		if v == nil {
			return nil
		}
		return v.SendActivityLogs
	}).(pulumi.BoolPtrOutput)
}

func (o LogRulesPtrOutput) SendSubscriptionLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRules) *bool {
		if v == nil {
			return nil
		}
		return v.SendSubscriptionLogs
	}).(pulumi.BoolPtrOutput)
}

type LogRulesResponse struct {
	FilteringTags        []FilteringTagResponse `pulumi:"filteringTags"`
	SendAadLogs          *bool                  `pulumi:"sendAadLogs"`
	SendActivityLogs     *bool                  `pulumi:"sendActivityLogs"`
	SendSubscriptionLogs *bool                  `pulumi:"sendSubscriptionLogs"`
}

type LogRulesResponseOutput struct{ *pulumi.OutputState }

func (LogRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogRulesResponse)(nil)).Elem()
}

func (o LogRulesResponseOutput) ToLogRulesResponseOutput() LogRulesResponseOutput {
	return o
}

func (o LogRulesResponseOutput) ToLogRulesResponseOutputWithContext(ctx context.Context) LogRulesResponseOutput {
	return o
}

func (o LogRulesResponseOutput) FilteringTags() FilteringTagResponseArrayOutput {
	return o.ApplyT(func(v LogRulesResponse) []FilteringTagResponse { return v.FilteringTags }).(FilteringTagResponseArrayOutput)
}

func (o LogRulesResponseOutput) SendAadLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRulesResponse) *bool { return v.SendAadLogs }).(pulumi.BoolPtrOutput)
}

func (o LogRulesResponseOutput) SendActivityLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRulesResponse) *bool { return v.SendActivityLogs }).(pulumi.BoolPtrOutput)
}

func (o LogRulesResponseOutput) SendSubscriptionLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRulesResponse) *bool { return v.SendSubscriptionLogs }).(pulumi.BoolPtrOutput)
}

type LogRulesResponsePtrOutput struct{ *pulumi.OutputState }

func (LogRulesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogRulesResponse)(nil)).Elem()
}

func (o LogRulesResponsePtrOutput) ToLogRulesResponsePtrOutput() LogRulesResponsePtrOutput {
	return o
}

func (o LogRulesResponsePtrOutput) ToLogRulesResponsePtrOutputWithContext(ctx context.Context) LogRulesResponsePtrOutput {
	return o
}

func (o LogRulesResponsePtrOutput) Elem() LogRulesResponseOutput {
	return o.ApplyT(func(v *LogRulesResponse) LogRulesResponse {
		if v != nil {
			return *v
		}
		var ret LogRulesResponse
		return ret
	}).(LogRulesResponseOutput)
}

func (o LogRulesResponsePtrOutput) FilteringTags() FilteringTagResponseArrayOutput {
	return o.ApplyT(func(v *LogRulesResponse) []FilteringTagResponse {
		if v == nil {
			return nil
		}
		return v.FilteringTags
	}).(FilteringTagResponseArrayOutput)
}

func (o LogRulesResponsePtrOutput) SendAadLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRulesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendAadLogs
	}).(pulumi.BoolPtrOutput)
}

func (o LogRulesResponsePtrOutput) SendActivityLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRulesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendActivityLogs
	}).(pulumi.BoolPtrOutput)
}

func (o LogRulesResponsePtrOutput) SendSubscriptionLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRulesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendSubscriptionLogs
	}).(pulumi.BoolPtrOutput)
}

type MonitorProperties struct {
	MonitoringStatus  *string   `pulumi:"monitoringStatus"`
	ProvisioningState *string   `pulumi:"provisioningState"`
	UserInfo          *UserInfo `pulumi:"userInfo"`
}





type MonitorPropertiesInput interface {
	pulumi.Input

	ToMonitorPropertiesOutput() MonitorPropertiesOutput
	ToMonitorPropertiesOutputWithContext(context.Context) MonitorPropertiesOutput
}

type MonitorPropertiesArgs struct {
	MonitoringStatus  pulumi.StringPtrInput `pulumi:"monitoringStatus"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	UserInfo          UserInfoPtrInput      `pulumi:"userInfo"`
}

func (MonitorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProperties)(nil)).Elem()
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesOutput() MonitorPropertiesOutput {
	return i.ToMonitorPropertiesOutputWithContext(context.Background())
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesOutputWithContext(ctx context.Context) MonitorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesOutput)
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return i.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesOutput).ToMonitorPropertiesPtrOutputWithContext(ctx)
}









type MonitorPropertiesPtrInput interface {
	pulumi.Input

	ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput
	ToMonitorPropertiesPtrOutputWithContext(context.Context) MonitorPropertiesPtrOutput
}

type monitorPropertiesPtrType MonitorPropertiesArgs

func MonitorPropertiesPtr(v *MonitorPropertiesArgs) MonitorPropertiesPtrInput {
	return (*monitorPropertiesPtrType)(v)
}

func (*monitorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProperties)(nil)).Elem()
}

func (i *monitorPropertiesPtrType) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return i.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (i *monitorPropertiesPtrType) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesPtrOutput)
}

type MonitorPropertiesOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProperties)(nil)).Elem()
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesOutput() MonitorPropertiesOutput {
	return o
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesOutputWithContext(ctx context.Context) MonitorPropertiesOutput {
	return o
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return o.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorProperties) *MonitorProperties {
		return &v
	}).(MonitorPropertiesPtrOutput)
}

func (o MonitorPropertiesOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *string { return v.MonitoringStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesOutput) UserInfo() UserInfoPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *UserInfo { return v.UserInfo }).(UserInfoPtrOutput)
}

type MonitorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProperties)(nil)).Elem()
}

func (o MonitorPropertiesPtrOutput) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return o
}

func (o MonitorPropertiesPtrOutput) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return o
}

func (o MonitorPropertiesPtrOutput) Elem() MonitorPropertiesOutput {
	return o.ApplyT(func(v *MonitorProperties) MonitorProperties {
		if v != nil {
			return *v
		}
		var ret MonitorProperties
		return ret
	}).(MonitorPropertiesOutput)
}

func (o MonitorPropertiesPtrOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *string {
		if v == nil {
			return nil
		}
		return v.MonitoringStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesPtrOutput) UserInfo() UserInfoPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *UserInfo {
		if v == nil {
			return nil
		}
		return v.UserInfo
	}).(UserInfoPtrOutput)
}

type MonitorPropertiesResponse struct {
	ElasticProperties       *ElasticPropertiesResponse `pulumi:"elasticProperties"`
	LiftrResourceCategory   string                     `pulumi:"liftrResourceCategory"`
	LiftrResourcePreference int                        `pulumi:"liftrResourcePreference"`
	MonitoringStatus        *string                    `pulumi:"monitoringStatus"`
	ProvisioningState       *string                    `pulumi:"provisioningState"`
}

type MonitorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorPropertiesResponse)(nil)).Elem()
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponseOutput() MonitorPropertiesResponseOutput {
	return o
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponseOutputWithContext(ctx context.Context) MonitorPropertiesResponseOutput {
	return o
}

func (o MonitorPropertiesResponseOutput) ElasticProperties() ElasticPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *ElasticPropertiesResponse { return v.ElasticProperties }).(ElasticPropertiesResponsePtrOutput)
}

func (o MonitorPropertiesResponseOutput) LiftrResourceCategory() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) string { return v.LiftrResourceCategory }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseOutput) LiftrResourcePreference() pulumi.IntOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) int { return v.LiftrResourcePreference }).(pulumi.IntOutput)
}

func (o MonitorPropertiesResponseOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *string { return v.MonitoringStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type MonitoredResourceResponse struct {
	Id                  *string `pulumi:"id"`
	ReasonForLogsStatus *string `pulumi:"reasonForLogsStatus"`
	SendingLogs         *string `pulumi:"sendingLogs"`
}

type MonitoringTagRulesProperties struct {
	LogRules          *LogRules `pulumi:"logRules"`
	ProvisioningState *string   `pulumi:"provisioningState"`
}





type MonitoringTagRulesPropertiesInput interface {
	pulumi.Input

	ToMonitoringTagRulesPropertiesOutput() MonitoringTagRulesPropertiesOutput
	ToMonitoringTagRulesPropertiesOutputWithContext(context.Context) MonitoringTagRulesPropertiesOutput
}

type MonitoringTagRulesPropertiesArgs struct {
	LogRules          LogRulesPtrInput      `pulumi:"logRules"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (MonitoringTagRulesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringTagRulesProperties)(nil)).Elem()
}

func (i MonitoringTagRulesPropertiesArgs) ToMonitoringTagRulesPropertiesOutput() MonitoringTagRulesPropertiesOutput {
	return i.ToMonitoringTagRulesPropertiesOutputWithContext(context.Background())
}

func (i MonitoringTagRulesPropertiesArgs) ToMonitoringTagRulesPropertiesOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringTagRulesPropertiesOutput)
}

func (i MonitoringTagRulesPropertiesArgs) ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput {
	return i.ToMonitoringTagRulesPropertiesPtrOutputWithContext(context.Background())
}

func (i MonitoringTagRulesPropertiesArgs) ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringTagRulesPropertiesOutput).ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx)
}









type MonitoringTagRulesPropertiesPtrInput interface {
	pulumi.Input

	ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput
	ToMonitoringTagRulesPropertiesPtrOutputWithContext(context.Context) MonitoringTagRulesPropertiesPtrOutput
}

type monitoringTagRulesPropertiesPtrType MonitoringTagRulesPropertiesArgs

func MonitoringTagRulesPropertiesPtr(v *MonitoringTagRulesPropertiesArgs) MonitoringTagRulesPropertiesPtrInput {
	return (*monitoringTagRulesPropertiesPtrType)(v)
}

func (*monitoringTagRulesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringTagRulesProperties)(nil)).Elem()
}

func (i *monitoringTagRulesPropertiesPtrType) ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput {
	return i.ToMonitoringTagRulesPropertiesPtrOutputWithContext(context.Background())
}

func (i *monitoringTagRulesPropertiesPtrType) ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringTagRulesPropertiesPtrOutput)
}

type MonitoringTagRulesPropertiesOutput struct{ *pulumi.OutputState }

func (MonitoringTagRulesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringTagRulesProperties)(nil)).Elem()
}

func (o MonitoringTagRulesPropertiesOutput) ToMonitoringTagRulesPropertiesOutput() MonitoringTagRulesPropertiesOutput {
	return o
}

func (o MonitoringTagRulesPropertiesOutput) ToMonitoringTagRulesPropertiesOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesOutput {
	return o
}

func (o MonitoringTagRulesPropertiesOutput) ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput {
	return o.ToMonitoringTagRulesPropertiesPtrOutputWithContext(context.Background())
}

func (o MonitoringTagRulesPropertiesOutput) ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitoringTagRulesProperties) *MonitoringTagRulesProperties {
		return &v
	}).(MonitoringTagRulesPropertiesPtrOutput)
}

func (o MonitoringTagRulesPropertiesOutput) LogRules() LogRulesPtrOutput {
	return o.ApplyT(func(v MonitoringTagRulesProperties) *LogRules { return v.LogRules }).(LogRulesPtrOutput)
}

func (o MonitoringTagRulesPropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitoringTagRulesProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type MonitoringTagRulesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MonitoringTagRulesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringTagRulesProperties)(nil)).Elem()
}

func (o MonitoringTagRulesPropertiesPtrOutput) ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput {
	return o
}

func (o MonitoringTagRulesPropertiesPtrOutput) ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesPtrOutput {
	return o
}

func (o MonitoringTagRulesPropertiesPtrOutput) Elem() MonitoringTagRulesPropertiesOutput {
	return o.ApplyT(func(v *MonitoringTagRulesProperties) MonitoringTagRulesProperties {
		if v != nil {
			return *v
		}
		var ret MonitoringTagRulesProperties
		return ret
	}).(MonitoringTagRulesPropertiesOutput)
}

func (o MonitoringTagRulesPropertiesPtrOutput) LogRules() LogRulesPtrOutput {
	return o.ApplyT(func(v *MonitoringTagRulesProperties) *LogRules {
		if v == nil {
			return nil
		}
		return v.LogRules
	}).(LogRulesPtrOutput)
}

func (o MonitoringTagRulesPropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringTagRulesProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type MonitoringTagRulesPropertiesResponse struct {
	LogRules          *LogRulesResponse `pulumi:"logRules"`
	ProvisioningState *string           `pulumi:"provisioningState"`
}

type MonitoringTagRulesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MonitoringTagRulesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringTagRulesPropertiesResponse)(nil)).Elem()
}

func (o MonitoringTagRulesPropertiesResponseOutput) ToMonitoringTagRulesPropertiesResponseOutput() MonitoringTagRulesPropertiesResponseOutput {
	return o
}

func (o MonitoringTagRulesPropertiesResponseOutput) ToMonitoringTagRulesPropertiesResponseOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesResponseOutput {
	return o
}

func (o MonitoringTagRulesPropertiesResponseOutput) LogRules() LogRulesResponsePtrOutput {
	return o.ApplyT(func(v MonitoringTagRulesPropertiesResponse) *LogRulesResponse { return v.LogRules }).(LogRulesResponsePtrOutput)
}

func (o MonitoringTagRulesPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitoringTagRulesPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ResourceSku struct {
	Name string `pulumi:"name"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Name string `pulumi:"name"`
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type UserInfo struct {
	CompanyInfo  *CompanyInfo `pulumi:"companyInfo"`
	CompanyName  *string      `pulumi:"companyName"`
	EmailAddress *string      `pulumi:"emailAddress"`
	FirstName    *string      `pulumi:"firstName"`
	LastName     *string      `pulumi:"lastName"`
}





type UserInfoInput interface {
	pulumi.Input

	ToUserInfoOutput() UserInfoOutput
	ToUserInfoOutputWithContext(context.Context) UserInfoOutput
}

type UserInfoArgs struct {
	CompanyInfo  CompanyInfoPtrInput   `pulumi:"companyInfo"`
	CompanyName  pulumi.StringPtrInput `pulumi:"companyName"`
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	FirstName    pulumi.StringPtrInput `pulumi:"firstName"`
	LastName     pulumi.StringPtrInput `pulumi:"lastName"`
}

func (UserInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (i UserInfoArgs) ToUserInfoOutput() UserInfoOutput {
	return i.ToUserInfoOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput)
}

func (i UserInfoArgs) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput).ToUserInfoPtrOutputWithContext(ctx)
}









type UserInfoPtrInput interface {
	pulumi.Input

	ToUserInfoPtrOutput() UserInfoPtrOutput
	ToUserInfoPtrOutputWithContext(context.Context) UserInfoPtrOutput
}

type userInfoPtrType UserInfoArgs

func UserInfoPtr(v *UserInfoArgs) UserInfoPtrInput {
	return (*userInfoPtrType)(v)
}

func (*userInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (i *userInfoPtrType) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i *userInfoPtrType) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoPtrOutput)
}

type UserInfoOutput struct{ *pulumi.OutputState }

func (UserInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (o UserInfoOutput) ToUserInfoOutput() UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o.ToUserInfoPtrOutputWithContext(context.Background())
}

func (o UserInfoOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserInfo) *UserInfo {
		return &v
	}).(UserInfoPtrOutput)
}

func (o UserInfoOutput) CompanyInfo() CompanyInfoPtrOutput {
	return o.ApplyT(func(v UserInfo) *CompanyInfo { return v.CompanyInfo }).(CompanyInfoPtrOutput)
}

func (o UserInfoOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.CompanyName }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

type UserInfoPtrOutput struct{ *pulumi.OutputState }

func (UserInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) Elem() UserInfoOutput {
	return o.ApplyT(func(v *UserInfo) UserInfo {
		if v != nil {
			return *v
		}
		var ret UserInfo
		return ret
	}).(UserInfoOutput)
}

func (o UserInfoPtrOutput) CompanyInfo() CompanyInfoPtrOutput {
	return o.ApplyT(func(v *UserInfo) *CompanyInfo {
		if v == nil {
			return nil
		}
		return v.CompanyInfo
	}).(CompanyInfoPtrOutput)
}

func (o UserInfoPtrOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.CompanyName
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

type VMResourcesResponse struct {
	VmResourceId *string `pulumi:"vmResourceId"`
}

func init() {
	pulumi.RegisterOutputType(CompanyInfoOutput{})
	pulumi.RegisterOutputType(CompanyInfoPtrOutput{})
	pulumi.RegisterOutputType(ElasticCloudDeploymentResponseOutput{})
	pulumi.RegisterOutputType(ElasticCloudDeploymentResponsePtrOutput{})
	pulumi.RegisterOutputType(ElasticCloudUserResponseOutput{})
	pulumi.RegisterOutputType(ElasticCloudUserResponsePtrOutput{})
	pulumi.RegisterOutputType(ElasticPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ElasticPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(FilteringTagOutput{})
	pulumi.RegisterOutputType(FilteringTagArrayOutput{})
	pulumi.RegisterOutputType(FilteringTagResponseOutput{})
	pulumi.RegisterOutputType(FilteringTagResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LogRulesOutput{})
	pulumi.RegisterOutputType(LogRulesPtrOutput{})
	pulumi.RegisterOutputType(LogRulesResponseOutput{})
	pulumi.RegisterOutputType(LogRulesResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MonitoringTagRulesPropertiesOutput{})
	pulumi.RegisterOutputType(MonitoringTagRulesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MonitoringTagRulesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserInfoOutput{})
	pulumi.RegisterOutputType(UserInfoPtrOutput{})
}
