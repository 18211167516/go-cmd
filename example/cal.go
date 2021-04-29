package main

import (
	"log"

	cmd "github.com/18211167516/go-cmd"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

type GetNum struct {
	Num decimal.Decimal
	Opt string
}

func isSlash(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}

// calCmd represents the version command
var calCmd = &cmd.Command{
	Use:   "cal",
	Short: "计算器",
	Long:  `简单计算器`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(Command *cmd.Command, args []string) {
		s, _ := decimal.NewFromString(args[0])
		nu := &GetNum{
			Num: s,
			Opt: args[1],
		}

		for i := 2; i < len(args); i++ {
			nu.opt(args[i])
		}

		log.Println(nu.Num)

	},
}

func (nu *GetNum) opt(n string) {

	switch nu.Opt {
	case "+":
		news, _ := decimal.NewFromString(n)
		log.Println(news)
		nu.Num = nu.Num.Add(news)
		nu.Opt = ""
	case "-":
		news, _ := decimal.NewFromString(n)
		nu.Num = nu.Num.Sub(news)
		nu.Opt = ""
	case "x":
		news, _ := decimal.NewFromString(n)
		nu.Num = nu.Num.Mul(news)
		nu.Opt = ""
	case "%":
		news, _ := decimal.NewFromString(n)
		nu.Num = nu.Num.Div(news)
		nu.Opt = ""
	default:
		nu.Opt = gconv.String(n)
	}
}

func init() {
	cmd.RootCmd.AddCommand(calCmd)
}
