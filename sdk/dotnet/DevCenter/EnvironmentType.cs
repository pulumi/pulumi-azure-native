// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter
{
    /// <summary>
    /// Represents an environment type.
    /// 
    /// Uses Azure REST API version 2024-02-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
    /// 
    /// Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:devcenter:EnvironmentType")]
    public partial class EnvironmentType : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The display name of the environment type.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
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
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EnvironmentType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvironmentType(string name, EnvironmentTypeArgs args, CustomResourceOptions? options = null)
            : base("azure-native:devcenter:EnvironmentType", name, args ?? new EnvironmentTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnvironmentType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:devcenter:EnvironmentType", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20220801preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20220901preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20221012preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20221111preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20230101preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20230401:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20230801preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20231001preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20240201:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20240501preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20240601preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20240701preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20240801preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20241001preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20250201:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20250401preview:EnvironmentType" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20250701preview:EnvironmentType" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EnvironmentType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvironmentType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EnvironmentType(name, id, options);
        }
    }

    public sealed class EnvironmentTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the devcenter.
        /// </summary>
        [Input("devCenterName", required: true)]
        public Input<string> DevCenterName { get; set; } = null!;

        /// <summary>
        /// The display name of the environment type.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the environment type.
        /// </summary>
        [Input("environmentTypeName")]
        public Input<string>? EnvironmentTypeName { get; set; }

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

        public EnvironmentTypeArgs()
        {
        }
        public static new EnvironmentTypeArgs Empty => new EnvironmentTypeArgs();
    }
}
