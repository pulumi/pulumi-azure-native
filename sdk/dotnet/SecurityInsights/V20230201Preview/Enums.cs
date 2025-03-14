// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.SecurityInsights.V20230201Preview
{
    /// <summary>
    /// The entity query kind
    /// </summary>
    [EnumType]
    public readonly struct EntityTimelineKind : IEquatable<EntityTimelineKind>
    {
        private readonly string _value;

        private EntityTimelineKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// activity
        /// </summary>
        public static EntityTimelineKind Activity { get; } = new EntityTimelineKind("Activity");
        /// <summary>
        /// bookmarks
        /// </summary>
        public static EntityTimelineKind Bookmark { get; } = new EntityTimelineKind("Bookmark");
        /// <summary>
        /// security alerts
        /// </summary>
        public static EntityTimelineKind SecurityAlert { get; } = new EntityTimelineKind("SecurityAlert");
        /// <summary>
        /// anomaly
        /// </summary>
        public static EntityTimelineKind Anomaly { get; } = new EntityTimelineKind("Anomaly");

        public static bool operator ==(EntityTimelineKind left, EntityTimelineKind right) => left.Equals(right);
        public static bool operator !=(EntityTimelineKind left, EntityTimelineKind right) => !left.Equals(right);

        public static explicit operator string(EntityTimelineKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityTimelineKind other && Equals(other);
        public bool Equals(EntityTimelineKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The reason the incident was closed
    /// </summary>
    [EnumType]
    public readonly struct IncidentClassification : IEquatable<IncidentClassification>
    {
        private readonly string _value;

        private IncidentClassification(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Incident classification was undetermined
        /// </summary>
        public static IncidentClassification Undetermined { get; } = new IncidentClassification("Undetermined");
        /// <summary>
        /// Incident was true positive
        /// </summary>
        public static IncidentClassification TruePositive { get; } = new IncidentClassification("TruePositive");
        /// <summary>
        /// Incident was benign positive
        /// </summary>
        public static IncidentClassification BenignPositive { get; } = new IncidentClassification("BenignPositive");
        /// <summary>
        /// Incident was false positive
        /// </summary>
        public static IncidentClassification FalsePositive { get; } = new IncidentClassification("FalsePositive");

        public static bool operator ==(IncidentClassification left, IncidentClassification right) => left.Equals(right);
        public static bool operator !=(IncidentClassification left, IncidentClassification right) => !left.Equals(right);

        public static explicit operator string(IncidentClassification value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IncidentClassification other && Equals(other);
        public bool Equals(IncidentClassification other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The classification reason the incident was closed with
    /// </summary>
    [EnumType]
    public readonly struct IncidentClassificationReason : IEquatable<IncidentClassificationReason>
    {
        private readonly string _value;

        private IncidentClassificationReason(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Classification reason was suspicious activity
        /// </summary>
        public static IncidentClassificationReason SuspiciousActivity { get; } = new IncidentClassificationReason("SuspiciousActivity");
        /// <summary>
        /// Classification reason was suspicious but expected
        /// </summary>
        public static IncidentClassificationReason SuspiciousButExpected { get; } = new IncidentClassificationReason("SuspiciousButExpected");
        /// <summary>
        /// Classification reason was incorrect alert logic
        /// </summary>
        public static IncidentClassificationReason IncorrectAlertLogic { get; } = new IncidentClassificationReason("IncorrectAlertLogic");
        /// <summary>
        /// Classification reason was inaccurate data
        /// </summary>
        public static IncidentClassificationReason InaccurateData { get; } = new IncidentClassificationReason("InaccurateData");

        public static bool operator ==(IncidentClassificationReason left, IncidentClassificationReason right) => left.Equals(right);
        public static bool operator !=(IncidentClassificationReason left, IncidentClassificationReason right) => !left.Equals(right);

        public static explicit operator string(IncidentClassificationReason value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IncidentClassificationReason other && Equals(other);
        public bool Equals(IncidentClassificationReason other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The severity of the incident
    /// </summary>
    [EnumType]
    public readonly struct IncidentSeverity : IEquatable<IncidentSeverity>
    {
        private readonly string _value;

        private IncidentSeverity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// High severity
        /// </summary>
        public static IncidentSeverity High { get; } = new IncidentSeverity("High");
        /// <summary>
        /// Medium severity
        /// </summary>
        public static IncidentSeverity Medium { get; } = new IncidentSeverity("Medium");
        /// <summary>
        /// Low severity
        /// </summary>
        public static IncidentSeverity Low { get; } = new IncidentSeverity("Low");
        /// <summary>
        /// Informational severity
        /// </summary>
        public static IncidentSeverity Informational { get; } = new IncidentSeverity("Informational");

        public static bool operator ==(IncidentSeverity left, IncidentSeverity right) => left.Equals(right);
        public static bool operator !=(IncidentSeverity left, IncidentSeverity right) => !left.Equals(right);

        public static explicit operator string(IncidentSeverity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IncidentSeverity other && Equals(other);
        public bool Equals(IncidentSeverity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the incident
    /// </summary>
    [EnumType]
    public readonly struct IncidentStatus : IEquatable<IncidentStatus>
    {
        private readonly string _value;

        private IncidentStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// An active incident which isn't being handled currently
        /// </summary>
        public static IncidentStatus New { get; } = new IncidentStatus("New");
        /// <summary>
        /// An active incident which is being handled
        /// </summary>
        public static IncidentStatus Active { get; } = new IncidentStatus("Active");
        /// <summary>
        /// A non-active incident
        /// </summary>
        public static IncidentStatus Closed { get; } = new IncidentStatus("Closed");

        public static bool operator ==(IncidentStatus left, IncidentStatus right) => left.Equals(right);
        public static bool operator !=(IncidentStatus left, IncidentStatus right) => !left.Equals(right);

        public static explicit operator string(IncidentStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IncidentStatus other && Equals(other);
        public bool Equals(IncidentStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The kind of content the metadata is for.
    /// </summary>
    [EnumType]
    public readonly struct Kind : IEquatable<Kind>
    {
        private readonly string _value;

        private Kind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Kind DataConnector { get; } = new Kind("DataConnector");
        public static Kind DataType { get; } = new Kind("DataType");
        public static Kind Workbook { get; } = new Kind("Workbook");
        public static Kind WorkbookTemplate { get; } = new Kind("WorkbookTemplate");
        public static Kind Playbook { get; } = new Kind("Playbook");
        public static Kind PlaybookTemplate { get; } = new Kind("PlaybookTemplate");
        public static Kind AnalyticsRuleTemplate { get; } = new Kind("AnalyticsRuleTemplate");
        public static Kind AnalyticsRule { get; } = new Kind("AnalyticsRule");
        public static Kind HuntingQuery { get; } = new Kind("HuntingQuery");
        public static Kind InvestigationQuery { get; } = new Kind("InvestigationQuery");
        public static Kind Parser { get; } = new Kind("Parser");
        public static Kind Watchlist { get; } = new Kind("Watchlist");
        public static Kind WatchlistTemplate { get; } = new Kind("WatchlistTemplate");
        public static Kind Solution { get; } = new Kind("Solution");
        public static Kind AzureFunction { get; } = new Kind("AzureFunction");
        public static Kind LogicAppsCustomConnector { get; } = new Kind("LogicAppsCustomConnector");
        public static Kind AutomationRule { get; } = new Kind("AutomationRule");

        public static bool operator ==(Kind left, Kind right) => left.Equals(right);
        public static bool operator !=(Kind left, Kind right) => !left.Equals(right);

        public static explicit operator string(Kind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Kind other && Equals(other);
        public bool Equals(Kind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Operator used for list of dependencies in criteria array.
    /// </summary>
    [EnumType]
    public readonly struct Operator : IEquatable<Operator>
    {
        private readonly string _value;

        private Operator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Operator AND { get; } = new Operator("AND");
        public static Operator OR { get; } = new Operator("OR");

        public static bool operator ==(Operator left, Operator right) => left.Equals(right);
        public static bool operator !=(Operator left, Operator right) => !left.Equals(right);

        public static explicit operator string(Operator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Operator other && Equals(other);
        public bool Equals(Operator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the owner the incident is assigned to.
    /// </summary>
    [EnumType]
    public readonly struct OwnerType : IEquatable<OwnerType>
    {
        private readonly string _value;

        private OwnerType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The incident owner type is unknown
        /// </summary>
        public static OwnerType Unknown { get; } = new OwnerType("Unknown");
        /// <summary>
        /// The incident owner type is an AAD user
        /// </summary>
        public static OwnerType User { get; } = new OwnerType("User");
        /// <summary>
        /// The incident owner type is an AAD group
        /// </summary>
        public static OwnerType Group { get; } = new OwnerType("Group");

        public static bool operator ==(OwnerType left, OwnerType right) => left.Equals(right);
        public static bool operator !=(OwnerType left, OwnerType right) => !left.Equals(right);

        public static explicit operator string(OwnerType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OwnerType other && Equals(other);
        public bool Equals(OwnerType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Source type of the content
    /// </summary>
    [EnumType]
    public readonly struct SourceKind : IEquatable<SourceKind>
    {
        private readonly string _value;

        private SourceKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SourceKind LocalWorkspace { get; } = new SourceKind("LocalWorkspace");
        public static SourceKind Community { get; } = new SourceKind("Community");
        public static SourceKind Solution { get; } = new SourceKind("Solution");
        public static SourceKind SourceRepository { get; } = new SourceKind("SourceRepository");

        public static bool operator ==(SourceKind left, SourceKind right) => left.Equals(right);
        public static bool operator !=(SourceKind left, SourceKind right) => !left.Equals(right);

        public static explicit operator string(SourceKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SourceKind other && Equals(other);
        public bool Equals(SourceKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of support for content item
    /// </summary>
    [EnumType]
    public readonly struct SupportTier : IEquatable<SupportTier>
    {
        private readonly string _value;

        private SupportTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SupportTier Microsoft { get; } = new SupportTier("Microsoft");
        public static SupportTier Partner { get; } = new SupportTier("Partner");
        public static SupportTier Community { get; } = new SupportTier("Community");

        public static bool operator ==(SupportTier left, SupportTier right) => left.Equals(right);
        public static bool operator !=(SupportTier left, SupportTier right) => !left.Equals(right);

        public static explicit operator string(SupportTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SupportTier other && Equals(other);
        public bool Equals(SupportTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
