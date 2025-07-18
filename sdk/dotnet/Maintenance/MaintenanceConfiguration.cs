// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Maintenance
{
    /// <summary>
    /// Maintenance configuration record type
    /// 
    /// Uses Azure REST API version 2023-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-11-01-preview.
    /// 
    /// Other available API versions: 2022-11-01-preview, 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:maintenance:MaintenanceConfiguration")]
    public partial class MaintenanceConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Duration of the maintenance window in HH:mm format. If not provided, default value will be used based on maintenance scope provided. Example: 05:00.
        /// </summary>
        [Output("duration")]
        public Output<string?> Duration { get; private set; } = null!;

        /// <summary>
        /// Effective expiration date of the maintenance window in YYYY-MM-DD hh:mm format. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone. Expiration date must be set to a future date. If not provided, it will be set to the maximum datetime 9999-12-31 23:59:59.
        /// </summary>
        [Output("expirationDateTime")]
        public Output<string?> ExpirationDateTime { get; private set; } = null!;

        /// <summary>
        /// Gets or sets extensionProperties of the maintenanceConfiguration
        /// </summary>
        [Output("extensionProperties")]
        public Output<ImmutableDictionary<string, string>?> ExtensionProperties { get; private set; } = null!;

        /// <summary>
        /// The input parameters to be passed to the patch run operation.
        /// </summary>
        [Output("installPatches")]
        public Output<Outputs.InputPatchConfigurationResponse?> InstallPatches { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Gets or sets maintenanceScope of the configuration
        /// </summary>
        [Output("maintenanceScope")]
        public Output<string?> MaintenanceScope { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets namespace of the resource
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Rate at which a Maintenance window is expected to recur. The rate can be expressed as daily, weekly, or monthly schedules. Daily schedule are formatted as recurEvery: [Frequency as integer]['Day(s)']. If no frequency is provided, the default frequency is 1. Daily schedule examples are recurEvery: Day, recurEvery: 3Days.  Weekly schedule are formatted as recurEvery: [Frequency as integer]['Week(s)'] [Optional comma separated list of weekdays Monday-Sunday]. Weekly schedule examples are recurEvery: 3Weeks, recurEvery: Week Saturday,Sunday. Monthly schedules are formatted as [Frequency as integer]['Month(s)'] [Comma separated list of month days] or [Frequency as integer]['Month(s)'] [Week of Month (First, Second, Third, Fourth, Last)] [Weekday Monday-Sunday] [Optional Offset(No. of days)]. Offset value must be between -6 to 6 inclusive. Monthly schedule examples are recurEvery: Month, recurEvery: 2Months, recurEvery: Month day23,day24, recurEvery: Month Last Sunday, recurEvery: Month Fourth Monday, recurEvery: Month Last Sunday Offset-3, recurEvery: Month Third Sunday Offset6.
        /// </summary>
        [Output("recurEvery")]
        public Output<string?> RecurEvery { get; private set; } = null!;

        /// <summary>
        /// Effective start date of the maintenance window in YYYY-MM-DD hh:mm format. The start date can be set to either the current date or future date. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone.
        /// </summary>
        [Output("startDateTime")]
        public Output<string?> StartDateTime { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Name of the timezone. List of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell. Example: Pacific Standard Time, UTC, W. Europe Standard Time, Korea Standard Time, Cen. Australia Standard Time.
        /// </summary>
        [Output("timeZone")]
        public Output<string?> TimeZone { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the visibility of the configuration. The default value is 'Custom'
        /// </summary>
        [Output("visibility")]
        public Output<string?> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a MaintenanceConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MaintenanceConfiguration(string name, MaintenanceConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azure-native:maintenance:MaintenanceConfiguration", name, args ?? new MaintenanceConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MaintenanceConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:maintenance:MaintenanceConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20180601preview:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20200401:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20200701preview:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20210401preview:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20210501:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20210901preview:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20220701preview:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20221101preview:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20230401:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20230901preview:MaintenanceConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20231001preview:MaintenanceConfiguration" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MaintenanceConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MaintenanceConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MaintenanceConfiguration(name, id, options);
        }
    }

    public sealed class MaintenanceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Duration of the maintenance window in HH:mm format. If not provided, default value will be used based on maintenance scope provided. Example: 05:00.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// Effective expiration date of the maintenance window in YYYY-MM-DD hh:mm format. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone. Expiration date must be set to a future date. If not provided, it will be set to the maximum datetime 9999-12-31 23:59:59.
        /// </summary>
        [Input("expirationDateTime")]
        public Input<string>? ExpirationDateTime { get; set; }

        [Input("extensionProperties")]
        private InputMap<string>? _extensionProperties;

        /// <summary>
        /// Gets or sets extensionProperties of the maintenanceConfiguration
        /// </summary>
        public InputMap<string> ExtensionProperties
        {
            get => _extensionProperties ?? (_extensionProperties = new InputMap<string>());
            set => _extensionProperties = value;
        }

        /// <summary>
        /// The input parameters to be passed to the patch run operation.
        /// </summary>
        [Input("installPatches")]
        public Input<Inputs.InputPatchConfigurationArgs>? InstallPatches { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Gets or sets maintenanceScope of the configuration
        /// </summary>
        [Input("maintenanceScope")]
        public InputUnion<string, Pulumi.AzureNative.Maintenance.MaintenanceScope>? MaintenanceScope { get; set; }

        /// <summary>
        /// Gets or sets namespace of the resource
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Rate at which a Maintenance window is expected to recur. The rate can be expressed as daily, weekly, or monthly schedules. Daily schedule are formatted as recurEvery: [Frequency as integer]['Day(s)']. If no frequency is provided, the default frequency is 1. Daily schedule examples are recurEvery: Day, recurEvery: 3Days.  Weekly schedule are formatted as recurEvery: [Frequency as integer]['Week(s)'] [Optional comma separated list of weekdays Monday-Sunday]. Weekly schedule examples are recurEvery: 3Weeks, recurEvery: Week Saturday,Sunday. Monthly schedules are formatted as [Frequency as integer]['Month(s)'] [Comma separated list of month days] or [Frequency as integer]['Month(s)'] [Week of Month (First, Second, Third, Fourth, Last)] [Weekday Monday-Sunday] [Optional Offset(No. of days)]. Offset value must be between -6 to 6 inclusive. Monthly schedule examples are recurEvery: Month, recurEvery: 2Months, recurEvery: Month day23,day24, recurEvery: Month Last Sunday, recurEvery: Month Fourth Monday, recurEvery: Month Last Sunday Offset-3, recurEvery: Month Third Sunday Offset6.
        /// </summary>
        [Input("recurEvery")]
        public Input<string>? RecurEvery { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the MaintenanceConfiguration
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        /// <summary>
        /// Effective start date of the maintenance window in YYYY-MM-DD hh:mm format. The start date can be set to either the current date or future date. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone.
        /// </summary>
        [Input("startDateTime")]
        public Input<string>? StartDateTime { get; set; }

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

        /// <summary>
        /// Name of the timezone. List of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell. Example: Pacific Standard Time, UTC, W. Europe Standard Time, Korea Standard Time, Cen. Australia Standard Time.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// Gets or sets the visibility of the configuration. The default value is 'Custom'
        /// </summary>
        [Input("visibility")]
        public InputUnion<string, Pulumi.AzureNative.Maintenance.Visibility>? Visibility { get; set; }

        public MaintenanceConfigurationArgs()
        {
        }
        public static new MaintenanceConfigurationArgs Empty => new MaintenanceConfigurationArgs();
    }
}
