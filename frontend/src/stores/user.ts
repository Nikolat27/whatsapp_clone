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
        isAuthenticated: false,
    }),
    actions: {
        async checkAuth(): Promise<boolean> {
            try {
                const resp = await axiosInstance.get("/auth/check");
                this.isAuthenticated = resp.status === 200;
            } catch (err: any) {
                this.isAuthenticated = false;
            }
            return this.isAuthenticated;
        },
        async fetchUser() {
            await axiosInstance
                .post("/users/")
                .then((res) => {
                    if (res.status !== 200) {
                        return;
                    }

                    this.user = {
                        user_id: res.data.user_id,
                        username: res.data.username,
                        name: res.data.name,
                        about: res.data.about,
                        profile_url: res.data.profile_url,
                    };

                    this.user.profile_url =
                        import.meta.env.VITE_BASE_BACKEND_URL +
                        this.user.profile_url;
                })
                .catch((err) => {
                    console.error(err);
                });
        },
    },
});
