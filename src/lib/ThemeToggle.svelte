<script>
  let theme = $state(localStorage.getItem("theme") || "light");

  $effect(() => {
    document.documentElement.setAttribute("data-theme", theme);
    localStorage.setItem("theme", theme);
  });
</script>

{#if theme == "light"}
  <button
    class="theme-toggle"
    aria-label="Light mode"
    onclick={() => (theme = "dark")}
  >
    <i class="ri-sun-fill"></i>
  </button>
{:else}
  <button
    class="theme-toggle"
    aria-label="Dark mode"
    onclick={() => (theme = "light")}
  >
    <i class="ri-moon-fill"></i>
  </button>
{/if}
