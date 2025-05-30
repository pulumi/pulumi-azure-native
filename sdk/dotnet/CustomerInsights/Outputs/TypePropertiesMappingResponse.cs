// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CustomerInsights.Outputs
{

    /// <summary>
    /// Metadata for a Link's property mapping.
    /// </summary>
    [OutputType]
    public sealed class TypePropertiesMappingResponse
    {
        /// <summary>
        /// Link type.
        /// </summary>
        public readonly string? LinkType;
        /// <summary>
        ///  Property name on the source Entity Type.
        /// </summary>
        public readonly string SourcePropertyName;
        /// <summary>
        /// Property name on the target Entity Type.
        /// </summary>
        public readonly string TargetPropertyName;

        [OutputConstructor]
        private TypePropertiesMappingResponse(
            string? linkType,

            string sourcePropertyName,

            string targetPropertyName)
        {
            LinkType = linkType;
            SourcePropertyName = sourcePropertyName;
            TargetPropertyName = targetPropertyName;
        }
    }
}
