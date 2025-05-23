// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Databricks
{
    [EnumType]
    public readonly struct AutomaticClusterUpdateValue : IEquatable<AutomaticClusterUpdateValue>
    {
        private readonly string _value;

        private AutomaticClusterUpdateValue(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AutomaticClusterUpdateValue Enabled { get; } = new AutomaticClusterUpdateValue("Enabled");
        public static AutomaticClusterUpdateValue Disabled { get; } = new AutomaticClusterUpdateValue("Disabled");

        public static bool operator ==(AutomaticClusterUpdateValue left, AutomaticClusterUpdateValue right) => left.Equals(right);
        public static bool operator !=(AutomaticClusterUpdateValue left, AutomaticClusterUpdateValue right) => !left.Equals(right);

        public static explicit operator string(AutomaticClusterUpdateValue value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AutomaticClusterUpdateValue other && Equals(other);
        public bool Equals(AutomaticClusterUpdateValue other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ComplianceSecurityProfileValue : IEquatable<ComplianceSecurityProfileValue>
    {
        private readonly string _value;

        private ComplianceSecurityProfileValue(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComplianceSecurityProfileValue Enabled { get; } = new ComplianceSecurityProfileValue("Enabled");
        public static ComplianceSecurityProfileValue Disabled { get; } = new ComplianceSecurityProfileValue("Disabled");

        public static bool operator ==(ComplianceSecurityProfileValue left, ComplianceSecurityProfileValue right) => left.Equals(right);
        public static bool operator !=(ComplianceSecurityProfileValue left, ComplianceSecurityProfileValue right) => !left.Equals(right);

        public static explicit operator string(ComplianceSecurityProfileValue value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComplianceSecurityProfileValue other && Equals(other);
        public bool Equals(ComplianceSecurityProfileValue other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Compliance standard that can be associated with a workspace.
    /// </summary>
    [EnumType]
    public readonly struct ComplianceStandard : IEquatable<ComplianceStandard>
    {
        private readonly string _value;

        private ComplianceStandard(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComplianceStandard NONE { get; } = new ComplianceStandard("NONE");
        public static ComplianceStandard HIPAA { get; } = new ComplianceStandard("HIPAA");
        public static ComplianceStandard PCI_DSS { get; } = new ComplianceStandard("PCI_DSS");

        public static bool operator ==(ComplianceStandard left, ComplianceStandard right) => left.Equals(right);
        public static bool operator !=(ComplianceStandard left, ComplianceStandard right) => !left.Equals(right);

        public static explicit operator string(ComplianceStandard value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComplianceStandard other && Equals(other);
        public bool Equals(ComplianceStandard other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or Sets Default Storage Firewall configuration information
    /// </summary>
    [EnumType]
    public readonly struct DefaultStorageFirewall : IEquatable<DefaultStorageFirewall>
    {
        private readonly string _value;

        private DefaultStorageFirewall(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DefaultStorageFirewall Disabled { get; } = new DefaultStorageFirewall("Disabled");
        public static DefaultStorageFirewall Enabled { get; } = new DefaultStorageFirewall("Enabled");

        public static bool operator ==(DefaultStorageFirewall left, DefaultStorageFirewall right) => left.Equals(right);
        public static bool operator !=(DefaultStorageFirewall left, DefaultStorageFirewall right) => !left.Equals(right);

        public static explicit operator string(DefaultStorageFirewall value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DefaultStorageFirewall other && Equals(other);
        public bool Equals(DefaultStorageFirewall other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Keyvault
    /// </summary>
    [EnumType]
    public readonly struct EncryptionKeySource : IEquatable<EncryptionKeySource>
    {
        private readonly string _value;

        private EncryptionKeySource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EncryptionKeySource Microsoft_Keyvault { get; } = new EncryptionKeySource("Microsoft.Keyvault");

        public static bool operator ==(EncryptionKeySource left, EncryptionKeySource right) => left.Equals(right);
        public static bool operator !=(EncryptionKeySource left, EncryptionKeySource right) => !left.Equals(right);

        public static explicit operator string(EncryptionKeySource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EncryptionKeySource other && Equals(other);
        public bool Equals(EncryptionKeySource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EnhancedSecurityMonitoringValue : IEquatable<EnhancedSecurityMonitoringValue>
    {
        private readonly string _value;

        private EnhancedSecurityMonitoringValue(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnhancedSecurityMonitoringValue Enabled { get; } = new EnhancedSecurityMonitoringValue("Enabled");
        public static EnhancedSecurityMonitoringValue Disabled { get; } = new EnhancedSecurityMonitoringValue("Disabled");

        public static bool operator ==(EnhancedSecurityMonitoringValue left, EnhancedSecurityMonitoringValue right) => left.Equals(right);
        public static bool operator !=(EnhancedSecurityMonitoringValue left, EnhancedSecurityMonitoringValue right) => !left.Equals(right);

        public static explicit operator string(EnhancedSecurityMonitoringValue value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnhancedSecurityMonitoringValue other && Equals(other);
        public bool Equals(EnhancedSecurityMonitoringValue other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The identity type of the Access Connector Resource.
    /// </summary>
    [EnumType]
    public readonly struct IdentityType : IEquatable<IdentityType>
    {
        private readonly string _value;

        private IdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IdentityType SystemAssigned { get; } = new IdentityType("SystemAssigned");
        public static IdentityType UserAssigned { get; } = new IdentityType("UserAssigned");

        public static bool operator ==(IdentityType left, IdentityType right) => left.Equals(right);
        public static bool operator !=(IdentityType left, IdentityType right) => !left.Equals(right);

        public static explicit operator string(IdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IdentityType other && Equals(other);
        public bool Equals(IdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Defines the initial type of the default catalog. Possible values (case-insensitive):  HiveMetastore, UnityCatalog
    /// </summary>
    [EnumType]
    public readonly struct InitialType : IEquatable<InitialType>
    {
        private readonly string _value;

        private InitialType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InitialType HiveMetastore { get; } = new InitialType("HiveMetastore");
        public static InitialType UnityCatalog { get; } = new InitialType("UnityCatalog");

        public static bool operator ==(InitialType left, InitialType right) => left.Equals(right);
        public static bool operator !=(InitialType left, InitialType right) => !left.Equals(right);

        public static explicit operator string(InitialType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InitialType other && Equals(other);
        public bool Equals(InitialType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The encryption keySource (provider). Possible values (case-insensitive):  Default, Microsoft.Keyvault
    /// </summary>
    [EnumType]
    public readonly struct KeySource : IEquatable<KeySource>
    {
        private readonly string _value;

        private KeySource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeySource Default { get; } = new KeySource("Default");
        public static KeySource Microsoft_Keyvault { get; } = new KeySource("Microsoft.Keyvault");

        public static bool operator ==(KeySource left, KeySource right) => left.Equals(right);
        public static bool operator !=(KeySource left, KeySource right) => !left.Equals(right);

        public static explicit operator string(KeySource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeySource other && Equals(other);
        public bool Equals(KeySource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
    /// </summary>
    [EnumType]
    public readonly struct ManagedServiceIdentityType : IEquatable<ManagedServiceIdentityType>
    {
        private readonly string _value;

        private ManagedServiceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ManagedServiceIdentityType None { get; } = new ManagedServiceIdentityType("None");
        public static ManagedServiceIdentityType SystemAssigned { get; } = new ManagedServiceIdentityType("SystemAssigned");
        public static ManagedServiceIdentityType UserAssigned { get; } = new ManagedServiceIdentityType("UserAssigned");
        public static ManagedServiceIdentityType SystemAssigned_UserAssigned { get; } = new ManagedServiceIdentityType("SystemAssigned,UserAssigned");

        public static bool operator ==(ManagedServiceIdentityType left, ManagedServiceIdentityType right) => left.Equals(right);
        public static bool operator !=(ManagedServiceIdentityType left, ManagedServiceIdentityType right) => !left.Equals(right);

        public static explicit operator string(ManagedServiceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ManagedServiceIdentityType other && Equals(other);
        public bool Equals(ManagedServiceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of a private endpoint connection
    /// </summary>
    [EnumType]
    public readonly struct PrivateLinkServiceConnectionStatus : IEquatable<PrivateLinkServiceConnectionStatus>
    {
        private readonly string _value;

        private PrivateLinkServiceConnectionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PrivateLinkServiceConnectionStatus Pending { get; } = new PrivateLinkServiceConnectionStatus("Pending");
        public static PrivateLinkServiceConnectionStatus Approved { get; } = new PrivateLinkServiceConnectionStatus("Approved");
        public static PrivateLinkServiceConnectionStatus Rejected { get; } = new PrivateLinkServiceConnectionStatus("Rejected");
        public static PrivateLinkServiceConnectionStatus Disconnected { get; } = new PrivateLinkServiceConnectionStatus("Disconnected");

        public static bool operator ==(PrivateLinkServiceConnectionStatus left, PrivateLinkServiceConnectionStatus right) => left.Equals(right);
        public static bool operator !=(PrivateLinkServiceConnectionStatus left, PrivateLinkServiceConnectionStatus right) => !left.Equals(right);

        public static explicit operator string(PrivateLinkServiceConnectionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PrivateLinkServiceConnectionStatus other && Equals(other);
        public bool Equals(PrivateLinkServiceConnectionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The network access type for accessing workspace. Set value to disabled to access workspace only via private link.
    /// </summary>
    [EnumType]
    public readonly struct PublicNetworkAccess : IEquatable<PublicNetworkAccess>
    {
        private readonly string _value;

        private PublicNetworkAccess(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PublicNetworkAccess Enabled { get; } = new PublicNetworkAccess("Enabled");
        public static PublicNetworkAccess Disabled { get; } = new PublicNetworkAccess("Disabled");

        public static bool operator ==(PublicNetworkAccess left, PublicNetworkAccess right) => left.Equals(right);
        public static bool operator !=(PublicNetworkAccess left, PublicNetworkAccess right) => !left.Equals(right);

        public static explicit operator string(PublicNetworkAccess value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PublicNetworkAccess other && Equals(other);
        public bool Equals(PublicNetworkAccess other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets a value indicating whether data plane (clusters) to control plane communication happen over private endpoint. Supported values are 'AllRules' and 'NoAzureDatabricksRules'. 'NoAzureServiceRules' value is for internal use only.
    /// </summary>
    [EnumType]
    public readonly struct RequiredNsgRules : IEquatable<RequiredNsgRules>
    {
        private readonly string _value;

        private RequiredNsgRules(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RequiredNsgRules AllRules { get; } = new RequiredNsgRules("AllRules");
        public static RequiredNsgRules NoAzureDatabricksRules { get; } = new RequiredNsgRules("NoAzureDatabricksRules");
        public static RequiredNsgRules NoAzureServiceRules { get; } = new RequiredNsgRules("NoAzureServiceRules");

        public static bool operator ==(RequiredNsgRules left, RequiredNsgRules right) => left.Equals(right);
        public static bool operator !=(RequiredNsgRules left, RequiredNsgRules right) => !left.Equals(right);

        public static explicit operator string(RequiredNsgRules value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RequiredNsgRules other && Equals(other);
        public bool Equals(RequiredNsgRules other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
