// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Blueprint
{
    public static class GetAssignment
    {
        /// <summary>
        /// Get a blueprint assignment.
        /// 
        /// Uses Azure REST API version 2018-11-01-preview.
        /// </summary>
        public static Task<GetAssignmentResult> InvokeAsync(GetAssignmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAssignmentResult>("azure-native:blueprint:getAssignment", args ?? new GetAssignmentArgs(), options.WithDefaults());

        /// <summary>
        /// Get a blueprint assignment.
        /// 
        /// Uses Azure REST API version 2018-11-01-preview.
        /// </summary>
        public static Output<GetAssignmentResult> Invoke(GetAssignmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssignmentResult>("azure-native:blueprint:getAssignment", args ?? new GetAssignmentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a blueprint assignment.
        /// 
        /// Uses Azure REST API version 2018-11-01-preview.
        /// </summary>
        public static Output<GetAssignmentResult> Invoke(GetAssignmentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssignmentResult>("azure-native:blueprint:getAssignment", args ?? new GetAssignmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAssignmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the blueprint assignment.
        /// </summary>
        [Input("assignmentName", required: true)]
        public string AssignmentName { get; set; } = null!;

        /// <summary>
        /// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
        /// </summary>
        [Input("resourceScope", required: true)]
        public string ResourceScope { get; set; } = null!;

        public GetAssignmentArgs()
        {
        }
        public static new GetAssignmentArgs Empty => new GetAssignmentArgs();
    }

    public sealed class GetAssignmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the blueprint assignment.
        /// </summary>
        [Input("assignmentName", required: true)]
        public Input<string> AssignmentName { get; set; } = null!;

        /// <summary>
        /// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
        /// </summary>
        [Input("resourceScope", required: true)]
        public Input<string> ResourceScope { get; set; } = null!;

        public GetAssignmentInvokeArgs()
        {
        }
        public static new GetAssignmentInvokeArgs Empty => new GetAssignmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetAssignmentResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// ID of the published version of a blueprint definition.
        /// </summary>
        public readonly string? BlueprintId;
        /// <summary>
        /// Multi-line explain this resource.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// One-liner string explain this resource.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// String Id used to locate any resource on Azure.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Managed identity for this blueprint assignment.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse Identity;
        /// <summary>
        /// The location of this blueprint assignment.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Defines how resources deployed by a blueprint assignment are locked.
        /// </summary>
        public readonly Outputs.AssignmentLockSettingsResponse? Locks;
        /// <summary>
        /// Name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Blueprint assignment parameter values.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterValueResponse> Parameters;
        /// <summary>
        /// State of the blueprint assignment.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Names and locations of resource group placeholders.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ResourceGroupValueResponse> ResourceGroups;
        /// <summary>
        /// The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// Status of blueprint assignment. This field is readonly.
        /// </summary>
        public readonly Outputs.AssignmentStatusResponse Status;
        /// <summary>
        /// Type of this resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAssignmentResult(
            string azureApiVersion,

            string? blueprintId,

            string? description,

            string? displayName,

            string id,

            Outputs.ManagedServiceIdentityResponse identity,

            string location,

            Outputs.AssignmentLockSettingsResponse? locks,

            string name,

            ImmutableDictionary<string, Outputs.ParameterValueResponse> parameters,

            string provisioningState,

            ImmutableDictionary<string, Outputs.ResourceGroupValueResponse> resourceGroups,

            string? scope,

            Outputs.AssignmentStatusResponse status,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            BlueprintId = blueprintId;
            Description = description;
            DisplayName = displayName;
            Id = id;
            Identity = identity;
            Location = location;
            Locks = locks;
            Name = name;
            Parameters = parameters;
            ProvisioningState = provisioningState;
            ResourceGroups = resourceGroups;
            Scope = scope;
            Status = status;
            Type = type;
        }
    }
}
