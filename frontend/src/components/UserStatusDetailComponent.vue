<template>
    <div class="fixed inset-0 justify-start items-start flex flex-col w-screen h-screen flex flex-col bg-black z-[999]">
        <div class="flex flex-row w-full pt-6 pl-8">
            <button>
                <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet" class="text-white"
                    version="1.1" x="0px" y="0px" enable-background="new 0 0 24 24">
                    <title>back</title>
                    <path fill="currentColor" d="M12,4l1.4,1.4L7.8,11H20v2H7.8l5.6,5.6L12,20l-8-8L12,4z"></path>
                </svg>
            </button>
            <button class="ml-auto mr-12">
                <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet" class="text-white"
                    version="1.1" x="0px" y="0px" enable-background="new 0 0 24 24">
                    <title>x-viewer</title>
                    <path fill="currentColor"
                        d="M19.8,5.8l-1.6-1.6L12,10.4L5.8,4.2L4.2,5.8l6.2,6.2l-6.2,6.2l1.6,1.6l6.2-6.2l6.2,6.2l1.6-1.6L13.6,12 L19.8,5.8z">
                    </path>
                </svg>
            </button>
        </div>
        <div class="my-auto flex flex-row w-full pl-8">
            <button class="cursor-pointer w-[45px] h-[45px] flex justify-center items-center rounded-full">
                <img draggable="false" class="w-full h-full" src="..\assets\icons\svg-icons\icons8-arrow-left-52.png"
                    alt="">
            </button>
            <button class="cursor-pointer w-[45px] h-[45px] flex justify-center items-center rounded-full
             ml-auto mr-12">
                <img draggable="false" class="w-full h-full rotate-180"
                    src="..\assets\icons\svg-icons\icons8-arrow-left-52.png" alt="">
            </button>
        </div>
        <div class="h-full justify-start items-center overflow-hidden w-[49%] absolute top-1/2 left-1/2 bg-white flex flex-col"
            style="transform: translate(-50%, -50%);">
            <img class="w-full h-full inset-0 fixed" src="../../background-image.jpg" alt="">
            <div class="flex flex-col items-center justify-center w-[55.55%] z-50 mt-10">
                <div class="w-full flex flex-row items-center justify-center gap-x-1">
                    <div class="w-[51.55px] h-[7px] bg-white rounded-2xl mb-4"></div>
                    <div class="w-[51.55px] h-[7px] bg-white rounded-2xl mb-4"></div>
                    <div class="w-[51.55px] h-[7px] bg-white rounded-2xl mb-4"></div>
                    <div class="w-[51.55px] h-[7px] bg-white rounded-2xl mb-4"></div>
                </div>
                <div class="flex flex-row w-full items-center justify-start">
                    <img draggable="false" class="w-10 h-10 rounded-full cursor-pointer" src="../../barcelona-logo.jpg" alt="">
                    <div class="flex flex-col text-white ml-2 select-none cursor-pointer">
                        <span class="text-[16px] font-normal">Friend...</span>
                        <span class="text-[14px] font-normal text-gray-700">Today at 10:15 AM</span>
                    </div>
                    <button class="w-[24px] h-[24px] ml-auto cursor-pointer">
                        <svg viewBox="0 0 24 24" height="24" width="24" preserveAspectRatio="xMidYMid meet"
                            class="text-white" fill="none">
                            <title>status-media-controls-play</title>
                            <path
                                d="M8 6.81976V17.1798C8 17.9698 8.87 18.4498 9.54 18.0198L17.68 12.8398C18.3 12.4498 18.3 11.5498 17.68 11.1498L9.54 5.97976C8.87 5.54976 8 6.02976 8 6.81976Z"
                                fill="currentColor"></path>
                        </svg>
                    </button>
                    <!-- <button class="w-[12px] h-[12px] ml-auto cursor-pointer">
                        <img class="w-full h-full" src="..\assets\icons\svg-icons\pause-icon.png" alt="">
                    </button> -->
                </div>
            </div>
        </div>
    </div>
</template>