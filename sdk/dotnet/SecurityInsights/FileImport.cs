// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights
{
    /// <summary>
    /// Represents a file import in Azure Security Insights.
    /// 
    /// Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.
    /// 
    /// Other available API versions: 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:securityinsights:FileImport")]
    public partial class FileImport : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The content type of this file.
        /// </summary>
        [Output("contentType")]
        public Output<string> ContentType { get; private set; } = null!;

        /// <summary>
        /// The time the file was imported.
        /// </summary>
        [Output("createdTimeUTC")]
        public Output<string> CreatedTimeUTC { get; private set; } = null!;

        /// <summary>
        /// Represents the error file (if the import was ingested with errors or failed the validation).
        /// </summary>
        [Output("errorFile")]
        public Output<Outputs.FileMetadataResponse> ErrorFile { get; private set; } = null!;

        /// <summary>
        /// An ordered list of some of the errors that were encountered during validation.
        /// </summary>
        [Output("errorsPreview")]
        public Output<ImmutableArray<Outputs.ValidationErrorResponse>> ErrorsPreview { get; private set; } = null!;

        /// <summary>
        /// The time the files associated with this import are deleted from the storage account.
        /// </summary>
        [Output("filesValidUntilTimeUTC")]
        public Output<string> FilesValidUntilTimeUTC { get; private set; } = null!;

        /// <summary>
        /// Represents the imported file.
        /// </summary>
        [Output("importFile")]
        public Output<Outputs.FileMetadataResponse> ImportFile { get; private set; } = null!;

        /// <summary>
        /// The time the file import record is soft deleted from the database and history.
        /// </summary>
        [Output("importValidUntilTimeUTC")]
        public Output<string> ImportValidUntilTimeUTC { get; private set; } = null!;

        /// <summary>
        /// The number of records that have been successfully ingested.
        /// </summary>
        [Output("ingestedRecordCount")]
        public Output<int> IngestedRecordCount { get; private set; } = null!;

        /// <summary>
        /// Describes how to ingest the records in the file.
        /// </summary>
        [Output("ingestionMode")]
        public Output<string> IngestionMode { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The source for the data in the file.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// The state of the file import.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The number of records in the file.
        /// </summary>
        [Output("totalRecordCount")]
        public Output<int> TotalRecordCount { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The number of records that have passed validation.
        /// </summary>
        [Output("validRecordCount")]
        public Output<int> ValidRecordCount { get; private set; } = null!;


        /// <summary>
        /// Create a FileImport resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FileImport(string name, FileImportArgs args, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights:FileImport", name, args ?? new FileImportArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FileImport(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights:FileImport", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220801preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220901preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221001preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221101preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221201preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230201preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230301preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230401preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230501preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230601preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230701preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230801preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230901preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231001preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231201preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240101preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240401preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20241001preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250101preview:FileImport" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250401preview:FileImport" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FileImport resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FileImport Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FileImport(name, id, options);
        }
    }

    public sealed class FileImportArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content type of this file.
        /// </summary>
        [Input("contentType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.FileImportContentType> ContentType { get; set; } = null!;

        /// <summary>
        /// File import ID
        /// </summary>
        [Input("fileImportId")]
        public Input<string>? FileImportId { get; set; }

        /// <summary>
        /// Represents the imported file.
        /// </summary>
        [Input("importFile", required: true)]
        public Input<Inputs.FileMetadataArgs> ImportFile { get; set; } = null!;

        /// <summary>
        /// Describes how to ingest the records in the file.
        /// </summary>
        [Input("ingestionMode", required: true)]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.IngestionMode> IngestionMode { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The source for the data in the file.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public FileImportArgs()
        {
        }
        public static new FileImportArgs Empty => new FileImportArgs();
    }
}
