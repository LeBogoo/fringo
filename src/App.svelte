<script lang="ts">
  import { images, goals } from "./data";
  import { Rng } from "./lib/rng";
  import Board from "./lib/Board.svelte";
  import Sidebar from "./lib/Sidebar.svelte";
  import ThemeList from "./lib/ThemeList.svelte";
  import DebugDisplay from "./lib/DebugDisplay.svelte";

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
  const headlingLength = $derived(headingText.length);

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
    let vowles = ["a", "e", "i", "o", "u"];

    let multipliers: { [key: number]: string } = {
      1: "",
      2: "Double ",
      3: "Triple ",
      4: "Quadruple ",
      5: "Quintuple ",
      6: "Sextuple ",
      7: "Septuple ",
      8: "Octuple ",
      9: "Nonuple ",
      10: "Decuple ",
      11: "Undecuple ",
      12: "Duodecuple ",
    };

    let multiplier = multipliers[amount];
    let article =
      multiplier && vowles.includes(multiplier[0].toLowerCase()) ? "an" : "a";

    isFringo = amount != 0;
    if (isFringo) headingText = `It's ${article} ${multiplier}Fringo!`;
    else headingText = "Fringo";
  }
</script>

<svelte:head>
  <link rel="icon" href={image} />
</svelte:head>

<Sidebar side="left">
  <ThemeList></ThemeList>
</Sidebar>

<DebugDisplay></DebugDisplay>

<div class="app">
  <h1 class="fringo-heading">
    {#each headingCharacters as char, index}
      <span
        class:space={char === " "}
        class:fringo={isFringo && char !== " "}
        style="animation-delay: -{(headlingLength - index) * 0.1}s;"
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
