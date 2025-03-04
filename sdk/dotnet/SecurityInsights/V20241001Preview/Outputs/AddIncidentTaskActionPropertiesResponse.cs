// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20241001Preview.Outputs
{

    [OutputType]
    public sealed class AddIncidentTaskActionPropertiesResponse
    {
        /// <summary>
        /// The description of the task.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The title of the task.
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private AddIncidentTaskActionPropertiesResponse(
            string? description,

            string title)
        {
            Description = description;
            Title = title;
        }
    }
}
