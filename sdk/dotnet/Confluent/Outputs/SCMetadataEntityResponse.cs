// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Confluent.Outputs
{

    /// <summary>
    /// Metadata of the data record
    /// </summary>
    [OutputType]
    public sealed class SCMetadataEntityResponse
    {
        /// <summary>
        /// Created Date Time
        /// </summary>
        public readonly string? CreatedTimestamp;
        /// <summary>
        /// Deleted Date time
        /// </summary>
        public readonly string? DeletedTimestamp;
        /// <summary>
        /// Resource name of the record
        /// </summary>
        public readonly string? ResourceName;
        /// <summary>
        /// Self lookup url
        /// </summary>
        public readonly string? Self;
        /// <summary>
        /// Updated Date time
        /// </summary>
        public readonly string? UpdatedTimestamp;

        [OutputConstructor]
        private SCMetadataEntityResponse(
            string? createdTimestamp,

            string? deletedTimestamp,

            string? resourceName,

            string? self,

            string? updatedTimestamp)
        {
            CreatedTimestamp = createdTimestamp;
            DeletedTimestamp = deletedTimestamp;
            ResourceName = resourceName;
            Self = self;
            UpdatedTimestamp = updatedTimestamp;
        }
    }
}
