import { reactive } from "vue";
import type { Reactive } from "vue";

const sharedState: Reactive<any> = reactive({
    "backendUrl": "http://localhost:8000",
    "isArchiveChatOpen": false,
    "isNewChatPageOpen": false,
    "isTextStatusCreationOpen": false,
})

export default sharedState