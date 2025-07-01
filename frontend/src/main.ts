import { createApp } from "vue";
import App from "./App.vue";
import { createPinia } from "pinia";
import "./assets/css/style.css";
import router from "./router/index";
import Toast from "vue-toastification";
import "../node_modules/vue3-emoji-picker/dist/style.css";
import "vue-toastification/dist/index.css";
import { useUserStore } from "./stores/user";

const app = createApp(App);

app.use(createPinia());

const userStore = useUserStore();
await userStore.fetchUser();

app.use(router);
app.use(Toast, {
    position: "top-right",
    timeout: 2000,
    closeOnClick: true,
    pauseOnHover: true,
});
app.mount("#app");
