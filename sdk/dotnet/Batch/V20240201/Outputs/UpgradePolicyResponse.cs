// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.V20240201.Outputs
{

    /// <summary>
    /// Describes an upgrade policy - automatic, manual, or rolling.
    /// </summary>
    [OutputType]
    public sealed class UpgradePolicyResponse
    {
        /// <summary>
        /// The configuration parameters used for performing automatic OS upgrade.
        /// </summary>
        public readonly Outputs.AutomaticOSUpgradePolicyResponse? AutomaticOSUpgradePolicy;
        public readonly string Mode;
        /// <summary>
        /// This property is only supported on Pools with the virtualMachineConfiguration property.
        /// </summary>
        public readonly Outputs.RollingUpgradePolicyResponse? RollingUpgradePolicy;

        [OutputConstructor]
        private UpgradePolicyResponse(
            Outputs.AutomaticOSUpgradePolicyResponse? automaticOSUpgradePolicy,

            string mode,

            Outputs.RollingUpgradePolicyResponse? rollingUpgradePolicy)
        {
            AutomaticOSUpgradePolicy = automaticOSUpgradePolicy;
            Mode = mode;
            RollingUpgradePolicy = rollingUpgradePolicy;
        }
    }
}
