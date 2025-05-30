// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiCenter.Outputs
{

    /// <summary>
    /// API specification details.
    /// </summary>
    [OutputType]
    public sealed class ApiDefinitionPropertiesSpecificationResponse
    {
        /// <summary>
        /// Specification name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Specification version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private ApiDefinitionPropertiesSpecificationResponse(
            string? name,

            string? version)
        {
            Name = name;
            Version = version;
        }
    }
}
