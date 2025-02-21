<script setup lang="ts">
import sharedState from '../sharedState';

const closeArchiveChats = () => {
    sharedState.isArchiveChatOpen = false
}
</script>
<template>
    <div class="flex flex-row border-b w-full h-[60px] border-b-gray-200 items-center pl-6 pt-1 gap-x-6">
        <button @click="closeArchiveChats" class="w-[24px] h-[24px] cursor-pointer">
            <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet" class="" version="1.1"
                x="0px" y="0px" enable-background="new 0 0 24 24">
                <title>back</title>
                <path fill="currentColor" d="M12,4l1.4,1.4L7.8,11H20v2H7.8l5.6,5.6L12,20l-8-8L12,4z"></path>
            </svg>
        </button>
        <span class="text-[16px] font-normal">Archived</span>
    </div>
    <div class="select-none flex flex-col w-full">
        <div class="flex border-b border-gray-200 flex-row w-full h-[72px] cursor-pointer
                justify-start items-center hover:bg-[#f5f6f6] pl-3">
            <img class="w-[49px] h-[49px] rounded-full object-cover" src="../../barcelona-logo.jpg" alt="">
            <div class="flex flex-col ml-4 w-full">
                <div class="flex flex-row w-full">
                    <p class="font-normal text-[17px]">Friend..</p>
                    <span class="ml-auto mr-3 font-normal text-[12px] text-gray-500">Yesterday</span>
                </div>
                <div class="flex flex-row items-center w-full">
                    <p class="font-normal text-[14px] text-gray-600">Your last chat...</p>
                    <button class="flex cursor-pointer justify-center ml-auto mr-3 items-center w-[24px] h-[24px]">
                        <img draggable="false" class="w-full h-full" src="../../arrow-down-icon.svg" alt="">
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>