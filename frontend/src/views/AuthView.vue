<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from "vue";
import { useRouter } from "vue-router";
import type { Ref } from "vue";
import axiosInstance from "../utils/axiosInstance";
import { useToast } from "vue-toastification";

const toast = useToast();
const router = useRouter();


const MAX_LOGIN_ATTEMPTS = 5;
const LOGIN_DELAY = 2000;
let loginAttempts = 0;
let lastFailedAttempt = 0;

const isRegisterPageShown: Ref<boolean> = ref(true);
const toggleAuthPage = () => {
    isRegisterPageShown.value = !isRegisterPageShown.value;
    document.title = isRegisterPageShown.value ? "Register" : "Login";
};


const sanitizeInput = (input: string): string => {
    return input.replace(/[<>"'`=]/g, '');
};


const loginUsernameText: Ref<string> = ref("");
const loginPassword: Ref<string> = ref("");
const isLoggingIn: Ref<boolean> = ref(false);

const submitLogin = async () => {
    if (isLoggingIn.value) return;
    
    const now = Date.now();
    if (now - lastFailedAttempt < LOGIN_DELAY) {
        toast.warning("Please wait before trying again");
        return;
    }

    if (loginAttempts >= MAX_LOGIN_ATTEMPTS) {
        toast.error("Too many failed attempts. Please try again later.");
        return;
    }

    isLoggingIn.value = true;
    
    try {
        const username = sanitizeInput(loginUsernameText.value);
        const password = sanitizeInput(loginPassword.value);

        if (username.length < 3 || password.length < 3) {
            throw new Error("Invalid credentials");
        }

        const response = await axiosInstance.post(
            "/auth/login/",
            {
                username,
                password
            },
            {
                headers: {
                    "Content-Type": "application/json",
                    "X-Requested-With": "XMLHttpRequest"
                },
                timeout: 5000
            }
        );

        if (response.status === 200) {
            loginUsernameText.value = "";
            loginPassword.value = "";
            
            loginAttempts = 0;
            
            toast.success("User logged in successfully!");
            router.push({ path: "/" });
        }
    } catch (error) {
        loginAttempts++;
        lastFailedAttempt = Date.now();
        
        if (axios.isAxiosError(error)) {
            const errorMsg = error.response?.data?.message || "Login failed. Please try again.";
            toast.error(errorMsg);
        } else {
            toast.error("An unexpected error occurred");
        }
    } finally {
        isLoggingIn.value = false;
    }
};

const canUserLogin = computed(() => {
    return loginUsernameText.value.length >= 3 && 
           loginPassword.value.length >= 3 &&
           !isLoggingIn.value;
});

const registerUsernameText: Ref<string> = ref("");
const registerPassword1: Ref<string> = ref("");
const registerPassword2: Ref<string> = ref("");
const isRegistering: Ref<boolean> = ref(false);

const submitRegister = async (event: Event) => {
    event.preventDefault();
    if (isRegistering.value) return;
    isRegistering.value = true;

    try {
        const username = sanitizeInput(registerUsernameText.value);
        const password1 = sanitizeInput(registerPassword1.value);
        const password2 = sanitizeInput(registerPassword2.value);

        if (password1 !== password2) {
            throw new Error("Passwords do not match");
        }

        if (username.length < 3 || password1.length < 3) {
            throw new Error("Username and password must be at least 3 characters");
        }

        const response = await axiosInstance.post(
            "/auth/register/",
            {
                username,
                password: password1
            },
            {
                headers: {
                    "Content-Type": "application/json",
                    "X-Requested-With": "XMLHttpRequest"
                },
                timeout: 5000
            }
        );

        if (response.status === 201) {
            // Clear sensitive data
            registerUsernameText.value = "";
            registerPassword1.value = "";
            registerPassword2.value = "";
            
            toast.success("User Registered Successfully!");
            toggleAuthPage();
        }
    } catch (error) {
        if (axios.isAxiosError(error)) {
            const errorMsg = error.response?.data?.message || "Registration failed. Please try again.";
            toast.error(errorMsg);
        } else {
            toast.error(error.message || "An unexpected error occurred");
        }
    } finally {
        isRegistering.value = false;
    }
};

const canUserRegister = computed(() => {
    return (
        registerUsernameText.value.length >= 3 &&
        registerPassword1.value.length >= 3 &&
        registerPassword2.value.length >= 3 &&
        registerPassword1.value === registerPassword2.value &&
        !isRegistering.value
    );
});

onMounted(() => {
    document.title = "Register";
    // Clear any sensitive data from memory when component mounts
    loginUsernameText.value = "";
    loginPassword.value = "";
    registerUsernameText.value = "";
    registerPassword1.value = "";
    registerPassword2.value = "";
});

onUnmounted(() => {
    document.title = "WhatsApp Web";
    loginUsernameText.value = "";
    loginPassword.value = "";
    registerUsernameText.value = "";
    registerPassword1.value = "";
    registerPassword2.value = "";
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
                alt="WhatsApp Logo"
            />
            <input
                v-model="registerUsernameText"
                name="username"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="text"
                placeholder="username..."
                maxlength="30"
                autocomplete="username"
                @input="registerUsernameText = sanitizeInput(registerUsernameText)"
            />
            <input
                v-model="registerPassword1"
                name="new-password"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="password"
                placeholder="password..."
                maxlength="100"
                autocomplete="new-password"
                @input="registerPassword1 = sanitizeInput(registerPassword1)"
            />
            <input
                v-model="registerPassword2"
                name="new-password-confirm"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="password"
                placeholder="repeat password..."
                maxlength="100"
                autocomplete="new-password"
                @input="registerPassword2 = sanitizeInput(registerPassword2)"
            />
            <button
                @click="submitRegister"
                :class="[
                    canUserRegister
                        ? 'cursor-pointer hover:shadow-2xs'
                        : 'cursor-not-allowed opacity-60',
                ]"
                :disabled="!canUserRegister"
                class="font-semibold text-white bg-green-800 my-6 rounded-3xl w-[110px] h-[40px] text-center transition-all"
            >
                <span v-if="!isRegistering">Register</span>
                <span v-else>Processing...</span>
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
                    Home page
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
                alt="WhatsApp Logo"
            />
            <input
                v-model="loginUsernameText"
                name="username"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="text"
                placeholder="username..."
                maxlength="30"
                autocomplete="username"
                @input="loginUsernameText = sanitizeInput(loginUsernameText)"
            />
            <input
                v-model="loginPassword"
                name="current-password"
                class="w-[80%] h-[40px] rounded-lg pl-2 border-black border"
                type="password"
                placeholder="password..."
                maxlength="100"
                autocomplete="current-password"
                @input="loginPassword = sanitizeInput(loginPassword)"
            />
            <button
                @click="submitLogin"
                :disabled="!canUserLogin || isLoggingIn"
                :class="[
                    canUserLogin && !isLoggingIn
                        ? 'cursor-pointer hover:shadow-2xs'
                        : 'cursor-not-allowed opacity-60',
                ]"
                class="font-semibold text-white bg-green-800 my-6 rounded-3xl w-[110px] h-[40px] text-center transition-all"
            >
                <span v-if="!isLoggingIn">Login</span>
                <span v-else>Logging in...</span>
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
                    Home page
                </p>
            </router-link>
        </div>
    </div>
</template>
