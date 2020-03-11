package utils

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

var localizerEn *i18n.Localizer
var localizerFa *i18n.Localizer

func InitI18n() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, errEn := bundle.LoadMessageFile("locales/en.json")
	_, errFa := bundle.LoadMessageFile("locales/fa.json")
	if errEn != nil || errFa != nil {
		log.Fatalf(fmt.Sprintf("Error fa => %v Error en => %v", errFa.Error(), errEn.Error()))
	}
	localizerEn = i18n.NewLocalizer(bundle, "en", "fa")
	localizerFa = i18n.NewLocalizer(bundle, "fa")
}

func Translate(messageId string, data map[string]string, lang string) string {
	switch lang {
	case "fa":
		return localizerFa.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    messageId,
			TemplateData: data,
		})
	case "en":
		return localizerEn.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    messageId,
			TemplateData: data,
		})
	default:
		return localizerEn.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    messageId,
			TemplateData: data,
		})

	}

}
