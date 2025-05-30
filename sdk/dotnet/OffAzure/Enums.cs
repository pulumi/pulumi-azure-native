// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.OffAzure
{
    /// <summary>
    /// Gets or sets the state of public network access.
    /// </summary>
    [EnumType]
    public readonly struct MasterSitePropertiesPublicNetworkAccess : IEquatable<MasterSitePropertiesPublicNetworkAccess>
    {
        private readonly string _value;

        private MasterSitePropertiesPublicNetworkAccess(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// NotSpecified value.
        /// </summary>
        public static MasterSitePropertiesPublicNetworkAccess NotSpecified { get; } = new MasterSitePropertiesPublicNetworkAccess("NotSpecified");
        /// <summary>
        /// Enabled value.
        /// </summary>
        public static MasterSitePropertiesPublicNetworkAccess Enabled { get; } = new MasterSitePropertiesPublicNetworkAccess("Enabled");
        /// <summary>
        /// Disabled value.
        /// </summary>
        public static MasterSitePropertiesPublicNetworkAccess Disabled { get; } = new MasterSitePropertiesPublicNetworkAccess("Disabled");

        public static bool operator ==(MasterSitePropertiesPublicNetworkAccess left, MasterSitePropertiesPublicNetworkAccess right) => left.Equals(right);
        public static bool operator !=(MasterSitePropertiesPublicNetworkAccess left, MasterSitePropertiesPublicNetworkAccess right) => !left.Equals(right);

        public static explicit operator string(MasterSitePropertiesPublicNetworkAccess value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MasterSitePropertiesPublicNetworkAccess other && Equals(other);
        public bool Equals(MasterSitePropertiesPublicNetworkAccess other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// state status
    /// </summary>
    [EnumType]
    public readonly struct PrivateLinkServiceConnectionStateStatus : IEquatable<PrivateLinkServiceConnectionStateStatus>
    {
        private readonly string _value;

        private PrivateLinkServiceConnectionStateStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Approved value.
        /// </summary>
        public static PrivateLinkServiceConnectionStateStatus Approved { get; } = new PrivateLinkServiceConnectionStateStatus("Approved");
        /// <summary>
        /// Pending value.
        /// </summary>
        public static PrivateLinkServiceConnectionStateStatus Pending { get; } = new PrivateLinkServiceConnectionStateStatus("Pending");
        /// <summary>
        /// Rejected value.
        /// </summary>
        public static PrivateLinkServiceConnectionStateStatus Rejected { get; } = new PrivateLinkServiceConnectionStateStatus("Rejected");
        /// <summary>
        /// Disconnected value.
        /// </summary>
        public static PrivateLinkServiceConnectionStateStatus Disconnected { get; } = new PrivateLinkServiceConnectionStateStatus("Disconnected");

        public static bool operator ==(PrivateLinkServiceConnectionStateStatus left, PrivateLinkServiceConnectionStateStatus right) => left.Equals(right);
        public static bool operator !=(PrivateLinkServiceConnectionStateStatus left, PrivateLinkServiceConnectionStateStatus right) => !left.Equals(right);

        public static explicit operator string(PrivateLinkServiceConnectionStateStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PrivateLinkServiceConnectionStateStatus other && Equals(other);
        public bool Equals(PrivateLinkServiceConnectionStateStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the last operation.
    /// </summary>
    [EnumType]
    public readonly struct ProvisioningState : IEquatable<ProvisioningState>
    {
        private readonly string _value;

        private ProvisioningState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Created value.
        /// </summary>
        public static ProvisioningState Created { get; } = new ProvisioningState("Created");
        /// <summary>
        /// Updated value.
        /// </summary>
        public static ProvisioningState Updated { get; } = new ProvisioningState("Updated");
        /// <summary>
        /// Running value.
        /// </summary>
        public static ProvisioningState Running { get; } = new ProvisioningState("Running");
        /// <summary>
        /// Completed value.
        /// </summary>
        public static ProvisioningState Completed { get; } = new ProvisioningState("Completed");
        /// <summary>
        /// Failed value.
        /// </summary>
        public static ProvisioningState Failed { get; } = new ProvisioningState("Failed");
        /// <summary>
        /// Succeeded value.
        /// </summary>
        public static ProvisioningState Succeeded { get; } = new ProvisioningState("Succeeded");
        /// <summary>
        /// Canceled value.
        /// </summary>
        public static ProvisioningState Canceled { get; } = new ProvisioningState("Canceled");

        public static bool operator ==(ProvisioningState left, ProvisioningState right) => left.Equals(right);
        public static bool operator !=(ProvisioningState left, ProvisioningState right) => !left.Equals(right);

        public static explicit operator string(ProvisioningState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProvisioningState other && Equals(other);
        public bool Equals(ProvisioningState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the discovery scenario.
    /// </summary>
    [EnumType]
    public readonly struct SqlSitePropertiesDiscoveryScenario : IEquatable<SqlSitePropertiesDiscoveryScenario>
    {
        private readonly string _value;

        private SqlSitePropertiesDiscoveryScenario(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Migrate value.
        /// </summary>
        public static SqlSitePropertiesDiscoveryScenario Migrate { get; } = new SqlSitePropertiesDiscoveryScenario("Migrate");
        /// <summary>
        /// DR value.
        /// </summary>
        public static SqlSitePropertiesDiscoveryScenario DR { get; } = new SqlSitePropertiesDiscoveryScenario("DR");

        public static bool operator ==(SqlSitePropertiesDiscoveryScenario left, SqlSitePropertiesDiscoveryScenario right) => left.Equals(right);
        public static bool operator !=(SqlSitePropertiesDiscoveryScenario left, SqlSitePropertiesDiscoveryScenario right) => !left.Equals(right);

        public static explicit operator string(SqlSitePropertiesDiscoveryScenario value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SqlSitePropertiesDiscoveryScenario other && Equals(other);
        public bool Equals(SqlSitePropertiesDiscoveryScenario other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the discovery scenario.
    /// </summary>
    [EnumType]
    public readonly struct WebAppSitePropertiesDiscoveryScenario : IEquatable<WebAppSitePropertiesDiscoveryScenario>
    {
        private readonly string _value;

        private WebAppSitePropertiesDiscoveryScenario(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Migrate value.
        /// </summary>
        public static WebAppSitePropertiesDiscoveryScenario Migrate { get; } = new WebAppSitePropertiesDiscoveryScenario("Migrate");
        /// <summary>
        /// DR value.
        /// </summary>
        public static WebAppSitePropertiesDiscoveryScenario DR { get; } = new WebAppSitePropertiesDiscoveryScenario("DR");

        public static bool operator ==(WebAppSitePropertiesDiscoveryScenario left, WebAppSitePropertiesDiscoveryScenario right) => left.Equals(right);
        public static bool operator !=(WebAppSitePropertiesDiscoveryScenario left, WebAppSitePropertiesDiscoveryScenario right) => !left.Equals(right);

        public static explicit operator string(WebAppSitePropertiesDiscoveryScenario value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WebAppSitePropertiesDiscoveryScenario other && Equals(other);
        public bool Equals(WebAppSitePropertiesDiscoveryScenario other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
