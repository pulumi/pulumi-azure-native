// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class ListStaticSiteAppSettings
    {
        /// <summary>
        /// Description for Gets the application settings of a static site.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListStaticSiteAppSettingsResult> InvokeAsync(ListStaticSiteAppSettingsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListStaticSiteAppSettingsResult>("azure-native:web:listStaticSiteAppSettings", args ?? new ListStaticSiteAppSettingsArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Gets the application settings of a static site.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListStaticSiteAppSettingsResult> Invoke(ListStaticSiteAppSettingsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListStaticSiteAppSettingsResult>("azure-native:web:listStaticSiteAppSettings", args ?? new ListStaticSiteAppSettingsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Gets the application settings of a static site.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListStaticSiteAppSettingsResult> Invoke(ListStaticSiteAppSettingsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListStaticSiteAppSettingsResult>("azure-native:web:listStaticSiteAppSettings", args ?? new ListStaticSiteAppSettingsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListStaticSiteAppSettingsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the static site.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListStaticSiteAppSettingsArgs()
        {
        }
        public static new ListStaticSiteAppSettingsArgs Empty => new ListStaticSiteAppSettingsArgs();
    }

    public sealed class ListStaticSiteAppSettingsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the static site.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListStaticSiteAppSettingsInvokeArgs()
        {
        }
        public static new ListStaticSiteAppSettingsInvokeArgs Empty => new ListStaticSiteAppSettingsInvokeArgs();
    }


    [OutputType]
    public sealed class ListStaticSiteAppSettingsResult
    {
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Settings.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListStaticSiteAppSettingsResult(
            string id,

            string? kind,

            string name,

            ImmutableDictionary<string, string> properties,

            string type)
        {
            Id = id;
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
