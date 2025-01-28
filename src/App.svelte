<script lang="ts">
  import { images, goals } from "./data.js";
  import { Rng } from "./lib/rng.js";

  let userSeed =
    parseInt(localStorage.getItem("seed")) ||
    Math.floor(Math.random() * 1000000);

  localStorage.setItem("seed", userSeed.toString());

  let seed = Math.floor(userSeed + Date.now() / 1000 / 60 / 60 / 24);

  let rng = new Rng(seed);

  let randomGoals = $state(
    (() => {
      let g = [...goals].sort(() => rng.next() - 0.5);
      g.splice(12, 0, "");
      return g;
    })()
  );

  // pick the first 24 goals
  let board = $derived(randomGoals.slice(0, 25));

  let image = $state(images[Math.floor(rng.next() * images.length)]);

  let checked = $state(new Array(25).map((e) => false)); // one extra false for the center

  // Ensure the center image is always checked and can't be unchecked
  checked[12] = true;

  let isFringo = $state(false);
  let isDoubleFringo = $state(false);

  let fringoIndexes = $state([]);

  let headingText = $state("Fringo");
  let headingCharacters = $derived(headingText.split(""));

  $effect(() => {
    if (isDoubleFringo) {
      headingText = "It's a double Fringo!";
    } else if (isFringo) {
      headingText = "It's a Fringo!";
    } else {
      headingText = "Fringo";
    }
  });

  // Function to update favicon
  function updateFavicon() {
    const link = (document.querySelector("link[rel~='icon']") ||
      document.createElement("link")) as HTMLLinkElement;
    link.rel = "icon";
    link.href = image;
    document.getElementsByTagName("head")[0].appendChild(link);
  }

  // Ensure "Gay" is always checked and can't be unchecked
  let gayIndex = $derived(board.findIndex((goal) => goal === "Gay"));
  $effect(() => {
    if (gayIndex !== -1) {
      checked[gayIndex] = true;
    }
  });

  function toggleField(index: number) {
    if (index === gayIndex) return; // Prevent toggling "Gay" field

    checked[index] = !checked[index];

    checkFringo(index);
  }

  function checkFringo(index: number) {
    fringoIndexes = [];
    isFringo = false;
    isDoubleFringo = false;

    const row = Math.floor(index / 5);
    const col = index % 5;

    // check row
    if (
      checked[row * 5] &&
      checked[row * 5 + 1] &&
      checked[row * 5 + 2] &&
      checked[row * 5 + 3] &&
      checked[row * 5 + 4]
    ) {
      isFringo = true;
      [row * 5, row * 5 + 1, row * 5 + 2, row * 5 + 3, row * 5 + 4].forEach(
        (i) => fringoIndexes.push(i)
      );
    }

    // check column
    if (
      checked[col] &&
      checked[col + 5] &&
      checked[col + 10] &&
      checked[col + 15] &&
      checked[col + 20]
    ) {
      isFringo = true;
      [col, col + 5, col + 10, col + 15, col + 20].forEach((i) =>
        fringoIndexes.push(i)
      );
    }

    // check diagonals
    if (
      (checked[0] && checked[6] && checked[12] && checked[18] && checked[24]) ||
      (checked[4] && checked[8] && checked[12] && checked[16] && checked[20])
    ) {
      isFringo = true;
      if (
        checked[0] &&
        checked[6] &&
        checked[12] &&
        checked[18] &&
        checked[24]
      ) {
        [0, 6, 12, 18, 24].forEach((i) => fringoIndexes.push(i));
      }
      if (
        checked[4] &&
        checked[8] &&
        checked[12] &&
        checked[16] &&
        checked[20]
      ) {
        [4, 8, 12, 16, 20].forEach((i) => fringoIndexes.push(i));
      }
    }

    if (fringoIndexes.length > 5) {
      isDoubleFringo = true;
    }
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

    checked = new Array(25).map((e) => false);

    // Ensure the center image is always checked and can't be unchecked
    checked[12] = true;

    isFringo = false;
    isDoubleFringo = false;
    fringoIndexes = [];

    updateFavicon();
  }
</script>

<svelte:head>
  <link rel="icon" href={image} />
</svelte:head>

<div class="app">
  <h1>
    {#each headingCharacters as char, index}
      <span
        class:space={char === " "}
        class:fringo={isFringo && char !== " "}
        style="animation-delay: {index * 0.1}s;">{char}</span
      >
    {/each}
  </h1>
  <div class="board">
    {#each board.slice(0, 12) as text, index (text)}
      <button
        disabled={isFringo && !fringoIndexes.includes(index)}
        class="field"
        onclick={() => toggleField(index)}
        class:checked={checked[index]}
        class:gay={index == gayIndex}
      >
        {text}
      </button>
    {/each}
    <img src={image} alt="Funny Fringo" class="field" />
    {#each board.slice(13) as text, index (text)}
      <button
        disabled={isFringo && !fringoIndexes.includes(index + 13)}
        class="field"
        onclick={() => toggleField(index + 13)}
        class:checked={checked[index + 13]}
        class:gay={index + 13 == gayIndex}
      >
        {text}
      </button>
    {/each}
  </div>
  <button
    id="new-game-button"
    class="btn btn-outline"
    class:hidden={!isFringo}
    onclick={() => newBoard()}>New Board</button
  >
</div>
