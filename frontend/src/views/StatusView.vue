<script setup lang="ts">
import { ref, watch } from 'vue';
import type { Ref } from 'vue';
import sharedState from '../sharedState';
import TextStatusCreationComponent from '../components/TextStatusCreationComponent.vue';

const isStatusCreatingDivOpen: Ref<boolean> = ref(false)
const toggleStatusCreationDiv = () => {
    isStatusCreatingDivOpen.value = !isStatusCreatingDivOpen.value
}

const isTextStatusCreationOpen: Ref<boolean> = ref(false)
const openTextStatusCreation = () => {
    isTextStatusCreationOpen.value = !isTextStatusCreationOpen.value
    sharedState.isTextStatusCreationOpen = true
}

watch(() => sharedState.isTextStatusCreationOpen, (newVal) => {
    isTextStatusCreationOpen.value = newVal
})

const statusUploadedFile = ref(null)
const fileInputRef = ref(null)

const handleStatusFileUpload = (event: any) => {
    statusUploadedFile.value = event.target.files[0]
}
</script>

<template>
    <TextStatusCreationComponent v-if="isTextStatusCreationOpen"></TextStatusCreationComponent>
    <div class="flex flex-row w-full pt-4 pl-4 select-none">
        <h1 class="text-[22px] font-bold">Status</h1>
        <div class="relative flex flex-row gap-x-2 ml-auto mr-2">
            <button @click="toggleStatusCreationDiv"
                class="w-10 h-10 rounded-full flex justify-center items-center cursor-pointer hover:bg-[#f2f2f2]">
                <span aria-hidden="true" data-icon="plus" class=""><svg viewBox="0 0 24 24" height="24" width="24"
                        preserveAspectRatio="xMidYMid meet" class="">
                        <title>plus</title>
                        <path fill="currentColor" d="M19,13h-6v6h-2v-6H5v-2h6V5h2v6h6V13z"></path>
                    </svg></span>
            </button>
            <div v-if="isStatusCreatingDivOpen"
                class="absolute top-10 right-0 flex flex-col w-[169px] h-auto py-3 text-[15px] font-normal bg-white custom-shadow-inset z-50">
                <div @click="fileInputRef.click()"
                    class="flex flex-row items-center h-[40px] gap-x-3 pl-3 hover:bg-[#f5f6f6] cursor-pointer">
                    <svg viewBox="0 0 20 20" height="20" width="20" preserveAspectRatio="xMidYMid meet"
                        class="text-gray-500" fill="none">
                        <title>media-multiple</title>
                        <path fill-rule="evenodd" clip-rule="evenodd"
                            d="M20 14V2C20 0.9 19.1 0 18 0H6C4.9 0 4 0.9 4 2V14C4 15.1 4.9 16 6 16H18C19.1 16 20 15.1 20 14ZM9.4 10.53L11.03 12.71L13.61 9.49C13.81 9.24 14.19 9.24 14.39 9.49L17.35 13.19C17.61 13.52 17.38 14 16.96 14H7C6.59 14 6.35 13.53 6.6 13.2L8.6 10.53C8.8 10.27 9.2 10.27 9.4 10.53ZM0 18V5C0 4.45 0.45 4 1 4C1.55 4 2 4.45 2 5V17C2 17.55 2.45 18 3 18H15C15.55 18 16 18.45 16 19C16 19.55 15.55 20 15 20H2C0.9 20 0 19.1 0 18Z"
                            fill="currentColor"></path>
                    </svg>
                    <span>Photos</span>
                </div>
                <div @click="openTextStatusCreation"
                    class="flex flex-row items-center h-[40px] gap-x-3 pl-3 hover:bg-[#f5f6f6] cursor-pointer">
                    <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet"
                        class="x1xp8n7a text-gray-500" fill="none">
                        <title>media-editor-drawing</title>
                        <path fill-rule="evenodd" clip-rule="evenodd"
                            d="M1.35999 18.0342V22.6393H5.99278L18.9143 9.66526L14.2817 5.06126L1.35999 18.0342ZM16.1582 9.67289L5.18355 20.6911H3.30815V18.8389L14.2887 7.81484L16.1582 9.67289Z"
                            fill="currentColor"></path>
                        <path fill-rule="evenodd" clip-rule="evenodd"
                            d="M17.5747 1.78958L15.3879 3.96288L20.0211 8.56752L22.2067 6.39544C22.4885 6.1154 22.6393 5.78963 22.6393 5.45357C22.6393 5.11751 22.4885 4.79175 22.2067 4.51171L19.4676 1.78958C19.1858 1.50949 18.8584 1.36 18.5212 1.36C18.1839 1.36 17.8565 1.50949 17.5747 1.78958ZM20.0211 5.82089L18.1516 3.96288L18.5212 3.59557L20.3907 5.45357L20.0211 5.82089Z"
                            fill="currentColor"></path>
                    </svg>
                    <span>Text</span>
                </div>
            </div>
            <input accept="img/" @change="handleStatusFileUpload" ref="fileInputRef" type="file" class="hidden">
        </div>
    </div>
    <div class="flex flex-row w-full pl-4 mt-4 select-none">
        <button class="w-10 h-10 rounded-full overflow-hidden relative">
            <img class="w-full h-full" src="../../barcelona-logo.jpg" alt="">
            <button class="relative ml-2 top-[20px] right-5 w-[21px] h-[21px] rounded-full flex
                justify-center items-center bg-green-600">
                <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="100" height="100" viewBox="0 0 24 24"
                    style="fill:#FFFFFF;">
                    <path
                        d="M 12 2 C 6.4860328 2 2 6.4860368 2 12 C 2 17.513963 6.4860328 22 12 22 C 17.513967 22 22 17.513963 22 12 C 22 6.4860368 17.513967 2 12 2 z M 12 3.5 C 16.703308 3.5 20.5 7.2966955 20.5 12 C 20.5 16.703304 16.703308 20.5 12 20.5 C 7.2966924 20.5 3.5 16.703304 3.5 12 C 3.5 7.2966955 7.2966924 3.5 12 3.5 z M 11.988281 6.9902344 A 0.750075 0.750075 0 0 0 11.25 7.75 L 11.25 11.25 L 7.75 11.25 A 0.750075 0.750075 0 1 0 7.75 12.75 L 11.25 12.75 L 11.25 16.25 A 0.750075 0.750075 0 1 0 12.75 16.25 L 12.75 12.75 L 16.25 12.75 A 0.750075 0.750075 0 1 0 16.25 11.25 L 12.75 11.25 L 12.75 7.75 A 0.750075 0.750075 0 0 0 11.988281 6.9902344 z">
                    </path>
                </svg>
            </button>
        </button>
        <div class="flex flex-col ml-2">
            <p class="font-normal text-[16px] text-gray-700">My status</p>
            <p class="font-normal text-[13px] text-gray-400">Click to add status update</p>
        </div>
    </div>
    <div class="w-full h-2.5 bg-[#f0f2f5] my-4" style="box-shadow: 0 0 4px rgba(0, 0, 0, 0.25);"></div>
    <div class="flex flex-col pt-2 select-none">
        <h1 class="text-[16px] pl-8 font-normal text-green-700 w-full border-b border-gray-200 pb-4">VIEWED
        </h1>
        <div class="flex flex-row pl-4 items-center cursor-pointer w-full mt-2 gap-x-2 hover:bg-[#f5f6f6] h-[72px]">
            <button class="w-10 h-10 rounded-full overflow-hidden">
                <img class="w-full h-full" src="../../barcelona-logo.jpg" alt="">
            </button>
            <div class="flex flex-col">
                <p class="font-normal text-[16px] textblack">Lionel messi</p>
                <p class="font-normal text-[13px] text-gray-600">Today at <span>3:45 PM</span></p>
            </div>
        </div>
    </div>
    <div class="select-none flex flex-row self-center mt-12 text-[12px] text-gray-500 items-center
     font-normal cursor-default">
        <span aria-hidden="true" data-icon="lock-small-v2" class="xvijh9v xeyog9w x1rg5ohu x16cd2qt x10kle22"><svg
                viewBox="0 0 13 12" height="12" width="13" preserveAspectRatio="xMidYMid meet" class="text-gray-400">
                <title>lock-small-v2</title>
                <path
                    d="M9.54004 3.4668C9.54004 2.87891 9.39421 2.33887 9.10254 1.84668C8.81543 1.34993 8.4235 0.958008 7.92676 0.670898C7.43457 0.379232 6.89681 0.233398 6.31348 0.233398C5.72559 0.233398 5.18327 0.379232 4.68652 0.670898C4.19434 0.958008 3.80241 1.34993 3.51074 1.84668C3.22363 2.33887 3.08008 2.87891 3.08008 3.4668V4.7041C3.05273 4.71322 2.99805 4.73828 2.91602 4.7793C2.61979 4.9388 2.39421 5.16439 2.23926 5.45605C2.15267 5.61556 2.09115 5.79102 2.05469 5.98242C2.01823 6.17383 2 6.45866 2 6.83691V9.25C2 9.62826 2.01823 9.91309 2.05469 10.1045C2.09115 10.2959 2.15267 10.4714 2.23926 10.6309C2.39421 10.9225 2.61979 11.1481 2.91602 11.3076C3.07096 11.3942 3.24414 11.4557 3.43555 11.4922C3.63151 11.5286 3.91634 11.5469 4.29004 11.5469H8.33008C8.70378 11.5469 8.98633 11.5286 9.17773 11.4922C9.3737 11.4557 9.54915 11.3942 9.7041 11.3076C9.99577 11.1527 10.2214 10.9271 10.3809 10.6309C10.4674 10.4714 10.529 10.2959 10.5654 10.1045C10.6019 9.91309 10.6201 9.62826 10.6201 9.25V6.83691C10.6201 6.45866 10.6019 6.17383 10.5654 5.98242C10.529 5.79102 10.4674 5.61556 10.3809 5.45605C10.2214 5.15983 9.99577 4.93424 9.7041 4.7793C9.62207 4.73828 9.56738 4.71322 9.54004 4.7041V3.4668ZM4.37207 3.4668C4.37207 3.11589 4.45866 2.79232 4.63184 2.49609C4.80501 2.19531 5.03971 1.95833 5.33594 1.78516C5.63672 1.61198 5.96257 1.52539 6.31348 1.52539C6.66439 1.52539 6.98796 1.61198 7.28418 1.78516C7.5804 1.95833 7.8151 2.19531 7.98828 2.49609C8.16146 2.79232 8.24805 3.11589 8.24805 3.4668V4.54004H4.37207V3.4668Z">
                </path>
            </svg></span>
        <p class="ml-1">Your status are <span
                class="text-green-700 hover:underline hover:underline-offset-2 cursor-pointer">end-to-end
                encrypted</span></p>
    </div>
</template>