// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    /// <summary>
    /// A2A provider specific settings.
    /// </summary>
    [OutputType]
    public sealed class A2AProtectionContainerMappingDetailsResponse
    {
        /// <summary>
        /// A value indicating whether the auto update is enabled.
        /// </summary>
        public readonly string? AgentAutoUpdateStatus;
        /// <summary>
        /// The automation account arm id.
        /// </summary>
        public readonly string? AutomationAccountArmId;
        /// <summary>
        /// A value indicating the type authentication to use for automation Account.
        /// </summary>
        public readonly string? AutomationAccountAuthenticationType;
        /// <summary>
        /// Gets the class type. Overridden in derived classes.
        /// Expected value is 'A2A'.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The job schedule arm name.
        /// </summary>
        public readonly string? JobScheduleName;
        /// <summary>
        /// The schedule arm name.
        /// </summary>
        public readonly string? ScheduleName;

        [OutputConstructor]
        private A2AProtectionContainerMappingDetailsResponse(
            string? agentAutoUpdateStatus,

            string? automationAccountArmId,

            string? automationAccountAuthenticationType,

            string instanceType,

            string? jobScheduleName,

            string? scheduleName)
        {
            AgentAutoUpdateStatus = agentAutoUpdateStatus;
            AutomationAccountArmId = automationAccountArmId;
            AutomationAccountAuthenticationType = automationAccountAuthenticationType;
            InstanceType = instanceType;
            JobScheduleName = jobScheduleName;
            ScheduleName = scheduleName;
        }
    }
}
