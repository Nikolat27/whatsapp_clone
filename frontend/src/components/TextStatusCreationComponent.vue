<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { Ref } from 'vue';
import sharedState from '../sharedState';


const closeTextStatusCreation = () => {
    console.log("hi")
    sharedState.isTextStatusCreationOpen = false
}

const statusTextInputRef: Ref = ref(null)
const statusText: Ref<string> = ref('')
const submitStatusText = () => {
    if (statusText.value.length < 1) {
        return;
    }

    // API Endpoint...
}

onMounted(() => {
    statusTextInputRef.value.focus()
})
</script>
<template>
    <div class="fixed inset-0 w-screen h-screen flex flex-col bg-[#a62c71] z-[999]">
        <div class="flex w-full items-center items-center p-10">
            <button @click="closeTextStatusCreation" class="w-[24px] h-[24px] cursor-pointer rounded-full z-50">
                <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet" class="text-white"
                    version="1.1" x="0px" y="0px" enable-background="new 0 0 24 24">
                    <title>x-viewer</title>
                    <path fill="currentColor"
                        d="M19.8,5.8l-1.6-1.6L12,10.4L5.8,4.2L4.2,5.8l6.2,6.2l-6.2,6.2l1.6,1.6l6.2-6.2l6.2,6.2l1.6-1.6L13.6,12 L19.8,5.8z">
                    </path>
                </svg>
            </button>
        </div>
        <div class="absolute inset-0 flex justify-center items-center">
            <input ref="statusTextInputRef" v-model="statusText" placeholder="Type a status"
                class="text-[60px] font-normal text-white outline-none text-center w-[450px] bg-transparent"
                type="text">
        </div>
        <div class="fixed bottom-0 w-full items-center bg-gray-100 bg-opacity-60 flex flex-row h-[96px]">
            <button :class="[statusText.length < 1 ? 'opacity-40 bg-gray-500 cursor-default' : 'cursor-pointer']"
                @click="submitStatusText" class="flex ml-auto mr-8 items-center justify-center w-[60px] h-[60px] bg-[#00a884]
             rounded-full shadow-lg">
                <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet"
                    class="text-white ml-2" version="1.1" x="0px" y="0px" enable-background="new 0 0 24 24">
                    <title>send</title>
                    <path fill="currentColor"
                        d="M1.101,21.757L23.8,12.028L1.101,2.3l0.011,7.912l13.623,1.816L1.112,13.845 L1.101,21.757z">
                    </path>
                </svg>
            </button>
        </div>
    </div>
</template>