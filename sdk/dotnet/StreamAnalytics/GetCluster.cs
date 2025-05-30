// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StreamAnalytics
{
    public static class GetCluster
    {
        /// <summary>
        /// Gets information about the specified cluster.
        /// 
        /// Uses Azure REST API version 2020-03-01.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azure-native:streamanalytics:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified cluster.
        /// 
        /// Uses Azure REST API version 2020-03-01.
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("azure-native:streamanalytics:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified cluster.
        /// 
        /// Uses Azure REST API version 2020-03-01.
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("azure-native:streamanalytics:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Represents the number of streaming units currently being used on the cluster.
        /// </summary>
        public readonly int CapacityAllocated;
        /// <summary>
        /// Represents the sum of the SUs of all streaming jobs associated with the cluster. If all of the jobs were running, this would be the capacity allocated.
        /// </summary>
        public readonly int CapacityAssigned;
        /// <summary>
        /// Unique identifier for the cluster.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The date this cluster was created.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// The current entity tag for the cluster. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of the cluster provisioning. The three terminal states are: Succeeded, Failed and Canceled
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The SKU of the cluster. This determines the size/capacity of the cluster. Required on PUT (CreateOrUpdate) requests.
        /// </summary>
        public readonly Outputs.ClusterSkuResponse? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetClusterResult(
            string azureApiVersion,

            int capacityAllocated,

            int capacityAssigned,

            string clusterId,

            string createdDate,

            string etag,

            string id,

            string? location,

            string name,

            string provisioningState,

            Outputs.ClusterSkuResponse? sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            CapacityAllocated = capacityAllocated;
            CapacityAssigned = capacityAssigned;
            ClusterId = clusterId;
            CreatedDate = createdDate;
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}
