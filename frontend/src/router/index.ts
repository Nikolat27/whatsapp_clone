import { createWebHistory, createRouter } from 'vue-router'

import ChatsView from '../views/ChatsView.vue'

const routes: any = [
  { path: "/",
    component: ChatsView
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router