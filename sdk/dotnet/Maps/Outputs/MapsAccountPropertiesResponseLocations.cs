// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Maps.Outputs
{

    /// <summary>
    /// Data processing location.
    /// </summary>
    [OutputType]
    public sealed class MapsAccountPropertiesResponseLocations
    {
        /// <summary>
        /// The location name.
        /// </summary>
        public readonly string LocationName;

        [OutputConstructor]
        private MapsAccountPropertiesResponseLocations(string locationName)
        {
            LocationName = locationName;
        }
    }
}
