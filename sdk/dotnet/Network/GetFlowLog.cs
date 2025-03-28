// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetFlowLog
    {
        /// <summary>
        /// Gets a flow log resource by name.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Task<GetFlowLogResult> InvokeAsync(GetFlowLogArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFlowLogResult>("azure-native:network:getFlowLog", args ?? new GetFlowLogArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a flow log resource by name.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetFlowLogResult> Invoke(GetFlowLogInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFlowLogResult>("azure-native:network:getFlowLog", args ?? new GetFlowLogInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a flow log resource by name.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetFlowLogResult> Invoke(GetFlowLogInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFlowLogResult>("azure-native:network:getFlowLog", args ?? new GetFlowLogInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFlowLogArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the flow log resource.
        /// </summary>
        [Input("flowLogName", required: true)]
        public string FlowLogName { get; set; } = null!;

        /// <summary>
        /// The name of the network watcher.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public string NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFlowLogArgs()
        {
        }
        public static new GetFlowLogArgs Empty => new GetFlowLogArgs();
    }

    public sealed class GetFlowLogInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the flow log resource.
        /// </summary>
        [Input("flowLogName", required: true)]
        public Input<string> FlowLogName { get; set; } = null!;

        /// <summary>
        /// The name of the network watcher.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public Input<string> NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetFlowLogInvokeArgs()
        {
        }
        public static new GetFlowLogInvokeArgs Empty => new GetFlowLogInvokeArgs();
    }


    [OutputType]
    public sealed class GetFlowLogResult
    {
        /// <summary>
        /// Flag to enable/disable flow logging.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Parameters that define the configuration of traffic analytics.
        /// </summary>
        public readonly Outputs.TrafficAnalyticsPropertiesResponse? FlowAnalyticsConfiguration;
        /// <summary>
        /// Parameters that define the flow log format.
        /// </summary>
        public readonly Outputs.FlowLogFormatParametersResponse? Format;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the flow log.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Parameters that define the retention policy for flow log.
        /// </summary>
        public readonly Outputs.RetentionPolicyParametersResponse? RetentionPolicy;
        /// <summary>
        /// ID of the storage account which is used to store the flow log.
        /// </summary>
        public readonly string StorageId;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Guid of network security group to which flow log will be applied.
        /// </summary>
        public readonly string TargetResourceGuid;
        /// <summary>
        /// ID of network security group to which flow log will be applied.
        /// </summary>
        public readonly string TargetResourceId;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetFlowLogResult(
            bool? enabled,

            string etag,

            Outputs.TrafficAnalyticsPropertiesResponse? flowAnalyticsConfiguration,

            Outputs.FlowLogFormatParametersResponse? format,

            string? id,

            string? location,

            string name,

            string provisioningState,

            Outputs.RetentionPolicyParametersResponse? retentionPolicy,

            string storageId,

            ImmutableDictionary<string, string>? tags,

            string targetResourceGuid,

            string targetResourceId,

            string type)
        {
            Enabled = enabled;
            Etag = etag;
            FlowAnalyticsConfiguration = flowAnalyticsConfiguration;
            Format = format;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            RetentionPolicy = retentionPolicy;
            StorageId = storageId;
            Tags = tags;
            TargetResourceGuid = targetResourceGuid;
            TargetResourceId = targetResourceId;
            Type = type;
        }
    }
}
