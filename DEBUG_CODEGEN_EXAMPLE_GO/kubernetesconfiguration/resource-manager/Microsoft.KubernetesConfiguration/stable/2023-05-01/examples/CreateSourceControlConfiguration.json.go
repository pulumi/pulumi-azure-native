package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kubernetesconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kubernetesconfiguration.NewSourceControlConfiguration(ctx, "sourceControlConfiguration", &kubernetesconfiguration.SourceControlConfigurationArgs{
			ClusterName:         pulumi.String("clusterName1"),
			ClusterResourceName: pulumi.String("connectedClusters"),
			ClusterRp:           pulumi.String("Microsoft.Kubernetes"),
			ConfigurationProtectedSettings: pulumi.StringMap{
				"protectedSetting1Key": pulumi.String("protectedSetting1Value"),
			},
			EnableHelmOperator: pulumi.Bool(true),
			HelmOperatorProperties: &kubernetesconfiguration.HelmOperatorPropertiesArgs{
				ChartValues:  pulumi.String("--set git.ssh.secretName=flux-git-deploy --set tillerNamespace=kube-system"),
				ChartVersion: pulumi.String("0.3.0"),
			},
			OperatorInstanceName:           pulumi.String("SRSGitHubFluxOp-01"),
			OperatorNamespace:              pulumi.String("SRS_Namespace"),
			OperatorParams:                 pulumi.String("--git-email=xyzgituser@users.srs.github.com"),
			OperatorScope:                  pulumi.String("namespace"),
			OperatorType:                   pulumi.String("Flux"),
			RepositoryUrl:                  pulumi.String("git@github.com:k8sdeveloper425/flux-get-started"),
			ResourceGroupName:              pulumi.String("rg1"),
			SourceControlConfigurationName: pulumi.String("SRS_GitHubConfig"),
			SshKnownHostsContents:          pulumi.String("c3NoLmRldi5henVyZS5jb20gc3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFDN0hyMW9UV3FOcU9sekdKT2ZHSjROYWtWeUl6ZjFyWFlkNGQ3d282akJsa0x2Q0E0b2RCbEwwbURVeVowL1FVZlRUcWV1K3RtMjJnT3N2K1ZyVlRNazZ2d1JVNzVnWS95OXV0NU1iM2JSNUJWNThkS1h5cTlBOVVlQjVDYWtlaG41WmdtNngxbUtvVnlmK0ZGbjI2aVlxWEpSZ3pJWlpjWjVWNmhyRTBRZzM5a1ptNGF6NDhvMEFVYmY2U3A0U0xkdm51TWEyc1ZOd0hCYm9TN0VKa201N1hRUFZVMy9RcHlOTEhiV0Rkend0cmxTK2V6MzBTM0FkWWhMS0VPeEFHOHdlT255cnRMSkFVZW45bVRrb2w4b0lJMWVkZjdtV1diV1ZmMG5CbWx5MjErblpjbUNUSVNRQnRkY3lQYUVubzdmRlFNREQyNi9zMGxmS29iNEt3OEg="),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
