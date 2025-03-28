// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConfidentialLedger
{
    /// <summary>
    /// Confidential Ledger. Contains the properties of Confidential Ledger Resource.
    /// 
    /// Uses Azure REST API version 2022-05-13. In version 1.x of the Azure Native provider, it used API version 2020-12-01-preview.
    /// 
    /// Other available API versions: 2023-01-26-preview, 2023-06-28-preview, 2024-07-09-preview, 2024-09-19-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:confidentialledger:Ledger")]
    public partial class Ledger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure location where the Confidential Ledger is running.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the Resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of Confidential Ledger Resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.LedgerPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Additional tags for Confidential Ledger
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Ledger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ledger(string name, LedgerArgs args, CustomResourceOptions? options = null)
            : base("azure-native:confidentialledger:Ledger", name, args ?? new LedgerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ledger(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:confidentialledger:Ledger", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:confidentialledger/v20201201preview:Ledger" },
                    new global::Pulumi.Alias { Type = "azure-native:confidentialledger/v20210513preview:Ledger" },
                    new global::Pulumi.Alias { Type = "azure-native:confidentialledger/v20220513:Ledger" },
                    new global::Pulumi.Alias { Type = "azure-native:confidentialledger/v20220908preview:Ledger" },
                    new global::Pulumi.Alias { Type = "azure-native:confidentialledger/v20230126preview:Ledger" },
                    new global::Pulumi.Alias { Type = "azure-native:confidentialledger/v20230628preview:Ledger" },
                    new global::Pulumi.Alias { Type = "azure-native:confidentialledger/v20240709preview:Ledger" },
                    new global::Pulumi.Alias { Type = "azure-native:confidentialledger/v20240919preview:Ledger" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Ledger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ledger Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Ledger(name, id, options);
        }
    }

    public sealed class LedgerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Confidential Ledger
        /// </summary>
        [Input("ledgerName")]
        public Input<string>? LedgerName { get; set; }

        /// <summary>
        /// The Azure location where the Confidential Ledger is running.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Properties of Confidential Ledger Resource.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.LedgerPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Additional tags for Confidential Ledger
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public LedgerArgs()
        {
        }
        public static new LedgerArgs Empty => new LedgerArgs();
    }
}
