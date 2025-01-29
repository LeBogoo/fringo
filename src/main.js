import { mount } from 'svelte'
import 'remixicon/fonts/remixicon.css'
import './app.scss'
import './css/field.scss'
import './css/button.scss'
import App from './App.svelte'
const app = mount(App, {
  target: document.getElementById('app'),
})

export default app
