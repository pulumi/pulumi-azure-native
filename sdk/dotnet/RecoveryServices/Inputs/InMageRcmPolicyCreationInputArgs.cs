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
    /// InMageRcm policy creation input.
    /// </summary>
    public sealed class InMageRcmPolicyCreationInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The app consistent snapshot frequency (in minutes).
        /// </summary>
        [Input("appConsistentFrequencyInMinutes")]
        public Input<int>? AppConsistentFrequencyInMinutes { get; set; }

        /// <summary>
        /// The crash consistent snapshot frequency (in minutes).
        /// </summary>
        [Input("crashConsistentFrequencyInMinutes")]
        public Input<int>? CrashConsistentFrequencyInMinutes { get; set; }

        /// <summary>
        /// A value indicating whether multi-VM sync has to be enabled.
        /// </summary>
        [Input("enableMultiVmSync")]
        public Input<string>? EnableMultiVmSync { get; set; }

        /// <summary>
        /// The class type.
        /// Expected value is 'InMageRcm'.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The duration in minutes until which the recovery points need to be stored.
        /// </summary>
        [Input("recoveryPointHistoryInMinutes")]
        public Input<int>? RecoveryPointHistoryInMinutes { get; set; }

        public InMageRcmPolicyCreationInputArgs()
        {
        }
        public static new InMageRcmPolicyCreationInputArgs Empty => new InMageRcmPolicyCreationInputArgs();
    }
}
