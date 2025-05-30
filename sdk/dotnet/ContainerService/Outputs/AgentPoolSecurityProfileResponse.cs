// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// The security settings of an agent pool.
    /// </summary>
    [OutputType]
    public sealed class AgentPoolSecurityProfileResponse
    {
        /// <summary>
        /// Secure Boot is a feature of Trusted Launch which ensures that only signed operating systems and drivers can boot. For more details, see aka.ms/aks/trustedlaunch.  If not specified, the default is false.
        /// </summary>
        public readonly bool? EnableSecureBoot;
        /// <summary>
        /// vTPM is a Trusted Launch feature for configuring a dedicated secure vault for keys and measurements held locally on the node. For more details, see aka.ms/aks/trustedlaunch. If not specified, the default is false.
        /// </summary>
        public readonly bool? EnableVTPM;

        [OutputConstructor]
        private AgentPoolSecurityProfileResponse(
            bool? enableSecureBoot,

            bool? enableVTPM)
        {
            EnableSecureBoot = enableSecureBoot;
            EnableVTPM = enableVTPM;
        }
    }
}
