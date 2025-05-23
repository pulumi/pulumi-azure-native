// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsMQ
{
    /// <summary>
    /// MQ mqttBridgeTopicMap resource
    /// 
    /// Uses Azure REST API version 2023-10-04-preview. In version 2.x of the Azure Native provider, it used API version 2023-10-04-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:iotoperationsmq:MqttBridgeTopicMap")]
    public partial class MqttBridgeTopicMap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Extended Location
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationPropertyResponse> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The MqttBridgeConnector CRD it refers to.
        /// </summary>
        [Output("mqttBridgeConnectorRef")]
        public Output<string> MqttBridgeConnectorRef { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The status of the last operation.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The route details for MqttBridge connector.
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<Outputs.MqttBridgeRoutesResponse>> Routes { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a MqttBridgeTopicMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MqttBridgeTopicMap(string name, MqttBridgeTopicMapArgs args, CustomResourceOptions? options = null)
            : base("azure-native:iotoperationsmq:MqttBridgeTopicMap", name, args ?? new MqttBridgeTopicMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MqttBridgeTopicMap(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:iotoperationsmq:MqttBridgeTopicMap", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:iotoperationsmq/v20231004preview:MqttBridgeTopicMap" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MqttBridgeTopicMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MqttBridgeTopicMap Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MqttBridgeTopicMap(name, id, options);
        }
    }

    public sealed class MqttBridgeTopicMapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Extended Location
        /// </summary>
        [Input("extendedLocation", required: true)]
        public Input<Inputs.ExtendedLocationPropertyArgs> ExtendedLocation { get; set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of MQ resource
        /// </summary>
        [Input("mqName", required: true)]
        public Input<string> MqName { get; set; } = null!;

        /// <summary>
        /// Name of MQ mqttBridgeConnector resource
        /// </summary>
        [Input("mqttBridgeConnectorName", required: true)]
        public Input<string> MqttBridgeConnectorName { get; set; } = null!;

        /// <summary>
        /// The MqttBridgeConnector CRD it refers to.
        /// </summary>
        [Input("mqttBridgeConnectorRef", required: true)]
        public Input<string> MqttBridgeConnectorRef { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("routes")]
        private InputList<Inputs.MqttBridgeRoutesArgs>? _routes;

        /// <summary>
        /// The route details for MqttBridge connector.
        /// </summary>
        public InputList<Inputs.MqttBridgeRoutesArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.MqttBridgeRoutesArgs>());
            set => _routes = value;
        }

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
        /// Name of MQ mqttBridgeTopicMap resource
        /// </summary>
        [Input("topicMapName")]
        public Input<string>? TopicMapName { get; set; }

        public MqttBridgeTopicMapArgs()
        {
        }
        public static new MqttBridgeTopicMapArgs Empty => new MqttBridgeTopicMapArgs();
    }
}
