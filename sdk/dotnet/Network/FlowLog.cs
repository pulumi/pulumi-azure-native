// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    /// <summary>
    /// A flow log resource.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:FlowLog")]
    public partial class FlowLog : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Flag to enable/disable flow logging.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Optional field to filter network traffic logs based on SrcIP, SrcPort, DstIP, DstPort, Protocol, Encryption, Direction and Action. If not specified, all network traffic will be logged.
        /// </summary>
        [Output("enabledFilteringCriteria")]
        public Output<string?> EnabledFilteringCriteria { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Parameters that define the configuration of traffic analytics.
        /// </summary>
        [Output("flowAnalyticsConfiguration")]
        public Output<Outputs.TrafficAnalyticsPropertiesResponse?> FlowAnalyticsConfiguration { get; private set; } = null!;

        /// <summary>
        /// Parameters that define the flow log format.
        /// </summary>
        [Output("format")]
        public Output<Outputs.FlowLogFormatParametersResponse?> Format { get; private set; } = null!;

        /// <summary>
        /// FlowLog resource Managed Identity
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the flow log.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Parameters that define the retention policy for flow log.
        /// </summary>
        [Output("retentionPolicy")]
        public Output<Outputs.RetentionPolicyParametersResponse?> RetentionPolicy { get; private set; } = null!;

        /// <summary>
        /// ID of the storage account which is used to store the flow log.
        /// </summary>
        [Output("storageId")]
        public Output<string> StorageId { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Guid of network security group to which flow log will be applied.
        /// </summary>
        [Output("targetResourceGuid")]
        public Output<string> TargetResourceGuid { get; private set; } = null!;

        /// <summary>
        /// ID of network security group to which flow log will be applied.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FlowLog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlowLog(string name, FlowLogArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:FlowLog", name, args ?? new FlowLogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlowLog(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:FlowLog", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191101:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191201:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200301:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200501:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200601:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200701:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200801:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20201101:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210301:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210801:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:FlowLog" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:FlowLog" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FlowLog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlowLog Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FlowLog(name, id, options);
        }
    }

    public sealed class FlowLogArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag to enable/disable flow logging.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Optional field to filter network traffic logs based on SrcIP, SrcPort, DstIP, DstPort, Protocol, Encryption, Direction and Action. If not specified, all network traffic will be logged.
        /// </summary>
        [Input("enabledFilteringCriteria")]
        public Input<string>? EnabledFilteringCriteria { get; set; }

        /// <summary>
        /// Parameters that define the configuration of traffic analytics.
        /// </summary>
        [Input("flowAnalyticsConfiguration")]
        public Input<Inputs.TrafficAnalyticsPropertiesArgs>? FlowAnalyticsConfiguration { get; set; }

        /// <summary>
        /// The name of the flow log.
        /// </summary>
        [Input("flowLogName")]
        public Input<string>? FlowLogName { get; set; }

        /// <summary>
        /// Parameters that define the flow log format.
        /// </summary>
        [Input("format")]
        public Input<Inputs.FlowLogFormatParametersArgs>? Format { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// FlowLog resource Managed Identity
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

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

        /// <summary>
        /// Parameters that define the retention policy for flow log.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.RetentionPolicyParametersArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// ID of the storage account which is used to store the flow log.
        /// </summary>
        [Input("storageId", required: true)]
        public Input<string> StorageId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of network security group to which flow log will be applied.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        public FlowLogArgs()
        {
        }
        public static new FlowLogArgs Empty => new FlowLogArgs();
    }
}
