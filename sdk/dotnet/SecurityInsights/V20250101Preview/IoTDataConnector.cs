// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20250101Preview
{
    /// <summary>
    /// Represents IoT data connector.
    /// </summary>
    [AzureNativeResourceType("azure-native:securityinsights/v20250101preview:IoTDataConnector")]
    public partial class IoTDataConnector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The available data types for the connector.
        /// </summary>
        [Output("dataTypes")]
        public Output<Outputs.AlertsDataTypeOfDataConnectorResponse?> DataTypes { get; private set; } = null!;

        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The kind of the data connector
        /// Expected value is 'IOT'.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The subscription id to connect to, and get the data from.
        /// </summary>
        [Output("subscriptionId")]
        public Output<string?> SubscriptionId { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IoTDataConnector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IoTDataConnector(string name, IoTDataConnectorArgs args, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights/v20250101preview:IoTDataConnector", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private IoTDataConnector(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights/v20250101preview:IoTDataConnector", name, null, MakeResourceOptions(options, id))
        {
        }

        private static IoTDataConnectorArgs MakeArgs(IoTDataConnectorArgs args)
        {
            args ??= new IoTDataConnectorArgs();
            args.Kind = "IOT";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20190101preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20200101:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20210301preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20210901preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20211001:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20211001preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220101preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220401preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220501preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220601preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220701preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220801:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220801preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220901preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221001preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221101:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221101preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221201preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230201:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230201preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230301preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230401preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230501preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230601preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230701preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230801preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230901preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231001preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231101:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231201preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240101preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240301:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240401preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240901:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20241001preview:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250301:IoTDataConnector" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights:IoTDataConnector" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IoTDataConnector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IoTDataConnector Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IoTDataConnector(name, id, options);
        }
    }

    public sealed class IoTDataConnectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connector ID
        /// </summary>
        [Input("dataConnectorId")]
        public Input<string>? DataConnectorId { get; set; }

        /// <summary>
        /// The available data types for the connector.
        /// </summary>
        [Input("dataTypes")]
        public Input<Inputs.AlertsDataTypeOfDataConnectorArgs>? DataTypes { get; set; }

        /// <summary>
        /// The kind of the data connector
        /// Expected value is 'IOT'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The subscription id to connect to, and get the data from.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public IoTDataConnectorArgs()
        {
        }
        public static new IoTDataConnectorArgs Empty => new IoTDataConnectorArgs();
    }
}
