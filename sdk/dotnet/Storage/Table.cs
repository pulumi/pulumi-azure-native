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
    /// Properties of the table, including Id, resource name, resource type.
    /// 
    /// Uses Azure REST API version 2024-01-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
    /// 
    /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:storage:Table")]
    public partial class Table : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of stored access policies specified on the table.
        /// </summary>
        [Output("signedIdentifiers")]
        public Output<ImmutableArray<Outputs.TableSignedIdentifierResponse>> SignedIdentifiers { get; private set; } = null!;

        /// <summary>
        /// Table name under the specified account
        /// </summary>
        [Output("tableName")]
        public Output<string> TableName { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Table resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Table(string name, TableArgs args, CustomResourceOptions? options = null)
            : base("azure-native:storage:Table", name, args ?? new TableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Table(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:storage:Table", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20190601:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20200801preview:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210101:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210201:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210401:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210601:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210801:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210901:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20220501:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20220901:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20230101:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20230401:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20230501:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20240101:Table" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20250101:Table" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Table resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Table Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Table(name, id, options);
        }
    }

    public sealed class TableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("signedIdentifiers")]
        private InputList<Inputs.TableSignedIdentifierArgs>? _signedIdentifiers;

        /// <summary>
        /// List of stored access policies specified on the table.
        /// </summary>
        public InputList<Inputs.TableSignedIdentifierArgs> SignedIdentifiers
        {
            get => _signedIdentifiers ?? (_signedIdentifiers = new InputList<Inputs.TableSignedIdentifierArgs>());
            set => _signedIdentifiers = value;
        }

        /// <summary>
        /// A table name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of only alphanumeric characters and it cannot begin with a numeric character.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public TableArgs()
        {
        }
        public static new TableArgs Empty => new TableArgs();
    }
}
