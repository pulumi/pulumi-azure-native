// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Resources
{
    public static class GetTemplateSpecVersion
    {
        /// <summary>
        /// Gets a Template Spec version from a specific Template Spec.
        /// 
        /// Uses Azure REST API version 2022-02-01.
        /// 
        /// Other available API versions: 2021-03-01-preview, 2021-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetTemplateSpecVersionResult> InvokeAsync(GetTemplateSpecVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTemplateSpecVersionResult>("azure-native:resources:getTemplateSpecVersion", args ?? new GetTemplateSpecVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Template Spec version from a specific Template Spec.
        /// 
        /// Uses Azure REST API version 2022-02-01.
        /// 
        /// Other available API versions: 2021-03-01-preview, 2021-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTemplateSpecVersionResult> Invoke(GetTemplateSpecVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemplateSpecVersionResult>("azure-native:resources:getTemplateSpecVersion", args ?? new GetTemplateSpecVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Template Spec version from a specific Template Spec.
        /// 
        /// Uses Azure REST API version 2022-02-01.
        /// 
        /// Other available API versions: 2021-03-01-preview, 2021-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTemplateSpecVersionResult> Invoke(GetTemplateSpecVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemplateSpecVersionResult>("azure-native:resources:getTemplateSpecVersion", args ?? new GetTemplateSpecVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTemplateSpecVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the Template Spec.
        /// </summary>
        [Input("templateSpecName", required: true)]
        public string TemplateSpecName { get; set; } = null!;

        /// <summary>
        /// The version of the Template Spec.
        /// </summary>
        [Input("templateSpecVersion", required: true)]
        public string TemplateSpecVersion { get; set; } = null!;

        public GetTemplateSpecVersionArgs()
        {
        }
        public static new GetTemplateSpecVersionArgs Empty => new GetTemplateSpecVersionArgs();
    }

    public sealed class GetTemplateSpecVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the Template Spec.
        /// </summary>
        [Input("templateSpecName", required: true)]
        public Input<string> TemplateSpecName { get; set; } = null!;

        /// <summary>
        /// The version of the Template Spec.
        /// </summary>
        [Input("templateSpecVersion", required: true)]
        public Input<string> TemplateSpecVersion { get; set; } = null!;

        public GetTemplateSpecVersionInvokeArgs()
        {
        }
        public static new GetTemplateSpecVersionInvokeArgs Empty => new GetTemplateSpecVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetTemplateSpecVersionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Template Spec version description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// String Id used to locate any resource on Azure.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An array of linked template artifacts.
        /// </summary>
        public readonly ImmutableArray<Outputs.LinkedTemplateArtifactResponse> LinkedTemplates;
        /// <summary>
        /// The location of the Template Spec Version. It must match the location of the parent Template Spec.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The main Azure Resource Manager template content.
        /// </summary>
        public readonly object? MainTemplate;
        /// <summary>
        /// The version metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
        /// </summary>
        public readonly object? Metadata;
        /// <summary>
        /// Name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Type of this resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The Azure Resource Manager template UI definition content.
        /// </summary>
        public readonly object? UiFormDefinition;

        [OutputConstructor]
        private GetTemplateSpecVersionResult(
            string azureApiVersion,

            string? description,

            string id,

            ImmutableArray<Outputs.LinkedTemplateArtifactResponse> linkedTemplates,

            string location,

            object? mainTemplate,

            object? metadata,

            string name,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            object? uiFormDefinition)
        {
            AzureApiVersion = azureApiVersion;
            Description = description;
            Id = id;
            LinkedTemplates = linkedTemplates;
            Location = location;
            MainTemplate = mainTemplate;
            Metadata = metadata;
            Name = name;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            UiFormDefinition = uiFormDefinition;
        }
    }
}
