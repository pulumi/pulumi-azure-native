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
    /// Append value for a Variable of type Array.
    /// </summary>
    [OutputType]
    public sealed class AppendVariableActivityResponse
    {
        /// <summary>
        /// Activity depends on condition.
        /// </summary>
        public readonly ImmutableArray<Outputs.ActivityDependencyResponse> DependsOn;
        /// <summary>
        /// Activity description.
        /// </summary>
        public readonly string? Description;
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
        /// Expected value is 'AppendVariable'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Activity user properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserPropertyResponse> UserProperties;
        /// <summary>
        /// Value to be appended. Type: could be a static value matching type of the variable item or Expression with resultType matching type of the variable item
        /// </summary>
        public readonly object? Value;
        /// <summary>
        /// Name of the variable whose value needs to be appended to.
        /// </summary>
        public readonly string? VariableName;

        [OutputConstructor]
        private AppendVariableActivityResponse(
            ImmutableArray<Outputs.ActivityDependencyResponse> dependsOn,

            string? description,

            string name,

            string? onInactiveMarkAs,

            string? state,

            string type,

            ImmutableArray<Outputs.UserPropertyResponse> userProperties,

            object? value,

            string? variableName)
        {
            DependsOn = dependsOn;
            Description = description;
            Name = name;
            OnInactiveMarkAs = onInactiveMarkAs;
            State = state;
            Type = type;
            UserProperties = userProperties;
            Value = value;
            VariableName = variableName;
        }
    }
}
