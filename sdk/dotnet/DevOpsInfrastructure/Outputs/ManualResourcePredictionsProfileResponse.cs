// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevOpsInfrastructure.Outputs
{

    /// <summary>
    /// Customer provides the stand-by agent scheme.
    /// </summary>
    [OutputType]
    public sealed class ManualResourcePredictionsProfileResponse
    {
        /// <summary>
        /// Determines how the stand-by scheme should be provided.
        /// Expected value is 'Manual'.
        /// </summary>
        public readonly string Kind;

        [OutputConstructor]
        private ManualResourcePredictionsProfileResponse(string kind)
        {
            Kind = kind;
        }
    }
}
