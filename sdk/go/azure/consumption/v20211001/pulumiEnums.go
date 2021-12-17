


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
	CultureCode_En_us = CultureCode("en-us")
	CultureCode_Ja_jp = CultureCode("ja-jp")
	CultureCode_Zh_cn = CultureCode("zh-cn")
	CultureCode_De_de = CultureCode("de-de")
	CultureCode_Es_es = CultureCode("es-es")
	CultureCode_Fr_fr = CultureCode("fr-fr")
	CultureCode_It_it = CultureCode("it-it")
	CultureCode_Ko_kr = CultureCode("ko-kr")
	CultureCode_Pt_br = CultureCode("pt-br")
	CultureCode_Ru_ru = CultureCode("ru-ru")
	CultureCode_Zh_tw = CultureCode("zh-tw")
	CultureCode_Cs_cz = CultureCode("cs-cz")
	CultureCode_Pl_pl = CultureCode("pl-pl")
	CultureCode_Tr_tr = CultureCode("tr-tr")
	CultureCode_Da_dk = CultureCode("da-dk")
	CultureCode_En_gb = CultureCode("en-gb")
	CultureCode_Hu_hu = CultureCode("hu-hu")
	CultureCode_Nb_no = CultureCode("nb-no")
	CultureCode_Nl_nl = CultureCode("nl-nl")
	CultureCode_Pt_pt = CultureCode("pt-pt")
	CultureCode_Sv_se = CultureCode("sv-se")
)

type OperatorType string

const (
	OperatorTypeEqualTo              = OperatorType("EqualTo")
	OperatorTypeGreaterThan          = OperatorType("GreaterThan")
	OperatorTypeGreaterThanOrEqualTo = OperatorType("GreaterThanOrEqualTo")
)

type ThresholdType string

const (
	ThresholdTypeActual     = ThresholdType("Actual")
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
