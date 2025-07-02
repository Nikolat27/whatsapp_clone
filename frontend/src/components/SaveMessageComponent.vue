<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted } from "vue";
import type { Ref, Reactive } from "vue";
import arrowLeftIcon from "../assets/icons/svg-icons/arrow-left-green.svg";
import { useToast } from "vue-toastification";
import EmojiPicker from "vue3-emoji-picker";
import axiosInstance from "../utils/axiosInstance";
import { useUserStore } from "../stores/user";
import { useChatStore } from "../stores/chat";

const toast = useToast();

const isChatOptionOpen = ref(false);
const toggleChatOption = () => {
    isChatOptionOpen.value = !isChatOptionOpen.value;
};

const messages: Reactive<any> = reactive([]);

const openChatMessageOptions: Reactive<any> = reactive([]);
const toggleChatMessageOptions = (chatId: Number) => {
    if (openChatMessageOptions.includes(chatId)) {
        openChatMessageOptions.pop();
    } else {
        openChatMessageOptions.pop();
        openChatMessageOptions.push(chatId);
    }
};

const dialogType: Ref<String | null> = ref(null);
const handleLayerClick = (event: Event) => {
    if (event.target === event.currentTarget) {
        console.log("hi");
    }
};

const searchText: Ref<String> = ref("");
const isSearchModeOn: Ref<Boolean> = ref(false);
const toggleSearchMode = () => {
    isSearchModeOn.value = !isSearchModeOn.value;
};

const closeSearchBar = () => {
    isChatOptionOpen.value = false;
    searchText.value = "";
    isSearchModeOn.value = false;
};

const messageText: Ref<String> = ref("");
const createNewMessage = (text: String) => {
    messageText.value = "";
};

const handleKeyDown = (event: any) => {
    if (event.key === "Enter") {
        if (messageText.value.length < 1) {
            toast.info("You must enter at least 1 word");
            return;
        }
        createNewMessage(messageText.value);
    }
};

const copyText = (messageText: string) => {
    navigator.clipboard.writeText(messageText);
    toast.info("Message copied to your clipboard");
    openChatMessageOptions.pop();
};

const editModeEnable: Ref<boolean> = ref(false);
const editMessageText: Ref<string> = ref(null);
const editMessageId: Ref<number> = ref(null);
const editText = (messageText: string, messageId: number) => {
    editMessageText.value = messageText;
    editMessageId.value = messageId;
    editModeEnable.value = !editModeEnable.value;
    openChatMessageOptions.pop();
};

const submitEditMessage = (messageText: string) => {
    editMessageText.value = messageText;
    if (editMessageText.value.length < 1) {
        toast.error("You must enter at least 1 word (or Delete it)");
        return;
    }
    // API endpoint...

    // End endpoint.
    editModeEnable.value = false;
    editMessageText.value = null;

    toast.success("Messaged edited successfully");
};

const deleteMessage = (messageId: number) => {
    toast.success("Message deleted successfully!");
    openChatMessageOptions.pop();
};

const isEmojiPickerOpen: Ref<boolean> = ref(false);
const toggleEmojiPicker = () => {
    isEmojiPickerOpen.value = !isEmojiPickerOpen.value;
};

function onSelectEmoji(emoji: any) {
    messageText.value += emoji.i;
}

function openSaveMessage() {
    axiosInstance
        .post("/users/get/save-messages/")
        .then((resp) => {
            if (resp.status === 200) {
                Object.assign(messages, resp.data);
            }
        })
        .catch((err) => {
            console.error(err);
        });
}

const userStore = useUserStore();
const chatStore = useChatStore();

onMounted(async () => {
    openSaveMessage();
    window.addEventListener("keydown", handleKeyDown);
});

