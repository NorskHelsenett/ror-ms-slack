package slackconfig

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

var (
	Version string = "1.1.0"
	Commit  string = "FFFFF"
)

func init() {
	viper.AutomaticEnv()

	viper.SetDefault(configconsts.ENVIRONMENT, "production")
	viper.SetDefault(rlog.LOG_LEVEL, "info")
	viper.SetDefault(configconsts.ROLE, "ror-ms-slack")
	viper.SetDefault(configconsts.VERSION, Version)
	viper.SetDefault(configconsts.COMMIT, Commit)
}

func Load() {
	environment := viper.GetString(configconsts.ENVIRONMENT)
	rlog.Info("loaded environment", rlog.String("Environment", environment))

	_ = viper.WriteConfig()
}

func GetRorVersion() rorversion.RorVersion {
	return rorversion.NewRorVersion(viper.GetString(configconsts.VERSION), viper.GetString(configconsts.COMMIT))
}
