<script setup lang="ts">
import { ref, onMounted, Ref } from "vue";
import sharedState from "../sharedState";
import axiosInstance from "../utils/axiosInstance";
import { useToast } from "vue-toastification";
import { useUserStore } from "../stores/user";

const toast = useToast();

const userProfileUrl = ref("");
const fileInputRef = ref(null);
const userProfileFile: Ref = ref(null);
const temporaryProfileImage: Ref<string> = ref(null);
const uploadUserProfilePicture = (event: any) => {
    const confirm = window.confirm("Change your profile picture?");
    if (!confirm) return;

    userProfileFile.value = event.target.files[0];
    temporaryProfileImage.value = URL.createObjectURL(userProfileFile.value);

    const formData = new FormData();
    formData.append("profile_picture", userProfileFile.value);

    axiosInstance.put("/users/update-profile", formData).then((resp) => {
        window.location.reload();
    });
};

const editUsernameText: Ref<string> = ref("");
const isEditUsernameModeOn: Ref<boolean> = ref(false);
const toggleEditUsernameMode = () => {
    if (isEditUsernameModeOn.value && editUsernameText.value.length > 1) {
        console.log("you changed your username");
        isEditUsernameModeOn.value = !isEditUsernameModeOn.value;
    } else {
        isEditUsernameModeOn.value = !isEditUsernameModeOn.value;
    }
};

const editNameText: Ref<string> = ref("");
const isEditNameModeOn: Ref<boolean> = ref(false);
const toggleEditNameMode = () => {
    if (isEditNameModeOn.value && editNameText.value.length > 1) {
        console.log("you changed your name");
        isEditNameModeOn.value = !isEditNameModeOn.value;
    } else {
        isEditNameModeOn.value = !isEditNameModeOn.value;
    }
};

const editAboutText: Ref<string> = ref("");
const isEditAboutModeOn: Ref<boolean> = ref(false);
const toggleEditAboutMode = () => {
    if (isEditAboutModeOn.value && editAboutText.value.length > 1) {
        console.log("you changed your about text");
        isEditAboutModeOn.value = !isEditAboutModeOn.value;
    } else {
        isEditAboutModeOn.value = !isEditAboutModeOn.value;
    }
};

const closeUserEditPage = () => {
    sharedState.isUserEditOpen = false;
};

async function applyChanges() {
    const formData = new FormData();

    formData.append("username", editUsernameText.value);
    formData.append("name", editNameText.value);
    formData.append("about", editAboutText.value);

    await axiosInstance
        .put("/users/update/", formData, {
            headers: {
                "Content-Type": "application/json",
            },
        })
        .then((resp) => {
            if (resp.status === 200) {
                toast.success("user profile updated successfully!");
            }
        })
        .catch((err) => {
            toast.error(err);
        });
}

async function getUserInfo() {
    let username = userStore.user.username;
    if (username) {
        editUsernameText.value = username;
    }

    let name = userStore.user.name;
    if (name) {
        editNameText.value = name;
    }

    let about = userStore.user.about;
    if (about) {
        editAboutText.value = about;
    }

    let profileUrl = userStore.user.profile_url;
    if (profileUrl) {
        userProfileUrl.value = profileUrl;
    }
}

const userStore = useUserStore();

onMounted(() => {
    getUserInfo();
});
</script>

