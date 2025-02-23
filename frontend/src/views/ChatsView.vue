<script setup lang="ts">
import { ref, watch, reactive, Reactive } from 'vue'
import type { Ref } from 'vue';
import ClipLoader from 'vue-spinner/src/ClipLoader.vue'
import sharedState from '../sharedState';

import ArchiveChatsComponent from '/src/components/ArchiveChatsComponent.vue';
import NewChatPageComponent from '/src/components/NewChatComponent.vue';

// Icons
import arrowLeftIcon from '/src/assets/icons/svg-icons/arrow-left-green.svg'

const isChatOptionsDivOpen: Ref<Boolean> = ref(false)
const toggleChatOptionsDiv = () => {
    isChatOptionsDivOpen.value = !isChatOptionsDivOpen.value
}

const searchText: Ref<String> = ref('')
const isSearchModeOn: Ref<Boolean> = ref(false)
const toggleSearchMode = () => {
    searchText.value = ''
    isSearchModeOn.value = !isSearchModeOn.value
}

const closeSearchBar = () => {
    toggleSearchMode()
}

const currentChatFilter: Ref<String> = ref('all')
const chatShowingFilter = (type: string) => {
    currentChatFilter.value = type
}

const isChatLoading: Ref<Boolean> = ref(false)


const isArchiveChatOpen: Ref<Boolean> = ref(false)
const toggleArchiveChat = () => {
    isArchiveChatOpen.value = !isArchiveChatOpen.value
    sharedState.isArchiveChatOpen = !sharedState.isArchiveChatOpen
}

watch(() => sharedState.isArchiveChatOpen, (newVal) => {
    isArchiveChatOpen.value = newVal
})

const openChatOptions = reactive([])
const toggleChatOptions = (chatId: Number) => {
    if (openChatOptions.includes(chatId)) {
        openChatOptions.pop()
    } else {
        openChatOptions.pop()
        openChatOptions.push(chatId)
    }
}

watch(() => sharedState.isNewChatPageOpen, (newVal) => {
    isNewChatPageOpen.value = newVal
})

