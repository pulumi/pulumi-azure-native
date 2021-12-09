


package v20220101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceControlConfiguration(ctx *pulumi.Context, args *LookupSourceControlConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlConfigurationResult, error) {
	var rv LookupSourceControlConfigurationResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration/v20220101preview:getSourceControlConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceControlConfigurationArgs struct {
	ClusterName                    string `pulumi:"clusterName"`
	ClusterResourceName            string `pulumi:"clusterResourceName"`
	ClusterRp                      string `pulumi:"clusterRp"`
	ResourceGroupName              string `pulumi:"resourceGroupName"`
	SourceControlConfigurationName string `pulumi:"sourceControlConfigurationName"`
}


type LookupSourceControlConfigurationResult struct {
	ComplianceStatus               ComplianceStatusResponse        `pulumi:"complianceStatus"`
	ConfigurationProtectedSettings map[string]string               `pulumi:"configurationProtectedSettings"`
	EnableHelmOperator             *bool                           `pulumi:"enableHelmOperator"`
	HelmOperatorProperties         *HelmOperatorPropertiesResponse `pulumi:"helmOperatorProperties"`
	Id                             string                          `pulumi:"id"`
	Name                           string                          `pulumi:"name"`
	OperatorInstanceName           *string                         `pulumi:"operatorInstanceName"`
	OperatorNamespace              *string                         `pulumi:"operatorNamespace"`
	OperatorParams                 *string                         `pulumi:"operatorParams"`
	OperatorScope                  *string                         `pulumi:"operatorScope"`
	OperatorType                   *string                         `pulumi:"operatorType"`
	ProvisioningState              string                          `pulumi:"provisioningState"`
	RepositoryPublicKey            string                          `pulumi:"repositoryPublicKey"`
	RepositoryUrl                  *string                         `pulumi:"repositoryUrl"`
	SshKnownHostsContents          *string                         `pulumi:"sshKnownHostsContents"`
	SystemData                     SystemDataResponse              `pulumi:"systemData"`
	Type                           string                          `pulumi:"type"`
}
