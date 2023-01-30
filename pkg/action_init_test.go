package pkg

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"
)

func TestGeneratePrompt(t *testing.T) {

	t.Run("Default", func(t *testing.T) {
		w := &bytes.Buffer{}
		options := PromptOptions{
			PromptIntervalSeconds: 97,
			CliName:               "test",
			ThisPath:              "",
			PromptArguments:       "notify",
			Alias:                 "tt",
		}
		if err := GeneratePrompt(w, options); err != nil {
			t.Errorf("GeneratePrompt() error = %v", err)
			return
		}
		fmt.Println(w.String())		
	})

}
