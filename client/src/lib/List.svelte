<script lang="ts">
    export let macros = [];
</script>

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
