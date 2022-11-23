<script lang="ts">
  import Input from "./lib/Input.svelte";
import List from "./lib/List.svelte";

  export let state = false;
  export let macros = [];

  fetch('http://localhost:8000/macros').then((response) => response.json()).then(res => {
    res.map(macro => {
      macro.description = macro.description.split(/\r?\n/);
    })

    macros = res;
  });

  function handleState(newState) {
    state = newState;
  }

</script>

<main>
  <button on:click={() => handleState(false)}>Add</button>
  <button on:click={() => handleState(true)}>Search</button>

  <br/>

  {#if !state}
    <Input />
  {:else}
    <List macros={macros} />
  {/if}
</main>

<style>
</style>