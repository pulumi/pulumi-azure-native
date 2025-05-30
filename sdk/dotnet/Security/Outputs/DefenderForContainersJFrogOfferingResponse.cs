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
    /// The Defender for Containers for JFrog Artifactory offering
    /// </summary>
    [OutputType]
    public sealed class DefenderForContainersJFrogOfferingResponse
    {
        /// <summary>
        /// The offering description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The type of the security offering.
        /// Expected value is 'DefenderForContainersJFrog'.
        /// </summary>
        public readonly string OfferingType;

        [OutputConstructor]
        private DefenderForContainersJFrogOfferingResponse(
            string description,

            string offeringType)
        {
            Description = description;
            OfferingType = offeringType;
        }
    }
}
