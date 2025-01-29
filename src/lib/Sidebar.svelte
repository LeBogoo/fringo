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
</style>
