// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse
{
    /// <summary>
    /// Class representing an iot hub data connection.
    /// 
    /// Uses Azure REST API version 2021-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-06-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:synapse:IotHubDataConnection")]
    public partial class IotHubDataConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The iot hub consumer group.
        /// </summary>
        [Output("consumerGroup")]
        public Output<string> ConsumerGroup { get; private set; } = null!;

        /// <summary>
        /// The data format of the message. Optionally the data format can be added to each message.
        /// </summary>
        [Output("dataFormat")]
        public Output<string?> DataFormat { get; private set; } = null!;

        /// <summary>
        /// System properties of the iot hub
        /// </summary>
        [Output("eventSystemProperties")]
        public Output<ImmutableArray<string>> EventSystemProperties { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the Iot hub to be used to create a data connection.
        /// </summary>
        [Output("iotHubResourceId")]
        public Output<string> IotHubResourceId { get; private set; } = null!;

        /// <summary>
        /// Kind of the endpoint for the data connection
        /// Expected value is 'IotHub'.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
        /// </summary>
        [Output("mappingRuleName")]
        public Output<string?> MappingRuleName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The name of the share access policy
        /// </summary>
        [Output("sharedAccessPolicyName")]
        public Output<string> SharedAccessPolicyName { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The table where the data should be ingested. Optionally the table information can be added to each message.
        /// </summary>
        [Output("tableName")]
        public Output<string?> TableName { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IotHubDataConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IotHubDataConnection(string name, IotHubDataConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:synapse:IotHubDataConnection", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private IotHubDataConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:synapse:IotHubDataConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static IotHubDataConnectionArgs MakeArgs(IotHubDataConnectionArgs args)
        {
            args ??= new IotHubDataConnectionArgs();
            args.Kind = "IotHub";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:synapse/v20210401preview:IotHubDataConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:synapse/v20210601preview:EventGridDataConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:synapse/v20210601preview:EventHubDataConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:synapse/v20210601preview:IotHubDataConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:synapse:EventGridDataConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:synapse:EventHubDataConnection" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IotHubDataConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IotHubDataConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IotHubDataConnection(name, id, options);
        }
    }

    public sealed class IotHubDataConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The iot hub consumer group.
        /// </summary>
        [Input("consumerGroup", required: true)]
        public Input<string> ConsumerGroup { get; set; } = null!;

        /// <summary>
        /// The name of the data connection.
        /// </summary>
        [Input("dataConnectionName")]
        public Input<string>? DataConnectionName { get; set; }

        /// <summary>
        /// The data format of the message. Optionally the data format can be added to each message.
        /// </summary>
        [Input("dataFormat")]
        public InputUnion<string, Pulumi.AzureNative.Synapse.IotHubDataFormat>? DataFormat { get; set; }

        /// <summary>
        /// The name of the database in the Kusto pool.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("eventSystemProperties")]
        private InputList<string>? _eventSystemProperties;

        /// <summary>
        /// System properties of the iot hub
        /// </summary>
        public InputList<string> EventSystemProperties
        {
            get => _eventSystemProperties ?? (_eventSystemProperties = new InputList<string>());
            set => _eventSystemProperties = value;
        }

        /// <summary>
        /// The resource ID of the Iot hub to be used to create a data connection.
        /// </summary>
        [Input("iotHubResourceId", required: true)]
        public Input<string> IotHubResourceId { get; set; } = null!;

        /// <summary>
        /// Kind of the endpoint for the data connection
        /// Expected value is 'IotHub'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto pool.
        /// </summary>
        [Input("kustoPoolName", required: true)]
        public Input<string> KustoPoolName { get; set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
        /// </summary>
        [Input("mappingRuleName")]
        public Input<string>? MappingRuleName { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share access policy
        /// </summary>
        [Input("sharedAccessPolicyName", required: true)]
        public Input<string> SharedAccessPolicyName { get; set; } = null!;

        /// <summary>
        /// The table where the data should be ingested. Optionally the table information can be added to each message.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public IotHubDataConnectionArgs()
        {
        }
        public static new IotHubDataConnectionArgs Empty => new IotHubDataConnectionArgs();
    }
}
