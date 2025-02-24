<script setup lang="ts">
import { ref } from 'vue';
import { Ref } from 'vue';

const isChannelOptionsDivOpen: Ref<boolean> = ref(false)
const toggleChannelOptionsDiv = () => isChannelOptionsDivOpen.value = !isChannelOptionsDivOpen.value
</script>
<template>
    <div class="custom-shadow-inset4 cursor-pointer select-none flex flex-row w-full h-[64px] bg-[#f0f2f5]
     justify-start items-center pl-4 top-0 sticky">
        <img class="w-10 h-10 rounded-full" src="../../barcelona-logo.jpg" alt="">
        <div class="flex flex-col ml-3">
            <p class="text-[16px] font-normal">FC Barcelona</p>
            <span class="text-[13px] font-normal text-gray-500">43.5M followers</span>
        </div>
        <div class="relative flex flex-row ml-auto mr-6 gap-x-3">
            <button @click="toggleChannelOptionsDiv"
                class="flex justify-center items-center w-10 h-10 rounded-full cursor-pointer">
                <span aria-hidden="true" data-icon="menu" class="xr9ek0c"><svg viewBox="0 0 24 24" height="24"
                        width="24" preserveAspectRatio="xMidYMid meet" class="" version="1.1" x="0px" y="0px"
                        enable-background="new 0 0 24 24">
                        <title>menu</title>
                        <path fill="currentColor"
                            d="M12,7c1.104,0,2-0.896,2-2c0-1.105-0.895-2-2-2c-1.104,0-2,0.894-2,2 C10,6.105,10.895,7,12,7z M12,9c-1.104,0-2,0.894-2,2c0,1.104,0.895,2,2,2c1.104,0,2-0.896,2-2C13.999,9.895,13.104,9,12,9z M12,15 c-1.104,0-2,0.894-2,2c0,1.104,0.895,2,2,2c1.104,0,2-0.896,2-2C13.999,15.894,13.104,15,12,15z">
                        </path>
                    </svg></span>
            </button>
            <div v-if="isChannelOptionsDivOpen"
                class="absolute top-10 right-0 w-[182px] h-auto flex flex-col bg-white z-30 custom-shadow-inset">
                <div class="flex items-center h-10 w-full hover:bg-[#f5f6f6] text-[14px] font-medium pl-4">
                    Unfollow
                </div>
            </div>
        </div>
    </div>
    <div class="w-full h-full gap-y-6 px-16 py-12 flex flex-col items-center overflow-auto scroll-smooth"
        style="background-image: url('../../whatsapp_default_background.jpg'); scrollbar-width: thin;">
        <div class="flex flex-row gap-x-2 w-full h-full items-center justify-center">
            <div class="custom-shadow-inset3 w-[62.87%] h-[518px] rounded-lg flex flex-col bg-white
             justify-start items-center">
                <img class="w-[98.75%] h-[92.66%] object-cover rounded-lg pt-1" src="../../random-img.jpg" alt="">
                <div class="w-full h-auto flex flex-row justify-between px-2 pb-1 mt-1 items-center">
                    <span class="text-[14.2px] font-normal">React to this pic with a smile</span>
                    <span class="text-[11px] text-gray-500 font-normal ml-auto">8:08 PM</span>
                </div>
            </div>
            <button class="w-[25px] h-[25px] bg-[#c1bfba] rounded-full cursor-pointer">
                <svg viewBox="0 0 25 25" height="25" width="25" preserveAspectRatio="xMidYMid meet" class="text-white"
                    version="1.1" x="0px" y="0px" enable-background="new 0 0 25 25">
                    <title>forward-chat</title>
                    <path fill-rule="evenodd" clip-rule="evenodd" fill="currentColor"
                        d="M14.248,6.973c0-0.614,0.741-0.921,1.174-0.488l5.131,5.136 c0.269,0.269,0.269,0.704,0,0.973l-5.131,5.136c-0.433,0.433-1.174,0.126-1.174-0.488v-2.319c-4.326,0-7.495,1.235-9.85,3.914 c-0.209,0.237-0.596,0.036-0.511-0.268c1.215-4.391,4.181-8.492,10.361-9.376V6.973z">
                    </path>
                </svg>
            </button>
        </div>
        <div class="flex flex-row gap-x-2 w-full h-full items-center justify-center">
            <div
                class="custom-shadow-inset3 w-[62.87%] h-[518px] rounded-lg flex flex-col bg-white justify-start items-center">
                <img class="w-[98.75%] h-[92.66%] object-cover rounded-lg pt-1" src="../../random-img.jpg" alt="">
                <div class="w-full h-auto items-center pl-1 pb-1 mt-1">
                    <span class="text-[14.2px] font-normal">React to this pic with a smile</span>
                </div>
            </div>
            <button class="w-[25px] h-[25px] bg-[#c1bfba] rounded-full cursor-pointer">
                <svg viewBox="0 0 25 25" height="25" width="25" preserveAspectRatio="xMidYMid meet" class="text-white"
                    version="1.1" x="0px" y="0px" enable-background="new 0 0 25 25">
                    <title>forward-chat</title>
                    <path fill-rule="evenodd" clip-rule="evenodd" fill="currentColor"
                        d="M14.248,6.973c0-0.614,0.741-0.921,1.174-0.488l5.131,5.136 c0.269,0.269,0.269,0.704,0,0.973l-5.131,5.136c-0.433,0.433-1.174,0.126-1.174-0.488v-2.319c-4.326,0-7.495,1.235-9.85,3.914 c-0.209,0.237-0.596,0.036-0.511-0.268c1.215-4.391,4.181-8.492,10.361-9.376V6.973z">
                    </path>
                </svg>
            </button>
        </div>
    </div>
</template>