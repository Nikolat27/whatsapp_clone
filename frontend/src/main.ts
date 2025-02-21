import { createApp } from 'vue'
import App from './App.vue'
import './assets/css/style.css'
import router from './router/index.ts'

createApp(App)

const app = createApp(App)

app.use(router)
app.mount("#app")
