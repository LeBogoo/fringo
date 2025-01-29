<script lang="ts">
  import { images, goals } from "./data.js";
  import { themes } from "./themes";
  import ThemeToggle from "./lib/ThemeToggle.svelte";
  import Board from "./lib/Board.svelte";
  import Sidebar from "./lib/Sidebar.svelte";
  import ThemeList from "./lib/ThemeList.svelte";

  let userSeed =
    parseInt(localStorage.getItem("seed")) ||
    Math.floor(Math.random() * 1000000);

  localStorage.setItem("seed", userSeed.toString());
  const now = Math.floor(Date.now() / 1000 / 60 / 60 / 24);

  let seed = userSeed + now;

  let rng = new Rng(seed);

  let randomGoals = $state(
    (() => {
      let g = [...goals].sort(() => rng.next() - 0.5);
      g.splice(12, 0, "");
      return g;
    })()
  );

  let fringoBoard: Board | null = null;

  const board = $derived(randomGoals.slice(0, 25));
  let image = $state(images[Math.floor(rng.next() * images.length)]);
  let isFringo = $state(false);
  let headingText = $state("Fringo");
  const headingCharacters = $derived(headingText.split(""));

  function updateFavicon() {
    const link = (document.querySelector("link[rel~='icon']") ||
      document.createElement("link")) as HTMLLinkElement;
    link.rel = "icon";
    link.href = image;
    document.getElementsByTagName("head")[0].appendChild(link);
  }

  updateFavicon();

  function newBoard() {
    userSeed++;
    localStorage.setItem("seed", userSeed.toString());
    seed++;

    rng = new Rng(seed);

    randomGoals = [...goals].sort(() => rng.next() - 0.5);
    randomGoals.splice(12, 0, "");

    image = images[Math.floor(rng.next() * images.length)];
    isFringo = false;

    fringoBoard?.reset();

    headingText = "Fringo";

    updateFavicon();
  }

  function onFringo(amount: number) {
    let multipliers = {
      1: "",
      2: "Double ",
      3: "Triple ",
      4: "Quadruple ",
    };

    let multiplier = multipliers[amount];

    isFringo = amount != 0;
    if (isFringo) headingText = `It's a ${multiplier}Fringo!`;
    else headingText = "Fringo";
  }
</script>

<svelte:head>
  <link rel="icon" href={image} />
</svelte:head>

<Sidebar side="left">
  <ThemeList {themes}></ThemeList>
</Sidebar>

<div class="app">
  <h1 class="fringo-heading">
    {#each headingCharacters as char, index}
      <span
        class:space={char === " "}
        class:fringo={isFringo && char !== " "}
        style="animation-delay: {index * 0.1}s;"
      >
        {char}
      </span>
    {/each}
  </h1>
  <Board bind:this={fringoBoard} {board} {onFringo} {image} />
  <button
    id="new-game-button"
    class="btn"
    class:hidden={!isFringo}
    onclick={() => newBoard()}
  >
    New Board
  </button>
</div>
