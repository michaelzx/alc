package alc_i18n

func Msg(locale Locale, keyPath string) *I18n {
	return &I18n{
		// 建议不要超过3级
		keyPath: keyPath,
		locale:  locale,
	}
}
