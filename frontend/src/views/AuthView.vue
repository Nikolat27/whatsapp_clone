<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import type { Ref } from 'vue';

const isRegisterPageShown: Ref<boolean> = ref(true)
const toggleAuthPage = () => {
    isRegisterPageShown.value = !isRegisterPageShown.value
}

const submit

onMounted(() => {
    document.title = "Register"
})

onUnmounted(() => {
    document.title = "WhatsApp Web"
})
</script>
<template>
    <div class="fixed inset-0 flex items-center justify-center w-full h-full bg-[#fcf5eb] z-[500]">
        <div v-if="isRegisterPageShown" class="flex flex-col gap-y-3 items-center justify-start w-[25%] h-[85%] rounded-2xl border-black
         border bg-white">
            <img class="w-[200px] h-[200px] my-14" src="../assets/WhatsApp.svg.webp" alt="">
            <input name="email" class="w-[80%] h-[40px] rounded-lg pl-2 border-black border" type="email"
                placeholder="email...">
            <input name="username" class="w-[80%] h-[40px] rounded-lg pl-2 border-black border" type="text"
                placeholder="username...">
            <input name="password" class="w-[80%] h-[40px] rounded-lg pl-2 border-black border" type="password"
                placeholder="password...">
            <input name="password2" class="w-[80%] h-[40px] rounded-lg pl-2 border-black border" type="password"
                placeholder="repeat password...">
            <button
                class="cursor-pointer hover:shadow-2xs font-semibold text-white bg-green-800 my-6 rounded-3xl w-[110px] h-[40px] text-center">
                Register
            </button>
            <p @click="toggleAuthPage" class="cursor-pointer text-[15px] font-semibold text-green-700">Already have an
                account? <span class="underline underline-offset-2">Login</span></p>
            <router-link to="/" class="cursor-pointer">
                <p class="text-[16px] font-semibold text-green-700 underline underline-offset-2">Home page?</p>
            </router-link>
        </div>
        <div v-else class="flex flex-col gap-y-3 items-center justify-start w-[25%] h-[85%] rounded-2xl border-black
         border bg-white">
            <img class="w-[200px] h-[200px] my-14" src="../assets/WhatsApp.svg.webp" alt="">
            <input name="username" class="w-[80%] h-[40px] rounded-lg pl-2 border-black border" type="text"
                placeholder="username...">
            <input name="password" class="w-[80%] h-[40px] rounded-lg pl-2 border-black border" type="password"
                placeholder="password...">
            <button class="cursor-pointer hover:shadow-2xs font-semibold text-white bg-green-800 my-6
             rounded-3xl w-[110px] h-[40px] text-center">
                Login
            </button>
            <p @click="toggleAuthPage" class="text-[15px] font-semibold text-green-700 cursor-pointer">No account? <span
                    class="underline underline-offset-2">Register</span></p>
            <router-link to="/" class="cursor-pointer">
                <p class="text-[16px] font-semibold text-green-700 underline underline-offset-2">Home page?</p>
            </router-link>
        </div>
    </div>
</template>