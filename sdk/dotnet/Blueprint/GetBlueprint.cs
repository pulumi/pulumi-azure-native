// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Blueprint
{
    public static class GetBlueprint
    {
        /// <summary>
        /// Get a blueprint definition.
        /// 
        /// Uses Azure REST API version 2018-11-01-preview.
        /// </summary>
        public static Task<GetBlueprintResult> InvokeAsync(GetBlueprintArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBlueprintResult>("azure-native:blueprint:getBlueprint", args ?? new GetBlueprintArgs(), options.WithDefaults());

        /// <summary>
        /// Get a blueprint definition.
        /// 
        /// Uses Azure REST API version 2018-11-01-preview.
        /// </summary>
        public static Output<GetBlueprintResult> Invoke(GetBlueprintInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBlueprintResult>("azure-native:blueprint:getBlueprint", args ?? new GetBlueprintInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a blueprint definition.
        /// 
        /// Uses Azure REST API version 2018-11-01-preview.
        /// </summary>
        public static Output<GetBlueprintResult> Invoke(GetBlueprintInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBlueprintResult>("azure-native:blueprint:getBlueprint", args ?? new GetBlueprintInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBlueprintArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the blueprint definition.
        /// </summary>
        [Input("blueprintName", required: true)]
        public string BlueprintName { get; set; } = null!;

        /// <summary>
        /// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
        /// </summary>
        [Input("resourceScope", required: true)]
        public string ResourceScope { get; set; } = null!;

        public GetBlueprintArgs()
        {
        }
        public static new GetBlueprintArgs Empty => new GetBlueprintArgs();
    }

    public sealed class GetBlueprintInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the blueprint definition.
        /// </summary>
        [Input("blueprintName", required: true)]
        public Input<string> BlueprintName { get; set; } = null!;

        /// <summary>
        /// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
        /// </summary>
        [Input("resourceScope", required: true)]
        public Input<string> ResourceScope { get; set; } = null!;

        public GetBlueprintInvokeArgs()
        {
        }
        public static new GetBlueprintInvokeArgs Empty => new GetBlueprintInvokeArgs();
    }


    [OutputType]
    public sealed class GetBlueprintResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
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
        /// Layout view of the blueprint definition for UI reference.
        /// </summary>
        public readonly object Layout;
        /// <summary>
        /// Name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Parameters required by this blueprint definition.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterDefinitionResponse>? Parameters;
        /// <summary>
        /// Resource group placeholders defined by this blueprint definition.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ResourceGroupDefinitionResponse>? ResourceGroups;
        /// <summary>
        /// Status of the blueprint. This field is readonly.
        /// </summary>
        public readonly Outputs.BlueprintStatusResponse Status;
        /// <summary>
        /// The scope where this blueprint definition can be assigned.
        /// </summary>
        public readonly string TargetScope;
        /// <summary>
        /// Type of this resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Published versions of this blueprint definition.
        /// </summary>
        public readonly object? Versions;

        [OutputConstructor]
        private GetBlueprintResult(
            string azureApiVersion,

            string? description,

            string? displayName,

            string id,

            object layout,

            string name,

            ImmutableDictionary<string, Outputs.ParameterDefinitionResponse>? parameters,

            ImmutableDictionary<string, Outputs.ResourceGroupDefinitionResponse>? resourceGroups,

            Outputs.BlueprintStatusResponse status,

            string targetScope,

            string type,

            object? versions)
        {
            AzureApiVersion = azureApiVersion;
            Description = description;
            DisplayName = displayName;
            Id = id;
            Layout = layout;
            Name = name;
            Parameters = parameters;
            ResourceGroups = resourceGroups;
            Status = status;
            TargetScope = targetScope;
            Type = type;
            Versions = versions;
        }
    }
}
