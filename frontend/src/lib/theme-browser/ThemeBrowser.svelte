<script lang="ts">
  import type { Theme } from "../Theme";
  import { customThemes, themeInfo } from "../../themes.svelte.js";

  interface ThemeSummary {
    id: string;
    name: string;
    author: string;
    fileName: string;
    hash: string;
  }

  const baseUrl = import.meta.env.PROD ? "" : "http://localhost:8080";

  const { load } = $props();

  let failed = $state(false);

  let themeSummary: ThemeSummary[] = $state([]);
  let themeRepo = $state("");

  async function fetchThemes() {
    const repoResponse = await fetch(`${baseUrl}/repo`);
    themeRepo = await repoResponse.text();

    const summaryResponse = await fetch(`${baseUrl}/user-themes`);
    themeSummary = (await summaryResponse.json()) as ThemeSummary[];
  }

  async function downloadTheme(summary: ThemeSummary) {
    const response = await fetch(`${baseUrl}/user-themes/${summary.fileName}`);
    const theme = (await response.json()) as Theme;
    theme.hash = summary.hash;

    return theme;
  }

  async function addTheme(summary: ThemeSummary) {
    const theme = await downloadTheme(summary);
    customThemes.push(theme);
    themeInfo.themeId = theme.id;
  }

  async function updateTheme(summary: ThemeSummary) {
    const theme = await downloadTheme(summary);
    const index = customThemes.findIndex((t) => t.id === summary.id);
    customThemes[index] = theme;
    themeInfo.themeId = theme.id;
  }

  $effect(() => {
    console.log("load", load);

    if (load) fetchThemes();
    else
      setTimeout(() => {
        themeSummary = [];
      }, 1000);
  });
</script>

<div class="theme-browser">
  <h3>Theme Browser</h3>
  <h4>Community Themes</h4>
  <br />

  {#if themeSummary.length === 0 && !failed}
    <p>Loading themes...</p>
  {:else if failed}
    <p>Failed to load community themes. Please try again later.</p>
  {:else}
    <table class="theme-list">
      <thead>
        <tr>
          <th>Name</th>
          <th>Author</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {#each themeSummary as summary (summary.id)}
          <tr>
            <td>{summary.name}</td>
            <td>{summary.author}</td>
            <td>
              {#if customThemes.find((t) => t.id === summary.id)}
                {#if customThemes.find((t) => t.id === summary.id).hash != summary.hash}
                  <button
                    class="icon-button"
                    aria-label="Update Theme"
                    onclick={() => updateTheme(summary)}
                  >
                    <i class="ri-loop-right-line"></i>
                  </button>
                {:else}
                  <span class="icon-button">
                    <i class="ri-check-line"></i>
                  </span>
                {/if}
              {:else}
                <button
                  class="icon-button"
                  aria-label="Add Theme"
                  onclick={() => addTheme(summary)}
                >
                  <i class="ri-add-line"></i>
                </button>
              {/if}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}

  <small class="submit-theme">
    Want to share your own theme?
    <br />
    <a target="_blank" href="https://github.com/{themeRepo}">
      Submit it on GitHub!
    </a>
  </small>
</div>

<style>
  span.icon-button {
    cursor: default;
  }

  .theme-list {
    border-collapse: collapse;
  }

  .submit-theme {
    display: block;
    margin-top: 3rem;
    text-align: center;
  }

  .theme-browser {
    padding: 1rem;
    background: var(--theme-background-color);
    border-radius: 1rem;
    margin-bottom: 0.5rem;
  }

  td {
    padding: 0.1rem;
  }
</style>
