<script>
  let { active, side, ...props } = $props();
</script>

<button
  {...props}
  class="sidebar-toggle {side}"
  class:active
  aria-label="Toggle sidebar"
>
  {#if side === "left"}
    <i class="ri-arrow-right-line"></i>
  {:else}
    <i class="ri-arrow-left-line"></i>
  {/if}
</button>
