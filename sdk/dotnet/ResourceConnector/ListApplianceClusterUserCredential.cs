// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ResourceConnector
{
    public static class ListApplianceClusterUserCredential
    {
        /// <summary>
        /// Returns the cluster user credentials for the dedicated appliance.
        /// 
        /// Uses Azure REST API version 2022-10-27.
        /// 
        /// Other available API versions: 2022-04-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resourceconnector [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListApplianceClusterUserCredentialResult> InvokeAsync(ListApplianceClusterUserCredentialArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListApplianceClusterUserCredentialResult>("azure-native:resourceconnector:listApplianceClusterUserCredential", args ?? new ListApplianceClusterUserCredentialArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the cluster user credentials for the dedicated appliance.
        /// 
        /// Uses Azure REST API version 2022-10-27.
        /// 
        /// Other available API versions: 2022-04-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resourceconnector [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListApplianceClusterUserCredentialResult> Invoke(ListApplianceClusterUserCredentialInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListApplianceClusterUserCredentialResult>("azure-native:resourceconnector:listApplianceClusterUserCredential", args ?? new ListApplianceClusterUserCredentialInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the cluster user credentials for the dedicated appliance.
        /// 
        /// Uses Azure REST API version 2022-10-27.
        /// 
        /// Other available API versions: 2022-04-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resourceconnector [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListApplianceClusterUserCredentialResult> Invoke(ListApplianceClusterUserCredentialInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListApplianceClusterUserCredentialResult>("azure-native:resourceconnector:listApplianceClusterUserCredential", args ?? new ListApplianceClusterUserCredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListApplianceClusterUserCredentialArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Appliances name.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public ListApplianceClusterUserCredentialArgs()
        {
        }
        public static new ListApplianceClusterUserCredentialArgs Empty => new ListApplianceClusterUserCredentialArgs();
    }

    public sealed class ListApplianceClusterUserCredentialInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Appliances name.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public ListApplianceClusterUserCredentialInvokeArgs()
        {
        }
        public static new ListApplianceClusterUserCredentialInvokeArgs Empty => new ListApplianceClusterUserCredentialInvokeArgs();
    }


    [OutputType]
    public sealed class ListApplianceClusterUserCredentialResult
    {
        /// <summary>
        /// Contains the REP (rendezvous endpoint) and “Listener” access token from notification service (NS).
        /// </summary>
        public readonly Outputs.HybridConnectionConfigResponse HybridConnectionConfig;
        /// <summary>
        /// The list of appliance kubeconfigs.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplianceCredentialKubeconfigResponse> Kubeconfigs;

        [OutputConstructor]
        private ListApplianceClusterUserCredentialResult(
            Outputs.HybridConnectionConfigResponse hybridConnectionConfig,

            ImmutableArray<Outputs.ApplianceCredentialKubeconfigResponse> kubeconfigs)
        {
            HybridConnectionConfig = hybridConnectionConfig;
            Kubeconfigs = kubeconfigs;
        }
    }
}
