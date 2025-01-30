<script lang="ts">
  let {
    label,
    value = $bindable(),
    expertMode,
    expertOnly,
  }: {
    label: string;
    value: string;
    expertMode: boolean;
    expertOnly: boolean;
  } = $props();

  const useExpertMode = $derived(
    expertMode || !value || !value.startsWith("#") || value.length !== 7
  );
</script>

{#if !expertOnly || (expertOnly && expertMode)}
  <tr class="theme-setting">
    <td>
      <label for="input">{label}:</label>
    </td>
    <td>
      <input id="input" type={useExpertMode ? "text" : "color"} bind:value />
    </td>
  </tr>
{/if}
