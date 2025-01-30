<script>
  let version = import.meta.env.VITE_VERSION;
  let commit = import.meta.env.VITE_GIT_COMMIT;
</script>

<small class="debug-display">
  {version} - {commit}
</small>

<style>
  .debug-display {
    position: absolute;
    bottom: 0;
    right: 0;
    color: var(--theme-text-color);
    opacity: 0.5;
  }
</style>
