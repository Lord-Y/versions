import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createHead } from '@vueuse/head'
import { createI18n } from './i18n'
import './tailwind.css'

const i18n = createI18n()

const app = createApp(App)
const head = createHead()
app.use(router)
app.use(i18n)
app.use(head)
app.mount('#app')
