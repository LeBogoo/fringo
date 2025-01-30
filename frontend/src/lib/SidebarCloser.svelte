<script>
  let { active, side, ...props } = $props();
</script>

<button
  {...props}
  class="sidebar-closer {side}"
  class:active
  aria-label="Close sidebar"
>
  <i class="ri-close-line"></i>
</button>

<style>
  .sidebar-closer {
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

  .sidebar-closer.left {
    left: -2.5rem;
    transition:
      left 0.3s ease-in-out,
      transform 0.3s;
  }

  .sidebar-closer.right {
    right: -2.5rem;
    transition:
      right 0.3s ease-in-out,
      transform 0.3s;
  }

  .sidebar-closer.left.active {
    transform: rotate(90deg);
    left: calc(max(25vw, 400px) - 4rem);
  }

  .sidebar-closer.right.active {
    right: calc(max(25vw, 400px) - 4rem);
    transform: rotate(-90deg);
  }

  @media (max-width: 750px) {
    .sidebar-closer.left.active {
      left: calc(90vw - 4rem);
    }

    .sidebar-closer.right.active {
      right: calc(90vw - 4rem);
    }
  }
</style>
