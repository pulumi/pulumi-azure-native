// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20160801
{
    /// <summary>
    /// A web app, a mobile app backend, or an API app.
    /// </summary>
    [AzureNativeResourceType("azure-native:web/v20160801:WebApp")]
    public partial class WebApp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Management information availability state for the app.
        /// </summary>
        [Output("availabilityState")]
        public Output<string> AvailabilityState { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to enable client affinity; &lt;code&gt;false&lt;/code&gt; to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is &lt;code&gt;true&lt;/code&gt;.
        /// </summary>
        [Output("clientAffinityEnabled")]
        public Output<bool?> ClientAffinityEnabled { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to enable client certificate authentication (TLS mutual authentication); otherwise, &lt;code&gt;false&lt;/code&gt;. Default is &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Output("clientCertEnabled")]
        public Output<bool?> ClientCertEnabled { get; private set; } = null!;

        /// <summary>
        /// Size of the function container.
        /// </summary>
        [Output("containerSize")]
        public Output<int?> ContainerSize { get; private set; } = null!;

        /// <summary>
        /// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
        /// </summary>
        [Output("dailyMemoryTimeQuota")]
        public Output<int?> DailyMemoryTimeQuota { get; private set; } = null!;

        /// <summary>
        /// Default hostname of the app. Read-only.
        /// </summary>
        [Output("defaultHostName")]
        public Output<string> DefaultHostName { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if the app is enabled; otherwise, &lt;code&gt;false&lt;/code&gt;. Setting this value to false disables the app (takes the app offline).
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
        /// the app is not served on those hostnames.
        /// </summary>
        [Output("enabledHostNames")]
        public Output<ImmutableArray<string>> EnabledHostNames { get; private set; } = null!;

        /// <summary>
        /// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
        /// </summary>
        [Output("hostNameSslStates")]
        public Output<ImmutableArray<Outputs.HostNameSslStateResponse>> HostNameSslStates { get; private set; } = null!;

        /// <summary>
        /// Hostnames associated with the app.
        /// </summary>
        [Output("hostNames")]
        public Output<ImmutableArray<string>> HostNames { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to disable the public hostnames of the app; otherwise, &lt;code&gt;false&lt;/code&gt;.
        ///  If &lt;code&gt;true&lt;/code&gt;, the app is only accessible via API management process.
        /// </summary>
        [Output("hostNamesDisabled")]
        public Output<bool?> HostNamesDisabled { get; private set; } = null!;

        /// <summary>
        /// App Service Environment to use for the app.
        /// </summary>
        [Output("hostingEnvironmentProfile")]
        public Output<Outputs.HostingEnvironmentProfileResponse?> HostingEnvironmentProfile { get; private set; } = null!;

        /// <summary>
        /// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
        /// http requests
        /// </summary>
        [Output("httpsOnly")]
        public Output<bool?> HttpsOnly { get; private set; } = null!;

        /// <summary>
        /// Managed service identity.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if the app is a default container; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Output("isDefaultContainer")]
        public Output<bool> IsDefaultContainer { get; private set; } = null!;

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Last time the app was modified, in UTC. Read-only.
        /// </summary>
        [Output("lastModifiedTimeUtc")]
        public Output<string> LastModifiedTimeUtc { get; private set; } = null!;

        /// <summary>
        /// Resource Location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Maximum number of workers.
        /// This only applies to Functions container.
        /// </summary>
        [Output("maxNumberOfWorkers")]
        public Output<int> MaxNumberOfWorkers { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings. Read-only.
        /// </summary>
        [Output("outboundIpAddresses")]
        public Output<string> OutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants. Read-only.
        /// </summary>
        [Output("possibleOutboundIpAddresses")]
        public Output<string> PossibleOutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// Name of the repository site.
        /// </summary>
        [Output("repositorySiteName")]
        public Output<string> RepositorySiteName { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if reserved; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Output("reserved")]
        public Output<bool?> Reserved { get; private set; } = null!;

        /// <summary>
        /// Name of the resource group the app belongs to. Read-only.
        /// </summary>
        [Output("resourceGroup")]
        public Output<string> ResourceGroup { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to stop SCM (KUDU) site when the app is stopped; otherwise, &lt;code&gt;false&lt;/code&gt;. The default is &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Output("scmSiteAlsoStopped")]
        public Output<bool?> ScmSiteAlsoStopped { get; private set; } = null!;

        /// <summary>
        /// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
        /// </summary>
        [Output("serverFarmId")]
        public Output<string?> ServerFarmId { get; private set; } = null!;

        /// <summary>
        /// Configuration of the app.
        /// </summary>
        [Output("siteConfig")]
        public Output<Outputs.SiteConfigResponse?> SiteConfig { get; private set; } = null!;

        /// <summary>
        /// Status of the last deployment slot swap operation.
        /// </summary>
        [Output("slotSwapStatus")]
        public Output<Outputs.SlotSwapStatusResponse> SlotSwapStatus { get; private set; } = null!;

        /// <summary>
        /// Current state of the app.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// App suspended till in case memory-time quota is exceeded.
        /// </summary>
        [Output("suspendedTill")]
        public Output<string> SuspendedTill { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies which deployment slot this app will swap into. Read-only.
        /// </summary>
        [Output("targetSwapSlot")]
        public Output<string> TargetSwapSlot { get; private set; } = null!;

        /// <summary>
        /// Azure Traffic Manager hostnames associated with the app. Read-only.
        /// </summary>
        [Output("trafficManagerHostNames")]
        public Output<ImmutableArray<string>> TrafficManagerHostNames { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// State indicating whether the app has exceeded its quota usage. Read-only.
        /// </summary>
        [Output("usageState")]
        public Output<string> UsageState { get; private set; } = null!;


        /// <summary>
        /// Create a WebApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebApp(string name, WebAppArgs args, CustomResourceOptions? options = null)
            : base("azure-native:web/v20160801:WebApp", name, args ?? new WebAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebApp(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:web/v20160801:WebApp", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:web/v20150801:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20180201:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20181101:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20190801:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20200601:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20200901:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20201001:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20201201:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210101:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210115:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210201:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210301:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20220301:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20220901:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20230101:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20231201:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20240401:WebApp" },
                    new global::Pulumi.Alias { Type = "azure-native:web:WebApp" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WebApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebApp Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebApp(name, id, options);
        }
    }

    public sealed class WebAppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to enable client affinity; &lt;code&gt;false&lt;/code&gt; to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is &lt;code&gt;true&lt;/code&gt;.
        /// </summary>
        [Input("clientAffinityEnabled")]
        public Input<bool>? ClientAffinityEnabled { get; set; }

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to enable client certificate authentication (TLS mutual authentication); otherwise, &lt;code&gt;false&lt;/code&gt;. Default is &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Input("clientCertEnabled")]
        public Input<bool>? ClientCertEnabled { get; set; }

        /// <summary>
        /// If specified during app creation, the app is cloned from a source app.
        /// </summary>
        [Input("cloningInfo")]
        public Input<Inputs.CloningInfoArgs>? CloningInfo { get; set; }

        /// <summary>
        /// Size of the function container.
        /// </summary>
        [Input("containerSize")]
        public Input<int>? ContainerSize { get; set; }

        /// <summary>
        /// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
        /// </summary>
        [Input("dailyMemoryTimeQuota")]
        public Input<int>? DailyMemoryTimeQuota { get; set; }

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if the app is enabled; otherwise, &lt;code&gt;false&lt;/code&gt;. Setting this value to false disables the app (takes the app offline).
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("hostNameSslStates")]
        private InputList<Inputs.HostNameSslStateArgs>? _hostNameSslStates;

        /// <summary>
        /// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
        /// </summary>
        public InputList<Inputs.HostNameSslStateArgs> HostNameSslStates
        {
            get => _hostNameSslStates ?? (_hostNameSslStates = new InputList<Inputs.HostNameSslStateArgs>());
            set => _hostNameSslStates = value;
        }

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to disable the public hostnames of the app; otherwise, &lt;code&gt;false&lt;/code&gt;.
        ///  If &lt;code&gt;true&lt;/code&gt;, the app is only accessible via API management process.
        /// </summary>
        [Input("hostNamesDisabled")]
        public Input<bool>? HostNamesDisabled { get; set; }

        /// <summary>
        /// App Service Environment to use for the app.
        /// </summary>
        [Input("hostingEnvironmentProfile")]
        public Input<Inputs.HostingEnvironmentProfileArgs>? HostingEnvironmentProfile { get; set; }

        /// <summary>
        /// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
        /// http requests
        /// </summary>
        [Input("httpsOnly")]
        public Input<bool>? HttpsOnly { get; set; }

        /// <summary>
        /// Managed service identity.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Unique name of the app to create or update. To create or update a deployment slot, use the {slot} parameter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if reserved; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Input("reserved")]
        public Input<bool>? Reserved { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to stop SCM (KUDU) site when the app is stopped; otherwise, &lt;code&gt;false&lt;/code&gt;. The default is &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Input("scmSiteAlsoStopped")]
        public Input<bool>? ScmSiteAlsoStopped { get; set; }

        /// <summary>
        /// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
        /// </summary>
        [Input("serverFarmId")]
        public Input<string>? ServerFarmId { get; set; }

        /// <summary>
        /// Configuration of the app.
        /// </summary>
        [Input("siteConfig")]
        public Input<Inputs.SiteConfigArgs>? SiteConfig { get; set; }

        /// <summary>
        /// If specified during app creation, the app is created from a previous snapshot.
        /// </summary>
        [Input("snapshotInfo")]
        public Input<Inputs.SnapshotRecoveryRequestArgs>? SnapshotInfo { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public WebAppArgs()
        {
            Reserved = false;
            ScmSiteAlsoStopped = false;
        }
        public static new WebAppArgs Empty => new WebAppArgs();
    }
}
