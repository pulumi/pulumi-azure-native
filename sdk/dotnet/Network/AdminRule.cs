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
    /// Network admin rule.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2021-02-01-preview, 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-01-01-preview, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:AdminRule")]
    public partial class AdminRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates the access allowed for this particular rule
        /// </summary>
        [Output("access")]
        public Output<string> Access { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// A description for this rule. Restricted to 140 chars.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The destination port ranges.
        /// </summary>
        [Output("destinationPortRanges")]
        public Output<ImmutableArray<string>> DestinationPortRanges { get; private set; } = null!;

        /// <summary>
        /// The destination address prefixes. CIDR or destination IP ranges.
        /// </summary>
        [Output("destinations")]
        public Output<ImmutableArray<Outputs.AddressPrefixItemResponse>> Destinations { get; private set; } = null!;

        /// <summary>
        /// Indicates if the traffic matched against the rule in inbound or outbound.
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Whether the rule is custom or default.
        /// Expected value is 'Custom'.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The priority of the rule. The value can be between 1 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Network protocol this rule applies to.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for this resource.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string> ResourceGuid { get; private set; } = null!;

        /// <summary>
        /// The source port ranges.
        /// </summary>
        [Output("sourcePortRanges")]
        public Output<ImmutableArray<string>> SourcePortRanges { get; private set; } = null!;

        /// <summary>
        /// The CIDR or source IP ranges.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.AddressPrefixItemResponse>> Sources { get; private set; } = null!;

        /// <summary>
        /// The system metadata related to this resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AdminRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AdminRule(string name, AdminRuleArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:AdminRule", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private AdminRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:AdminRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static AdminRuleArgs MakeArgs(AdminRuleArgs args)
        {
            args ??= new AdminRuleArgs();
            args.Kind = "Custom";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201preview:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501preview:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501preview:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220201preview:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220401preview:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101preview:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101preview:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:DefaultAdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:AdminRule" },
                    new global::Pulumi.Alias { Type = "azure-native:network:DefaultAdminRule" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AdminRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AdminRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AdminRule(name, id, options);
        }
    }

    public sealed class AdminRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the access allowed for this particular rule
        /// </summary>
        [Input("access", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.SecurityConfigurationRuleAccess> Access { get; set; } = null!;

        /// <summary>
        /// The name of the network manager Security Configuration.
        /// </summary>
        [Input("configurationName", required: true)]
        public Input<string> ConfigurationName { get; set; } = null!;

        /// <summary>
        /// A description for this rule. Restricted to 140 chars.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationPortRanges")]
        private InputList<string>? _destinationPortRanges;

        /// <summary>
        /// The destination port ranges.
        /// </summary>
        public InputList<string> DestinationPortRanges
        {
            get => _destinationPortRanges ?? (_destinationPortRanges = new InputList<string>());
            set => _destinationPortRanges = value;
        }

        [Input("destinations")]
        private InputList<Inputs.AddressPrefixItemArgs>? _destinations;

        /// <summary>
        /// The destination address prefixes. CIDR or destination IP ranges.
        /// </summary>
        public InputList<Inputs.AddressPrefixItemArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.AddressPrefixItemArgs>());
            set => _destinations = value;
        }

        /// <summary>
        /// Indicates if the traffic matched against the rule in inbound or outbound.
        /// </summary>
        [Input("direction", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.SecurityConfigurationRuleDirection> Direction { get; set; } = null!;

        /// <summary>
        /// Whether the rule is custom or default.
        /// Expected value is 'Custom'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The name of the network manager.
        /// </summary>
        [Input("networkManagerName", required: true)]
        public Input<string> NetworkManagerName { get; set; } = null!;

        /// <summary>
        /// The priority of the rule. The value can be between 1 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// Network protocol this rule applies to.
        /// </summary>
        [Input("protocol", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.SecurityConfigurationRuleProtocol> Protocol { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the network manager security Configuration rule collection.
        /// </summary>
        [Input("ruleCollectionName", required: true)]
        public Input<string> RuleCollectionName { get; set; } = null!;

        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        [Input("sourcePortRanges")]
        private InputList<string>? _sourcePortRanges;

        /// <summary>
        /// The source port ranges.
        /// </summary>
        public InputList<string> SourcePortRanges
        {
            get => _sourcePortRanges ?? (_sourcePortRanges = new InputList<string>());
            set => _sourcePortRanges = value;
        }

        [Input("sources")]
        private InputList<Inputs.AddressPrefixItemArgs>? _sources;

        /// <summary>
        /// The CIDR or source IP ranges.
        /// </summary>
        public InputList<Inputs.AddressPrefixItemArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.AddressPrefixItemArgs>());
            set => _sources = value;
        }

        public AdminRuleArgs()
        {
        }
        public static new AdminRuleArgs Empty => new AdminRuleArgs();
    }
}
