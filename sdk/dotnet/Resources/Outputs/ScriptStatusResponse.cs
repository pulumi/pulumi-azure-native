// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Resources.Outputs
{

    /// <summary>
    /// Generic object modeling results of script execution.
    /// </summary>
    [OutputType]
    public sealed class ScriptStatusResponse
    {
        /// <summary>
        /// ACI resource Id.
        /// </summary>
        public readonly string ContainerInstanceId;
        /// <summary>
        /// End time of the script execution.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Error that is relayed from the script execution.
        /// </summary>
        public readonly Outputs.ErrorResponseResponse? Error;
        /// <summary>
        /// Time the deployment script resource will expire.
        /// </summary>
        public readonly string ExpirationTime;
        /// <summary>
        /// Start time of the script execution.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Storage account resource Id.
        /// </summary>
        public readonly string StorageAccountId;

        [OutputConstructor]
        private ScriptStatusResponse(
            string containerInstanceId,

            string endTime,

            Outputs.ErrorResponseResponse? error,

            string expirationTime,

            string startTime,

            string storageAccountId)
        {
            ContainerInstanceId = containerInstanceId;
            EndTime = endTime;
            Error = error;
            ExpirationTime = expirationTime;
            StartTime = startTime;
            StorageAccountId = storageAccountId;
        }
    }
}
