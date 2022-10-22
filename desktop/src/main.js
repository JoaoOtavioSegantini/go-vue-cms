import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import "bootstrap/dist/css/bootstrap.css";
import "animate.css/animate.css";

import "./axios";
import "./ckeditor";

createApp(App).use(router).use(store).mount("#app");
