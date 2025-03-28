// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql
{
    public static class GetInstanceFailoverGroup
    {
        /// <summary>
        /// Gets a failover group.
        /// 
        /// Uses Azure REST API version 2021-11-01.
        /// 
        /// Other available API versions: 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01, 2023-08-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Task<GetInstanceFailoverGroupResult> InvokeAsync(GetInstanceFailoverGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceFailoverGroupResult>("azure-native:sql:getInstanceFailoverGroup", args ?? new GetInstanceFailoverGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a failover group.
        /// 
        /// Uses Azure REST API version 2021-11-01.
        /// 
        /// Other available API versions: 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01, 2023-08-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Output<GetInstanceFailoverGroupResult> Invoke(GetInstanceFailoverGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceFailoverGroupResult>("azure-native:sql:getInstanceFailoverGroup", args ?? new GetInstanceFailoverGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a failover group.
        /// 
        /// Uses Azure REST API version 2021-11-01.
        /// 
        /// Other available API versions: 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01, 2023-08-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Output<GetInstanceFailoverGroupResult> Invoke(GetInstanceFailoverGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceFailoverGroupResult>("azure-native:sql:getInstanceFailoverGroup", args ?? new GetInstanceFailoverGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceFailoverGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the failover group.
        /// </summary>
        [Input("failoverGroupName", required: true)]
        public string FailoverGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the region where the resource is located.
        /// </summary>
        [Input("locationName", required: true)]
        public string LocationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetInstanceFailoverGroupArgs()
        {
        }
        public static new GetInstanceFailoverGroupArgs Empty => new GetInstanceFailoverGroupArgs();
    }

    public sealed class GetInstanceFailoverGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the failover group.
        /// </summary>
        [Input("failoverGroupName", required: true)]
        public Input<string> FailoverGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the region where the resource is located.
        /// </summary>
        [Input("locationName", required: true)]
        public Input<string> LocationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetInstanceFailoverGroupInvokeArgs()
        {
        }
        public static new GetInstanceFailoverGroupInvokeArgs Empty => new GetInstanceFailoverGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceFailoverGroupResult
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of managed instance pairs in the failover group.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedInstancePairInfoResponse> ManagedInstancePairs;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Partner region information for the failover group.
        /// </summary>
        public readonly ImmutableArray<Outputs.PartnerRegionInfoResponse> PartnerRegions;
        /// <summary>
        /// Read-only endpoint of the failover group instance.
        /// </summary>
        public readonly Outputs.InstanceFailoverGroupReadOnlyEndpointResponse? ReadOnlyEndpoint;
        /// <summary>
        /// Read-write endpoint of the failover group instance.
        /// </summary>
        public readonly Outputs.InstanceFailoverGroupReadWriteEndpointResponse ReadWriteEndpoint;
        /// <summary>
        /// Local replication role of the failover group instance.
        /// </summary>
        public readonly string ReplicationRole;
        /// <summary>
        /// Replication state of the failover group instance.
        /// </summary>
        public readonly string ReplicationState;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInstanceFailoverGroupResult(
            string id,

            ImmutableArray<Outputs.ManagedInstancePairInfoResponse> managedInstancePairs,

            string name,

            ImmutableArray<Outputs.PartnerRegionInfoResponse> partnerRegions,

            Outputs.InstanceFailoverGroupReadOnlyEndpointResponse? readOnlyEndpoint,

            Outputs.InstanceFailoverGroupReadWriteEndpointResponse readWriteEndpoint,

            string replicationRole,

            string replicationState,

            string type)
        {
            Id = id;
            ManagedInstancePairs = managedInstancePairs;
            Name = name;
            PartnerRegions = partnerRegions;
            ReadOnlyEndpoint = readOnlyEndpoint;
            ReadWriteEndpoint = readWriteEndpoint;
            ReplicationRole = replicationRole;
            ReplicationState = replicationState;
            Type = type;
        }
    }
}
