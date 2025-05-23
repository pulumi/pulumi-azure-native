// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Blueprint
{
    public static class GetPublishedBlueprint
    {
        /// <summary>
        /// Get a published version of a blueprint definition.
        /// 
        /// Uses Azure REST API version 2018-11-01-preview.
        /// </summary>
        public static Task<GetPublishedBlueprintResult> InvokeAsync(GetPublishedBlueprintArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPublishedBlueprintResult>("azure-native:blueprint:getPublishedBlueprint", args ?? new GetPublishedBlueprintArgs(), options.WithDefaults());

        /// <summary>
        /// Get a published version of a blueprint definition.
        /// 
        /// Uses Azure REST API version 2018-11-01-preview.
        /// </summary>
        public static Output<GetPublishedBlueprintResult> Invoke(GetPublishedBlueprintInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPublishedBlueprintResult>("azure-native:blueprint:getPublishedBlueprint", args ?? new GetPublishedBlueprintInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a published version of a blueprint definition.
        /// 
        /// Uses Azure REST API version 2018-11-01-preview.
        /// </summary>
        public static Output<GetPublishedBlueprintResult> Invoke(GetPublishedBlueprintInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPublishedBlueprintResult>("azure-native:blueprint:getPublishedBlueprint", args ?? new GetPublishedBlueprintInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPublishedBlueprintArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// Version of the published blueprint definition.
        /// </summary>
        [Input("versionId", required: true)]
        public string VersionId { get; set; } = null!;

        public GetPublishedBlueprintArgs()
        {
        }
        public static new GetPublishedBlueprintArgs Empty => new GetPublishedBlueprintArgs();
    }

    public sealed class GetPublishedBlueprintInvokeArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// Version of the published blueprint definition.
        /// </summary>
        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        public GetPublishedBlueprintInvokeArgs()
        {
        }
        public static new GetPublishedBlueprintInvokeArgs Empty => new GetPublishedBlueprintInvokeArgs();
    }


    [OutputType]
    public sealed class GetPublishedBlueprintResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Name of the published blueprint definition.
        /// </summary>
        public readonly string? BlueprintName;
        /// <summary>
        /// Version-specific change notes.
        /// </summary>
        public readonly string? ChangeNotes;
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
        public readonly string? TargetScope;
        /// <summary>
        /// Type of this resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPublishedBlueprintResult(
            string azureApiVersion,

            string? blueprintName,

            string? changeNotes,

            string? description,

            string? displayName,

            string id,

            string name,

            ImmutableDictionary<string, Outputs.ParameterDefinitionResponse>? parameters,

            ImmutableDictionary<string, Outputs.ResourceGroupDefinitionResponse>? resourceGroups,

            Outputs.BlueprintStatusResponse status,

            string? targetScope,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            BlueprintName = blueprintName;
            ChangeNotes = changeNotes;
            Description = description;
            DisplayName = displayName;
            Id = id;
            Name = name;
            Parameters = parameters;
            ResourceGroups = resourceGroups;
            Status = status;
            TargetScope = targetScope;
            Type = type;
        }
    }
}
