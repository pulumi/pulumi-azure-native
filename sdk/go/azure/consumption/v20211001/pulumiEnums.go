


package v20211001

type BudgetOperatorType string

const (
	BudgetOperatorTypeIn = BudgetOperatorType("In")
)

type CategoryType string

const (
	CategoryTypeCost = CategoryType("Cost")
)

type CultureCode string

const (
	CultureCode_En_Us = CultureCode("en-us")
	CultureCode_Ja_Jp = CultureCode("ja-jp")
	CultureCode_Zh_Cn = CultureCode("zh-cn")
	CultureCode_De_De = CultureCode("de-de")
	CultureCode_Es_Es = CultureCode("es-es")
	CultureCode_Fr_Fr = CultureCode("fr-fr")
	CultureCode_It_It = CultureCode("it-it")
	CultureCode_Ko_Kr = CultureCode("ko-kr")
	CultureCode_Pt_Br = CultureCode("pt-br")
	CultureCode_Ru_Ru = CultureCode("ru-ru")
	CultureCode_Zh_Tw = CultureCode("zh-tw")
	CultureCode_Cs_Cz = CultureCode("cs-cz")
	CultureCode_Pl_Pl = CultureCode("pl-pl")
	CultureCode_Tr_Tr = CultureCode("tr-tr")
	CultureCode_Da_Dk = CultureCode("da-dk")
	CultureCode_En_Gb = CultureCode("en-gb")
	CultureCode_Hu_Hu = CultureCode("hu-hu")
	CultureCode_Nb_No = CultureCode("nb-no")
	CultureCode_Nl_Nl = CultureCode("nl-nl")
	CultureCode_Pt_Pt = CultureCode("pt-pt")
	CultureCode_Sv_Se = CultureCode("sv-se")
)

type OperatorType string

const (
	// Alert will be triggered if the evaluated cost is the same as threshold value. Note: It’s not recommended to use this OperatorType as there’s low chance of cost being exactly the same as threshold value, leading to missing of your alert. This OperatorType will be deprecated in future.
	OperatorTypeEqualTo = OperatorType("EqualTo")
	// Alert will be triggered if the evaluated cost is greater than the threshold value. Note: This is the recommended OperatorType while configuring Budget Alert.
	OperatorTypeGreaterThan = OperatorType("GreaterThan")
	// Alert will be triggered if the evaluated cost is greater than or equal to the threshold value.
	OperatorTypeGreaterThanOrEqualTo = OperatorType("GreaterThanOrEqualTo")
)

type ThresholdType string

const (
	// Actual costs budget alerts notify when the actual accrued cost exceeds the allocated budget .
	ThresholdTypeActual = ThresholdType("Actual")
	// Forecasted costs budget alerts provide advanced notification that your spending trends are likely to exceed your allocated budget, as it relies on forecasted cost predictions.
	ThresholdTypeForecasted = ThresholdType("Forecasted")
)

type TimeGrainType string

const (
	TimeGrainTypeMonthly        = TimeGrainType("Monthly")
	TimeGrainTypeQuarterly      = TimeGrainType("Quarterly")
	TimeGrainTypeAnnually       = TimeGrainType("Annually")
	TimeGrainTypeBillingMonth   = TimeGrainType("BillingMonth")
	TimeGrainTypeBillingQuarter = TimeGrainType("BillingQuarter")
	TimeGrainTypeBillingAnnual  = TimeGrainType("BillingAnnual")
)

func init() {
}
