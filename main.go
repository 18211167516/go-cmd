package gocmd

import (
	"os"

	"github.com/gookit/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type Command = cobra.Command

var RootCmd = &Command{
	Use:   "gocmd",
	Short: "这是一款简单神奇易用的终端工具",
	Long: `欢迎使用gcmd终端工具
  ________       ___________     ___________
 / ——————       //\\      //\\     ||\\
/ /            //  \\	 //  \\    ||   \\
\ \           //    \\  //    \\   ||     \\
 \ \______    \\     \\//     //   ||    //
  \——————      \\             //   ||___//
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		color.Warn.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	//全局持久化config
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			color.Warn.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gva" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		color.Warn.Println("Using config file:", viper.ConfigFileUsed())
	}
}
