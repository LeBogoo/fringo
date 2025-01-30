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

<style>
  .sidebar-toggle {
    position: fixed;
    top: 0.5rem;
    z-index: 1000;
    background: transparent;
    color: var(--theme-text-color);
    border: none;
    padding: 0.5rem;
    font-size: 2rem;
    cursor: pointer;
  }

  .sidebar-toggle.left {
    left: 0.5rem;
    transition:
      left 0.3s ease-in-out,
      transform 0.3s;
  }

  .sidebar-toggle.right {
    right: 0.5rem;
    transition:
      right 0.3s ease-in-out,
      transform 0.3s;
  }

  .sidebar-toggle.left.active {
    transform: rotate(180deg);
    left: calc(max(25vw, 400px) - 4rem);
  }

  .sidebar-toggle.right.active {
    right: calc(max(25vw, 400px) - 4rem);
    transform: rotate(-180deg);
  }

  @media (max-width: 750px) {
    .sidebar-toggle.left.active {
      left: calc(90vw - 4rem);
    }

    .sidebar-toggle.right.active {
      right: calc(90vw - 4rem);
    }
  }
</style>
