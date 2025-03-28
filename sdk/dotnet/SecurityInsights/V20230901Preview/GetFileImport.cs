// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20230901Preview
{
    public static class GetFileImport
    {
        /// <summary>
        /// Gets a file import.
        /// </summary>
        public static Task<GetFileImportResult> InvokeAsync(GetFileImportArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFileImportResult>("azure-native:securityinsights/v20230901preview:getFileImport", args ?? new GetFileImportArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a file import.
        /// </summary>
        public static Output<GetFileImportResult> Invoke(GetFileImportInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFileImportResult>("azure-native:securityinsights/v20230901preview:getFileImport", args ?? new GetFileImportInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a file import.
        /// </summary>
        public static Output<GetFileImportResult> Invoke(GetFileImportInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFileImportResult>("azure-native:securityinsights/v20230901preview:getFileImport", args ?? new GetFileImportInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileImportArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File import ID
        /// </summary>
        [Input("fileImportId", required: true)]
        public string FileImportId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetFileImportArgs()
        {
        }
        public static new GetFileImportArgs Empty => new GetFileImportArgs();
    }

    public sealed class GetFileImportInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File import ID
        /// </summary>
        [Input("fileImportId", required: true)]
        public Input<string> FileImportId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetFileImportInvokeArgs()
        {
        }
        public static new GetFileImportInvokeArgs Empty => new GetFileImportInvokeArgs();
    }


    [OutputType]
    public sealed class GetFileImportResult
    {
        /// <summary>
        /// The content type of this file.
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// The time the file was imported.
        /// </summary>
        public readonly string CreatedTimeUTC;
        /// <summary>
        /// Represents the error file (if the import was ingested with errors or failed the validation).
        /// </summary>
        public readonly Outputs.FileMetadataResponse ErrorFile;
        /// <summary>
        /// An ordered list of some of the errors that were encountered during validation.
        /// </summary>
        public readonly ImmutableArray<Outputs.ValidationErrorResponse> ErrorsPreview;
        /// <summary>
        /// The time the files associated with this import are deleted from the storage account.
        /// </summary>
        public readonly string FilesValidUntilTimeUTC;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Represents the imported file.
        /// </summary>
        public readonly Outputs.FileMetadataResponse ImportFile;
        /// <summary>
        /// The time the file import record is soft deleted from the database and history.
        /// </summary>
        public readonly string ImportValidUntilTimeUTC;
        /// <summary>
        /// The number of records that have been successfully ingested.
        /// </summary>
        public readonly int IngestedRecordCount;
        /// <summary>
        /// Describes how to ingest the records in the file.
        /// </summary>
        public readonly string IngestionMode;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The source for the data in the file.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// The state of the file import.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The number of records in the file.
        /// </summary>
        public readonly int TotalRecordCount;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The number of records that have passed validation.
        /// </summary>
        public readonly int ValidRecordCount;

        [OutputConstructor]
        private GetFileImportResult(
            string contentType,

            string createdTimeUTC,

            Outputs.FileMetadataResponse errorFile,

            ImmutableArray<Outputs.ValidationErrorResponse> errorsPreview,

            string filesValidUntilTimeUTC,

            string id,

            Outputs.FileMetadataResponse importFile,

            string importValidUntilTimeUTC,

            int ingestedRecordCount,

            string ingestionMode,

            string name,

            string source,

            string state,

            Outputs.SystemDataResponse systemData,

            int totalRecordCount,

            string type,

            int validRecordCount)
        {
            ContentType = contentType;
            CreatedTimeUTC = createdTimeUTC;
            ErrorFile = errorFile;
            ErrorsPreview = errorsPreview;
            FilesValidUntilTimeUTC = filesValidUntilTimeUTC;
            Id = id;
            ImportFile = importFile;
            ImportValidUntilTimeUTC = importValidUntilTimeUTC;
            IngestedRecordCount = ingestedRecordCount;
            IngestionMode = ingestionMode;
            Name = name;
            Source = source;
            State = state;
            SystemData = systemData;
            TotalRecordCount = totalRecordCount;
            Type = type;
            ValidRecordCount = validRecordCount;
        }
    }
}
