<script setup lang="ts">
import { ref, reactive } from "vue";
import axiosInstance from "../utils/axiosInstance";
import { useToast } from "vue-toastification";
import sharedState from "../sharedState";

const toast = useToast();

const closeNewChatPage = () => {
    sharedState.isNewChatPageOpen = false;
};

const userInfo = reactive({
    username: "",
    profile_url: "",
});

const searchText = ref("");
const search = async () => {
    let formData = new FormData();
    formData.append("username", searchText.value);

    await axiosInstance
        .post("/users/search/", formData, {
            headers: {
                "Content-Type": "application/json",
            },
        })
        .then((resp) => {
            if (resp.status === 200) {
                userInfo["username"] = resp.data.username;
                userInfo["profile_url"] =
                    import.meta.env.VITE_BASE_BACKEND_URL +
                    resp.data.profile_url;
            }
        })
        .catch(() => {
            toast.error("Username doesnt exist");
        });
};

async function openChat() {
    sharedState.isNewChatPageOpen = true;
    sharedState.NewChatUsername = userInfo.username;
}
</script>
<template>
    <div
        class="flex flex-row border-b w-full h-[60px] border-b-gray-200 items-center pl-6 pt-1 gap-x-6"
    >
        <button
            @click="closeNewChatPage"
            class="w-[24px] h-[24px] cursor-pointer"
        >
            <svg
                viewBox="0 0 24 24"
                height="24"
                width="24"
                preserveAspectRatio="xMidYMid meet"
                class=""
                version="1.1"
                x="0px"
                y="0px"
                enable-background="new 0 0 24 24"
            >
                <title>back</title>
                <path
                    fill="currentColor"
                    d="M12,4l1.4,1.4L7.8,11H20v2H7.8l5.6,5.6L12,20l-8-8L12,4z"
                ></path>
            </svg>
        </button>
        <span class="text-[16px] font-normal">New chat</span>
    </div>
    <div
        class="flex flex-col w-full h-[72px] items-start bg-white pl-6 border-b border-gray-200"
    >
        <span class="text-[16px] font-normal text-green-700"
            >Search by Username</span
        >
        <div class="flex flex-row gap-x-4 items-center">
            <input
                v-model="searchText"
                type="text"
                placeholder="username"
                class="border-black rounded border-2 my-2 pl-1"
            />
            <button
                @click="search"
                class="cursor-pointer text-white bg-green-700 font-bold h-7 px-2 rounded"
            >
                Search
            </button>
        </div>
    </div>
    <div class="select-none flex flex-col w-full">
        <div
            @click="openChat"
            v-if="userInfo.profile_url"
            class="flex border-b border-gray-200 flex-row w-full h-[72px] cursor-pointer justify-start items-center hover:bg-[#f5f6f6] pl-3"
        >
            <img
                class="w-[49px] h-[49px] rounded-full object-cover"
                :src="userInfo.profile_url"
                alt=""
            />
            <div class="flex flex-col ml-4 w-full">
                <div class="flex flex-row w-full">
                    <p class="font-normal text-[17px]">
                        {{ userInfo.username }}
                    </p>
                    <!-- <span class="ml-auto mr-3 font-normal text-[12px] text-gray-500">Yesterday</span> -->
                </div>
                <!-- <p class="font-normal text-[14px] text-gray-600">Your last chat...</p> -->
            </div>
        </div>
    </div>
</template>
