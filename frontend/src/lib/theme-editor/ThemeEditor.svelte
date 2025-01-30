<script lang="ts">
  import type { Theme } from "../Theme";

  import ToggleSwitch from "../ToggleSwitch.svelte";

  import ThemeColorSetting from "./settings/ThemeColorSetting.svelte";
  import ThemeButtonColorsSetting from "./settings/ThemeButtonColorsSetting.svelte";
  import ThemeTextSetting from "./settings/ThemeTextSetting.svelte";

  const { theme = $bindable() }: { theme: Theme } = $props();

  let expertMode = $state(false);
</script>

<div class="theme-editor sidebar-panel">
  <h3>Theme Editor</h3>
  <div class="editor-settings">
    <ToggleSwitch bind:checked={expertMode}>Expert Mode</ToggleSwitch>
  </div>
  <hr />

  <table>
    <tbody>
      <tr>
        <td colspan="2">
          <h4>General</h4>
        </td>
      </tr>
      <ThemeTextSetting
        label="Theme Name"
        {expertMode}
        expertOnly={false}
        bind:value={theme.name}
      />
      <ThemeTextSetting
        label="Author"
        {expertMode}
        expertOnly={false}
        bind:value={theme.author}
      />
      <ThemeColorSetting
        label={"Text Color"}
        expertOnly={false}
        {expertMode}
        bind:value={theme.textColor}
      />
      <ThemeColorSetting
        label={"Background Color"}
        expertOnly={false}
        {expertMode}
        bind:value={theme.backgroundColor}
      />
      <ThemeColorSetting
        label={"Grid Color"}
        expertOnly={false}
        {expertMode}
        bind:value={theme.grid.color}
      />
      <ThemeButtonColorsSetting
        label="Field - Default"
        {expertMode}
        bind:value={theme.field.default}
      />

      <ThemeButtonColorsSetting
        label="Field - Checked"
        {expertMode}
        bind:value={theme.field.checked}
      />

      <ThemeButtonColorsSetting
        label="Field - Gay"
        {expertMode}
        bind:value={theme.field.gay}
      />

      <ThemeButtonColorsSetting
        label="Button"
        {expertMode}
        bind:value={theme.button}
      />

      <tr>
        <td colspan="2">
          <h4>Sidebar</h4>
        </td>
      </tr>

      <ThemeColorSetting
        label="Color"
        {expertMode}
        expertOnly={false}
        bind:value={theme.sidebar.color}
      />

      <ThemeColorSetting
        label="Background Blur"
        {expertMode}
        expertOnly={false}
        bind:value={theme.sidebar.backgroundBlur}
      />

      <tr>
        <td colspan="2">
          <h4>Inputs</h4>
        </td>
      </tr>

      <ThemeColorSetting
        label="Color"
        {expertMode}
        expertOnly={false}
        bind:value={theme.input.color}
      />

      <ThemeColorSetting
        label="Active Color"
        {expertMode}
        expertOnly={false}
        bind:value={theme.input.activeColor}
      />

      <ThemeColorSetting
        label="Indicator Color"
        {expertMode}
        expertOnly={false}
        bind:value={theme.input.indicatorColor}
      />
    </tbody>
  </table>
</div>

<style>
  .editor-settings {
    width: max-content;
    margin: 0.5rem auto;
  }

  hr {
    margin: 0.5rem 0;
    border: none;
    border-top: 1px solid var(--theme-text-color);
  }
</style>
