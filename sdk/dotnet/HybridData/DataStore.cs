// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridData
{
    /// <summary>
    /// Data store.
    /// 
    /// Uses Azure REST API version 2019-06-01. In version 2.x of the Azure Native provider, it used API version 2019-06-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:hybriddata:DataStore")]
    public partial class DataStore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
        /// </summary>
        [Output("customerSecrets")]
        public Output<ImmutableArray<Outputs.CustomerSecretResponse>> CustomerSecrets { get; private set; } = null!;

        /// <summary>
        /// The arm id of the data store type.
        /// </summary>
        [Output("dataStoreTypeId")]
        public Output<string> DataStoreTypeId { get; private set; } = null!;

        /// <summary>
        /// A generic json used differently by each data source type.
        /// </summary>
        [Output("extendedProperties")]
        public Output<object?> ExtendedProperties { get; private set; } = null!;

        /// <summary>
        /// Name of the object.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Arm Id for the manager resource to which the data source is associated. This is optional.
        /// </summary>
        [Output("repositoryId")]
        public Output<string?> RepositoryId { get; private set; } = null!;

        /// <summary>
        /// State of the data source.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DataStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataStore(string name, DataStoreArgs args, CustomResourceOptions? options = null)
            : base("azure-native:hybriddata:DataStore", name, args ?? new DataStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataStore(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:hybriddata:DataStore", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:hybriddata/v20160601:DataStore" },
                    new global::Pulumi.Alias { Type = "azure-native:hybriddata/v20190601:DataStore" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataStore Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataStore(name, id, options);
        }
    }

    public sealed class DataStoreArgs : global::Pulumi.ResourceArgs
    {
        [Input("customerSecrets")]
        private InputList<Inputs.CustomerSecretArgs>? _customerSecrets;

        /// <summary>
        /// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
        /// </summary>
        public InputList<Inputs.CustomerSecretArgs> CustomerSecrets
        {
            get => _customerSecrets ?? (_customerSecrets = new InputList<Inputs.CustomerSecretArgs>());
            set => _customerSecrets = value;
        }

        /// <summary>
        /// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
        /// </summary>
        [Input("dataManagerName", required: true)]
        public Input<string> DataManagerName { get; set; } = null!;

        /// <summary>
        /// The data store/repository name to be created or updated.
        /// </summary>
        [Input("dataStoreName")]
        public Input<string>? DataStoreName { get; set; }

        /// <summary>
        /// The arm id of the data store type.
        /// </summary>
        [Input("dataStoreTypeId", required: true)]
        public Input<string> DataStoreTypeId { get; set; } = null!;

        /// <summary>
        /// A generic json used differently by each data source type.
        /// </summary>
        [Input("extendedProperties")]
        public Input<object>? ExtendedProperties { get; set; }

        /// <summary>
        /// Arm Id for the manager resource to which the data source is associated. This is optional.
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        /// <summary>
        /// The Resource Group Name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// State of the data source.
        /// </summary>
        [Input("state", required: true)]
        public Input<Pulumi.AzureNative.HybridData.State> State { get; set; } = null!;

        public DataStoreArgs()
        {
        }
        public static new DataStoreArgs Empty => new DataStoreArgs();
    }
}
