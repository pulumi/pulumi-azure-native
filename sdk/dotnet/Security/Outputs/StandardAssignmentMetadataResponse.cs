// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// The standard assignment metadata
    /// </summary>
    [OutputType]
    public sealed class StandardAssignmentMetadataResponse
    {
        /// <summary>
        /// Standard assignment Created by object id (GUID)
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// Standard assignment creation date
        /// </summary>
        public readonly string CreatedOn;
        /// <summary>
        /// Standard assignment last updated by object id (GUID)
        /// </summary>
        public readonly string LastUpdatedBy;
        /// <summary>
        /// Standard assignment last update date
        /// </summary>
        public readonly string LastUpdatedOn;

        [OutputConstructor]
        private StandardAssignmentMetadataResponse(
            string createdBy,

            string createdOn,

            string lastUpdatedBy,

            string lastUpdatedOn)
        {
            CreatedBy = createdBy;
            CreatedOn = createdOn;
            LastUpdatedBy = lastUpdatedBy;
            LastUpdatedOn = lastUpdatedOn;
        }
    }
}
