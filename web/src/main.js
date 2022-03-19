import { createApp } from "vue";
import App from "./App.vue";
// element-plus
import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'

import store from "@/store";
import router from '@/router'

const app = createApp(App);
app
    .use(store)
    .use(router)
    .use(ElementPlus, { locale: zhCn })
    .mount("#app");
