<script lang="ts">
  import Field from "./Field.svelte";

  const { board, onFringo, image } = $props();

  const now = Math.floor(Date.now() / 1000 / 60 / 60 / 24);

  let savedState: boolean[] = [];
  let savedDay = localStorage.getItem("saveDay");
  if (savedDay && parseInt(savedDay, 36) == now) {
    savedState = JSON.parse(localStorage.getItem("checked"));
  } else {
    localStorage.removeItem("checked");
    localStorage.removeItem("saveDay");
  }

  let checked = $state(
    savedState.length
      ? savedState
      : new Array(25).fill(false).map((_, i) => i == 12 || board[i] == "Gay")
  );
  let isFringo = $state(false);
  let fringoCount = $state(0);
  let fringoIndexes = $state<number[]>([]);

  function checkFringo() {
    isFringo = false;
    fringoIndexes = [];
    fringoCount = 0;

    for (let row = 0; row < 5; row++) {
      if (checkRow(row)) {
        isFringo = true;
        fringoCount++;
        for (let i = 0; i < 5; i++) {
          fringoIndexes.push(row * 5 + i);
        }
      }
    }

    for (let col = 0; col < 5; col++) {
      if (checkCol(col)) {
        isFringo = true;
        fringoCount++;
        for (let i = 0; i < 5; i++) {
          fringoIndexes.push(col + i * 5);
        }
      }
    }

    if (checkDiagonal()) {
      isFringo = true;
      fringoCount++;
      [0, 6, 18, 24].forEach((i) => fringoIndexes.push(i));
    }

    if (checkAntiDiagonal()) {
      isFringo = true;
      fringoCount++;
      [4, 8, 16, 20].forEach((i) => fringoIndexes.push(i));
    }

    onFringo(fringoCount);
  }

  function checkRow(row: number) {
    return (
      checked[row * 5] &&
      checked[row * 5 + 1] &&
      checked[row * 5 + 2] &&
      checked[row * 5 + 3] &&
      checked[row * 5 + 4]
    );
  }

  function checkCol(col: number) {
    return (
      checked[col] &&
      checked[col + 5] &&
      checked[col + 10] &&
      checked[col + 15] &&
      checked[col + 20]
    );
  }

  function checkDiagonal() {
    return checked[0] && checked[6] && checked[18] && checked[24];
  }

  function checkAntiDiagonal() {
    return checked[4] && checked[8] && checked[16] && checked[20];
  }

  function toggle(index: number) {
    checked[index] = !checked[index];
    checkFringo();

    saveState();
  }

  function saveState() {
    localStorage.setItem("saveDay", now.toString(36));
    localStorage.setItem("checked", JSON.stringify(checked));
  }

  export function reset() {
    checked = new Array(25)
      .fill(false)
      .map((_, i) => i == 12 || board[i] == "Gay");
    localStorage.removeItem("checked");
    localStorage.removeItem("saveDay");
    isFringo = false;
    fringoCount = 0;
    fringoIndexes = [];
  }

  checkFringo();
</script>

<div class="board">
  {#each board as text, index (text)}
    {#if index != 12}
      <Field
        {text}
        disabled={isFringo && !fringoIndexes.includes(index)}
        onclick={() => toggle(index)}
        checked={checked[index]}
      ></Field>
    {:else}
      <Field url={image}></Field>
    {/if}
  {/each}
</div>
