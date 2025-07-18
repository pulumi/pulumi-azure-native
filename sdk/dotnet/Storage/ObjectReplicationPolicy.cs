// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage
{
    /// <summary>
    /// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
    /// 
    /// Uses Azure REST API version 2024-01-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
    /// 
    /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:storage:ObjectReplicationPolicy")]
    public partial class ObjectReplicationPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Required. Destination account name. It should be full resource id if allowCrossTenantReplication set to false.
        /// </summary>
        [Output("destinationAccount")]
        public Output<string> DestinationAccount { get; private set; } = null!;

        /// <summary>
        /// Indicates when the policy is enabled on the source account.
        /// </summary>
        [Output("enabledTime")]
        public Output<string> EnabledTime { get; private set; } = null!;

        /// <summary>
        /// Optional. The object replication policy metrics feature options.
        /// </summary>
        [Output("metrics")]
        public Output<Outputs.ObjectReplicationPolicyPropertiesResponseMetrics?> Metrics { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A unique id for object replication policy.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The storage account object replication rules.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.ObjectReplicationPolicyRuleResponse>> Rules { get; private set; } = null!;

        /// <summary>
        /// Required. Source account name. It should be full resource id if allowCrossTenantReplication set to false.
        /// </summary>
        [Output("sourceAccount")]
        public Output<string> SourceAccount { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectReplicationPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectReplicationPolicy(string name, ObjectReplicationPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure-native:storage:ObjectReplicationPolicy", name, args ?? new ObjectReplicationPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectReplicationPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:storage:ObjectReplicationPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20190601:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20200801preview:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210101:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210201:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210401:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210601:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210801:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210901:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20220501:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20220901:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20230101:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20230401:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20230501:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20240101:ObjectReplicationPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20250101:ObjectReplicationPolicy" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ObjectReplicationPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectReplicationPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ObjectReplicationPolicy(name, id, options);
        }
    }

    public sealed class ObjectReplicationPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Required. Destination account name. It should be full resource id if allowCrossTenantReplication set to false.
        /// </summary>
        [Input("destinationAccount", required: true)]
        public Input<string> DestinationAccount { get; set; } = null!;

        /// <summary>
        /// Optional. The object replication policy metrics feature options.
        /// </summary>
        [Input("metrics")]
        public Input<Inputs.ObjectReplicationPolicyPropertiesMetricsArgs>? Metrics { get; set; }

        /// <summary>
        /// For the destination account, provide the value 'default'. Configure the policy on the destination account first. For the source account, provide the value of the policy ID that is returned when you download the policy that was defined on the destination account. The policy is downloaded as a JSON file.
        /// </summary>
        [Input("objectReplicationPolicyId")]
        public Input<string>? ObjectReplicationPolicyId { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("rules")]
        private InputList<Inputs.ObjectReplicationPolicyRuleArgs>? _rules;

        /// <summary>
        /// The storage account object replication rules.
        /// </summary>
        public InputList<Inputs.ObjectReplicationPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ObjectReplicationPolicyRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Required. Source account name. It should be full resource id if allowCrossTenantReplication set to false.
        /// </summary>
        [Input("sourceAccount", required: true)]
        public Input<string> SourceAccount { get; set; } = null!;

        public ObjectReplicationPolicyArgs()
        {
        }
        public static new ObjectReplicationPolicyArgs Empty => new ObjectReplicationPolicyArgs();
    }
}
