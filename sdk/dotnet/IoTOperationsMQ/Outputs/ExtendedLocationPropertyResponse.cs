// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsMQ.Outputs
{

    /// <summary>
    /// ExtendedLocation properties
    /// </summary>
    [OutputType]
    public sealed class ExtendedLocationPropertyResponse
    {
        /// <summary>
        /// The name of the extended location.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of ExtendedLocation.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ExtendedLocationPropertyResponse(
            string name,

            string type)
        {
            Name = name;
            Type = type;
        }
    }
}
