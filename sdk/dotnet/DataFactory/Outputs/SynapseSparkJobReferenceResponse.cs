// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Synapse spark job reference type.
    /// </summary>
    [OutputType]
    public sealed class SynapseSparkJobReferenceResponse
    {
        /// <summary>
        /// Reference spark job name. Expression with resultType string.
        /// </summary>
        public readonly object ReferenceName;
        /// <summary>
        /// Synapse spark job reference type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SynapseSparkJobReferenceResponse(
            object referenceName,

            string type)
        {
            ReferenceName = referenceName;
            Type = type;
        }
    }
}
