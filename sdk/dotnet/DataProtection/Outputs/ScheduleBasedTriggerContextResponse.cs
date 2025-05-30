// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.Outputs
{

    /// <summary>
    /// Schedule based trigger context
    /// </summary>
    [OutputType]
    public sealed class ScheduleBasedTriggerContextResponse
    {
        /// <summary>
        /// Type of the specific object - used for deserializing
        /// Expected value is 'ScheduleBasedTriggerContext'.
        /// </summary>
        public readonly string ObjectType;
        /// <summary>
        /// Schedule for this backup
        /// </summary>
        public readonly Outputs.BackupScheduleResponse Schedule;
        /// <summary>
        /// List of tags that can be applicable for given schedule.
        /// </summary>
        public readonly ImmutableArray<Outputs.TaggingCriteriaResponse> TaggingCriteria;

        [OutputConstructor]
        private ScheduleBasedTriggerContextResponse(
            string objectType,

            Outputs.BackupScheduleResponse schedule,

            ImmutableArray<Outputs.TaggingCriteriaResponse> taggingCriteria)
        {
            ObjectType = objectType;
            Schedule = schedule;
            TaggingCriteria = taggingCriteria;
        }
    }
}
