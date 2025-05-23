// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App
{
    /// <summary>
    /// Dapr Component.
    /// 
    /// Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2022-10-01.
    /// 
    /// Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:app:DaprComponent")]
    public partial class DaprComponent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Component type
        /// </summary>
        [Output("componentType")]
        public Output<string?> ComponentType { get; private set; } = null!;

        /// <summary>
        /// Boolean describing if the component errors are ignores
        /// </summary>
        [Output("ignoreErrors")]
        public Output<bool?> IgnoreErrors { get; private set; } = null!;

        /// <summary>
        /// Initialization timeout
        /// </summary>
        [Output("initTimeout")]
        public Output<string?> InitTimeout { get; private set; } = null!;

        /// <summary>
        /// Component metadata
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableArray<Outputs.DaprMetadataResponse>> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Names of container apps that can use this Dapr component
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Name of a Dapr component to retrieve component secrets from
        /// </summary>
        [Output("secretStoreComponent")]
        public Output<string?> SecretStoreComponent { get; private set; } = null!;

        /// <summary>
        /// Collection of secrets used by a Dapr component
        /// </summary>
        [Output("secrets")]
        public Output<ImmutableArray<Outputs.SecretResponse>> Secrets { get; private set; } = null!;

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
        /// Component version
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a DaprComponent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DaprComponent(string name, DaprComponentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:app:DaprComponent", name, args ?? new DaprComponentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DaprComponent(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:app:DaprComponent", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:app/v20220101preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20220301:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20220601preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20221001:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20221101preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20230401preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20230501:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20230502preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20230801preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20231102preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20240202preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20240301:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20240802preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20241002preview:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20250101:DaprComponent" },
                    new global::Pulumi.Alias { Type = "azure-native:app/v20250202preview:DaprComponent" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DaprComponent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DaprComponent Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DaprComponent(name, id, options);
        }
    }

    public sealed class DaprComponentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Dapr Component.
        /// </summary>
        [Input("componentName")]
        public Input<string>? ComponentName { get; set; }

        /// <summary>
        /// Component type
        /// </summary>
        [Input("componentType")]
        public Input<string>? ComponentType { get; set; }

        /// <summary>
        /// Name of the Managed Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Boolean describing if the component errors are ignores
        /// </summary>
        [Input("ignoreErrors")]
        public Input<bool>? IgnoreErrors { get; set; }

        /// <summary>
        /// Initialization timeout
        /// </summary>
        [Input("initTimeout")]
        public Input<string>? InitTimeout { get; set; }

        [Input("metadata")]
        private InputList<Inputs.DaprMetadataArgs>? _metadata;

        /// <summary>
        /// Component metadata
        /// </summary>
        public InputList<Inputs.DaprMetadataArgs> Metadata
        {
            get => _metadata ?? (_metadata = new InputList<Inputs.DaprMetadataArgs>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// Names of container apps that can use this Dapr component
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// Name of a Dapr component to retrieve component secrets from
        /// </summary>
        [Input("secretStoreComponent")]
        public Input<string>? SecretStoreComponent { get; set; }

        [Input("secrets")]
        private InputList<Inputs.SecretArgs>? _secrets;

        /// <summary>
        /// Collection of secrets used by a Dapr component
        /// </summary>
        public InputList<Inputs.SecretArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.SecretArgs>());
            set => _secrets = value;
        }

        /// <summary>
        /// Component version
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public DaprComponentArgs()
        {
            IgnoreErrors = false;
        }
        public static new DaprComponentArgs Empty => new DaprComponentArgs();
    }
}
