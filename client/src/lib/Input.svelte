<script lang="ts">
    function PostToMacroBackend(body) {
        console.log(body);

        fetch('http://localhost:8000/macro', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
        }).then((res) => res.json()).then((data) => console.log(data));
    }

    function handleSubmit(e) {
        const formData = new FormData(e.target);

        const data = {};
        for (let field of formData) {
        const [key, value] = field;
        data[key] = value;
        }
        
        PostToMacroBackend(data);
    }
</script>

<div>
    <form class="macro-form" on:submit|preventDefault={handleSubmit}>
        <label class="label">
            <p>Macro Name</p>
            <input type="text" class="macro-input" id="name" name="name">
        </label>
        <label class="label">
            <p>Macro content</p>
            <textarea class="macro-textarea" id="description" name="description"></textarea>
        </label>
        <button class="macro-button">Submit</button>
    </form>
</div>

<style>
    .label {
        display: block;
        margin-bottom: 12px;
    }

    .label input, .label textarea {
        width: 100%;
        padding: 12px;
        background: var(--background-color);
        border: none;
    }

    p {
        margin: 0;
    }

    .macro-form {
        margin: 0 auto;
        width: 500px;
    }

    .macro-textarea {
        height: 300px;
    }

    .macro-button {
        background: var(--primary-color);
        border: none;
        padding: 8px 50px;
        color: var(--text-color-dark);
        font-weight: 700;
        border-radius: 3px;
    }
</style>