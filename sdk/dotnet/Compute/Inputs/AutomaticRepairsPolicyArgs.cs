// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// Specifies the configuration parameters for automatic repairs on the virtual machine scale set.
    /// </summary>
    public sealed class AutomaticRepairsPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether automatic repairs should be enabled on the virtual machine scale set. The default value is false.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The amount of time for which automatic repairs are suspended due to a state change on VM. The grace time starts after the state change has completed. This helps avoid premature or accidental repairs. The time duration should be specified in ISO 8601 format. The minimum allowed grace period is 10 minutes (PT10M), which is also the default value. The maximum allowed grace period is 90 minutes (PT90M).
        /// </summary>
        [Input("gracePeriod")]
        public Input<string>? GracePeriod { get; set; }

        /// <summary>
        /// Type of repair action (replace, restart, reimage) that will be used for repairing unhealthy virtual machines in the scale set. Default value is replace.
        /// </summary>
        [Input("repairAction")]
        public InputUnion<string, Pulumi.AzureNative.Compute.RepairAction>? RepairAction { get; set; }

        public AutomaticRepairsPolicyArgs()
        {
        }
        public static new AutomaticRepairsPolicyArgs Empty => new AutomaticRepairsPolicyArgs();
    }
}
