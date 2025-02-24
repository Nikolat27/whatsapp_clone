import { createApp } from 'vue'
import App from './App.vue'
import './assets/css/style.css'
import router from './router/index'
import Toast from 'vue-toastification'
import '../node_modules/vue3-emoji-picker/dist/style.css'
import "vue-toastification/dist/index.css";

const app = createApp(App)

app.use(router)
app.use(Toast, {
    position: "top-right",
    timeout: 2000,
    closeOnClick: true,
    pauseOnHover: true,
});
app.mount("#app")
