<script setup lang="ts">
import { onMounted, ref } from "vue";

import axiosInstance from "./utils/axiosInstance";

import SideBarComponent from "./components/SideBarComponent.vue";
import ChannelPageComponent from "./components/ChannelPageComponent.vue";

const isUserAuthenticated = ref(false);

function checkAuth() {
    axiosInstance.get("/auth/check").then((resp) => {
        if (resp.status === 200) {
            isUserAuthenticated.value = true;
        } else {
            isUserAuthenticated.value = false;
        }
    });
}

onMounted(() => {
    checkAuth();
});
</script>
<template>
    <div class="w-[83.3%] h-full absolute top-[28px] left-[110px]">
        <div class="flex flex-row w-full h-[94%] custom-shadow-inset">
            <div
                class="flex flex-col flex-basis-[58.8px] flex-shrink-0 min-w-[58.8px] w-[4%] h-full bg-[#f0f2f5] justify-start items-center pt-2 gap-y-4 border-r border-gray-300"
            >
                <SideBarComponent
                    :isUserAuthenticated="isUserAuthenticated"
                ></SideBarComponent>
            </div>
            <div class="flex flex-col max-w-[30%] w-[30%] h-full bg-white">
                <router-view></router-view>
            </div>
            <div class="flex flex-col w-[66%] h-full relative">
                <!-- <ChatPageComponent></ChatPageComponent> -->
                <ChannelPageComponent></ChannelPageComponent>
            </div>
        </div>
    </div>
</template>
