// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20250301
{
    /// <summary>
    /// Represents Premium Microsoft Defender for Threat Intelligence data connector.
    /// </summary>
    [AzureNativeResourceType("azure-native:securityinsights/v20250301:PremiumMicrosoftDefenderForThreatIntelligence")]
    public partial class PremiumMicrosoftDefenderForThreatIntelligence : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The available data types for the connector.
        /// </summary>
        [Output("dataTypes")]
        public Output<Outputs.PremiumMdtiDataConnectorDataTypesResponse> DataTypes { get; private set; } = null!;

        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The kind of the data connector
        /// Expected value is 'PremiumMicrosoftDefenderForThreatIntelligence'.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The lookback period for the feed to be imported. The date-time to begin importing the feed from, for example: 2024-01-01T00:00:00.000Z.
        /// </summary>
        [Output("lookbackPeriod")]
        public Output<string> LookbackPeriod { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The flag to indicate whether the tenant has the premium SKU required to access this connector.
        /// </summary>
        [Output("requiredSKUsPresent")]
        public Output<bool?> RequiredSKUsPresent { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The tenant id to connect to, and get the data from.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PremiumMicrosoftDefenderForThreatIntelligence resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PremiumMicrosoftDefenderForThreatIntelligence(string name, PremiumMicrosoftDefenderForThreatIntelligenceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights/v20250301:PremiumMicrosoftDefenderForThreatIntelligence", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private PremiumMicrosoftDefenderForThreatIntelligence(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights/v20250301:PremiumMicrosoftDefenderForThreatIntelligence", name, null, MakeResourceOptions(options, id))
        {
        }

        private static PremiumMicrosoftDefenderForThreatIntelligenceArgs MakeArgs(PremiumMicrosoftDefenderForThreatIntelligenceArgs args)
        {
            args ??= new PremiumMicrosoftDefenderForThreatIntelligenceArgs();
            args.Kind = "PremiumMicrosoftDefenderForThreatIntelligence";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20190101preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20200101:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20210301preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20210901preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20211001:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20211001preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220101preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220401preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220501preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220601preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220701preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220801:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220801preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220901preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221001preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221101:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221101preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221201preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230201:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230201preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230301preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230401preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230501preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230601preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230701preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230801preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230901preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231001preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231101:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231201preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240101preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240301:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240401preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240901:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20241001preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250101preview:PremiumMicrosoftDefenderForThreatIntelligence" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights:PremiumMicrosoftDefenderForThreatIntelligence" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PremiumMicrosoftDefenderForThreatIntelligence resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PremiumMicrosoftDefenderForThreatIntelligence Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PremiumMicrosoftDefenderForThreatIntelligence(name, id, options);
        }
    }

    public sealed class PremiumMicrosoftDefenderForThreatIntelligenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connector ID
        /// </summary>
        [Input("dataConnectorId")]
        public Input<string>? DataConnectorId { get; set; }

        /// <summary>
        /// The available data types for the connector.
        /// </summary>
        [Input("dataTypes", required: true)]
        public Input<Inputs.PremiumMdtiDataConnectorDataTypesArgs> DataTypes { get; set; } = null!;

        /// <summary>
        /// The kind of the data connector
        /// Expected value is 'PremiumMicrosoftDefenderForThreatIntelligence'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The lookback period for the feed to be imported. The date-time to begin importing the feed from, for example: 2024-01-01T00:00:00.000Z.
        /// </summary>
        [Input("lookbackPeriod", required: true)]
        public Input<string> LookbackPeriod { get; set; } = null!;

        /// <summary>
        /// The flag to indicate whether the tenant has the premium SKU required to access this connector.
        /// </summary>
        [Input("requiredSKUsPresent")]
        public Input<bool>? RequiredSKUsPresent { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The tenant id to connect to, and get the data from.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public PremiumMicrosoftDefenderForThreatIntelligenceArgs()
        {
        }
        public static new PremiumMicrosoftDefenderForThreatIntelligenceArgs Empty => new PremiumMicrosoftDefenderForThreatIntelligenceArgs();
    }
}
