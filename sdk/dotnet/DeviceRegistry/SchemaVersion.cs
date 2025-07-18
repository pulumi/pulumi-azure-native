// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeviceRegistry
{
    /// <summary>
    /// Schema version's definition.
    /// 
    /// Uses Azure REST API version 2024-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-09-01-preview.
    /// 
    /// Other available API versions: 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:deviceregistry:SchemaVersion")]
    public partial class SchemaVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Human-readable description of the schema.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Hash of the schema content.
        /// </summary>
        [Output("hash")]
        public Output<string> Hash { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Schema content.
        /// </summary>
        [Output("schemaContent")]
        public Output<string> SchemaContent { get; private set; } = null!;

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
        /// Globally unique, immutable, non-reusable id.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a SchemaVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SchemaVersion(string name, SchemaVersionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:deviceregistry:SchemaVersion", name, args ?? new SchemaVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SchemaVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:deviceregistry:SchemaVersion", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:deviceregistry/v20240901preview:SchemaVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:deviceregistry/v20250701preview:SchemaVersion" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SchemaVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SchemaVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SchemaVersion(name, id, options);
        }
    }

    public sealed class SchemaVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-readable description of the schema.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Schema content.
        /// </summary>
        [Input("schemaContent", required: true)]
        public Input<string> SchemaContent { get; set; } = null!;

        /// <summary>
        /// Schema name parameter.
        /// </summary>
        [Input("schemaName", required: true)]
        public Input<string> SchemaName { get; set; } = null!;

        /// <summary>
        /// Schema registry name parameter.
        /// </summary>
        [Input("schemaRegistryName", required: true)]
        public Input<string> SchemaRegistryName { get; set; } = null!;

        /// <summary>
        /// Schema version name parameter.
        /// </summary>
        [Input("schemaVersionName")]
        public Input<string>? SchemaVersionName { get; set; }

        public SchemaVersionArgs()
        {
        }
        public static new SchemaVersionArgs Empty => new SchemaVersionArgs();
    }
}
