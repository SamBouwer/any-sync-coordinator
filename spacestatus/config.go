package spacestatus

type Config struct {
	RunSeconds         int `yaml:"runSeconds"`
	DeletionPeriodDays int `yaml:"deletionPeriodDays"`
}
