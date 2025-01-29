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

<style>
  input[type="range"] {
    /* custom range input */
    appearance: none;
    width: 100%;
    height: 1.25rem;
    background: var(--theme-slider-color);
    border-radius: 10px;
    outline: none;
    opacity: 0.7;
    transition: opacity 0.2s;
  }

  input[type="range"]:hover {
    opacity: 1;
  }

  input[type="range"]::-webkit-slider-thumb {
    /* custom thumb */
    appearance: none;
    width: 1rem;
    height: 1rem;
    background: white;
    border-radius: 50%;
    cursor: pointer;
  }

  input[type="range"]::-moz-range-thumb {
    /* custom thumb */
    width: 1rem;
    height: 1rem;
    background: white;
    border-radius: 50%;
    cursor: pointer;
    border: none;
  }

  :global(.theme-setting input[type="number"]) {
    background-color: white;
    color: black;
    border: 1px solid var(--theme-text-color);
    padding: 0.2rem;
    border-radius: 5px;
  }
</style>
