// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.Outputs
{

    /// <summary>
    /// Contains information of current state for an Azure SQL Database, Server or Elastic Pool Recommended Action.
    /// </summary>
    [OutputType]
    public sealed class RecommendedActionStateInfoResponse
    {
        /// <summary>
        /// Gets who initiated the execution of this recommended action. Possible Value are: User    -&gt; When user explicity notified system to apply the recommended action. System  -&gt; When auto-execute status of this advisor was set to 'Enabled', in which case the system applied it.
        /// </summary>
        public readonly string ActionInitiatedBy;
        /// <summary>
        /// Current state the recommended action is in. Some commonly used states are: Active      -&gt; recommended action is active and no action has been taken yet. Pending     -&gt; recommended action is approved for and is awaiting execution. Executing   -&gt; recommended action is being applied on the user database. Verifying   -&gt; recommended action was applied and is being verified of its usefulness by the system. Success     -&gt; recommended action was applied and improvement found during verification. Pending Revert  -&gt; verification found little or no improvement so recommended action is queued for revert or user has manually reverted. Reverting   -&gt; changes made while applying recommended action are being reverted on the user database. Reverted    -&gt; successfully reverted the changes made by recommended action on user database. Ignored     -&gt; user explicitly ignored/discarded the recommended action. 
        /// </summary>
        public readonly string CurrentValue;
        /// <summary>
        /// Gets the time when the state was last modified
        /// </summary>
        public readonly string LastModified;

        [OutputConstructor]
        private RecommendedActionStateInfoResponse(
            string actionInitiatedBy,

            string currentValue,

            string lastModified)
        {
            ActionInitiatedBy = actionInitiatedBy;
            CurrentValue = currentValue;
            LastModified = lastModified;
        }
    }
}
