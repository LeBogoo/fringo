<script lang="ts">
  import type { Theme } from "./Theme";

  let { themes }: { themes: Theme[] } = $props();

  let themeName = $state(localStorage.getItem("theme") || themes[0].name);

  const theme = $derived(themes.find((t) => t.name === themeName));

  function applyTheme(theme: Theme) {
    console.log("Applying theme", themeName);
    localStorage.setItem("theme", theme.name);

    document.documentElement.style.setProperty(
      "--theme-background-color",
      theme.backgroundColor
    );
    document.documentElement.style.setProperty(
      "--theme-text-color",
      theme.textColor
    );
    document.documentElement.style.setProperty(
      "--theme-grid-color",
      theme.grid.color
    );
    document.documentElement.style.setProperty(
      "--theme-field-default-color",
      theme.field.default.color
    );
    document.documentElement.style.setProperty(
      "--theme-field-default-hover-color",
      theme.field.default.hoverColor
    );
    document.documentElement.style.setProperty(
      "--theme-field-default-text-color",
      theme.field.default.textColor || ""
    );
    document.documentElement.style.setProperty(
      "--theme-field-default-text-outline-color",
      theme.field.default.textOutlineColor || ""
    );
    document.documentElement.style.setProperty(
      "--theme-field-checked-color",
      theme.field.checked.color
    );
    document.documentElement.style.setProperty(
      "--theme-field-checked-hover-color",
      theme.field.checked.hoverColor
    );
    document.documentElement.style.setProperty(
      "--theme-field-checked-text-color",
      theme.field.checked.textColor || ""
    );
    document.documentElement.style.setProperty(
      "--theme-field-checked-text-outline-color",
      theme.field.checked.textOutlineColor || ""
    );
    document.documentElement.style.setProperty(
      "--theme-field-disabled-color",
      theme.field.disabled.color
    );
    document.documentElement.style.setProperty(
      "--theme-field-disabled-hover-color",
      theme.field.disabled.hoverColor
    );
    document.documentElement.style.setProperty(
      "--theme-field-disabled-text-color",
      theme.field.disabled.textColor || ""
    );
    document.documentElement.style.setProperty(
      "--theme-field-disabled-text-outline-color",
      theme.field.disabled.textOutlineColor || ""
    );
    document.documentElement.style.setProperty(
      "--theme-field-gay-color",
      theme.field.gay.color
    );
    document.documentElement.style.setProperty(
      "--theme-field-gay-hover-color",
      theme.field.gay.hoverColor
    );
    document.documentElement.style.setProperty(
      "--theme-field-gay-text-color",
      theme.field.gay.textColor || ""
    );
    document.documentElement.style.setProperty(
      "--theme-field-gay-text-outline-color",
      theme.field.gay.textOutlineColor || ""
    );
    document.documentElement.style.setProperty(
      "--theme-button-color",
      theme.button.color
    );
    document.documentElement.style.setProperty(
      "--theme-button-hover-color",
      theme.button.hoverColor
    );
    document.documentElement.style.setProperty(
      "--theme-button-text-color",
      theme.button.textColor || ""
    );
    document.documentElement.style.setProperty(
      "--theme-sidebar-color",
      theme.sidebar.color
    );
    document.documentElement.style.setProperty(
      "--theme-sidebar-background-blur",
      theme.sidebar.backgroundBlur
    );
  }

  $effect(() => {
    applyTheme(theme);
  });
</script>

<div class="theme-list">
  <h3>Themes</h3>
  {#each themes as theme (theme.name)}
    <label>
      <input
        type="radio"
        name="theme"
        value={theme.name}
        bind:group={themeName}
      />
      {theme.name}
    </label>
  {/each}
</div>

<style>
  label {
    display: block;
    margin: 0.5rem 0;
  }

  input {
    margin-right: 0.5rem;
  }

  .theme-list {
    padding: 1rem;
    background-color: var(--theme-background-color);
    border-radius: 1rem;
  }
</style>
