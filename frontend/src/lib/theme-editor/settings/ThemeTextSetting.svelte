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
</script>

{#if !expertOnly || (expertOnly && expertMode)}
  <tr class="theme-setting">
    <td>
      <label for="input">{label}:</label>
    </td>
    <td>
      <input id="input" type="text" bind:value />
    </td>
  </tr>
{/if}
