// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.AVS
{
    /// <summary>
    /// Addon type
    /// </summary>
    [EnumType]
    public readonly struct AddonType : IEquatable<AddonType>
    {
        private readonly string _value;

        private AddonType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AddonType SRM { get; } = new AddonType("SRM");
        public static AddonType VR { get; } = new AddonType("VR");
        public static AddonType HCX { get; } = new AddonType("HCX");
        public static AddonType Arc { get; } = new AddonType("Arc");

        public static bool operator ==(AddonType left, AddonType right) => left.Equals(right);
        public static bool operator !=(AddonType left, AddonType right) => !left.Equals(right);

        public static explicit operator string(AddonType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AddonType other && Equals(other);
        public bool Equals(AddonType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The availability strategy for the private cloud
    /// </summary>
    [EnumType]
    public readonly struct AvailabilityStrategy : IEquatable<AvailabilityStrategy>
    {
        private readonly string _value;

        private AvailabilityStrategy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// in single zone
        /// </summary>
        public static AvailabilityStrategy SingleZone { get; } = new AvailabilityStrategy("SingleZone");
        /// <summary>
        /// in two zones
        /// </summary>
        public static AvailabilityStrategy DualZone { get; } = new AvailabilityStrategy("DualZone");

        public static bool operator ==(AvailabilityStrategy left, AvailabilityStrategy right) => left.Equals(right);
        public static bool operator !=(AvailabilityStrategy left, AvailabilityStrategy right) => !left.Equals(right);

        public static explicit operator string(AvailabilityStrategy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AvailabilityStrategy other && Equals(other);
        public bool Equals(AvailabilityStrategy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of DHCP: SERVER or RELAY.
    /// </summary>
    [EnumType]
    public readonly struct DhcpTypeEnum : IEquatable<DhcpTypeEnum>
    {
        private readonly string _value;

        private DhcpTypeEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DhcpTypeEnum SERVER { get; } = new DhcpTypeEnum("SERVER");
        public static DhcpTypeEnum RELAY { get; } = new DhcpTypeEnum("RELAY");

        public static bool operator ==(DhcpTypeEnum left, DhcpTypeEnum right) => left.Equals(right);
        public static bool operator !=(DhcpTypeEnum left, DhcpTypeEnum right) => !left.Equals(right);

        public static explicit operator string(DhcpTypeEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DhcpTypeEnum other && Equals(other);
        public bool Equals(DhcpTypeEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// DNS Service log level.
    /// </summary>
    [EnumType]
    public readonly struct DnsServiceLogLevelEnum : IEquatable<DnsServiceLogLevelEnum>
    {
        private readonly string _value;

        private DnsServiceLogLevelEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// is debug
        /// </summary>
        public static DnsServiceLogLevelEnum DEBUG { get; } = new DnsServiceLogLevelEnum("DEBUG");
        /// <summary>
        /// is info
        /// </summary>
        public static DnsServiceLogLevelEnum INFO { get; } = new DnsServiceLogLevelEnum("INFO");
        /// <summary>
        /// is warning
        /// </summary>
        public static DnsServiceLogLevelEnum WARNING { get; } = new DnsServiceLogLevelEnum("WARNING");
        /// <summary>
        /// is error
        /// </summary>
        public static DnsServiceLogLevelEnum ERROR { get; } = new DnsServiceLogLevelEnum("ERROR");
        /// <summary>
        /// is fatal
        /// </summary>
        public static DnsServiceLogLevelEnum FATAL { get; } = new DnsServiceLogLevelEnum("FATAL");

        public static bool operator ==(DnsServiceLogLevelEnum left, DnsServiceLogLevelEnum right) => left.Equals(right);
        public static bool operator !=(DnsServiceLogLevelEnum left, DnsServiceLogLevelEnum right) => !left.Equals(right);

        public static explicit operator string(DnsServiceLogLevelEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DnsServiceLogLevelEnum other && Equals(other);
        public bool Equals(DnsServiceLogLevelEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of DNS zone to use.
    /// </summary>
    [EnumType]
    public readonly struct DnsZoneType : IEquatable<DnsZoneType>
    {
        private readonly string _value;

        private DnsZoneType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Primary DNS zone.
        /// </summary>
        public static DnsZoneType Public { get; } = new DnsZoneType("Public");
        /// <summary>
        /// Private DNS zone.
        /// </summary>
        public static DnsZoneType Private { get; } = new DnsZoneType("Private");

        public static bool operator ==(DnsZoneType left, DnsZoneType right) => left.Equals(right);
        public static bool operator !=(DnsZoneType left, DnsZoneType right) => !left.Equals(right);

        public static explicit operator string(DnsZoneType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DnsZoneType other && Equals(other);
        public bool Equals(DnsZoneType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Status of customer managed encryption key
    /// </summary>
    [EnumType]
    public readonly struct EncryptionState : IEquatable<EncryptionState>
    {
        private readonly string _value;

        private EncryptionState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// is enabled
        /// </summary>
        public static EncryptionState Enabled { get; } = new EncryptionState("Enabled");
        /// <summary>
        /// is disabled
        /// </summary>
        public static EncryptionState Disabled { get; } = new EncryptionState("Disabled");

        public static bool operator ==(EncryptionState left, EncryptionState right) => left.Equals(right);
        public static bool operator !=(EncryptionState left, EncryptionState right) => !left.Equals(right);

        public static explicit operator string(EncryptionState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EncryptionState other && Equals(other);
        public bool Equals(EncryptionState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Connectivity to internet is enabled or disabled
    /// </summary>
    [EnumType]
    public readonly struct InternetEnum : IEquatable<InternetEnum>
    {
        private readonly string _value;

        private InternetEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// is enabled
        /// </summary>
        public static InternetEnum Enabled { get; } = new InternetEnum("Enabled");
        /// <summary>
        /// is disabled
        /// </summary>
        public static InternetEnum Disabled { get; } = new InternetEnum("Disabled");

        public static bool operator ==(InternetEnum left, InternetEnum right) => left.Equals(right);
        public static bool operator !=(InternetEnum left, InternetEnum right) => !left.Equals(right);

        public static explicit operator string(InternetEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InternetEnum other && Equals(other);
        public bool Equals(InternetEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Mode that describes whether the LUN has to be mounted as a datastore or
    /// attached as a LUN
    /// </summary>
    [EnumType]
    public readonly struct MountOptionEnum : IEquatable<MountOptionEnum>
    {
        private readonly string _value;

        private MountOptionEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// is mount
        /// </summary>
        public static MountOptionEnum MOUNT { get; } = new MountOptionEnum("MOUNT");
        /// <summary>
        /// is attach
        /// </summary>
        public static MountOptionEnum ATTACH { get; } = new MountOptionEnum("ATTACH");

        public static bool operator ==(MountOptionEnum left, MountOptionEnum right) => left.Equals(right);
        public static bool operator !=(MountOptionEnum left, MountOptionEnum right) => !left.Equals(right);

        public static explicit operator string(MountOptionEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MountOptionEnum other && Equals(other);
        public bool Equals(MountOptionEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Whether the placement policy is enabled or disabled
    /// </summary>
    [EnumType]
    public readonly struct PlacementPolicyState : IEquatable<PlacementPolicyState>
    {
        private readonly string _value;

        private PlacementPolicyState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// is enabled
        /// </summary>
        public static PlacementPolicyState Enabled { get; } = new PlacementPolicyState("Enabled");
        /// <summary>
        /// is disabled
        /// </summary>
        public static PlacementPolicyState Disabled { get; } = new PlacementPolicyState("Disabled");

        public static bool operator ==(PlacementPolicyState left, PlacementPolicyState right) => left.Equals(right);
        public static bool operator !=(PlacementPolicyState left, PlacementPolicyState right) => !left.Equals(right);

        public static explicit operator string(PlacementPolicyState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PlacementPolicyState other && Equals(other);
        public bool Equals(PlacementPolicyState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Placement Policy type
    /// </summary>
    [EnumType]
    public readonly struct PlacementPolicyType : IEquatable<PlacementPolicyType>
    {
        private readonly string _value;

        private PlacementPolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PlacementPolicyType VmVm { get; } = new PlacementPolicyType("VmVm");
        public static PlacementPolicyType VmHost { get; } = new PlacementPolicyType("VmHost");

        public static bool operator ==(PlacementPolicyType left, PlacementPolicyType right) => left.Equals(right);
        public static bool operator !=(PlacementPolicyType left, PlacementPolicyType right) => !left.Equals(right);

        public static explicit operator string(PlacementPolicyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PlacementPolicyType other && Equals(other);
        public bool Equals(PlacementPolicyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Direction of port mirroring profile.
    /// </summary>
    [EnumType]
    public readonly struct PortMirroringDirectionEnum : IEquatable<PortMirroringDirectionEnum>
    {
        private readonly string _value;

        private PortMirroringDirectionEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// is ingress
        /// </summary>
        public static PortMirroringDirectionEnum INGRESS { get; } = new PortMirroringDirectionEnum("INGRESS");
        /// <summary>
        /// is egress
        /// </summary>
        public static PortMirroringDirectionEnum EGRESS { get; } = new PortMirroringDirectionEnum("EGRESS");
        /// <summary>
        /// is bidirectional
        /// </summary>
        public static PortMirroringDirectionEnum BIDIRECTIONAL { get; } = new PortMirroringDirectionEnum("BIDIRECTIONAL");

        public static bool operator ==(PortMirroringDirectionEnum left, PortMirroringDirectionEnum right) => left.Equals(right);
        public static bool operator !=(PortMirroringDirectionEnum left, PortMirroringDirectionEnum right) => !left.Equals(right);

        public static explicit operator string(PortMirroringDirectionEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PortMirroringDirectionEnum other && Equals(other);
        public bool Equals(PortMirroringDirectionEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// script execution parameter type
    /// </summary>
    [EnumType]
    public readonly struct ScriptExecutionParameterType : IEquatable<ScriptExecutionParameterType>
    {
        private readonly string _value;

        private ScriptExecutionParameterType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScriptExecutionParameterType Value { get; } = new ScriptExecutionParameterType("Value");
        public static ScriptExecutionParameterType SecureValue { get; } = new ScriptExecutionParameterType("SecureValue");
        public static ScriptExecutionParameterType Credential { get; } = new ScriptExecutionParameterType("Credential");

        public static bool operator ==(ScriptExecutionParameterType left, ScriptExecutionParameterType right) => left.Equals(right);
        public static bool operator !=(ScriptExecutionParameterType left, ScriptExecutionParameterType right) => !left.Equals(right);

        public static explicit operator string(ScriptExecutionParameterType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScriptExecutionParameterType other && Equals(other);
        public bool Equals(ScriptExecutionParameterType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
    /// </summary>
    [EnumType]
    public readonly struct SkuTier : IEquatable<SkuTier>
    {
        private readonly string _value;

        private SkuTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuTier Free { get; } = new SkuTier("Free");
        public static SkuTier Basic { get; } = new SkuTier("Basic");
        public static SkuTier Standard { get; } = new SkuTier("Standard");
        public static SkuTier Premium { get; } = new SkuTier("Premium");

        public static bool operator ==(SkuTier left, SkuTier right) => left.Equals(right);
        public static bool operator !=(SkuTier left, SkuTier right) => !left.Equals(right);

        public static explicit operator string(SkuTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuTier other && Equals(other);
        public bool Equals(SkuTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Protect LDAP communication using SSL certificate (LDAPS)
    /// </summary>
    [EnumType]
    public readonly struct SslEnum : IEquatable<SslEnum>
    {
        private readonly string _value;

        private SslEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// is enabled
        /// </summary>
        public static SslEnum Enabled { get; } = new SslEnum("Enabled");
        /// <summary>
        /// is disabled
        /// </summary>
        public static SslEnum Disabled { get; } = new SslEnum("Disabled");

        public static bool operator ==(SslEnum left, SslEnum right) => left.Equals(right);
        public static bool operator !=(SslEnum left, SslEnum right) => !left.Equals(right);

        public static explicit operator string(SslEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SslEnum other && Equals(other);
        public bool Equals(SslEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of managed service identity (either system assigned, or none).
    /// </summary>
    [EnumType]
    public readonly struct SystemAssignedServiceIdentityType : IEquatable<SystemAssignedServiceIdentityType>
    {
        private readonly string _value;

        private SystemAssignedServiceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SystemAssignedServiceIdentityType None { get; } = new SystemAssignedServiceIdentityType("None");
        public static SystemAssignedServiceIdentityType SystemAssigned { get; } = new SystemAssignedServiceIdentityType("SystemAssigned");

        public static bool operator ==(SystemAssignedServiceIdentityType left, SystemAssignedServiceIdentityType right) => left.Equals(right);
        public static bool operator !=(SystemAssignedServiceIdentityType left, SystemAssignedServiceIdentityType right) => !left.Equals(right);

        public static explicit operator string(SystemAssignedServiceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SystemAssignedServiceIdentityType other && Equals(other);
        public bool Equals(SystemAssignedServiceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
