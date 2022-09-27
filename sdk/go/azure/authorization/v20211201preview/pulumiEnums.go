


package v20211201preview

type AccessReviewRecurrencePatternType string

const (
	AccessReviewRecurrencePatternTypeWeekly          = AccessReviewRecurrencePatternType("weekly")
	AccessReviewRecurrencePatternTypeAbsoluteMonthly = AccessReviewRecurrencePatternType("absoluteMonthly")
)

type AccessReviewRecurrenceRangeType string

const (
	AccessReviewRecurrenceRangeTypeEndDate  = AccessReviewRecurrenceRangeType("endDate")
	AccessReviewRecurrenceRangeTypeNoEnd    = AccessReviewRecurrenceRangeType("noEnd")
	AccessReviewRecurrenceRangeTypeNumbered = AccessReviewRecurrenceRangeType("numbered")
)

type AccessReviewResult string

const (
	AccessReviewResultApprove     = AccessReviewResult("Approve")
	AccessReviewResultDeny        = AccessReviewResult("Deny")
	AccessReviewResultNotReviewed = AccessReviewResult("NotReviewed")
	AccessReviewResultDontKnow    = AccessReviewResult("DontKnow")
	AccessReviewResultNotNotified = AccessReviewResult("NotNotified")
)

type DefaultDecisionType string

const (
	DefaultDecisionTypeApprove        = DefaultDecisionType("Approve")
	DefaultDecisionTypeDeny           = DefaultDecisionType("Deny")
	DefaultDecisionTypeRecommendation = DefaultDecisionType("Recommendation")
)

func init() {
}
