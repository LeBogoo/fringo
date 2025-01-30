import { mount } from "svelte";
import "remixicon/fonts/remixicon.css";
import "./css/app.scss";
import "./css/theme-setting.scss";
import "./css/button.scss";
import App from "./App.svelte";
const app = mount(App, {
  target: document.getElementById("app"),
});

export default app;
