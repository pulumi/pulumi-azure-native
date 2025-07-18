// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw
{
    /// <summary>
    /// GlobalRulestack prefixList
    /// 
    /// Uses Azure REST API version 2025-02-06-preview. In version 2.x of the Azure Native provider, it used API version 2023-09-01.
    /// 
    /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:cloudngfw:PrefixListGlobalRulestack")]
    public partial class PrefixListGlobalRulestack : global::Pulumi.CustomResource
    {
        /// <summary>
        /// comment for this object
        /// </summary>
        [Output("auditComment")]
        public Output<string?> AuditComment { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// prefix description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// etag info
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// prefix list
        /// </summary>
        [Output("prefixList")]
        public Output<ImmutableArray<string>> PrefixList { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Create a PrefixListGlobalRulestack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrefixListGlobalRulestack(string name, PrefixListGlobalRulestackArgs args, CustomResourceOptions? options = null)
            : base("azure-native:cloudngfw:PrefixListGlobalRulestack", name, args ?? new PrefixListGlobalRulestackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrefixListGlobalRulestack(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:cloudngfw:PrefixListGlobalRulestack", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:cloudngfw/v20220829:PrefixListGlobalRulestack" },
                    new global::Pulumi.Alias { Type = "azure-native:cloudngfw/v20220829preview:PrefixListGlobalRulestack" },
                    new global::Pulumi.Alias { Type = "azure-native:cloudngfw/v20230901:PrefixListGlobalRulestack" },
                    new global::Pulumi.Alias { Type = "azure-native:cloudngfw/v20230901preview:PrefixListGlobalRulestack" },
                    new global::Pulumi.Alias { Type = "azure-native:cloudngfw/v20231010preview:PrefixListGlobalRulestack" },
                    new global::Pulumi.Alias { Type = "azure-native:cloudngfw/v20240119preview:PrefixListGlobalRulestack" },
                    new global::Pulumi.Alias { Type = "azure-native:cloudngfw/v20240207preview:PrefixListGlobalRulestack" },
                    new global::Pulumi.Alias { Type = "azure-native:cloudngfw/v20250206preview:PrefixListGlobalRulestack" },
                    new global::Pulumi.Alias { Type = "azure-native:cloudngfw/v20250523:PrefixListGlobalRulestack" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrefixListGlobalRulestack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrefixListGlobalRulestack Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrefixListGlobalRulestack(name, id, options);
        }
    }

    public sealed class PrefixListGlobalRulestackArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// comment for this object
        /// </summary>
        [Input("auditComment")]
        public Input<string>? AuditComment { get; set; }

        /// <summary>
        /// prefix description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// GlobalRulestack resource name
        /// </summary>
        [Input("globalRulestackName", required: true)]
        public Input<string> GlobalRulestackName { get; set; } = null!;

        /// <summary>
        /// Local Rule priority
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("prefixList", required: true)]
        private InputList<string>? _prefixList;

        /// <summary>
        /// prefix list
        /// </summary>
        public InputList<string> PrefixList
        {
            get => _prefixList ?? (_prefixList = new InputList<string>());
            set => _prefixList = value;
        }

        public PrefixListGlobalRulestackArgs()
        {
        }
        public static new PrefixListGlobalRulestackArgs Empty => new PrefixListGlobalRulestackArgs();
    }
}
