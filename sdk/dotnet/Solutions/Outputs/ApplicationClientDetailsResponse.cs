// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Solutions.Outputs
{

    /// <summary>
    /// The application client details to track the entity creating/updating the managed app resource.
    /// </summary>
    [OutputType]
    public sealed class ApplicationClientDetailsResponse
    {
        /// <summary>
        /// The client application Id.
        /// </summary>
        public readonly string? ApplicationId;
        /// <summary>
        /// The client Oid.
        /// </summary>
        public readonly string? Oid;
        /// <summary>
        /// The client Puid
        /// </summary>
        public readonly string? Puid;

        [OutputConstructor]
        private ApplicationClientDetailsResponse(
            string? applicationId,

            string? oid,

            string? puid)
        {
            ApplicationId = applicationId;
            Oid = oid;
            Puid = puid;
        }
    }
}
