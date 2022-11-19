<script lang="ts">
  export let macros = [];

  fetch('http://localhost:8000/macros').then((response) => response.json()).then(res => {
    res.map(macro => {
      macro.description = macro.description.split(/\r?\n/);
    })

    macros = res;
  });

</script>

<main>
  <h1>FF14 macro</h1>
  <ul>
    {#each macros as macro}
      <li>
        <p>{macro.name}</p>
        {#each macro.description as line}
          <pre class="ffxiv-text">{line}</pre>
        {/each}
      </li>
    {/each}
  </ul>
</main>

<style>
</style>