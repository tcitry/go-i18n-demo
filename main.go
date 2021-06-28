package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("active.zh.toml")
	localizer := i18n.NewLocalizer(bundle, "zh-CN")
	localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "HelloPerson",
			Other: "Hello {{.Name}}",
		},
	})

	cn := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "HelloPerson", TemplateData: map[string]string{"Name": "Nick"}})
	fmt.Println("Chinese: ", cn)

	enlocalizer := i18n.NewLocalizer(bundle, "en-US")
	en := enlocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "HelloPerson", TemplateData: map[string]string{"Name": "Nick"}})
	fmt.Println("English: ", en)

	// output:
	// Chinese:  你好 Nick
	// English:  Hello Nick
}
