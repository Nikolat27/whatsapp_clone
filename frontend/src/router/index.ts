import { createWebHistory, createRouter } from 'vue-router'

import ChatsView from '../views/ChatsView.vue'
import StatusView from '../views/StatusView.vue'

const routes: any = [
  {
    path: "/",
    component: ChatsView
  },
  {
    path: "/status",
    component: StatusView
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router