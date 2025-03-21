// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Datadog.V20220801
{
    public static class GetMonitor
    {
        public static Task<GetMonitorResult> InvokeAsync(GetMonitorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitorResult>("azure-native:datadog/v20220801:getMonitor", args ?? new GetMonitorArgs(), options.WithDefaults());

        public static Output<GetMonitorResult> Invoke(GetMonitorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitorResult>("azure-native:datadog/v20220801:getMonitor", args ?? new GetMonitorInvokeArgs(), options.WithDefaults());

        public static Output<GetMonitorResult> Invoke(GetMonitorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitorResult>("azure-native:datadog/v20220801:getMonitor", args ?? new GetMonitorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMonitorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Monitor resource name
        /// </summary>
        [Input("monitorName", required: true)]
        public string MonitorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMonitorArgs()
        {
        }
        public static new GetMonitorArgs Empty => new GetMonitorArgs();
    }

    public sealed class GetMonitorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Monitor resource name
        /// </summary>
        [Input("monitorName", required: true)]
        public Input<string> MonitorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetMonitorInvokeArgs()
        {
        }
        public static new GetMonitorInvokeArgs Empty => new GetMonitorInvokeArgs();
    }


    [OutputType]
    public sealed class GetMonitorResult
    {
        /// <summary>
        /// ARM id of the monitor resource.
        /// </summary>
        public readonly string Id;
        public readonly Outputs.IdentityPropertiesResponse? Identity;
        public readonly string Location;
        /// <summary>
        /// Name of the monitor resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties specific to the monitor resource.
        /// </summary>
        public readonly Outputs.MonitorPropertiesResponse Properties;
        public readonly Outputs.ResourceSkuResponse? Sku;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the monitor resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMonitorResult(
            string id,

            Outputs.IdentityPropertiesResponse? identity,

            string location,

            string name,

            Outputs.MonitorPropertiesResponse properties,

            Outputs.ResourceSkuResponse? sku,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
