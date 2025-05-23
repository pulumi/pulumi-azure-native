// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of TotalLocalStorageGBRequest
    /// </summary>
    [OutputType]
    public sealed class TotalLocalStorageGBRequestResponse
    {
        /// <summary>
        /// The storage maximum in GB.
        /// </summary>
        public readonly int? Max;
        /// <summary>
        /// The storage minimum in GB.
        /// </summary>
        public readonly int? Min;

        [OutputConstructor]
        private TotalLocalStorageGBRequestResponse(
            int? max,

            int? min)
        {
            Max = max;
            Min = min;
        }
    }
}
