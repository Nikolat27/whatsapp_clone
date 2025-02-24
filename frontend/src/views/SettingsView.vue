<script setup lang="ts">
import UserEditComponent from '../components/UserEditProfileComponent.vue';
</script>
<template>
    <UserEditComponent v-if="true"></UserEditComponent>
    <div v-else class="flex flex-col w-full h-full">
        <h1 class="text-[22px] font-bold pl-3 pt-4">Settings</h1>
        <div class="flex flex-row gap-x-1 ml-3 mt-6 w-[94%] h-[35px] justify-center
         items-center bg-[#f0f2f5] rounded-lg">
            <button class="w-[24px] h-[24px]">
                <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet" class=""
                    version="1.1" x="0px" y="0px" enable-background="new 0 0 24 24">
                    <title>search</title>
                    <path fill="currentColor"
                        d="M15.009,13.805h-0.636l-0.22-0.219c0.781-0.911,1.256-2.092,1.256-3.386 c0-2.876-2.332-5.207-5.207-5.207c-2.876,0-5.208,2.331-5.208,5.207s2.331,5.208,5.208,5.208c1.293,0,2.474-0.474,3.385-1.255 l0.221,0.22v0.635l4.004,3.999l1.194-1.195L15.009,13.805z M10.201,13.805c-1.991,0-3.605-1.614-3.605-3.605 s1.614-3.605,3.605-3.605s3.605,1.614,3.605,3.605S12.192,13.805,10.201,13.805z">
                    </path>
                </svg>
            </button>
            <input type="text" class="w-[90%] h-full outline-none" placeholder="Search settings..">
        </div>

        <div class="flex flex-row w-full h-[98px] mt-4 pl-3 justify-start items-center gap-x-3
            cursor-pointer hover:bg-[#f5f6f6]">
            <img class="w-[82px] h-[82px] rounded-full" src="../../barcelona-logo.jpg" alt="">
            <div class="flex flex-col items-start justify-center">
                <p class="text-[17px] font-normal">Sam</p>
                <span class="text-[14px] font-normal text-gray-600">User`s biography</span>
            </div>
        </div>

        <div class="flex flex-col w-full">
            <!-- <div class="flex flex-row gap-x-4 h-[60px] w-full justify-start items-center pl-5 cursor-pointer
         select-none hover:bg-[#f5f6f6]">
            <span aria-hidden="true" data-icon="account-circle" class=""><svg viewBox="0 0 24 24" height="24" width="24"
                    preserveAspectRatio="xMidYMid meet" class="" fill="none">
                    <title>account-circle</title>
                    <path
                        d="M12 2C6.48 2 2 6.48 2 12C2 17.52 6.48 22 12 22C17.52 22 22 17.52 22 12C22 6.48 17.52 2 12 2ZM12 5C13.66 5 15 6.34 15 8C15 9.66 13.66 11 12 11C10.34 11 9 9.66 9 8C9 6.34 10.34 5 12 5ZM12 19.2C9.5 19.2 7.29 17.92 6 15.98C6.03 13.99 10 12.9 12 12.9C13.99 12.9 17.97 13.99 18 15.98C16.71 17.92 14.5 19.2 12 19.2Z"
                        fill="currentColor"></path>
                </svg></span>
            <p class="font-normal text-[17px]">Account</p>
        </div> -->
            <!-- <div class="flex flex-row gap-x-4 h-[60px] w-full justify-start items-center pl-5 cursor-pointer select-none
         hover:bg-[#f5f6f6]">
            <span aria-hidden="true" data-icon="security-lock" class=""><svg viewBox="0 0 28 35" height="23" width="23"
                    preserveAspectRatio="xMidYMid meet" class="" version="1.1">
                    <title>security-lock</title>
                    <path
                        d="M14,1.10204082 C18.5689011,1.10204082 22.2727273,4.80586698 22.2727273,9.37476809 L22.272,12.1790408 L22.3564837,12.181606 C24.9401306,12.294858 27,14.4253101 27,17.0368705 L27,29.4665309 C27,32.1506346 24.824104,34.3265306 22.1400003,34.3265306 L5.85999974,34.3265306 C3.175896,34.3265306 1,32.1506346 1,29.4665309 L1,17.0368705 C1,14.3970988 3.10461313,12.2488858 5.72742704,12.178644 L5.72727273,9.37476809 C5.72727273,4.80586698 9.43109889,1.10204082 14,1.10204082 Z M14,19.5600907 C12.0418995,19.5600907 10.4545455,21.2128808 10.4545455,23.2517007 C10.4545455,25.2905206 12.0418995,26.9433107 14,26.9433107 C15.9581005,26.9433107 17.5454545,25.2905206 17.5454545,23.2517007 C17.5454545,21.2128808 15.9581005,19.5600907 14,19.5600907 Z M14,4.79365079 C11.4617216,4.79365079 9.39069048,6.79417418 9.27759175,9.30453585 L9.27272727,9.52092352 L9.272,12.1760408 L18.727,12.1760408 L18.7272727,9.52092352 C18.7272727,6.91012289 16.6108006,4.79365079 14,4.79365079 Z"
                        fill="currentColor"></path>
                </svg></span>
            <p class="font-normal text-[17px]">Privacy</p>
        </div>
        <div class="flex flex-row gap-x-4 h-[60px] w-full justify-start items-center pl-5 cursor-pointer
         select-none hover:bg-[#f5f6f6]">
            <span aria-hidden="true" data-icon="chats-filled" class=""><svg viewBox="0 0 24 24" height="24" width="24"
                    preserveAspectRatio="xMidYMid meet" class="" fill="none">
                    <title>chats-filled</title>
                    <path fill-rule="evenodd" clip-rule="evenodd"
                        d="M4.8384 8.45501L5 8.70356V9V17.8333C5 18.7538 5.7462 19.5 6.6667 19.5H20.3333C21.2538 19.5 22 18.7538 22 17.8333V6.66667C22 5.74619 21.2538 5 20.3333 5H2.5927L4.8384 8.45501ZM8 14.5C8 13.6716 8.67157 13 9.5 13H14.5C15.3284 13 16 13.6716 16 14.5C16 15.3284 15.3284 16 14.5 16H9.5C8.67157 16 8 15.3284 8 14.5ZM9.5 8C8.67157 8 8 8.67157 8 9.5C8 10.3284 8.67157 11 9.5 11H16.5C17.3284 11 18 10.3284 18 9.5C18 8.67157 17.3284 8 16.5 8H9.5Z"
                        fill="currentColor"></path>
                    <path
                        d="M5 8.70356L5.41919 8.43101L5.5 8.55531V8.70356H5ZM4.8384 8.45501L4.41921 8.72756L4.41917 8.7275L4.8384 8.45501ZM2.5927 5L2.17347 5.27249L1.67137 4.5H2.5927V5ZM4.58081 8.9761L4.41921 8.72756L5.25759 8.18247L5.41919 8.43101L4.58081 8.9761ZM4.5 9V8.70356H5.5V9H4.5ZM4.5 17.8333V9H5.5V17.8333H4.5ZM6.6667 20C5.47006 20 4.5 19.0299 4.5 17.8333H5.5C5.5 18.4777 6.02234 19 6.6667 19V20ZM20.3333 20H6.6667V19H20.3333V20ZM22.5 17.8333C22.5 19.0299 21.53 20 20.3333 20V19C20.9777 19 21.5 18.4777 21.5 17.8333H22.5ZM22.5 6.66667V17.8333H21.5V6.66667H22.5ZM20.3333 4.5C21.53 4.5 22.5 5.47005 22.5 6.66667H21.5C21.5 6.02233 20.9777 5.5 20.3333 5.5V4.5ZM2.5927 4.5H20.3333V5.5H2.5927V4.5ZM4.41917 8.7275L2.17347 5.27249L3.01192 4.72751L5.25762 8.18252L4.41917 8.7275ZM9.5 13.5C8.94772 13.5 8.5 13.9477 8.5 14.5H7.5C7.5 13.3954 8.39543 12.5 9.5 12.5V13.5ZM14.5 13.5H9.5V12.5H14.5V13.5ZM15.5 14.5C15.5 13.9477 15.0523 13.5 14.5 13.5V12.5C15.6046 12.5 16.5 13.3954 16.5 14.5H15.5ZM14.5 15.5C15.0523 15.5 15.5 15.0523 15.5 14.5H16.5C16.5 15.6046 15.6046 16.5 14.5 16.5V15.5ZM9.5 15.5H14.5V16.5H9.5V15.5ZM8.5 14.5C8.5 15.0523 8.94772 15.5 9.5 15.5V16.5C8.39543 16.5 7.5 15.6046 7.5 14.5H8.5ZM7.5 9.5C7.5 8.39543 8.39543 7.5 9.5 7.5V8.5C8.94772 8.5 8.5 8.94772 8.5 9.5H7.5ZM9.5 11.5C8.39543 11.5 7.5 10.6046 7.5 9.5H8.5C8.5 10.0523 8.94772 10.5 9.5 10.5V11.5ZM16.5 11.5H9.5V10.5H16.5V11.5ZM18.5 9.5C18.5 10.6046 17.6046 11.5 16.5 11.5V10.5C17.0523 10.5 17.5 10.0523 17.5 9.5H18.5ZM16.5 7.5C17.6046 7.5 18.5 8.39543 18.5 9.5H17.5C17.5 8.94772 17.0523 8.5 16.5 8.5V7.5ZM9.5 7.5H16.5V8.5H9.5V7.5Z"
                        fill="currentColor"></path>
                </svg></span>
            <p class="font-normal text-[17px]">Chats</p>
        </div>
        <div class="flex flex-row gap-x-4 h-[60px] w-full justify-start items-center pl-5 cursor-pointer
         select-none hover:bg-[#f5f6f6]">
            <span aria-hidden="true" data-icon="settings-notifications" class=""><svg viewBox="0 0 24 24" height="27"
                    width="27" preserveAspectRatio="xMidYMid meet" class="" version="1.1" x="0px" y="0px"
                    enable-background="new 0 0 24 24">
                    <title>settings-notifications</title>
                    <path fill="currentColor"
                        d="M12,21.7c0.9,0,1.7-0.8,1.7-1.7h-3.4C10.3,20.9,11.1,21.7,12,21.7z M17.6,16.5v-4.7 c0-2.7-1.8-4.8-4.3-5.4V5.8c0-0.7-0.6-1.3-1.3-1.3s-1.3,0.6-1.3,1.3v0.6C8.2,7,6.4,9.1,6.4,11.8v4.7l-1.7,1.7v0.9h14.6v-0.9 L17.6,16.5z">
                    </path>
                </svg></span>
            <p class="font-normal text-[17px]">Notifications</p>
        </div> -->
            <div class="flex flex-row gap-x-4 h-[60px] w-full justify-start items-center pl-5 cursor-pointer
         select-none hover:bg-[#f5f6f6]">
                <span aria-hidden="true" data-icon="exit" class="text-red-600"><svg viewBox="0 0 24 24" height="24"
                        width="24" preserveAspectRatio="xMidYMid meet" class="" version="1.1" x="0px" y="0px"
                        enable-background="new 0 0 24 24">
                        <title>exit</title>
                        <path fill="currentColor"
                            d="M16.6,8.1l1.2-1.2l5.1,5.1l-5.1,5.1l-1.2-1.2l3-3H8.7v-1.8h10.9L16.6,8.1z M3.8,19.9h9.1 c1,0,1.8-0.8,1.8-1.8v-1.4h-1.8v1.4H3.8V5.8h9.1v1.4h1.8V5.8c0-1-0.8-1.8-1.8-1.8H3.8C2.8,4,2,4.8,2,5.8v12.4 C2,19.1,2.8,19.9,3.8,19.9z">
                        </path>
                    </svg></span>
                <p class="font-normal text-[17px] text-red-600">Log out</p>
            </div>
        </div>
    </div>
</template>