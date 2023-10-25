import { createI18n as _createI18n } from 'vue-i18n'
import en from './locales/en-US.json'

export function createI18n() {
  return _createI18n({
    legacy: false,
    globalInjection: true,
    locale: 'en',
    fallbackLocale: 'en',
    messages: {
      en,
    }
  })
}