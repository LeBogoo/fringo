<script>
  let theme = $state(localStorage.getItem("theme") || "light");

  $effect(() => {
    document.documentElement.setAttribute("data-theme", theme);
    localStorage.setItem("theme", theme);
  });
</script>

{#if theme == "light"}
  <button aria-label="Light mode" onclick={() => (theme = "dark")}>
    <i class="ri-sun-fill"></i>
  </button>
{:else}
  <button aria-label="Dark mode" onclick={() => (theme = "light")}>
    <i class="ri-moon-fill"></i>
  </button>
{/if}

<style>
  button {
    position: absolute;
    top: 0;
    right: 0;
    background: none;
    border: none;
    cursor: pointer;
    font-size: 3rem;
    padding: 0.5rem;
  }

  .ri-moon-fill {
    color: white;
  }
</style>
