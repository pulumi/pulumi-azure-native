// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.LabServices
{
    /// <summary>
    /// The enabled access level for Web Access over SSH.
    /// </summary>
    [EnumType]
    public readonly struct ConnectionType : IEquatable<ConnectionType>
    {
        private readonly string _value;

        private ConnectionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConnectionType Public { get; } = new ConnectionType("Public");
        public static ConnectionType Private { get; } = new ConnectionType("Private");
        public static ConnectionType None { get; } = new ConnectionType("None");

        public static bool operator ==(ConnectionType left, ConnectionType right) => left.Equals(right);
        public static bool operator !=(ConnectionType left, ConnectionType right) => !left.Equals(right);

        public static explicit operator string(ConnectionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConnectionType other && Equals(other);
        public bool Equals(ConnectionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates what lab virtual machines are created from.
    /// </summary>
    [EnumType]
    public readonly struct CreateOption : IEquatable<CreateOption>
    {
        private readonly string _value;

        private CreateOption(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// An image is used to create all lab user virtual machines. When this option is set, no template VM will be created.
        /// </summary>
        public static CreateOption Image { get; } = new CreateOption("Image");
        /// <summary>
        /// A template VM will be used to create all lab user virtual machines.
        /// </summary>
        public static CreateOption TemplateVM { get; } = new CreateOption("TemplateVM");

        public static bool operator ==(CreateOption left, CreateOption right) => left.Equals(right);
        public static bool operator !=(CreateOption left, CreateOption right) => !left.Equals(right);

        public static explicit operator string(CreateOption value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CreateOption other && Equals(other);
        public bool Equals(CreateOption other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Whether a VM will get shutdown when it hasn't been connected to after a period of time.
    /// </summary>
    [EnumType]
    public readonly struct EnableState : IEquatable<EnableState>
    {
        private readonly string _value;

        private EnableState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnableState Enabled { get; } = new EnableState("Enabled");
        public static EnableState Disabled { get; } = new EnableState("Disabled");

        public static bool operator ==(EnableState left, EnableState right) => left.Equals(right);
        public static bool operator !=(EnableState left, EnableState right) => !left.Equals(right);

        public static explicit operator string(EnableState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnableState other && Equals(other);
        public bool Equals(EnableState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The frequency of the recurrence.
    /// </summary>
    [EnumType]
    public readonly struct RecurrenceFrequency : IEquatable<RecurrenceFrequency>
    {
        private readonly string _value;

        private RecurrenceFrequency(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Schedule will run every days.
        /// </summary>
        public static RecurrenceFrequency Daily { get; } = new RecurrenceFrequency("Daily");
        /// <summary>
        /// Schedule will run every week on days specified in weekDays.
        /// </summary>
        public static RecurrenceFrequency Weekly { get; } = new RecurrenceFrequency("Weekly");

        public static bool operator ==(RecurrenceFrequency left, RecurrenceFrequency right) => left.Equals(right);
        public static bool operator !=(RecurrenceFrequency left, RecurrenceFrequency right) => !left.Equals(right);

        public static explicit operator string(RecurrenceFrequency value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RecurrenceFrequency other && Equals(other);
        public bool Equals(RecurrenceFrequency other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The identity type.
    /// </summary>
    [EnumType]
    public readonly struct ResourceIdentityType : IEquatable<ResourceIdentityType>
    {
        private readonly string _value;

        private ResourceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ResourceIdentityType SystemAssigned { get; } = new ResourceIdentityType("SystemAssigned");

        public static bool operator ==(ResourceIdentityType left, ResourceIdentityType right) => left.Equals(right);
        public static bool operator !=(ResourceIdentityType left, ResourceIdentityType right) => !left.Equals(right);

        public static explicit operator string(ResourceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourceIdentityType other && Equals(other);
        public bool Equals(ResourceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Whether a VM will get shutdown when it has idled for a period of time.
    /// </summary>
    [EnumType]
    public readonly struct ShutdownOnIdleMode : IEquatable<ShutdownOnIdleMode>
    {
        private readonly string _value;

        private ShutdownOnIdleMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The VM won't be shut down when it is idle.
        /// </summary>
        public static ShutdownOnIdleMode None { get; } = new ShutdownOnIdleMode("None");
        /// <summary>
        /// The VM will be considered as idle when there is no keyboard or mouse input.
        /// </summary>
        public static ShutdownOnIdleMode UserAbsence { get; } = new ShutdownOnIdleMode("UserAbsence");
        /// <summary>
        /// The VM will be considered as idle when user is absent and the resource (CPU and disk) consumption is low.
        /// </summary>
        public static ShutdownOnIdleMode LowUsage { get; } = new ShutdownOnIdleMode("LowUsage");

        public static bool operator ==(ShutdownOnIdleMode left, ShutdownOnIdleMode right) => left.Equals(right);
        public static bool operator !=(ShutdownOnIdleMode left, ShutdownOnIdleMode right) => !left.Equals(right);

        public static explicit operator string(ShutdownOnIdleMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ShutdownOnIdleMode other && Equals(other);
        public bool Equals(ShutdownOnIdleMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

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
    /// Days of the week.
    /// </summary>
    [EnumType]
    public readonly struct WeekDay : IEquatable<WeekDay>
    {
        private readonly string _value;

        private WeekDay(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Schedule will run on Sunday
        /// </summary>
        public static WeekDay Sunday { get; } = new WeekDay("Sunday");
        /// <summary>
        /// Schedule will run on Monday
        /// </summary>
        public static WeekDay Monday { get; } = new WeekDay("Monday");
        /// <summary>
        /// Schedule will run on Tuesday
        /// </summary>
        public static WeekDay Tuesday { get; } = new WeekDay("Tuesday");
        /// <summary>
        /// Schedule will run on Wednesday
        /// </summary>
        public static WeekDay Wednesday { get; } = new WeekDay("Wednesday");
        /// <summary>
        /// Schedule will run on Thursday
        /// </summary>
        public static WeekDay Thursday { get; } = new WeekDay("Thursday");
        /// <summary>
        /// Schedule will run on Friday
        /// </summary>
        public static WeekDay Friday { get; } = new WeekDay("Friday");
        /// <summary>
        /// Schedule will run on Saturday
        /// </summary>
        public static WeekDay Saturday { get; } = new WeekDay("Saturday");

        public static bool operator ==(WeekDay left, WeekDay right) => left.Equals(right);
        public static bool operator !=(WeekDay left, WeekDay right) => !left.Equals(right);

        public static explicit operator string(WeekDay value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WeekDay other && Equals(other);
        public bool Equals(WeekDay other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
