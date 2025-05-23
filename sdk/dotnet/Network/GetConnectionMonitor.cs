// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetConnectionMonitor
    {
        /// <summary>
        /// Gets a connection monitor by name.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetConnectionMonitorResult> InvokeAsync(GetConnectionMonitorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectionMonitorResult>("azure-native:network:getConnectionMonitor", args ?? new GetConnectionMonitorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a connection monitor by name.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetConnectionMonitorResult> Invoke(GetConnectionMonitorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectionMonitorResult>("azure-native:network:getConnectionMonitor", args ?? new GetConnectionMonitorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a connection monitor by name.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetConnectionMonitorResult> Invoke(GetConnectionMonitorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectionMonitorResult>("azure-native:network:getConnectionMonitor", args ?? new GetConnectionMonitorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectionMonitorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the connection monitor.
        /// </summary>
        [Input("connectionMonitorName", required: true)]
        public string ConnectionMonitorName { get; set; } = null!;

        /// <summary>
        /// The name of the Network Watcher resource.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public string NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing Network Watcher.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConnectionMonitorArgs()
        {
        }
        public static new GetConnectionMonitorArgs Empty => new GetConnectionMonitorArgs();
    }

    public sealed class GetConnectionMonitorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the connection monitor.
        /// </summary>
        [Input("connectionMonitorName", required: true)]
        public Input<string> ConnectionMonitorName { get; set; } = null!;

        /// <summary>
        /// The name of the Network Watcher resource.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public Input<string> NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing Network Watcher.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetConnectionMonitorInvokeArgs()
        {
        }
        public static new GetConnectionMonitorInvokeArgs Empty => new GetConnectionMonitorInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectionMonitorResult
    {
        /// <summary>
        /// Determines if the connection monitor will start automatically once created.
        /// </summary>
        public readonly bool? AutoStart;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Type of connection monitor.
        /// </summary>
        public readonly string ConnectionMonitorType;
        /// <summary>
        /// Describes the destination of connection monitor.
        /// </summary>
        public readonly Outputs.ConnectionMonitorDestinationResponse? Destination;
        /// <summary>
        /// List of connection monitor endpoints.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionMonitorEndpointResponse> Endpoints;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// ID of the connection monitor.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Connection monitor location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Monitoring interval in seconds.
        /// </summary>
        public readonly int? MonitoringIntervalInSeconds;
        /// <summary>
        /// The monitoring status of the connection monitor.
        /// </summary>
        public readonly string MonitoringStatus;
        /// <summary>
        /// Name of the connection monitor.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional notes to be associated with the connection monitor.
        /// </summary>
        public readonly string? Notes;
        /// <summary>
        /// List of connection monitor outputs.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionMonitorOutputResponse> Outputs;
        /// <summary>
        /// The provisioning state of the connection monitor.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Describes the source of connection monitor.
        /// </summary>
        public readonly Outputs.ConnectionMonitorSourceResponse? Source;
        /// <summary>
        /// The date and time when the connection monitor was started.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Connection monitor tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// List of connection monitor test configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionMonitorTestConfigurationResponse> TestConfigurations;
        /// <summary>
        /// List of connection monitor test groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionMonitorTestGroupResponse> TestGroups;
        /// <summary>
        /// Connection monitor type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConnectionMonitorResult(
            bool? autoStart,

            string azureApiVersion,

            string connectionMonitorType,

            Outputs.ConnectionMonitorDestinationResponse? destination,

            ImmutableArray<Outputs.ConnectionMonitorEndpointResponse> endpoints,

            string etag,

            string id,

            string? location,

            int? monitoringIntervalInSeconds,

            string monitoringStatus,

            string name,

            string? notes,

            ImmutableArray<Outputs.ConnectionMonitorOutputResponse> outputs,

            string provisioningState,

            Outputs.ConnectionMonitorSourceResponse? source,

            string startTime,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<Outputs.ConnectionMonitorTestConfigurationResponse> testConfigurations,

            ImmutableArray<Outputs.ConnectionMonitorTestGroupResponse> testGroups,

            string type)
        {
            AutoStart = autoStart;
            AzureApiVersion = azureApiVersion;
            ConnectionMonitorType = connectionMonitorType;
            Destination = destination;
            Endpoints = endpoints;
            Etag = etag;
            Id = id;
            Location = location;
            MonitoringIntervalInSeconds = monitoringIntervalInSeconds;
            MonitoringStatus = monitoringStatus;
            Name = name;
            Notes = notes;
            Outputs = outputs;
            ProvisioningState = provisioningState;
            Source = source;
            StartTime = startTime;
            Tags = tags;
            TestConfigurations = testConfigurations;
            TestGroups = testGroups;
            Type = type;
        }
    }
}
