// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DesktopVirtualization.Outputs
{

    /// <summary>
    /// The time for a scaling action to occur.
    /// </summary>
    [OutputType]
    public sealed class TimeResponse
    {
        /// <summary>
        /// The hour.
        /// </summary>
        public readonly int Hour;
        /// <summary>
        /// The minute.
        /// </summary>
        public readonly int Minute;

        [OutputConstructor]
        private TimeResponse(
            int hour,

            int minute)
        {
            Hour = hour;
            Minute = minute;
        }
    }
}
