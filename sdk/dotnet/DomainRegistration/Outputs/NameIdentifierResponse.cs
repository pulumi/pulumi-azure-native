// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DomainRegistration.Outputs
{

    /// <summary>
    /// Identifies an object.
    /// </summary>
    [OutputType]
    public sealed class NameIdentifierResponse
    {
        /// <summary>
        /// Name of the object.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private NameIdentifierResponse(string? name)
        {
            Name = name;
        }
    }
}