onUnmounted(() => {
    chatStore.chat.openSaveMessage = false;
});
</script>
<template>
    <div
        @click="handleLayerClick"
        v-if="dialogType"
        class="fixed w-screen h-screen inset-0 z-[200] backdrop-blur-xs flex justify-center items-center"
    >
        <div
            v-if="dialogType === 'clearChat'"
            class="w-[500px] h-[215px] pt-6 pl-6 bg-white z-[501] custom-shadow-inset flex flex-col select-none"
        >
            <h1 class="text-[20px] font-normal mb-4">Clear This chat?</h1>
            <span class="text-[14px] font-normal mb-4 text-gray-500"
                >This chat will be empty but will remain in your chat list</span
            >
            <div class="w-full flex flex-row gap-x-4 mt-auto mb-6">
                <button
                    class="cursor-pointer hover:shadow-2xs rounded-3xl w-[90px] ml-auto text-center h-[38px] bg-white text-[#008069] border border-gray-200"
                >
                    Cancel
                </button>
                <button
                    class="cursor-pointer rounded-3xl w-[114px] mr-8 text-center border-none h-[38px] bg-[#008069] text-white border border-gray-300"
                >
                    Clear chat
                </button>
            </div>
        </div>
        <div
            v-else-if="dialogType === 'blockUser'"
            class="w-[500px] h-[215px] pt-6 pl-6 bg-white z-[501] custom-shadow-inset flex flex-col select-none"
        >
            <h1 class="text-[20px] font-normal mb-4">Block user-name</h1>
            <span class="text-[14px] font-normal mb-4 text-gray-500"
                >Blocked contacts will no longer be able to call you or send you
                messages. This contact will not be notified</span
            >
            <div class="w-full flex flex-row gap-x-4 mt-auto mb-6">
                <button
                    class="cursor-pointer hover:shadow-2xs rounded-3xl w-[90px] ml-auto text-center h-[38px] bg-white text-[14px] text-[#008069] border border-gray-200"
                >
                    Cancel
                </button>
                <button
                    class="cursor-pointer rounded-3xl w-[83px] mr-8 text-center border-none h-[38px] bg-[#008069] text-white border text-[14px] border-gray-300"
                >
                    Block
                </button>
            </div>
        </div>
        <div
            v-else-if="dialogType === 'deleteChat'"
            class="w-[500px] h-[215px] pt-6 pl-6 bg-white z-[501] custom-shadow-inset flex flex-col select-none"
        >
            <h1 class="text-[20px] font-normal mb-4">Delete this chat?</h1>
            <span class="text-[14px] font-normal mb-4 text-gray-500"
                >Chat will be completely deleted from our database and you cant
                backup it again ever</span
            >
            <div class="w-full flex flex-row gap-x-4 mt-auto mb-6">
                <button
                    class="cursor-pointer hover:shadow-2xs rounded-3xl w-[90px] ml-auto text-center h-[38px] bg-white text-[14px] text-[#008069] border border-gray-200"
                >
                    Cancel
                </button>
                <button
                    class="cursor-pointer rounded-3xl w-[110px] mr-8 text-center border-none h-[38px] bg-[#008069] text-white border text-[14px] border-gray-300"
                >
                    Delete chat
                </button>
            </div>
        </div>
        <div
            v-else-if="dialogType === 'muteNotifications'"
            class="w-[500px] h-[215px] pt-6 pl-6 bg-white z-[501] custom-shadow-inset flex flex-col select-none"
        >
            <h1 class="text-[20px] font-normal mb-4">Mute notifications</h1>
            <span class="text-[14px] font-normal mb-4 text-gray-500"
                >No one else in this chat will see that you muted it, and you
                will still be notified if you are mentioned.</span
            >
            <div class="w-full flex flex-row gap-x-4 mt-auto mb-6">
                <button
                    class="cursor-pointer hover:shadow-2xs rounded-3xl w-[90px] ml-auto text-center h-[38px] bg-white text-[14px] text-[#008069] border border-gray-200"
                >
                    Cancel
                </button>
                <button
                    class="cursor-pointer rounded-3xl w-[83px] mr-8 text-center border-none h-[38px] bg-[#008069] text-white border text-[14px] border-gray-300"
                >
                    Mute
                </button>
            </div>
        </div>
    </div>
    <div
        v-if="!isSearchModeOn"
        class="flex flex-row w-full h-[64px] bg-[#f0f2f5] justify-start items-center pl-4 top-0 sticky"
    >
        <img
            class="w-10 h-10 rounded-full"
            :src="userStore.user.profile_url"
            alt=""
        />
        <div class="flex flex-col ml-2">
            <p class="text-[16px] font-normal">Saved Messages</p>
            <span class="text-[13px] font-normal text-gray-500"
                >Last seen on ...</span
            >
        </div>
        <div class="flex flex-row ml-auto mr-6 relative gap-x-3">
            <button
                @click="toggleSearchMode"
                class="flex justify-center items-center w-10 h-10 rounded-full cursor-pointer"
            >
                <span aria-hidden="true" data-icon="search-alt" class=""
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
                        <title>search-alt</title>
                        <path
                            fill="currentColor"
                            d="M15.9,14.3H15L14.7,14c1-1.1,1.6-2.7,1.6-4.3c0-3.7-3-6.7-6.7-6.7S3,6,3,9.7 s3,6.7,6.7,6.7c1.6,0,3.2-0.6,4.3-1.6l0.3,0.3v0.8l5.1,5.1l1.5-1.5L15.9,14.3z M9.7,14.3c-2.6,0-4.6-2.1-4.6-4.6s2.1-4.6,4.6-4.6 s4.6,2.1,4.6,4.6S12.3,14.3,9.7,14.3z"
                        ></path></svg
                ></span>
            </button>
            <button
                @click="toggleChatOption"
                class="flex justify-center items-center w-10 h-10 rounded-full cursor-pointer"
            >
                <span aria-hidden="true" data-icon="menu" class="xr9ek0c"
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
                        <title>menu</title>
                        <path
                            fill="currentColor"
                            d="M12,7c1.104,0,2-0.896,2-2c0-1.105-0.895-2-2-2c-1.104,0-2,0.894-2,2 C10,6.105,10.895,7,12,7z M12,9c-1.104,0-2,0.894-2,2c0,1.104,0.895,2,2,2c1.104,0,2-0.896,2-2C13.999,9.895,13.104,9,12,9z M12,15 c-1.104,0-2,0.894-2,2c0,1.104,0.895,2,2,2c1.104,0,2-0.896,2-2C13.999,15.894,13.104,15,12,15z"
                        ></path></svg
                ></span>
            </button>
            <div
                v-if="isChatOptionOpen"
                class="absolute top-10 right-0 w-[192px] h-auto py-4 flex flex-col bg-white z-50 custom-shadow-inset"
            >
                <div
                    class="cursor-pointer flex items-center pl-6 w-full h-[40px] hover:bg-[#f5f6f6]"
                >
                    <span>Select messages</span>
                </div>
                <div
                    class="cursor-pointer flex items-center pl-6 w-full h-[40px] hover:bg-[#f5f6f6]"
                >
                    <span>Mute notifiations</span>
                </div>
                <div
                    class="cursor-pointer flex items-center pl-6 w-full h-[40px] hover:bg-[#f5f6f6]"
                >
                    <span>Close chat</span>
                </div>
                <div
                    class="cursor-pointer flex items-center pl-6 w-full h-[40px] hover:bg-[#f5f6f6]"
                >
                    <span>Clear Chat</span>
                </div>
                <div
                    class="cursor-pointer flex items-center pl-6 w-full h-[40px] hover:bg-[#f5f6f6]"
                >
                    <span>Delete Chat</span>
                </div>
            </div>
        </div>
    </div>
    <div
        v-else
        class="z-50 flex flex-row w-full h-[64px] bg-[#f0f2f5] justify-start items-center pl-6 top-0 sticky"
    >
        <button
            @click="toggleSearchMode"
            class="w-[24px] h-[24px] mr-4 cursor-pointer"
        >
            <img class="w-[18px] h-[18px]" :src="arrowLeftIcon" alt="" />
        </button>
        <input
            v-model="searchText"
            class="w-[90%] h-[30px] outline-none text-[16px] font-normal"
            type="text"
            placeholder="Search messages..."
        />
        <button
            v-if="searchText.length >= 1"
            @click="closeSearchBar"
            class="w-[24px] h-[24px] cursor-pointer mx-4"
        >
            <svg
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
                <title>x-alt</title>
                <path
                    fill="currentColor"
                    d="M17.25,7.8L16.2,6.75l-4.2,4.2l-4.2-4.2L6.75,7.8l4.2,4.2l-4.2,4.2l1.05,1.05l4.2-4.2l4.2,4.2l1.05-1.05 l-4.2-4.2L17.25,7.8z"
                ></path>
            </svg>
        </button>
        <div
            class="w-full max-h-[330px] overflow-y-auto h-auto bg-white absolute top-13 left-0 z-50"
            style="scrollbar-width: thin"
        >
            <div
                class="w-full border-t border-t-gray-200 h-[55px] pl-4 hover:bg-[#f5f5f5] flex items-center cursor-pointer"
            >
                <p class="text-[16px] font-normal">hey babe</p>
            </div>
            <div
                class="w-full border-t border-t-gray-200 h-[55px] pl-4 hover:bg-[#f5f5f5] flex items-center cursor-pointer"
            >
                <p class="text-[16px] font-normal">hey babe</p>
            </div>
        </div>
    </div>
    <div
        class="w-full h-full max-h-[754px] overflow-y-auto overflow-x-hidden gap-y-2 px-16 py-12 flex flex-col"
        style="
            background-image: url('../../whatsapp_default_background.jpg');
            scrollbar-width: thin;
        "
    >
        <div
            v-for="message in messages"
            :key="message.id"
            class="message-out w-full h-auto flex mb-[1px]"
        >
            <div
                class="bg-white custom-shadow-inset2 pt-1 pb-2 flex flex-col flex-wrap max-w-[603px] w-auto font-normal px-2 h-full min-h-full rounded-2xl items-center"
            >
                <p class="user-text text-[14.2px] flex-wrap">
                    {{ message.msg_content }}
                </p>
                <div
                    class="relative w-full min-h-2 h-2 flex flex-row items-center mt-2"
                >
                    <span
                        class="text-[11px] mr-1 self-end font-normal text-gray-600 flex ml-2"
                        >3.19 PM</span
                    >
                    <button
                        @click="toggleChatMessageOptions(message.id)"
                        class="flex cursor-pointer justify-center mb-2 ml-auto items-center w-[20px] h-[20px]"
                    >
                        <img
                            draggable="false"
                            class="w-full h-full"
                            src="../../arrow-down-icon.svg"
                            alt=""
                        />
                    </button>
                    <div
                        v-if="openChatMessageOptions.includes(message.id)"
                        class="left-[74px] absolute top-4 w-[192px] custom-shadow-inset bg-white h-auto py-3 flex flex-col z-50"
                    >
                        <div
                            class="w-full h-[40px] cursor-pointer hover:bg-[#f5f6f6] pl-6 flex items-center"
                        >
                            <span>Reply</span>
                        </div>
                        <div
                            @click="editText(message.msg_content, message.id)"
                            class="w-full h-[40px] cursor-pointer hover:bg-[#f5f6f6] pl-6 flex items-center"
                        >
                            <span>Edit</span>
                        </div>
                        <div
                            @click="deleteMessage(message.id)"
                            class="w-full h-[40px] cursor-pointer hover:bg-[#f5f6f6] pl-6 flex items-center"
                        >
                            <span>Delete</span>
                        </div>
                        <div
                            class="w-full h-[40px] cursor-pointer hover:bg-[#f5f6f6] pl-6 flex items-center"
                        >
                            <span>Forward</span>
                        </div>
                        <div
                            @click="copyText(message.msg_content)"
                            class="w-full h-[40px] cursor-pointer hover:bg-[#f5f6f6] pl-6 flex items-center"
                        >
                            <span>Copy</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div
        class="w-full min-h-[62px] h-auto flex flex-row bg-[#f0f2f5] items-center"
    >
        <button class="pl-6 cursor-pointer mr-4">
            <svg
                viewBox="0 0 24 24"
                width="30"
                preserveAspectRatio="xMidYMid meet"
                class="x11xpdln x1d8287x x1h4ghdb"
            >
                <title>plus</title>
                <path
                    fill="currentColor"
                    d="M19,13h-6v6h-2v-6H5v-2h6V5h2v6h6V13z"
                ></path>
            </svg>
        </button>
        <div
            v-if="!editModeEnable"
            class="relative flex flex-row bg-white w-[920px] min-h-[44px] h-auto items-center rounded-lg"
        >
            <EmojiPicker
                v-if="isEmojiPickerOpen"
                class="absolute bottom-14 left-0"
                :native="true"
                @select="onSelectEmoji"
            />
            <button
                @click="toggleEmojiPicker"
                class="w-[24px] h-[24px] mx-2 cursor-pointer"
            >
                <svg
                    viewBox="0 0 24 24"
                    height="24"
                    width="24"
                    preserveAspectRatio="xMidYMid meet"
                    class=""
                    fill="none"
                >
                    <title>expressions</title>
                    <path
                        d="M8.49893 10.2521C9.32736 10.2521 9.99893 9.5805 9.99893 8.75208C9.99893 7.92365 9.32736 7.25208 8.49893 7.25208C7.6705 7.25208 6.99893 7.92365 6.99893 8.75208C6.99893 9.5805 7.6705 10.2521 8.49893 10.2521Z"
                        fill="currentColor"
                    ></path>
                    <path
                        d="M17.0011 8.75208C17.0011 9.5805 16.3295 10.2521 15.5011 10.2521C14.6726 10.2521 14.0011 9.5805 14.0011 8.75208C14.0011 7.92365 14.6726 7.25208 15.5011 7.25208C16.3295 7.25208 17.0011 7.92365 17.0011 8.75208Z"
                        fill="currentColor"
                    ></path>
                    <path
                        fill-rule="evenodd"
                        clip-rule="evenodd"
                        d="M16.8221 19.9799C15.5379 21.2537 13.8087 21.9781 12 22H9.27273C5.25611 22 2 18.7439 2 14.7273V9.27273C2 5.25611 5.25611 2 9.27273 2H14.7273C18.7439 2 22 5.25611 22 9.27273V11.8141C22 13.7532 21.2256 15.612 19.8489 16.9776L16.8221 19.9799ZM14.7273 4H9.27273C6.36068 4 4 6.36068 4 9.27273V14.7273C4 17.6393 6.36068 20 9.27273 20H11.3331C11.722 19.8971 12.0081 19.5417 12.0058 19.1204L11.9935 16.8564C11.9933 16.8201 11.9935 16.784 11.9941 16.7479C11.0454 16.7473 10.159 16.514 9.33502 16.0479C8.51002 15.5812 7.84752 14.9479 7.34752 14.1479C7.24752 13.9479 7.25585 13.7479 7.37252 13.5479C7.48919 13.3479 7.66419 13.2479 7.89752 13.2479L13.5939 13.2479C14.4494 12.481 15.5811 12.016 16.8216 12.0208L19.0806 12.0296C19.5817 12.0315 19.9889 11.6259 19.9889 11.1248V9.07648H19.9964C19.8932 6.25535 17.5736 4 14.7273 4ZM14.0057 19.1095C14.0066 19.2605 13.9959 19.4089 13.9744 19.5537C14.5044 19.3124 14.9926 18.9776 15.4136 18.5599L18.4405 15.5576C18.8989 15.1029 19.2653 14.5726 19.5274 13.996C19.3793 14.0187 19.2275 14.0301 19.0729 14.0295L16.8138 14.0208C15.252 14.0147 13.985 15.2837 13.9935 16.8455L14.0057 19.1095Z"
                        fill="currentColor"
                    ></path>
                </svg>
            </button>
            <input
                v-model="messageText"
                class="outline-none border-none w-full pr-2 user-text placeholder:pr-2"
                placeholder="Type a message"
                type="text"
            />
        </div>
        <div
            v-else
            :class="[
                editMessageText.length < 1 ? 'border-2 border-red-700' : '',
            ]"
            class="flex flex-row bg-white w-[920px] min-h-[44px] h-auto items-center rounded-lg"
        >
            <input
                v-model="editMessageText"
                class="outline-none border-none w-full pl-2 user-text placeholder:pr-2"
                placeholder="Edit"
                type="text"
            />
            <button
                @click="submitEditMessage(editMessageText)"
                class="cursor-pointer w-20 h-8 rounded-3xl bg-green-700 text-white text-[16px] font-semibold mr-4"
            >
                Edit
            </button>
        </div>
        <button class="w-[24px] h-[24px] mx-4 cursor-pointer">
            <svg
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
                <title>ptt</title>
                <path
                    fill="currentColor"
                    d="M11.999,14.942c2.001,0,3.531-1.53,3.531-3.531V4.35c0-2.001-1.53-3.531-3.531-3.531 S8.469,2.35,8.469,4.35v7.061C8.469,13.412,9.999,14.942,11.999,14.942z M18.237,11.412c0,3.531-2.942,6.002-6.237,6.002 s-6.237-2.471-6.237-6.002H3.761c0,4.001,3.178,7.297,7.061,7.885v3.884h2.354v-3.884c3.884-0.588,7.061-3.884,7.061-7.885 L18.237,11.412z"
                ></path>
            </svg>
        </button>
    </div>
</template>
