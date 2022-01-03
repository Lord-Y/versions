import { createRouter, createWebHistory } from 'vue-router'
import contents from './contents'

const router = createRouter({
  history: createWebHistory('/ui/'),
  routes: [...contents],
})

export default router
