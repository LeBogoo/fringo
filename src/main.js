import { mount } from "svelte";
import "remixicon/fonts/remixicon.css";
import "./css/app.scss";
import "./css/field.scss";
import "./css/button.scss";
import "./css/sidebar.scss";
import App from "./App.svelte";
const app = mount(App, {
  target: document.getElementById("app"),
});

export default app;
