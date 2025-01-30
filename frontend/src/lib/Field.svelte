<script lang="ts">
  let {
    text,
    url,
    checked,
    disabled,
    onclick,
  }: {
    text?: string;
    url?: string;
    checked?: boolean;
    disabled?: boolean;
    onclick?: () => void;
  } = $props();
</script>

{#if url}
  <img src={url} alt="Funny Fringo" class="field" />
{:else}
  <button {onclick} class="field" class:checked class:gay={text == "Gay"}>
    {text}
  </button>
{/if}

<style>
  .field {
    user-select: none;
    text-align: center;
    font-weight: 700;
    display: inline-block;
    padding: 5px;
    width: 100%;
    aspect-ratio: 1;
    color: var(--theme-field-default-text-color, var(--theme-text-color));
    background: var(--theme-field-default-color);
    text-shadow:
      -1px -1px 0 var(--theme-field-default-text-outline-color),
      1px -1px 0 var(--theme-field-default-text-outline-color),
      -1px 1px 0 var(--theme-field-default-text-outline-color),
      1px 1px 0 var(--theme-field-default-text-outline-color);
    border-radius: 10px;
    font-size: 2vmin;
  }

  img.field {
    padding: 0;
  }

  button.field {
    border: none;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  button.field.checked {
    background: var(--theme-field-checked-color);
  }

  button.field.checked:hover {
    background: var(--theme-field-checked-hover-color);
  }

  button.field.gay {
    background: var(--theme-field-gay-color);

    text-shadow:
      -1px -1px 0 var(--theme-field-gay-text-outline-color),
      1px -1px 0 var(--theme-field-gay-text-outline-color),
      -1px 1px 0 var(--theme-field-gay-text-outline-color),
      1px 1px 0 var(--theme-field-gay-text-outline-color);
  }

  button.field.gay:hover {
    cursor: not-allowed;
    background: var(
      --theme-field-gay-hover-color,
      var(--theme-field-default-hover-color)
    );
  }

  button.field.gay:disabled,
  button.field:disabled {
    background: none;
    text-shadow: none;
    background: var(--theme-field-disabled-color);
    cursor: not-allowed;
  }

  button.field:hover {
    background: var(--theme-field-default-hover-color);
  }
</style>
