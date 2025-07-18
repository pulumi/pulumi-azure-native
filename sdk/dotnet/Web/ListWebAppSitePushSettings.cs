// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class ListWebAppSitePushSettings
    {
        /// <summary>
        /// Description for Gets the Push settings associated with web app.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListWebAppSitePushSettingsResult> InvokeAsync(ListWebAppSitePushSettingsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWebAppSitePushSettingsResult>("azure-native:web:listWebAppSitePushSettings", args ?? new ListWebAppSitePushSettingsArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Gets the Push settings associated with web app.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWebAppSitePushSettingsResult> Invoke(ListWebAppSitePushSettingsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppSitePushSettingsResult>("azure-native:web:listWebAppSitePushSettings", args ?? new ListWebAppSitePushSettingsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Gets the Push settings associated with web app.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWebAppSitePushSettingsResult> Invoke(ListWebAppSitePushSettingsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppSitePushSettingsResult>("azure-native:web:listWebAppSitePushSettings", args ?? new ListWebAppSitePushSettingsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWebAppSitePushSettingsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of web app.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListWebAppSitePushSettingsArgs()
        {
        }
        public static new ListWebAppSitePushSettingsArgs Empty => new ListWebAppSitePushSettingsArgs();
    }

    public sealed class ListWebAppSitePushSettingsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of web app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListWebAppSitePushSettingsInvokeArgs()
        {
        }
        public static new ListWebAppSitePushSettingsInvokeArgs Empty => new ListWebAppSitePushSettingsInvokeArgs();
    }


    [OutputType]
    public sealed class ListWebAppSitePushSettingsResult
    {
        /// <summary>
        /// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
        /// </summary>
        public readonly string? DynamicTagsJson;
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Gets or sets a flag indicating whether the Push endpoint is enabled.
        /// </summary>
        public readonly bool IsPushEnabled;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
        /// </summary>
        public readonly string? TagWhitelistJson;
        /// <summary>
        /// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
        /// Tags can consist of alphanumeric characters and the following:
        /// '_', '@', '#', '.', ':', '-'. 
        /// Validation should be performed at the PushRequestHandler.
        /// </summary>
        public readonly string? TagsRequiringAuth;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListWebAppSitePushSettingsResult(
            string? dynamicTagsJson,

            string id,

            bool isPushEnabled,

            string? kind,

            string name,

            string? tagWhitelistJson,

            string? tagsRequiringAuth,

            string type)
        {
            DynamicTagsJson = dynamicTagsJson;
            Id = id;
            IsPushEnabled = isPushEnabled;
            Kind = kind;
            Name = name;
            TagWhitelistJson = tagWhitelistJson;
            TagsRequiringAuth = tagsRequiringAuth;
            Type = type;
        }
    }
}
