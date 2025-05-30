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
    /// Definition of InstanceStatusSummary
    /// </summary>
    public sealed class InstanceStatusSummaryArgs : global::Pulumi.ResourceArgs
    {
        [Input("details")]
        private InputList<Inputs.InstanceStatusDetailsArgs>? _details;

        /// <summary>
        /// &lt;p&gt;The system instance health or application instance health.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.InstanceStatusDetailsArgs> Details
        {
            get => _details ?? (_details = new InputList<Inputs.InstanceStatusDetailsArgs>());
            set => _details = value;
        }

        /// <summary>
        /// &lt;p&gt;The status.&lt;/p&gt;
        /// </summary>
        [Input("status")]
        public Input<Inputs.SummaryStatusEnumValueArgs>? Status { get; set; }

        public InstanceStatusSummaryArgs()
        {
        }
        public static new InstanceStatusSummaryArgs Empty => new InstanceStatusSummaryArgs();
    }
}
