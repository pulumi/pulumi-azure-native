// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Outputs
{

    [OutputType]
    public sealed class WorkloadInstanceModelPropertiesResponseCurrentJob
    {
        /// <summary>
        /// Gets or sets the workflow friendly display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Gets or sets end time of the workflow.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Gets or sets workflow Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Gets or sets workflow name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets workload scenario name.
        /// </summary>
        public readonly string ScenarioName;
        /// <summary>
        /// Gets or sets start time of the workflow.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Gets or sets workflow state.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private WorkloadInstanceModelPropertiesResponseCurrentJob(
            string displayName,

            string endTime,

            string id,

            string name,

            string scenarioName,

            string startTime,

            string state)
        {
            DisplayName = displayName;
            EndTime = endTime;
            Id = id;
            Name = name;
            ScenarioName = scenarioName;
            StartTime = startTime;
            State = state;
        }
    }
}
