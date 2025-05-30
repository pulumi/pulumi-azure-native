// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Filter and return results from input array based on the conditions.
    /// </summary>
    [OutputType]
    public sealed class FilterActivityResponse
    {
        /// <summary>
        /// Condition to be used for filtering the input.
        /// </summary>
        public readonly Outputs.ExpressionResponse Condition;
        /// <summary>
        /// Activity depends on condition.
        /// </summary>
        public readonly ImmutableArray<Outputs.ActivityDependencyResponse> DependsOn;
        /// <summary>
        /// Activity description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Input array on which filter should be applied.
        /// </summary>
        public readonly Outputs.ExpressionResponse Items;
        /// <summary>
        /// Activity name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Status result of the activity when the state is set to Inactive. This is an optional property and if not provided when the activity is inactive, the status will be Succeeded by default.
        /// </summary>
        public readonly string? OnInactiveMarkAs;
        /// <summary>
        /// Activity state. This is an optional property and if not provided, the state will be Active by default.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Type of activity.
        /// Expected value is 'Filter'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Activity user properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserPropertyResponse> UserProperties;

        [OutputConstructor]
        private FilterActivityResponse(
            Outputs.ExpressionResponse condition,

            ImmutableArray<Outputs.ActivityDependencyResponse> dependsOn,

            string? description,

            Outputs.ExpressionResponse items,

            string name,

            string? onInactiveMarkAs,

            string? state,

            string type,

            ImmutableArray<Outputs.UserPropertyResponse> userProperties)
        {
            Condition = condition;
            DependsOn = dependsOn;
            Description = description;
            Items = items;
            Name = name;
            OnInactiveMarkAs = onInactiveMarkAs;
            State = state;
            Type = type;
            UserProperties = userProperties;
        }
    }
}
