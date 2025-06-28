<script setup lang="ts">
</script>
<template>
    <div class="flex flex-row w-full pt-4 pl-4">
        <h1 class="text-[22px] font-bold">Channels</h1>
        <button class="w-10 h-10 rounded-full flex justify-center items-center cursor-pointer hover:bg-[#f2f2f2]
         ml-auto mr-4">
            <span aria-hidden="true" data-icon="plus" class=""><svg viewBox="0 0 24 24" height="24" width="24"
                    preserveAspectRatio="xMidYMid meet" class="">
                    <title>plus</title>
                    <path fill="currentColor" d="M19,13h-6v6h-2v-6H5v-2h6V5h2v6h6V13z"></path>
                </svg></span>
        </button>
    </div>
    <div class="flex flex-row gap-x-1 ml-3 mt-6 w-[94%] h-[35px] justify-center items-center bg-[#f0f2f5] rounded-lg">
        <button class="w-[24px] h-[24px] ml-2 cursor-pointer">
            <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet" class="" version="1.1"
                x="0px" y="0px" enable-background="new 0 0 24 24">
                <title>search</title>
                <path fill="currentColor"
                    d="M15.009,13.805h-0.636l-0.22-0.219c0.781-0.911,1.256-2.092,1.256-3.386 c0-2.876-2.332-5.207-5.207-5.207c-2.876,0-5.208,2.331-5.208,5.207s2.331,5.208,5.208,5.208c1.293,0,2.474-0.474,3.385-1.255 l0.221,0.22v0.635l4.004,3.999l1.194-1.195L15.009,13.805z M10.201,13.805c-1.991,0-3.605-1.614-3.605-3.605 s1.614-3.605,3.605-3.605s3.605,1.614,3.605,3.605S12.192,13.805,10.201,13.805z">
                </path>
            </svg>
        </button>
        <input type="text" class="ml-4 w-[90%] h-full outline-none text-[14px] text-gray-600 placeholder:text-[14px]"
            placeholder="Search">
    </div>
    <div class="flex flex-row pl-4 items-center cursor-pointer w-full mt-1 gap-x-3 hover:bg-[#f5f6f6] h-[72px]">
        <button class="min-w-[49px] w-[49px] h-[49px] rounded-full overflow-hidden">
            <img class="w-full h-full" src="../../barcelona-logo.jpg" alt="">
        </button>
        <div class="flex flex-col w-full">
            <div class="flex flex-row w-full items-center">
                <p class="font-normal text-[17px] text-black">FC Bercelona</p>
                <span class="font-normal text-[12px] text-gray-600 ml-auto mr-4">Yesterday</span>
            </div>
            <div class="flex flex-row w-full items-center">
                <p class="font-normal text-[13px] text-gray-600">last message</p>
                <div class="ml-auto mr-4 w-5 h-5 rounded-full bg-[#25d366] text-white text-[12px]
                 font-semibold text-center">
                    <span>2</span>
                </div>
            </div>
        </div>
    </div>
</template>