<template>
    <div class="w-full h-[64px] pl-3 bg-white flex justify-start items-center">
        <span
            @click="closeUserEditPage"
            class="text-[15px] font-medium mr-4 cursor-pointer underline underline-offset-3"
            >get Back</span
        >
        <h1 class="text-[22px] font-bold">Profile</h1>
    </div>
    <div
        class="w-full h-auto min-h-[200px] flex justify-center items-center mb-8 bg-[#f0f2f5] py-8"
    >
        <input
            accept="image/png, image/jpeg, image/jpg"
            ref="fileInputRef"
            @change="uploadUserProfilePicture"
            type="file"
            class="hidden"
        />
        <div
            @click="fileInputRef.click()"
            class="w-[200px] h-[200px] cursor-pointer"
        >
            <img
                draggable="false"
                class="w-full h-full rounded-full object-cover"
                :src="
                    userProfileUrl
                        ? userProfileUrl
                        : temporaryProfileImage
                        ? temporaryProfileImage
                        : '../../barcelona-logo.jpg'
                "
                alt=""
            />
        </div>
    </div>
    <div
        class="flex flex-col h-[90px] w-full gap-y-4 pl-6 bg-white justify-start items-start border-b-1 border-gray-400"
    >
        <span class="text-[14px] font-normal text-green-700">Username</span>
        <div class="w-[94%] flex flex-row">
            <span v-if="!isEditUsernameModeOn">{{ editUsernameText }}</span>
            <input
                v-else
                v-model="editUsernameText"
                :class="[
                    editUsernameText.length < 1
                        ? 'border-red-600 border-2'
                        : 'border-black border',
                ]"
                type="text"
                class="w-full rounded-xl outline-none p-2 text-[16px] font-normal mr-2"
            />
            <button
                @click="toggleEditUsernameMode"
                :class="[
                    isEditUsernameModeOn
                        ? editUsernameText.length < 1
                            ? 'text-red-600'
                            : 'text-green-600'
                        : 'text-gray-400',
                ]"
                class="cursor-pointer ml-auto"
                tabindex="0"
                title="Click to edit your name"
                type="button"
                aria-label="Click to edit your name"
                style="transform: scaleX(1) scaleY(1); opacity: 1"
            >
                <span aria-hidden="true" data-icon="pencil" class="_ald7"
                    ><svg
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
                        <title>Edit</title>
                        <path
                            fill="currentColor"
                            d="M3.95,16.7v3.4h3.4l9.8-9.9l-3.4-3.4L3.95,16.7z M19.75,7.6c0.4-0.4,0.4-0.9,0-1.3 l-2.1-2.1c-0.4-0.4-0.9-0.4-1.3,0l-1.6,1.6l3.4,3.4L19.75,7.6z"
                        ></path></svg
                ></span>
            </button>
        </div>
    </div>
    <div
        class="flex flex-col h-[90px] w-full gap-y-4 pl-6 bg-white justify-start items-start mt-6"
    >
        <span class="text-[14px] font-normal text-green-700">Your name</span>
        <div class="w-[94%] flex flex-row">
            <span v-if="!isEditNameModeOn">{{ editNameText }}</span>
            <input
                v-else
                v-model="editNameText"
                :class="[
                    editNameText.length < 1
                        ? 'border-red-600 border-2'
                        : 'border-black border',
                ]"
                type="text"
                class="w-full rounded-xl outline-none p-2 text-[16px] font-normal mr-2"
            />
            <button
                @click="toggleEditNameMode"
                :class="[
                    isEditNameModeOn
                        ? editNameText.length < 1
                            ? 'text-red-600'
                            : 'text-green-600'
                        : 'text-gray-400',
                ]"
                class="cursor-pointer ml-auto"
                tabindex="0"
                title="Click to edit your name"
                type="button"
                aria-label="Click to edit your name"
                style="transform: scaleX(1) scaleY(1); opacity: 1"
            >
                <span aria-hidden="true" data-icon="pencil" class="_ald7"
                    ><svg
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
                        <title>Edit</title>
                        <path
                            fill="currentColor"
                            d="M3.95,16.7v3.4h3.4l9.8-9.9l-3.4-3.4L3.95,16.7z M19.75,7.6c0.4-0.4,0.4-0.9,0-1.3 l-2.1-2.1c-0.4-0.4-0.9-0.4-1.3,0l-1.6,1.6l3.4,3.4L19.75,7.6z"
                        ></path></svg
                ></span>
            </button>
        </div>
    </div>
    <span
        class="select-none pl-4 py-4 mb-6 text-[14px] font-normal text-gray-600 bg-[#f0f2f5]"
        >This is not your username or PIN. This name will be visible to your
        WhatsApp contacts.</span
    >
    <div
        class="flex flex-col h-[90px] w-full gap-y-4 pl-6 bg-white justify-start items-start border-b border-gray-200"
    >
        <span class="text-[14px] font-normal text-green-700">About</span>
        <div class="w-[94%] flex flex-row">
            <span v-if="!isEditAboutModeOn">{{ editAboutText }}</span>
            <input
                v-else
                v-model="editAboutText"
                :class="[
                    editAboutText.length < 1
                        ? 'border-red-600 border-2'
                        : 'border-black border',
                ]"
                type="text"
                class="w-full rounded-xl outline-none p-2 text-[16px] font-normal mr-2"
            />
            <button
                @click="toggleEditAboutMode"
                :class="[
                    isEditAboutModeOn
                        ? editAboutText.length < 1
                            ? 'text-red-600'
                            : 'text-green-600'
                        : 'text-gray-400',
                ]"
                class="cursor-pointer ml-auto"
                tabindex="0"
                title="Click to edit your name"
                type="button"
                aria-label="Click to edit your name"
                style="transform: scaleX(1) scaleY(1); opacity: 1"
            >
                <span aria-hidden="true" data-icon="pencil" class="_ald7"
                    ><svg
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
                        <title>Edit</title>
                        <path
                            fill="currentColor"
                            d="M3.95,16.7v3.4h3.4l9.8-9.9l-3.4-3.4L3.95,16.7z M19.75,7.6c0.4-0.4,0.4-0.9,0-1.3 l-2.1-2.1c-0.4-0.4-0.9-0.4-1.3,0l-1.6,1.6l3.4,3.4L19.75,7.6z"
                        ></path></svg
                ></span>
            </button>
        </div>
    </div>
    <button
        @click="applyChanges"
        class="text-white bg-green-600 w-auto px-4 py-2 self-center rounded font-semibold mt-2 cursor-pointer"
    >
        Apply Changes
    </button>
</template>
