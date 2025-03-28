// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.V20241001Preview.Outputs
{

    /// <summary>
    /// An export run.
    /// </summary>
    [OutputType]
    public sealed class ExportRunResponse
    {
        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        public readonly string? ETag;
        /// <summary>
        /// The end datetime for the export.
        /// </summary>
        public readonly string? EndDate;
        /// <summary>
        /// The details of any error.
        /// </summary>
        public readonly Outputs.ErrorDetailsResponse? Error;
        /// <summary>
        /// The type of the export run.
        /// </summary>
        public readonly string? ExecutionType;
        /// <summary>
        /// The name of the exported file.
        /// </summary>
        public readonly string? FileName;
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The manifest file location(URI location) for the exported files.
        /// </summary>
        public readonly string? ManifestFile;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The time when the export run finished.
        /// </summary>
        public readonly string? ProcessingEndTime;
        /// <summary>
        /// The time when export was picked up to be run.
        /// </summary>
        public readonly string? ProcessingStartTime;
        /// <summary>
        /// The export settings that were in effect for this run.
        /// </summary>
        public readonly Outputs.CommonExportPropertiesResponse? RunSettings;
        /// <summary>
        /// The start datetime for the export.
        /// </summary>
        public readonly string? StartDate;
        /// <summary>
        /// The last known status of the export run.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The identifier for the entity that triggered the export. For on-demand runs it is the user email. For scheduled runs it is 'System'.
        /// </summary>
        public readonly string? SubmittedBy;
        /// <summary>
        /// The time when export was queued to be run.
        /// </summary>
        public readonly string? SubmittedTime;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ExportRunResponse(
            string? eTag,

            string? endDate,

            Outputs.ErrorDetailsResponse? error,

            string? executionType,

            string? fileName,

            string id,

            string? manifestFile,

            string name,

            string? processingEndTime,

            string? processingStartTime,

            Outputs.CommonExportPropertiesResponse? runSettings,

            string? startDate,

            string? status,

            string? submittedBy,

            string? submittedTime,

            string type)
        {
            ETag = eTag;
            EndDate = endDate;
            Error = error;
            ExecutionType = executionType;
            FileName = fileName;
            Id = id;
            ManifestFile = manifestFile;
            Name = name;
            ProcessingEndTime = processingEndTime;
            ProcessingStartTime = processingStartTime;
            RunSettings = runSettings;
            StartDate = startDate;
            Status = status;
            SubmittedBy = submittedBy;
            SubmittedTime = submittedTime;
            Type = type;
        }
    }
}
