// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Kubernetes
{
    public static class ListConnectedClusterUserCredentials
    {
        /// <summary>
        /// Gets cluster user credentials of the connected cluster with a specified resource group and name.
        /// 
        /// Uses Azure REST API version 2021-04-01-preview.
        /// </summary>
        public static Task<ListConnectedClusterUserCredentialsResult> InvokeAsync(ListConnectedClusterUserCredentialsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListConnectedClusterUserCredentialsResult>("azure-native:kubernetes:listConnectedClusterUserCredentials", args ?? new ListConnectedClusterUserCredentialsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets cluster user credentials of the connected cluster with a specified resource group and name.
        /// 
        /// Uses Azure REST API version 2021-04-01-preview.
        /// </summary>
        public static Output<ListConnectedClusterUserCredentialsResult> Invoke(ListConnectedClusterUserCredentialsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListConnectedClusterUserCredentialsResult>("azure-native:kubernetes:listConnectedClusterUserCredentials", args ?? new ListConnectedClusterUserCredentialsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets cluster user credentials of the connected cluster with a specified resource group and name.
        /// 
        /// Uses Azure REST API version 2021-04-01-preview.
        /// </summary>
        public static Output<ListConnectedClusterUserCredentialsResult> Invoke(ListConnectedClusterUserCredentialsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListConnectedClusterUserCredentialsResult>("azure-native:kubernetes:listConnectedClusterUserCredentials", args ?? new ListConnectedClusterUserCredentialsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListConnectedClusterUserCredentialsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The mode of client authentication.
        /// </summary>
        [Input("authenticationMethod", required: true)]
        public Union<string, Pulumi.AzureNative.Kubernetes.AuthenticationMethod> AuthenticationMethod { get; set; } = null!;

        /// <summary>
        /// Boolean value to indicate whether the request is for client side proxy or not
        /// </summary>
        [Input("clientProxy", required: true)]
        public bool ClientProxy { get; set; }

        /// <summary>
        /// The name of the Kubernetes cluster on which get is called.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListConnectedClusterUserCredentialsArgs()
        {
        }
        public static new ListConnectedClusterUserCredentialsArgs Empty => new ListConnectedClusterUserCredentialsArgs();
    }

    public sealed class ListConnectedClusterUserCredentialsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The mode of client authentication.
        /// </summary>
        [Input("authenticationMethod", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Kubernetes.AuthenticationMethod> AuthenticationMethod { get; set; } = null!;

        /// <summary>
        /// Boolean value to indicate whether the request is for client side proxy or not
        /// </summary>
        [Input("clientProxy", required: true)]
        public Input<bool> ClientProxy { get; set; } = null!;

        /// <summary>
        /// The name of the Kubernetes cluster on which get is called.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListConnectedClusterUserCredentialsInvokeArgs()
        {
        }
        public static new ListConnectedClusterUserCredentialsInvokeArgs Empty => new ListConnectedClusterUserCredentialsInvokeArgs();
    }


    [OutputType]
    public sealed class ListConnectedClusterUserCredentialsResult
    {
        /// <summary>
        /// Contains the REP (rendezvous endpoint) and “Sender” access token.
        /// </summary>
        public readonly Outputs.HybridConnectionConfigResponse HybridConnectionConfig;
        /// <summary>
        /// Base64-encoded Kubernetes configuration file.
        /// </summary>
        public readonly ImmutableArray<Outputs.CredentialResultResponse> Kubeconfigs;

        [OutputConstructor]
        private ListConnectedClusterUserCredentialsResult(
            Outputs.HybridConnectionConfigResponse hybridConnectionConfig,

            ImmutableArray<Outputs.CredentialResultResponse> kubeconfigs)
        {
            HybridConnectionConfig = hybridConnectionConfig;
            Kubeconfigs = kubeconfigs;
        }
    }
}
