// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.DataProtection
{
    [EnumType]
    public readonly struct AKSVolumeTypes : IEquatable<AKSVolumeTypes>
    {
        private readonly string _value;

        private AKSVolumeTypes(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AKSVolumeTypes AzureDisk { get; } = new AKSVolumeTypes("AzureDisk");
        public static AKSVolumeTypes AzureFileShareSMB { get; } = new AKSVolumeTypes("AzureFileShareSMB");

        public static bool operator ==(AKSVolumeTypes left, AKSVolumeTypes right) => left.Equals(right);
        public static bool operator !=(AKSVolumeTypes left, AKSVolumeTypes right) => !left.Equals(right);

        public static explicit operator string(AKSVolumeTypes value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AKSVolumeTypes other && Equals(other);
        public bool Equals(AKSVolumeTypes other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct AbsoluteMarker : IEquatable<AbsoluteMarker>
    {
        private readonly string _value;

        private AbsoluteMarker(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AbsoluteMarker AllBackup { get; } = new AbsoluteMarker("AllBackup");
        public static AbsoluteMarker FirstOfDay { get; } = new AbsoluteMarker("FirstOfDay");
        public static AbsoluteMarker FirstOfMonth { get; } = new AbsoluteMarker("FirstOfMonth");
        public static AbsoluteMarker FirstOfWeek { get; } = new AbsoluteMarker("FirstOfWeek");
        public static AbsoluteMarker FirstOfYear { get; } = new AbsoluteMarker("FirstOfYear");

        public static bool operator ==(AbsoluteMarker left, AbsoluteMarker right) => left.Equals(right);
        public static bool operator !=(AbsoluteMarker left, AbsoluteMarker right) => !left.Equals(right);

        public static explicit operator string(AbsoluteMarker value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AbsoluteMarker other && Equals(other);
        public bool Equals(AbsoluteMarker other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct AlertsState : IEquatable<AlertsState>
    {
        private readonly string _value;

        private AlertsState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertsState Enabled { get; } = new AlertsState("Enabled");
        public static AlertsState Disabled { get; } = new AlertsState("Disabled");

        public static bool operator ==(AlertsState left, AlertsState right) => left.Equals(right);
        public static bool operator !=(AlertsState left, AlertsState right) => !left.Equals(right);

        public static explicit operator string(AlertsState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertsState other && Equals(other);
        public bool Equals(AlertsState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// CrossRegionRestore state
    /// </summary>
    [EnumType]
    public readonly struct CrossRegionRestoreState : IEquatable<CrossRegionRestoreState>
    {
        private readonly string _value;

        private CrossRegionRestoreState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CrossRegionRestoreState Disabled { get; } = new CrossRegionRestoreState("Disabled");
        public static CrossRegionRestoreState Enabled { get; } = new CrossRegionRestoreState("Enabled");

        public static bool operator ==(CrossRegionRestoreState left, CrossRegionRestoreState right) => left.Equals(right);
        public static bool operator !=(CrossRegionRestoreState left, CrossRegionRestoreState right) => !left.Equals(right);

        public static explicit operator string(CrossRegionRestoreState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CrossRegionRestoreState other && Equals(other);
        public bool Equals(CrossRegionRestoreState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// CrossSubscriptionRestore state
    /// </summary>
    [EnumType]
    public readonly struct CrossSubscriptionRestoreState : IEquatable<CrossSubscriptionRestoreState>
    {
        private readonly string _value;

        private CrossSubscriptionRestoreState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CrossSubscriptionRestoreState Disabled { get; } = new CrossSubscriptionRestoreState("Disabled");
        public static CrossSubscriptionRestoreState PermanentlyDisabled { get; } = new CrossSubscriptionRestoreState("PermanentlyDisabled");
        public static CrossSubscriptionRestoreState Enabled { get; } = new CrossSubscriptionRestoreState("Enabled");

        public static bool operator ==(CrossSubscriptionRestoreState left, CrossSubscriptionRestoreState right) => left.Equals(right);
        public static bool operator !=(CrossSubscriptionRestoreState left, CrossSubscriptionRestoreState right) => !left.Equals(right);

        public static explicit operator string(CrossSubscriptionRestoreState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CrossSubscriptionRestoreState other && Equals(other);
        public bool Equals(CrossSubscriptionRestoreState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// type of datastore; Operational/Vault/Archive
    /// </summary>
    [EnumType]
    public readonly struct DataStoreTypes : IEquatable<DataStoreTypes>
    {
        private readonly string _value;

        private DataStoreTypes(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataStoreTypes OperationalStore { get; } = new DataStoreTypes("OperationalStore");
        public static DataStoreTypes VaultStore { get; } = new DataStoreTypes("VaultStore");
        public static DataStoreTypes ArchiveStore { get; } = new DataStoreTypes("ArchiveStore");

        public static bool operator ==(DataStoreTypes left, DataStoreTypes right) => left.Equals(right);
        public static bool operator !=(DataStoreTypes left, DataStoreTypes right) => !left.Equals(right);

        public static explicit operator string(DataStoreTypes value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataStoreTypes other && Equals(other);
        public bool Equals(DataStoreTypes other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DayOfWeek : IEquatable<DayOfWeek>
    {
        private readonly string _value;

        private DayOfWeek(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DayOfWeek Friday { get; } = new DayOfWeek("Friday");
        public static DayOfWeek Monday { get; } = new DayOfWeek("Monday");
        public static DayOfWeek Saturday { get; } = new DayOfWeek("Saturday");
        public static DayOfWeek Sunday { get; } = new DayOfWeek("Sunday");
        public static DayOfWeek Thursday { get; } = new DayOfWeek("Thursday");
        public static DayOfWeek Tuesday { get; } = new DayOfWeek("Tuesday");
        public static DayOfWeek Wednesday { get; } = new DayOfWeek("Wednesday");

        public static bool operator ==(DayOfWeek left, DayOfWeek right) => left.Equals(right);
        public static bool operator !=(DayOfWeek left, DayOfWeek right) => !left.Equals(right);

        public static explicit operator string(DayOfWeek value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DayOfWeek other && Equals(other);
        public bool Equals(DayOfWeek other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Encryption state of the Backup Vault.
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
        /// CMK encryption is enabled on the Backup Vault
        /// </summary>
        public static EncryptionState Enabled { get; } = new EncryptionState("Enabled");
        /// <summary>
        /// CMK encryption is disabled on the Backup Vault. User can not set this state once Encryption State is 'Enabled'.
        /// </summary>
        public static EncryptionState Disabled { get; } = new EncryptionState("Disabled");
        /// <summary>
        /// CMK encryption is in inconsistent state on the Backup Vault. This state indicates that user needs to retry the encryption settings operation immediately to correct the state.
        /// </summary>
        public static EncryptionState Inconsistent { get; } = new EncryptionState("Inconsistent");

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
    /// The identity type. 'SystemAssigned' and 'UserAssigned' are mutually exclusive. 'SystemAssigned' will use implicitly created managed identity.
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
    /// Immutability state
    /// </summary>
    [EnumType]
    public readonly struct ImmutabilityState : IEquatable<ImmutabilityState>
    {
        private readonly string _value;

        private ImmutabilityState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ImmutabilityState Disabled { get; } = new ImmutabilityState("Disabled");
        public static ImmutabilityState Unlocked { get; } = new ImmutabilityState("Unlocked");
        public static ImmutabilityState Locked { get; } = new ImmutabilityState("Locked");

        public static bool operator ==(ImmutabilityState left, ImmutabilityState right) => left.Equals(right);
        public static bool operator !=(ImmutabilityState left, ImmutabilityState right) => !left.Equals(right);

        public static explicit operator string(ImmutabilityState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ImmutabilityState other && Equals(other);
        public bool Equals(ImmutabilityState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Enabling/Disabling the Double Encryption state
    /// </summary>
    [EnumType]
    public readonly struct InfrastructureEncryptionState : IEquatable<InfrastructureEncryptionState>
    {
        private readonly string _value;

        private InfrastructureEncryptionState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InfrastructureEncryptionState Enabled { get; } = new InfrastructureEncryptionState("Enabled");
        public static InfrastructureEncryptionState Disabled { get; } = new InfrastructureEncryptionState("Disabled");

        public static bool operator ==(InfrastructureEncryptionState left, InfrastructureEncryptionState right) => left.Equals(right);
        public static bool operator !=(InfrastructureEncryptionState left, InfrastructureEncryptionState right) => !left.Equals(right);

        public static explicit operator string(InfrastructureEncryptionState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InfrastructureEncryptionState other && Equals(other);
        public bool Equals(InfrastructureEncryptionState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct Month : IEquatable<Month>
    {
        private readonly string _value;

        private Month(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Month April { get; } = new Month("April");
        public static Month August { get; } = new Month("August");
        public static Month December { get; } = new Month("December");
        public static Month February { get; } = new Month("February");
        public static Month January { get; } = new Month("January");
        public static Month July { get; } = new Month("July");
        public static Month June { get; } = new Month("June");
        public static Month March { get; } = new Month("March");
        public static Month May { get; } = new Month("May");
        public static Month November { get; } = new Month("November");
        public static Month October { get; } = new Month("October");
        public static Month September { get; } = new Month("September");

        public static bool operator ==(Month left, Month right) => left.Equals(right);
        public static bool operator !=(Month left, Month right) => !left.Equals(right);

        public static explicit operator string(Month value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Month other && Equals(other);
        public bool Equals(Month other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the specific object - used for deserializing
    /// </summary>
    [EnumType]
    public readonly struct ResourcePropertiesObjectType : IEquatable<ResourcePropertiesObjectType>
    {
        private readonly string _value;

        private ResourcePropertiesObjectType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ResourcePropertiesObjectType DefaultResourceProperties { get; } = new ResourcePropertiesObjectType("DefaultResourceProperties");

        public static bool operator ==(ResourcePropertiesObjectType left, ResourcePropertiesObjectType right) => left.Equals(right);
        public static bool operator !=(ResourcePropertiesObjectType left, ResourcePropertiesObjectType right) => !left.Equals(right);

        public static explicit operator string(ResourcePropertiesObjectType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourcePropertiesObjectType other && Equals(other);
        public bool Equals(ResourcePropertiesObjectType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the type of secret store
    /// </summary>
    [EnumType]
    public readonly struct SecretStoreType : IEquatable<SecretStoreType>
    {
        private readonly string _value;

        private SecretStoreType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecretStoreType Invalid { get; } = new SecretStoreType("Invalid");
        public static SecretStoreType AzureKeyVault { get; } = new SecretStoreType("AzureKeyVault");

        public static bool operator ==(SecretStoreType left, SecretStoreType right) => left.Equals(right);
        public static bool operator !=(SecretStoreType left, SecretStoreType right) => !left.Equals(right);

        public static explicit operator string(SecretStoreType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecretStoreType other && Equals(other);
        public bool Equals(SecretStoreType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// State of soft delete
    /// </summary>
    [EnumType]
    public readonly struct SoftDeleteState : IEquatable<SoftDeleteState>
    {
        private readonly string _value;

        private SoftDeleteState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Soft Delete is turned off for the BackupVault
        /// </summary>
        public static SoftDeleteState Off { get; } = new SoftDeleteState("Off");
        /// <summary>
        /// Soft Delete is enabled for the BackupVault but can be turned off
        /// </summary>
        public static SoftDeleteState On { get; } = new SoftDeleteState("On");
        /// <summary>
        /// Soft Delete is permanently enabled for the BackupVault and the setting cannot be changed
        /// </summary>
        public static SoftDeleteState AlwaysOn { get; } = new SoftDeleteState("AlwaysOn");

        public static bool operator ==(SoftDeleteState left, SoftDeleteState right) => left.Equals(right);
        public static bool operator !=(SoftDeleteState left, SoftDeleteState right) => !left.Equals(right);

        public static explicit operator string(SoftDeleteState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SoftDeleteState other && Equals(other);
        public bool Equals(SoftDeleteState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the type of the datastore.
    /// </summary>
    [EnumType]
    public readonly struct StorageSettingStoreTypes : IEquatable<StorageSettingStoreTypes>
    {
        private readonly string _value;

        private StorageSettingStoreTypes(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StorageSettingStoreTypes ArchiveStore { get; } = new StorageSettingStoreTypes("ArchiveStore");
        public static StorageSettingStoreTypes OperationalStore { get; } = new StorageSettingStoreTypes("OperationalStore");
        public static StorageSettingStoreTypes VaultStore { get; } = new StorageSettingStoreTypes("VaultStore");

        public static bool operator ==(StorageSettingStoreTypes left, StorageSettingStoreTypes right) => left.Equals(right);
        public static bool operator !=(StorageSettingStoreTypes left, StorageSettingStoreTypes right) => !left.Equals(right);

        public static explicit operator string(StorageSettingStoreTypes value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StorageSettingStoreTypes other && Equals(other);
        public bool Equals(StorageSettingStoreTypes other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the type.
    /// </summary>
    [EnumType]
    public readonly struct StorageSettingTypes : IEquatable<StorageSettingTypes>
    {
        private readonly string _value;

        private StorageSettingTypes(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StorageSettingTypes GeoRedundant { get; } = new StorageSettingTypes("GeoRedundant");
        public static StorageSettingTypes LocallyRedundant { get; } = new StorageSettingTypes("LocallyRedundant");
        public static StorageSettingTypes ZoneRedundant { get; } = new StorageSettingTypes("ZoneRedundant");

        public static bool operator ==(StorageSettingTypes left, StorageSettingTypes right) => left.Equals(right);
        public static bool operator !=(StorageSettingTypes left, StorageSettingTypes right) => !left.Equals(right);

        public static explicit operator string(StorageSettingTypes value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StorageSettingTypes other && Equals(other);
        public bool Equals(StorageSettingTypes other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies the type of validation. In case of DeepValidation, all validations from /validateForBackup API will run again.
    /// </summary>
    [EnumType]
    public readonly struct ValidationType : IEquatable<ValidationType>
    {
        private readonly string _value;

        private ValidationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ValidationType ShallowValidation { get; } = new ValidationType("ShallowValidation");
        public static ValidationType DeepValidation { get; } = new ValidationType("DeepValidation");

        public static bool operator ==(ValidationType left, ValidationType right) => left.Equals(right);
        public static bool operator !=(ValidationType left, ValidationType right) => !left.Equals(right);

        public static explicit operator string(ValidationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ValidationType other && Equals(other);
        public bool Equals(ValidationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WeekNumber : IEquatable<WeekNumber>
    {
        private readonly string _value;

        private WeekNumber(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WeekNumber First { get; } = new WeekNumber("First");
        public static WeekNumber Fourth { get; } = new WeekNumber("Fourth");
        public static WeekNumber Last { get; } = new WeekNumber("Last");
        public static WeekNumber Second { get; } = new WeekNumber("Second");
        public static WeekNumber Third { get; } = new WeekNumber("Third");

        public static bool operator ==(WeekNumber left, WeekNumber right) => left.Equals(right);
        public static bool operator !=(WeekNumber left, WeekNumber right) => !left.Equals(right);

        public static explicit operator string(WeekNumber value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WeekNumber other && Equals(other);
        public bool Equals(WeekNumber other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
