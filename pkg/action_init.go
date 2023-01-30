package pkg

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

//go:embed action_init.sh
var actionInitSh string

type (
	PromptOptions struct {
		PromptIntervalSeconds int
		CliName               string
		ThisPath              string
		PromptArguments       string
		Alias                 string
	}
)

func GeneratePrompt(w io.Writer, options PromptOptions) error {
	prefix := RandomString(6)
	prompt_command := fmt.Sprintf("__%s_%s", prefix, CreateSlug(options.CliName))
	timer_var := strings.ToUpper(fmt.Sprintf("__%s_timer", prefix))
	this_path := options.ThisPath
	if len(this_path) == 0 {
		this_path = os.Args[0]
	}
	alias := ""
	message := fmt.Sprintf("%s (%s)", options.CliName, path.Base(this_path))
	if len(options.Alias) > 0 {
		alias = fmt.Sprintf("alias %s=%s", options.Alias, this_path)
		message += fmt.Sprintf(" (alias %s)", options.Alias)
	}
	message += fmt.Sprintf(" will run every %v", time.Duration(options.PromptIntervalSeconds)*time.Second)
	replacer := strings.NewReplacer(
		"#PERIOD#", fmt.Sprintf("%d", options.PromptIntervalSeconds),
		"#CMD#", prompt_command,
		"#ARGS#", options.PromptArguments,
		"#THIS_PATH#", this_path,
		"#TIMER_VAR#", timer_var,
		"#ALIAS#", alias,
		"#MESSAGE#", message,
	)
	_, err := replacer.WriteString(w, actionInitSh)

	return err
}
