<script>
    import Scenes from './scenes.svelte'

    let comp;
    try {
        (async () => {
            const auth = await fetch("/api/auth/check");
            const resp = await auth.json();
            if (resp.hasOwnProperty("url") && resp.url !== "") {
                location.assign(resp.url)
            } else {
                comp = Scenes
            }
        })()
    } catch (e) {
        console.error(e)
    }
</script>

<main>
    <svelte:component this={comp}/>
</main>

<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 240px;
        margin: 0 auto;
    }

    h1 {
        color: #ff3e00;
        text-transform: uppercase;
        font-size: 4em;
        font-weight: 100;
    }

    @media (min-width: 640px) {
        main {
            max-width: none;
        }
    }
</style>