// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    /// <summary>
    /// Extended Location.
    /// </summary>
    [OutputType]
    public sealed class ExtendedLocationResponse
    {
        /// <summary>
        /// Name of extended location.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Type of extended location.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ExtendedLocationResponse(
            string? name,

            string type)
        {
            Name = name;
            Type = type;
        }
    }
}
