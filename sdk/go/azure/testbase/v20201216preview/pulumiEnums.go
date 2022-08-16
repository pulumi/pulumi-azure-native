


package v20201216preview

type Action string

const (
	ActionInstall   = Action("Install")
	ActionLaunch    = Action("Launch")
	ActionClose     = Action("Close")
	ActionUninstall = Action("Uninstall")
	ActionCustom    = Action("Custom")
)

type ContentType string

const (
	ContentTypeInline = ContentType("Inline")
	ContentTypeFile   = ContentType("File")
	ContentTypePath   = ContentType("Path")
)

type TestType string

const (
	TestTypeOutOfBoxTest   = TestType("OutOfBoxTest")
	TestTypeFunctionalTest = TestType("FunctionalTest")
)

type Tier string

const (
	TierStandard = Tier("Standard")
)

func init() {
}
