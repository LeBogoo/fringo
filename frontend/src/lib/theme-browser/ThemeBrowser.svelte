<script lang="ts">
  import type { Theme } from "../Theme";
  import { customThemes, themeInfo } from "../../themes.svelte.js";

  interface GithubRepoResponse {
    sha: string;
    url: string;
    tree: Tree[];
    truncated: boolean;
  }

  interface Tree {
    path: string;
    mode: string;
    type: string;
    sha: string;
    size?: number;
    url: string;
  }

  interface GithubThemeResponse {
    sha: string;
    node_id: string;
    size: number;
    url: string;
    content: string;
    encoding: string;
  }

  const themeRepo = import.meta.env.VITE_THEME_REPO;
  const repoUrl = `https://api.github.com/repos/${themeRepo}/git/trees/master?recursive=1`;

  const { load } = $props();

  let failed = $state(false);

  let userThemes = $state([]);

  async function fetchThemes() {
    const response = await fetch(repoUrl);

    if (response.status !== 200) {
      console.error("Failed to fetch themes:", response.statusText);
      failed = true;
      return;
    }

    const githubResponse = (await response.json()) as GithubRepoResponse;
    console.log(githubResponse);

    const themeFiles = githubResponse.tree.filter(
      (e) => e.path.startsWith("themes/") && e.path.endsWith(".json")
    );

    userThemes = (await Promise.all(
      themeFiles.map(async (theme) => {
        const fileResponse = await fetch(theme.url);
        const githubFile = (await fileResponse.json()) as GithubThemeResponse;
        return JSON.parse(atob(githubFile.content));
      })
    )) as Theme[];
  }

  function addTheme(theme: Theme) {
    customThemes.push(theme);
    themeInfo.themeId = theme.id;
  }

  $effect(() => {
    console.log("load", load);

    if (load) fetchThemes();
    else
      setTimeout(() => {
        userThemes = [];
      }, 1000);
  });
</script>

<div class="theme-browser">
  <h3>Theme Browser</h3>
  <h4>Community Themes</h4>
  <br />

  {#if userThemes.length === 0 && !failed}
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
        {#each userThemes as theme (theme.id)}
          <tr>
            <td>{theme.name}</td>
            <td>{theme.author}</td>
            <td>
              {#if customThemes.find((t) => t.id === theme.id)}
                <span class="icon-button">
                  <i class="ri-check-line"></i>
                </span>
              {:else}
                <button
                  class="icon-button"
                  aria-label="Add Theme"
                  onclick={() => addTheme(theme)}
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
    <a href="https://github.com/{themeRepo}">Submit it on GitHub!</a>
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
