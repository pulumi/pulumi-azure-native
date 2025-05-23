// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Inputs
{

    /// <summary>
    /// Details of an inquired protectable item.
    /// </summary>
    public sealed class WorkloadInquiryDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Inquiry validation such as permissions and other backup validations.
        /// </summary>
        [Input("inquiryValidation")]
        public Input<Inputs.InquiryValidationArgs>? InquiryValidation { get; set; }

        /// <summary>
        /// Contains the protectable item Count inside this Container.
        /// </summary>
        [Input("itemCount")]
        public Input<double>? ItemCount { get; set; }

        /// <summary>
        /// Type of the Workload such as SQL, Oracle etc.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public WorkloadInquiryDetailsArgs()
        {
        }
        public static new WorkloadInquiryDetailsArgs Empty => new WorkloadInquiryDetailsArgs();
    }
}
