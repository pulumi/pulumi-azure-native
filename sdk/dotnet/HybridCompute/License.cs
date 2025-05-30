// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute
{
    /// <summary>
    /// Describes a license in a hybrid machine.
    /// 
    /// Uses Azure REST API version 2024-07-10. In version 2.x of the Azure Native provider, it used API version 2023-06-20-preview.
    /// 
    /// Other available API versions: 2023-06-20-preview, 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-07-31-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13, 2025-02-19-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcompute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:hybridcompute:License")]
    public partial class License : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Describes the properties of a License.
        /// </summary>
        [Output("licenseDetails")]
        public Output<Outputs.LicenseDetailsResponse?> LicenseDetails { get; private set; } = null!;

        /// <summary>
        /// The type of the license resource.
        /// </summary>
        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Describes the tenant id.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a License resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public License(string name, LicenseArgs args, CustomResourceOptions? options = null)
            : base("azure-native:hybridcompute:License", name, args ?? new LicenseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private License(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:hybridcompute:License", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20230620preview:License" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20231003preview:License" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20240331preview:License" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20240520preview:License" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20240710:License" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20240731preview:License" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20240910preview:License" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20241110preview:License" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20250113:License" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcompute/v20250219preview:License" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing License resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static License Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new License(name, id, options);
        }
    }

    public sealed class LicenseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the properties of a License.
        /// </summary>
        [Input("licenseDetails")]
        public Input<Inputs.LicenseDetailsArgs>? LicenseDetails { get; set; }

        /// <summary>
        /// The name of the license.
        /// </summary>
        [Input("licenseName")]
        public Input<string>? LicenseName { get; set; }

        /// <summary>
        /// The type of the license resource.
        /// </summary>
        [Input("licenseType")]
        public InputUnion<string, Pulumi.AzureNative.HybridCompute.LicenseType>? LicenseType { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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
        /// Describes the tenant id.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public LicenseArgs()
        {
        }
        public static new LicenseArgs Empty => new LicenseArgs();
    }
}
