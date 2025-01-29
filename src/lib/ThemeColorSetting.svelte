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

<style>
  :global(.theme-setting input[type="text"]) {
    background-color: white;
    color: black;
    border: 1px solid var(--theme-text-color);
    padding: 0.2rem;
    border-radius: 5px;
  }
</style>
