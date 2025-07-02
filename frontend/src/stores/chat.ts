import { defineStore } from "pinia";

export const useChatStore = defineStore("chat", {
    state: () => ({
        chat: null as null | {
            username: string;
            openChat: boolean;
            openSaveMessage: boolean;
        },
    }),
    actions: {
        async openChat(username: string) {
            this.chat = {
                username: username,
                openChat: true,
                openSaveMessage: false,
            };
        },
        async openSaveMessage() {
            this.chat = {
                openSaveMessage: true,
                openChat: false,
            };
        },
    },
});
