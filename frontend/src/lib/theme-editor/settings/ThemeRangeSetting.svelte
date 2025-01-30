<script lang="ts">
  let {
    label,
    value = $bindable(),
    min,
    max,
    step,
    expertMode,
    expertOnly,
  }: {
    label: string;
    value: string;
    min: string;
    max: string;
    step: string;
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
      <input
        id="input"
        type={expertMode ? "number" : "range"}
        bind:value
        {min}
        {max}
        {step}
      />
    </td>
  </tr>
{/if}
