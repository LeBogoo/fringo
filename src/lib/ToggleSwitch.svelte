<script lang="ts">
  import type { Snippet } from "svelte";

  let {
    children,
    checked = $bindable(),
  }: {
    children: Snippet;
    checked: boolean;
  } = $props();
</script>

<label class="input-label" for="switch">
  {@render children?.()}:
</label>
<label class="switch">
  <input type="checkbox" id="switch" bind:checked />
  <span class="slider round"></span>
</label>

<style>
  :root {
    --slider-size: 2.5rem;
  }

  .switch {
    position: relative;
    display: inline-block;
    width: var(--slider-size);
    height: calc(var(--slider-size) / 2);
  }

  .switch input {
    opacity: 0;
    width: 0;
    height: 0;
  }

  .slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: var(--theme-slider-color);
    transition: 0.3s;
  }

  .slider:before {
    position: absolute;
    content: "";
    height: calc(var(--slider-size) / 2);
    width: calc(var(--slider-size) / 2);
    left: 0;
    bottom: 0;
    background-color: var(--theme-slider-color-indicator);
    transition: 0.3s;
  }

  input:checked + .slider {
    background-color: var(--theme-slider-color-active);
  }

  input:focus + .slider {
    box-shadow: 0 0 1px var(--theme-slider-color-active);
  }

  input:checked + .slider:before {
    transform: translateX(calc(var(--slider-size) / 2));
  }

  .slider.round {
    border-radius: 100px;
  }

  .slider.round:before {
    border-radius: 50%;
  }
</style>
