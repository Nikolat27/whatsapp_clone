import { defineStore } from "pinia";
import axiosInstance from "../utils/axiosInstance";

export const useUserStore = defineStore("user", {
    state: () => ({
        user: null as null | {
            user_id: string;
            username: string;
            name: string;
            about: string;
            profile_url: string;
        },
    }),
    actions: {
        async fetchUser() {
            const res = await axiosInstance.post("/users/");
            this.user = {
                user_id: res.data.user_id,
                username: res.data.username,
                name: res.data.name,
                about: res.data.about,
                profile_url: res.data.profile_url,
            };

            this.user.profile_url = import.meta.env.VITE_BASE_BACKEND_URL + this.user.profile_url
        },
    },
});