const isNewChatPageOpen = ref(false)
const toggleNewChatPage = () => {
    isNewChatPageOpen.value = !isNewChatPageOpen.value
    sharedState.isNewChatPageOpen = !sharedState.isNewChatPageOpen
}
</script>
<template>
    <div v-if="!isArchiveChatOpen && !isNewChatPageOpen" class="flex flex-col w-full h-full">
        <div class="select-none flex flex-row relative h-[64px] w-full pt-4 pl-4 justify-start items-center">
            <h1 class="text-[22px] font-bold">Chats</h1>
            <button @click="toggleNewChatPage"
                class="ml-auto w-10 h-10 rounded-full cursor-pointer flex justify-center items-center">
                <span><svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet" class=""
                        fill="none">
                        <title>new-chat-outline</title>
                        <path
                            d="M9.53277 12.9911H11.5086V14.9671C11.5086 15.3999 11.7634 15.8175 12.1762 15.9488C12.8608 16.1661 13.4909 15.6613 13.4909 15.009V12.9911H15.4672C15.9005 12.9911 16.3181 12.7358 16.449 12.3226C16.6659 11.6381 16.1606 11.0089 15.5086 11.0089H13.4909V9.03332C13.4909 8.60007 13.2361 8.18252 12.8233 8.05119C12.1391 7.83391 11.5086 8.33872 11.5086 8.991V11.0089H9.49088C8.83941 11.0089 8.33411 11.6381 8.55097 12.3226C8.68144 12.7358 9.09947 12.9911 9.53277 12.9911Z"
                            fill="currentColor"></path>
                        <path fill-rule="evenodd" clip-rule="evenodd"
                            d="M0.944298 5.52617L2.99998 8.84848V17.3333C2.99998 18.8061 4.19389 20 5.66665 20H19.3333C20.8061 20 22 18.8061 22 17.3333V6.66667C22 5.19391 20.8061 4 19.3333 4H1.79468C1.01126 4 0.532088 4.85997 0.944298 5.52617ZM4.99998 8.27977V17.3333C4.99998 17.7015 5.29845 18 5.66665 18H19.3333C19.7015 18 20 17.7015 20 17.3333V6.66667C20 6.29848 19.7015 6 19.3333 6H3.58937L4.99998 8.27977Z"
                            fill="currentColor"></path>
                    </svg></span>
            </button>
            <button @click="toggleChatOptionsDiv"
                class="w-10 hover:bg-[#e6e7e8] h-10 rounded-full cursor-pointer flex justify-center items-center mr-3">
                <span aria-hidden="true" data-icon="menu" class=""><svg viewBox="0 0 24 24" height="24" width="24"
                        preserveAspectRatio="xMidYMid meet" class="" version="1.1" x="0px" y="0px"
                        enable-background="new 0 0 24 24">
                        <title>menu</title>
                        <path fill="currentColor"
                            d="M12,7c1.104,0,2-0.896,2-2c0-1.105-0.895-2-2-2c-1.104,0-2,0.894-2,2 C10,6.105,10.895,7,12,7z M12,9c-1.104,0-2,0.894-2,2c0,1.104,0.895,2,2,2c1.104,0,2-0.896,2-2C13.999,9.895,13.104,9,12,9z M12,15 c-1.104,0-2,0.894-2,2c0,1.104,0.895,2,2,2c1.104,0,2-0.896,2-2C13.999,15.894,13.104,15,12,15z">
                        </path>
                    </svg></span>
            </button>
            <div v-if="isChatOptionsDivOpen" class="absolute custom-shadow-inset3 top-16 right-4 w-[225px] h-auto flex flex-col
         items-center py-2 bg-white z-40 text-[14.5px] font-normal">
                <div class="w-full h-[40px] hover:bg-[#f5f6f6] pl-6 flex items-center cursor-pointer">
                    <p>New group</p>
                </div>
                <div class="w-full h-[40px] hover:bg-[#f5f6f6] pl-6 flex items-center cursor-pointer">
                    <p>Select chats</p>
                </div>
                <div class="w-full h-[40px] hover:bg-[#f5f6f6] pl-6 flex items-center cursor-pointer">
                    <p>Log out</p>
                </div>
            </div>
        </div>
        <div class="select-none flex flex-row gap-x-1 ml-3 mt-6 w-[93%] h-[35px] justify-center
     items-center bg-[#f0f2f5] rounded-lg">
            <button v-if="!isSearchModeOn" @click="toggleSearchMode" class="w-[24px] h-[24px] ml-2 cursor-pointer">
                <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet" class=""
                    version="1.1" x="0px" y="0px" enable-background="new 0 0 24 24">
                    <title>search</title>
                    <path fill="currentColor"
                        d="M15.009,13.805h-0.636l-0.22-0.219c0.781-0.911,1.256-2.092,1.256-3.386 c0-2.876-2.332-5.207-5.207-5.207c-2.876,0-5.208,2.331-5.208,5.207s2.331,5.208,5.208,5.208c1.293,0,2.474-0.474,3.385-1.255 l0.221,0.22v0.635l4.004,3.999l1.194-1.195L15.009,13.805z M10.201,13.805c-1.991,0-3.605-1.614-3.605-3.605 s1.614-3.605,3.605-3.605s3.605,1.614,3.605,3.605S12.192,13.805,10.201,13.805z">
                    </path>
                </svg>
            </button>
            <button v-else @click="toggleSearchMode" class="w-[24px] h-[24px] ml-2 cursor-pointer">
                <img class="w-[18px] h-[18px]" :src="arrowLeftIcon" alt="">
            </button>
            <input @click="isSearchModeOn = true" v-model="searchText" type="text"
                class="w-[90%] ml-4 h-full outline-none" placeholder="Search">
            <button v-if="searchText.length >= 1" @click="closeSearchBar" class="w-[24px] h-[24px] cursor-pointer mr-2">
                <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet" class=""
                    version="1.1" x="0px" y="0px" enable-background="new 0 0 24 24">
                    <title>x-alt</title>
                    <path fill="currentColor"
                        d="M17.25,7.8L16.2,6.75l-4.2,4.2l-4.2-4.2L6.75,7.8l4.2,4.2l-4.2,4.2l1.05,1.05l4.2-4.2l4.2,4.2l1.05-1.05 l-4.2-4.2L17.25,7.8z">
                    </path>
                </svg>
            </button>
        </div>
        <div class="select-none flex flex-row gap-x-2 pl-3 mt-2">
            <button @click="chatShowingFilter('all')"
                :class="[currentChatFilter === 'all' ? 'bg-[#e7fce3] text-green-600' : 'bg-[#f0f2f5] text-gray-600']"
                class="w-auto px-3 h-[32px] rounded-3xl cursor-pointer">
                All
            </button>
            <button @click="chatShowingFilter('unread')"
                :class="[currentChatFilter === 'unread' ? 'bg-[#e7fce3] text-green-600' : 'bg-[#f0f2f5] text-gray-600']"
                class="w-auto px-3 h-[32px] rounded-3xl cursor-pointer">
                Unread
            </button>
            <button @click="chatShowingFilter('favorites')"
                :class="[currentChatFilter === 'favorites' ? 'bg-[#e7fce3] text-green-600' : 'bg-[#f0f2f5] text-gray-600']"
                class="w-auto px-3 h-[32px] rounded-3xl cursor-pointer">
                Favorites
            </button>
            <button @click="chatShowingFilter('groups')"
                :class="[currentChatFilter === 'groups' ? 'bg-[#e7fce3] text-green-600' : 'bg-[#f0f2f5] text-gray-600']"
                class="w-auto px-3 h-[32px] rounded-3xl cursor-pointer">
                Groups
            </button>
        </div>
        <div @click="toggleArchiveChat" class="cursor-pointer select-none w-full h-[49px] pl-6 mt-2 flex border-b border-gray-200
        flex-row justify-start items-center">
            <button class="w-[24px] h-[24px] mr-8 cursor-pointer">
                <span aria-hidden="true" data-icon="archived-outline" class="text-green-600"><svg viewBox="0 0 24 24"
                        height="24" width="24" preserveAspectRatio="xMidYMid meet" class="" fill="none">
                        <title>archived-outline</title>
                        <path
                            d="M21.4889 4.47778L19.9444 2.61111C19.6444 2.23333 19.1889 2 18.6667 2H5.33333C4.81111 2 4.35556 2.23333 4.04444 2.61111L2.51111 4.47778C2.18889 4.85556 2 5.35556 2 5.88889V19.7778C2 21 3 22 4.22222 22H19.7778C21 22 22 21 22 19.7778V5.88889C22 5.35556 21.8111 4.85556 21.4889 4.47778ZM5.6 4.22222H18.4L19.3 5.3H4.71111L5.6 4.22222ZM4.22222 19.7778V7.55556H19.7778V19.7778H4.22222ZM13.6111 9.77778H10.3889V13.1111H7.55556L12 17.5556L16.4444 13.1111H13.6111V9.77778Z"
                            fill="currentColor"></path>
                    </svg></span>
            </button>
            <p class="text-[17px] font-normal">Archived</p>
        </div>
        <div v-if="!isChatLoading" class="select-none flex flex-col w-full">
            <div class="flex border-b border-gray-200 flex-row w-full h-[72px] cursor-pointer
                justify-start items-center hover:bg-[#f5f6f6] pl-3">
                <img class="w-[49px] h-[49px] rounded-full object-cover" src="../../barcelona-logo.jpg" alt="">
                <div class="flex flex-col ml-4 w-full">
                    <div class="flex flex-row w-full">
                        <p class="font-normal text-[17px]">Friend..</p>
                        <span class="ml-auto mr-3 font-normal text-[12px] text-gray-500">Yesterday</span>
                    </div>
                    <div class="relative flex flex-row items-center w-full">
                        <p class="font-normal text-[14px] text-gray-600">Your last chat...</p>
                        <button @click="toggleChatOptions(1)"
                            class="flex cursor-pointer justify-center ml-auto mr-3 items-center w-[24px] h-[24px]">
                            <img draggable="false" class="w-full h-full" src="../../arrow-down-icon.svg" alt="">
                        </button>
                        <div v-if="openChatOptions.includes(1)" class="absolute top-8 right-2 flex w-[198px] h-auto py-2 flex-col items-center
                         bg-white z-[999] custom-shadow-inset cursor-default">
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Archive chat</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Mute notifications</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Pin chat</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Mark as unread</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Block</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Delete chat</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="flex border-b border-gray-200 flex-row w-full h-[72px] cursor-pointer
                justify-start items-center hover:bg-[#f5f6f6] pl-3">
                <img class="w-[49px] h-[49px] rounded-full object-cover" src="../../barcelona-logo.jpg" alt="">
                <div class="flex flex-col ml-4 w-full">
                    <div class="flex flex-row w-full">
                        <p class="font-normal text-[17px]">Friend..</p>
                        <span class="ml-auto mr-3 font-normal text-[12px] text-gray-500">Yesterday</span>
                    </div>
                    <div class="relative flex flex-row items-center w-full">
                        <p class="font-normal text-[14px] text-gray-600">Your last chat...</p>
                        <button @click="toggleChatOptions(2)"
                            class="flex cursor-pointer justify-center ml-auto mr-3 items-center w-[24px] h-[24px]">
                            <img draggable="false" class="w-full h-full" src="../../arrow-down-icon.svg" alt="">
                        </button>
                        <div v-if="openChatOptions.includes(2)" class="absolute top-8 right-2 flex w-[198px] h-auto py-2 flex-col items-center
                         bg-white z-[999] custom-shadow-inset cursor-default">
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Archive chat</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Mute notifications</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Pin chat</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Mark as unread</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Block</span>
                            </div>
                            <div class="w-full h-[40px] flex items-center pl-6 cursor-pointer hover:bg-[#f5f6f6]">
                                <span>Delete chat</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div v-else class="w-full h-full flex justify-center items-center">
            <ClipLoader size="45px"></ClipLoader>
        </div>
    </div>
    <div v-if="isArchiveChatOpen && !isNewChatPageOpen" class="flex flex-col w-full h-full">
        <ArchiveChatsComponent></ArchiveChatsComponent>
    </div>
    <div v-if="isNewChatPageOpen && !isArchiveChatOpen" class="flex flex-col w-full h-full">
        <NewChatPageComponent></NewChatPageComponent>
    </div>
</template>