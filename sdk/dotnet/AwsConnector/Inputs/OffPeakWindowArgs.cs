// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of OffPeakWindow
    /// </summary>
    public sealed class OffPeakWindowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;A custom start time for the off-peak window, in Coordinated Universal Time (UTC). The window length will always be 10 hours, so you can't specify an end time. For example, if you specify 11:00 P.M. UTC as a start time, the end time will automatically be set to 9:00 A.M.&lt;/p&gt;
        /// </summary>
        [Input("windowStartTime")]
        public Input<Inputs.WindowStartTimeArgs>? WindowStartTime { get; set; }

        public OffPeakWindowArgs()
        {
        }
        public static new OffPeakWindowArgs Empty => new OffPeakWindowArgs();
    }
}
