<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from "vue";
import { useRouter } from "vue-router";
import type { Ref } from "vue";
import axiosInstance from "../utils/axiosInstance";
import { useToast } from "vue-toastification";
import axios from "axios";

const toast = useToast();

const router = useRouter();

const isRegisterPageShown: Ref<boolean> = ref(true);
const toggleAuthPage = () => {
    isRegisterPageShown.value = !isRegisterPageShown.value;
    document.title = isRegisterPageShown.value ? "Register" : "Login";
};


// Login
const loginUsernameText: Ref<string> = ref("");
const loginPassword: Ref<string> = ref("");

const submitLogin = async () => {
    const formData = new FormData();

    formData.append("username", loginUsernameText.value);
    formData.append("password", loginPassword.value);

    await axiosInstance
        .post("/auth/login/", formData, {
            headers: {
                "Content-Type": "application/json",
            },
        })
        .then((resp) => {
            if (resp.status === 200) {
                toast.success("User logged in successfully!");
                router.push({ path: "/" });
            }
        })
        .catch((error) => {
            toast.error(error.response.data);
        });
};

const canUserLogin = computed(() => {
    if (loginUsernameText.value.length < 3 || loginPassword.value.length < 3) {
        return false;
    }
    return true;
});


// Register
const registerUsernameText: Ref<string> = ref("");
const registerPassword1: Ref<string> = ref("");
const registerPassword2: Ref<string> = ref("");

const submitRegister = async (event: Event) => {
    event.preventDefault();
    const formData = new FormData();

    formData.append("username", registerUsernameText.value);
    formData.append("password", registerPassword1.value);

    await axiosInstance
        .post("/auth/register/", formData, {
            headers: {
                "Content-Type": "application/json",
            },
        })
        .then((resp) => {
            if (resp.status === 201) {
                toast.success("User Registered Successfully!");
                toggleAuthPage();
            }
        })
        .catch((error) => toast.error(error.response.data));
};

const canUserRegister = computed(() => {
    if (
        !registerUsernameText.value ||
        !registerPassword1.value ||
        !registerPassword2.value
    ) {
        return false;
    }

    if (
        registerUsernameText.value.length < 3 ||
        registerPassword1.value.length < 3 ||
        registerPassword2.value.length < 3
    ) {
        return false;
    }

    if (registerPassword1.value !== registerPassword2.value) {
        return false;
    }

    return true;
});

onMounted(() => {
    document.title = "Register";
});

onUnmounted(() => {
    document.title = "WhatsApp Web";
});
</script>
<template>
    <div
        class="fixed inset-0 flex items-center justify-center w-full h-full bg-[#fcf5eb] z-[500]"
    >
        <div
            v-if="isRegisterPageShown"
            class="flex flex-col gap-y-3 items-center justify-start w-[25%] h-[85%] rounded-2xl border-black border bg-white"
        >
            <img
                class="w-[200px] h-[200px] my-14"
                src="../assets/WhatsApp.svg.webp"
                alt=""
            />
            <input
                v-model="registerUsernameText"
                name="username"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="text"
                placeholder="username..."
            />
            <input
                v-model="registerPassword1"
                name="password"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="password"
                placeholder="password..."
            />
            <input
                v-model="registerPassword2"
                name="password2"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="password"
                placeholder="repeat password..."
            />
            <button
                @click="submitRegister"
                :class="[
                    canUserRegister
                        ? 'cursor-pointer'
                        : 'cursor-default opacity-60',
                ]"
                :disabled="!canUserRegister"
                class="hover:shadow-2xs font-semibold text-white bg-green-800 my-6 rounded-3xl w-[110px] h-[40px] text-center"
            >
                Register
            </button>
            <p
                @click="toggleAuthPage"
                class="cursor-pointer text-[15px] font-semibold text-green-700"
            >
                Already have an account?
                <span class="underline underline-offset-2">Login</span>
            </p>
            <router-link to="/" class="cursor-pointer">
                <p
                    class="text-[16px] font-semibold text-green-700 underline underline-offset-2"
                >
                    Home page?
                </p>
            </router-link>
        </div>
        <div
            v-else
            class="flex flex-col gap-y-3 items-center justify-start w-[25%] h-[85%] rounded-2xl border-black border bg-white"
        >
            <img
                class="w-[200px] h-[200px] my-14"
                src="../assets/WhatsApp.svg.webp"
                alt=""
            />
            <input
                v-model="loginUsernameText"
                name="username"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="text"
                placeholder="username..."
            />
            <input
                v-model="loginPassword"
                name="password"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="password"
                placeholder="password..."
            />
            <button
                @click="submitLogin"
                :disabled="!canUserLogin"
                :class="[
                    canUserLogin
                        ? 'cursor-pointer'
                        : 'cursor-default opacity-60',
                ]"
                class="hover:shadow-2xs font-semibold text-white bg-green-800 my-6 rounded-3xl w-[110px] h-[40px] text-center"
            >
                Login
            </button>
            <p
                @click="toggleAuthPage"
                class="text-[15px] font-semibold text-green-700 cursor-pointer"
            >
                No account?
                <span class="underline underline-offset-2">Register</span>
            </p>
            <router-link to="/" class="cursor-pointer">
                <p
                    class="text-[16px] font-semibold text-green-700 underline underline-offset-2"
                >
                    Home page?
                </p>
            </router-link>
        </div>
    </div>
</template>
