// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20241002Preview.Outputs
{

    /// <summary>
    /// Configuration of datadog 
    /// </summary>
    [OutputType]
    public sealed class DataDogConfigurationResponse
    {
        /// <summary>
        /// The data dog api key
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The data dog site
        /// </summary>
        public readonly string? Site;

        [OutputConstructor]
        private DataDogConfigurationResponse(
            string? key,

            string? site)
        {
            Key = key;
            Site = site;
        }
    }
}
