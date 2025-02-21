import { reactive } from "vue";
import type { Reactive } from "vue";

const sharedState: Reactive<any> = reactive({
    "isArchiveChatOpen": false,
})

export default sharedState