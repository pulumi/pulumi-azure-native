// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataShare
{
    /// <summary>
    /// A Blob data set mapping.
    /// 
    /// Uses Azure REST API version 2021-08-01. In version 2.x of the Azure Native provider, it used API version 2021-08-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:datashare:BlobDataSetMapping")]
    public partial class BlobDataSetMapping : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Container that has the file path.
        /// </summary>
        [Output("containerName")]
        public Output<string> ContainerName { get; private set; } = null!;

        /// <summary>
        /// The id of the source data set.
        /// </summary>
        [Output("dataSetId")]
        public Output<string> DataSetId { get; private set; } = null!;

        /// <summary>
        /// Gets the status of the data set mapping.
        /// </summary>
        [Output("dataSetMappingStatus")]
        public Output<string> DataSetMappingStatus { get; private set; } = null!;

        /// <summary>
        /// File path within the source data set
        /// </summary>
        [Output("filePath")]
        public Output<string> FilePath { get; private set; } = null!;

        /// <summary>
        /// Kind of data set mapping.
        /// Expected value is 'Blob'.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the azure resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// File output type
        /// </summary>
        [Output("outputType")]
        public Output<string?> OutputType { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the data set mapping.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource group of storage account.
        /// </summary>
        [Output("resourceGroup")]
        public Output<string> ResourceGroup { get; private set; } = null!;

        /// <summary>
        /// Storage account name of the source data set.
        /// </summary>
        [Output("storageAccountName")]
        public Output<string> StorageAccountName { get; private set; } = null!;

        /// <summary>
        /// Subscription id of storage account.
        /// </summary>
        [Output("subscriptionId")]
        public Output<string> SubscriptionId { get; private set; } = null!;

        /// <summary>
        /// System Data of the Azure resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Type of the azure resource
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a BlobDataSetMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BlobDataSetMapping(string name, BlobDataSetMappingArgs args, CustomResourceOptions? options = null)
            : base("azure-native:datashare:BlobDataSetMapping", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private BlobDataSetMapping(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:datashare:BlobDataSetMapping", name, null, MakeResourceOptions(options, id))
        {
        }

        private static BlobDataSetMappingArgs MakeArgs(BlobDataSetMappingArgs args)
        {
            args ??= new BlobDataSetMappingArgs();
            args.Kind = "Blob";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20181101preview:BlobDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20191101:BlobDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20200901:BlobDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20201001preview:BlobDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20201001preview:BlobStorageAccountDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:ADLSGen2FileDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:ADLSGen2FileSystemDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:ADLSGen2FolderDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:BlobContainerDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:BlobDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:BlobFolderDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:KustoClusterDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:KustoDatabaseDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:KustoTableDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:SqlDBTableDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:SqlDWTableDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare/v20210801:SynapseWorkspaceSqlPoolTableDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:ADLSGen2FileDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:ADLSGen2FileSystemDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:ADLSGen2FolderDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:BlobContainerDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:BlobFolderDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:KustoClusterDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:KustoDatabaseDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:KustoTableDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:SqlDBTableDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:SqlDWTableDataSetMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:datashare:SynapseWorkspaceSqlPoolTableDataSetMapping" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BlobDataSetMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BlobDataSetMapping Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BlobDataSetMapping(name, id, options);
        }
    }

    public sealed class BlobDataSetMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Container that has the file path.
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// The id of the source data set.
        /// </summary>
        [Input("dataSetId", required: true)]
        public Input<string> DataSetId { get; set; } = null!;

        /// <summary>
        /// The name of the data set mapping to be created.
        /// </summary>
        [Input("dataSetMappingName")]
        public Input<string>? DataSetMappingName { get; set; }

        /// <summary>
        /// File path within the source data set
        /// </summary>
        [Input("filePath", required: true)]
        public Input<string> FilePath { get; set; } = null!;

        /// <summary>
        /// Kind of data set mapping.
        /// Expected value is 'Blob'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// File output type
        /// </summary>
        [Input("outputType")]
        public InputUnion<string, Pulumi.AzureNative.DataShare.OutputType>? OutputType { get; set; }

        /// <summary>
        /// Resource group of storage account.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public Input<string> ResourceGroup { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share subscription which will hold the data set sink.
        /// </summary>
        [Input("shareSubscriptionName", required: true)]
        public Input<string> ShareSubscriptionName { get; set; } = null!;

        /// <summary>
        /// Storage account name of the source data set.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        /// <summary>
        /// Subscription id of storage account.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        public BlobDataSetMappingArgs()
        {
        }
        public static new BlobDataSetMappingArgs Empty => new BlobDataSetMappingArgs();
    }
}
