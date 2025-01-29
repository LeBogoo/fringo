<script lang="ts">
  import type { ButtonColors } from "./Theme";
  import ThemeColorSetting from "./ThemeColorSetting.svelte";

  let {
    label,
    value = $bindable(),
    expertMode,
  }: {
    label: string;
    value: ButtonColors;
    expertMode: boolean;
  } = $props();
</script>

<tr>
  <td colspan="2">
    <h4>{label}</h4>
  </td>
</tr>
<ThemeColorSetting
  label={"Color"}
  expertOnly={false}
  {expertMode}
  bind:value={value.color}
/>
<ThemeColorSetting
  label={"Hover Color"}
  expertOnly={false}
  {expertMode}
  bind:value={value.hoverColor}
/>
<ThemeColorSetting
  label={"Text Color"}
  expertOnly={true}
  {expertMode}
  bind:value={value.textColor}
/>
<ThemeColorSetting
  label={"Text Outline Color"}
  expertOnly={true}
  {expertMode}
  bind:value={value.textOutlineColor}
/>
