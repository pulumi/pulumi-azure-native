// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric
{
    public static class GetManagedAzResiliencyStatus
    {
        /// <summary>
        /// Action to get Az Resiliency Status of all the Base resources constituting Service Fabric Managed Clusters.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01-preview, 2024-02-01-preview, 2024-06-01-preview, 2024-09-01-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicefabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetManagedAzResiliencyStatusResult> InvokeAsync(GetManagedAzResiliencyStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetManagedAzResiliencyStatusResult>("azure-native:servicefabric:getManagedAzResiliencyStatus", args ?? new GetManagedAzResiliencyStatusArgs(), options.WithDefaults());

        /// <summary>
        /// Action to get Az Resiliency Status of all the Base resources constituting Service Fabric Managed Clusters.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01-preview, 2024-02-01-preview, 2024-06-01-preview, 2024-09-01-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicefabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetManagedAzResiliencyStatusResult> Invoke(GetManagedAzResiliencyStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedAzResiliencyStatusResult>("azure-native:servicefabric:getManagedAzResiliencyStatus", args ?? new GetManagedAzResiliencyStatusInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Action to get Az Resiliency Status of all the Base resources constituting Service Fabric Managed Clusters.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01-preview, 2024-02-01-preview, 2024-06-01-preview, 2024-09-01-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicefabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetManagedAzResiliencyStatusResult> Invoke(GetManagedAzResiliencyStatusInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedAzResiliencyStatusResult>("azure-native:servicefabric:getManagedAzResiliencyStatus", args ?? new GetManagedAzResiliencyStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetManagedAzResiliencyStatusArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagedAzResiliencyStatusArgs()
        {
        }
        public static new GetManagedAzResiliencyStatusArgs Empty => new GetManagedAzResiliencyStatusArgs();
    }

    public sealed class GetManagedAzResiliencyStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetManagedAzResiliencyStatusInvokeArgs()
        {
        }
        public static new GetManagedAzResiliencyStatusInvokeArgs Empty => new GetManagedAzResiliencyStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetManagedAzResiliencyStatusResult
    {
        /// <summary>
        /// List of Managed VM Sizes for Service Fabric Managed Clusters.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceAzStatusResponse> BaseResourceStatus;
        /// <summary>
        /// URL to get the next set of Managed VM Sizes if there are any.
        /// </summary>
        public readonly bool IsClusterZoneResilient;

        [OutputConstructor]
        private GetManagedAzResiliencyStatusResult(
            ImmutableArray<Outputs.ResourceAzStatusResponse> baseResourceStatus,

            bool isClusterZoneResilient)
        {
            BaseResourceStatus = baseResourceStatus;
            IsClusterZoneResilient = isClusterZoneResilient;
        }
    }
}
