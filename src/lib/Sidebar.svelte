<script lang="ts">
  import SidebarToggle from "./SidebarToggle.svelte";
  let {
    side,
    children,
  }: {
    side: "left" | "right";
    children?: () => any;
  } = $props();

  let visible = $state(false);

  function toggle() {
    visible = !visible;
  }
</script>

<SidebarToggle onclick={toggle} active={visible} {side}></SidebarToggle>
<div class="sidebar {side}" class:visible>
  <div class="spacer"></div>
  <div class="content">
    {@render children?.()}
  </div>
</div>

<style>
  .spacer {
    width: 100%;
    height: 4rem;
  }

  .content {
    padding: 0 1rem;
  }

  .sidebar {
    position: absolute;
    top: 0;
    bottom: 0;
    width: max(25vw, 400px);
    max-height: 100vh;
    overflow: auto;
    scrollbar-width: none;
    z-index: 1;
    background: var(--theme-sidebar-color);
    backdrop-filter: blur(var(--theme-sidebar-background-blur));
  }

  .sidebar.left {
    left: min(-25vw, -400px);
    transition: left 0.3s ease-in-out;
    border-radius: 0 1rem 1rem 0;
  }

  .sidebar.right {
    right: min(-25vw, -400px);
    transition: right 0.3s ease-in-out;
    border-radius: 1rem 0 0 1rem;
  }

  .sidebar.left.visible {
    left: 0;
  }

  .sidebar.right.visible {
    right: 0;
  }

  :global(.sidebar-panel) {
    padding: 1rem;
    background: var(--theme-background-color);
    border-radius: 1rem;
    margin-bottom: 0.5rem;
  }

  @media (max-width: 750px) {
    .sidebar {
      width: 90vw;
    }

    .sidebar.left {
      left: -90vw;
    }

    .sidebar.right {
      right: -90vw;
    }
  }
</style>
