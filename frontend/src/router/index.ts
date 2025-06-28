import { createWebHistory, createRouter } from "vue-router";

import ChatsView from "../views/ChatsView.vue";
import StatusView from "../views/StatusView.vue";
import UserChannelsView from "../views/UserChannelsView.vue";
import SettingsView from "../views/SettingsView.vue";
import UserEditProfileComponent from "../components/UserEditProfileComponent.vue";
import AuthView from "../views/AuthView.vue";

const routes: any = [
    {
        path: "/",
        component: ChatsView,
        name: "home",
    },
    {
        path: "/status",
        component: StatusView,
        name: "status",
    },
    {
        path: "/channels",
        component: UserChannelsView,
        name: "channels",
    },
    {
        path: "/setting",
        component: SettingsView,
        name: "setting",
    },
    {
        path: "/profile",
        component: UserEditProfileComponent,
        name: "profile",
    },
    {
        path: "/auth",
        component: AuthView,
        name: "auth",
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
