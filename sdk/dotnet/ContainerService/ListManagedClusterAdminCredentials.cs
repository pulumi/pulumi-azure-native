// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService
{
    public static class ListManagedClusterAdminCredentials
    {
        /// <summary>
        /// The list credential result response.
        /// 
        /// Uses Azure REST API version 2023-04-01.
        /// 
        /// Other available API versions: 2019-06-01, 2021-05-01, 2023-05-02-preview, 2023-06-01, 2023-06-02-preview, 2023-07-01, 2023-07-02-preview, 2023-08-01, 2023-08-02-preview, 2023-09-01, 2023-09-02-preview, 2023-10-01, 2023-10-02-preview, 2023-11-01, 2023-11-02-preview, 2024-01-01, 2024-01-02-preview, 2024-02-01, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-01, 2024-05-02-preview, 2024-06-02-preview, 2024-07-01, 2024-07-02-preview, 2024-08-01, 2024-09-01, 2024-09-02-preview, 2024-10-01, 2024-10-02-preview, 2025-01-01.
        /// </summary>
        public static Task<ListManagedClusterAdminCredentialsResult> InvokeAsync(ListManagedClusterAdminCredentialsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListManagedClusterAdminCredentialsResult>("azure-native:containerservice:listManagedClusterAdminCredentials", args ?? new ListManagedClusterAdminCredentialsArgs(), options.WithDefaults());

        /// <summary>
        /// The list credential result response.
        /// 
        /// Uses Azure REST API version 2023-04-01.
        /// 
        /// Other available API versions: 2019-06-01, 2021-05-01, 2023-05-02-preview, 2023-06-01, 2023-06-02-preview, 2023-07-01, 2023-07-02-preview, 2023-08-01, 2023-08-02-preview, 2023-09-01, 2023-09-02-preview, 2023-10-01, 2023-10-02-preview, 2023-11-01, 2023-11-02-preview, 2024-01-01, 2024-01-02-preview, 2024-02-01, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-01, 2024-05-02-preview, 2024-06-02-preview, 2024-07-01, 2024-07-02-preview, 2024-08-01, 2024-09-01, 2024-09-02-preview, 2024-10-01, 2024-10-02-preview, 2025-01-01.
        /// </summary>
        public static Output<ListManagedClusterAdminCredentialsResult> Invoke(ListManagedClusterAdminCredentialsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListManagedClusterAdminCredentialsResult>("azure-native:containerservice:listManagedClusterAdminCredentials", args ?? new ListManagedClusterAdminCredentialsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The list credential result response.
        /// 
        /// Uses Azure REST API version 2023-04-01.
        /// 
        /// Other available API versions: 2019-06-01, 2021-05-01, 2023-05-02-preview, 2023-06-01, 2023-06-02-preview, 2023-07-01, 2023-07-02-preview, 2023-08-01, 2023-08-02-preview, 2023-09-01, 2023-09-02-preview, 2023-10-01, 2023-10-02-preview, 2023-11-01, 2023-11-02-preview, 2024-01-01, 2024-01-02-preview, 2024-02-01, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-01, 2024-05-02-preview, 2024-06-02-preview, 2024-07-01, 2024-07-02-preview, 2024-08-01, 2024-09-01, 2024-09-02-preview, 2024-10-01, 2024-10-02-preview, 2025-01-01.
        /// </summary>
        public static Output<ListManagedClusterAdminCredentialsResult> Invoke(ListManagedClusterAdminCredentialsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListManagedClusterAdminCredentialsResult>("azure-native:containerservice:listManagedClusterAdminCredentials", args ?? new ListManagedClusterAdminCredentialsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListManagedClusterAdminCredentialsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        /// <summary>
        /// server fqdn type for credentials to be returned
        /// </summary>
        [Input("serverFqdn")]
        public string? ServerFqdn { get; set; }

        public ListManagedClusterAdminCredentialsArgs()
        {
        }
        public static new ListManagedClusterAdminCredentialsArgs Empty => new ListManagedClusterAdminCredentialsArgs();
    }

    public sealed class ListManagedClusterAdminCredentialsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// server fqdn type for credentials to be returned
        /// </summary>
        [Input("serverFqdn")]
        public Input<string>? ServerFqdn { get; set; }

        public ListManagedClusterAdminCredentialsInvokeArgs()
        {
        }
        public static new ListManagedClusterAdminCredentialsInvokeArgs Empty => new ListManagedClusterAdminCredentialsInvokeArgs();
    }


    [OutputType]
    public sealed class ListManagedClusterAdminCredentialsResult
    {
        /// <summary>
        /// Base64-encoded Kubernetes configuration file.
        /// </summary>
        public readonly ImmutableArray<Outputs.CredentialResultResponse> Kubeconfigs;

        [OutputConstructor]
        private ListManagedClusterAdminCredentialsResult(ImmutableArray<Outputs.CredentialResultResponse> kubeconfigs)
        {
            Kubeconfigs = kubeconfigs;
        }
    }
}
