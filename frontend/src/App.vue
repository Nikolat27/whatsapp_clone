<script setup lang="ts">
import SideBarComponent from './components/SideBarComponent.vue';
import ChatPageComponent from './components/ChatPageComponent.vue';

</script>
<template>
  <div class="w-[83.3%] h-full absolute top-[28px] left-[110px]">
    <div class="flex flex-row w-full h-[94%] custom-shadow-inset">
      <div class="flex flex-col flex-basis-[58.8px] flex-shrink-0 min-w-[58.8px] w-[4%] h-full
       bg-[#f0f2f5] justify-start items-center pt-2 gap-y-4 border-rborder-gray-300">
        <SideBarComponent></SideBarComponent>
      </div>
      <div class="flex flex-col w-[30%] h-full bg-white">
        <router-view></router-view>
      </div>
      <div class="flex flex-col w-[66%] h-full relative">
        <ChatPageComponent></ChatPageComponent>
      </div>
    </div>
  </div>
</template>
