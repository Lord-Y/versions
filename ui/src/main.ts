import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createHead } from '@vueuse/head'
import { createI18n } from 'vue-i18n'
import messages from '@intlify/vite-plugin-vue-i18n/messages'
import './tailwind.css'
const i18n = createI18n({
  legacy: false,
  locale: 'en-US',
  fallbackLocale: 'en-US',
  globalInjection: true,
  messages,
})

const app = createApp(App)
const head = createHead()
app.use(router)
app.use(i18n)
app.use(head)
app.mount('#app')
