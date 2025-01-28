<script lang="ts">
  import { images, goals } from "./data.js";
  import { Rng } from "./lib/rng.js";

  let userSeed =
    parseInt(localStorage.getItem("seed")) ||
    Math.floor(Math.random() * 1000000);

  localStorage.setItem("seed", userSeed.toString());

  let seed = Math.floor(userSeed + Date.now() / 1000 / 60 / 60 / 24);

  const rng = new Rng(seed);

  let randomGoals = [...goals].sort(() => rng.next() - 0.5);
  // pick the first 24 goals
  let board = randomGoals.slice(0, 24);
  // insert an empty string at position 12
  board.splice(12, 0, "");

  let image = images[Math.floor(rng.next() * images.length)];

  let checked = $state([...randomGoals.map(() => false), false]); // one extra false for the center

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

  function toggleField(index: number) {
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

      console.log("row", fringoIndexes);
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

    if (fringoIndexes.length > 5) {
      isDoubleFringo = true;
    }
  }

  // Call the function to set the initial favicon
  updateFavicon();
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
    {#each board.slice(0, 12) as text, index}
      <button
        disabled={isFringo && !fringoIndexes.includes(index)}
        class="field"
        onclick={() => toggleField(index)}
        class:checked={checked[index]}>{text}</button
      >
    {/each}
    <img src={image} alt="Funny Fringo" class="field" />
    {#each board.slice(13) as text, index}
      <button
        disabled={isFringo && !fringoIndexes.includes(index + 13)}
        class="field"
        onclick={() => toggleField(index + 13)}
        class:checked={checked[index + 13]}>{text}</button
      >
    {/each}
  </div>
</div>

<style>
  .app {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 95vh;
  }

  h1 {
    font-family: "Nunito";
    text-align: center;
    font-weight: 700;
    font-size: 4rem;
    margin-bottom: 1rem;
    text-transform: uppercase;
  }

  .board {
    margin: 0 auto;
    display: grid;
    grid-template-columns: repeat(5, 1fr);
    grid-auto-rows: 1fr;
    gap: 5px;
    max-width: 80vmin;
    max-height: 80vmin;
    background-color: black;
    padding: 5px;
    border-radius: 15px;
  }

  .field {
    font-family: "Nunito";
    text-align: center;
    font-weight: 700;
    display: inline-block;
    padding: 5px;
    width: 100%;
    aspect-ratio: 1;
    color: black;
    background-color: white;
    border-radius: 10px;
    font-size: 2vmin;
  }

  img.field {
    padding: 0;
  }

  button.field {
    border: none;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  button.field.checked {
    background-color: #a1ff84;
  }
  button.field.checked:hover {
    background-color: #72cf56;
  }

  button.field:disabled {
    background-color: #777 !important;
    cursor: not-allowed;
  }

  button.field:hover {
    background-color: #d5d5d5;
  }

  @keyframes bounce {
    0%,
    20%,
    50%,
    80%,
    100% {
      transform: translateY(0);
    }
    40% {
      transform: translateY(-15px);
    }
    60% {
      transform: translateY(-7.5px);
    }
  }

  @keyframes rainbow {
    0% {
      color: red;
    }
    14% {
      color: orange;
    }
    28% {
      color: yellow;
    }
    42% {
      color: green;
    }
    57% {
      color: blue;
    }
    71% {
      color: indigo;
    }
    85% {
      color: violet;
    }
    100% {
      color: red;
    }
  }

  .fringo {
    display: inline-block;

    animation:
      bounce 1s infinite,
      rainbow 1s infinite;
  }

  .space {
    display: inline-block;
    padding: 0 1rem;
  }
</style>
