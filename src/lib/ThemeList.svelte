<script lang="ts">
  import type { Theme } from "./Theme";
  import { themes } from "../themes.svelte.js";
  import ThemeEditor from "./theme-editor/ThemeEditor.svelte";

  let customThemes: Theme[] = $state(
    JSON.parse(localStorage.getItem("custom-themes") || "[]")
  );

  let allThemes = $derived([...themes, ...customThemes]);

  let themeId = $state(localStorage.getItem("theme") || themes[0].id);

  let currentTheme = $state(null);

  $effect(() => {
    currentTheme = [...themes, ...customThemes].find((t) => t.id === themeId);
  });

  let isEditorOpen = $derived(currentTheme?.type == "custom");

  function applyTheme(theme: Theme) {
    if (!theme) {
      console.log(`Theme ${themeId} is in the deprecated format. Migrating...`);

      let newTheme: Theme;
      switch (themeId) {
        case "dark":
          newTheme = themes[1];
          break;
        default:
          newTheme = themes[0];
          break;
      }
      themeId = newTheme.id;
      return;
    }
    localStorage.setItem("theme", theme.id);
    localStorage.setItem("custom-themes", JSON.stringify(customThemes));

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

    document.documentElement.style.setProperty(
      "--theme-slider-color",
      theme.input.color
    );
    document.documentElement.style.setProperty(
      "--theme-slider-color-active",
      theme.input.activeColor
    );
    document.documentElement.style.setProperty(
      "--theme-slider-color-indicator",
      theme.input.indicatorColor
    );
  }

  $effect(() => {
    applyTheme(currentTheme);
  });

  function newCustomTheme() {
    let newTheme: Theme = JSON.parse(JSON.stringify(currentTheme));
    newTheme.type = "custom";
    let untitledThemeCounter = 1;
    while (
      customThemes.find(
        (t) => t.name === `Untitled Theme ${untitledThemeCounter}`
      )
    ) {
      untitledThemeCounter++;
    }
    newTheme.name = `Untitled Theme ${untitledThemeCounter}`;
    newTheme.id = crypto.randomUUID();

    customThemes.push(newTheme);

    themeId = newTheme.id;
  }

  function loadCustomTheme() {
    // open prompt to paste theme json
    let themeJson = prompt("Paste your theme data here:");
    if (!themeJson) {
      alert("Error: No theme data.");
      return;
    }

    try {
      let newTheme: Theme = JSON.parse(themeJson);
      if (!newTheme.id) {
        alert("Error: Invalid theme data.");
        return;
      }
      newTheme.type = "custom";

      // check if the id already exist. if so, replace the existing one
      let existingTheme = customThemes.find((t) => t.id === newTheme.id);
      if (existingTheme) {
        let index = customThemes.indexOf(existingTheme);
        customThemes[index] = newTheme;
      } else customThemes.push(newTheme);

      themeId = newTheme.id;
    } catch (e) {
      alert("Error: Invalid theme data.");
      return;
    }
  }

  function shareTheme(theme: Theme) {
    navigator.clipboard.writeText(JSON.stringify(theme));

    setTimeout(() => {
      alert("Theme copied to clipboard!");
    }, 10);
  }

  function delteTheme(theme: Theme) {
    let index = customThemes.indexOf(theme);
    customThemes.splice(index, 1);
    localStorage.setItem("custom-themes", JSON.stringify(customThemes));

    index += themes.length;

    if (themeId == theme.id) themeId = allThemes[index - 1].id;
  }
</script>

<div class="sidebar-panel">
  <h3>Themes</h3>
  {#each allThemes as theme (theme.id)}
    <label>
      <input type="radio" name="theme" value={theme.id} bind:group={themeId} />
      {theme.name}

      {#if theme.type === "custom"}
        <button
          class="icon-button"
          aria-label="Share Theme"
          onclick={() => shareTheme(theme)}
        >
          <i class="ri-share-line"></i>
        </button>

        <button
          class="icon-button"
          aria-label="Delete Theme"
          onclick={() => delteTheme(theme)}
        >
          <i class="ri-delete-bin-6-line"></i>
        </button>
      {/if}
    </label>
  {/each}

  <div class="btn-group">
    <button class="btn btn-sm" onclick={() => newCustomTheme()}>
      <i class="ri-add-line"></i> New Custom Theme
    </button>
    <button class="btn btn-sm" onclick={() => loadCustomTheme()}>
      <i class="ri-upload-2-line"></i> Load Custom Theme
    </button>
  </div>

  <small class="theme-storage-info">
    Custom Themes are stored in your browser.<br />No manual saving is required.
  </small>
</div>

{#if isEditorOpen}
  <ThemeEditor bind:theme={currentTheme}></ThemeEditor>
{/if}

<style>
  label {
    display: block;
    margin: 0.5rem 0;
    height: 100%;
  }

  input {
    margin-right: 0.5rem;
  }

  .theme-storage-info {
    margin-top: 1rem;
    display: block;
    font-size: 0.8rem;
    letter-spacing: 0px;
    text-align: center;
  }

  .icon-button {
    background: none;
    color: var(--theme-button-text-color);
    border: none;
    font-size: 1.2rem;
    cursor: pointer;
  }
</style>
