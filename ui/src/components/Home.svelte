<script>
    let error = null;
    let sessions = [];
    let activeSession = null;
    (async () => {
        try {
            const res = await fetch('/api/session');
            const body = await res.json();
            activeSession = body.activeSession
            sessions = body.sessions
        } catch (e) {
            console.error(e)
        }
    })()
    const onClickGoToActiveSession = () => {
        if (activeSession) {
            location.assign("/photos/" + activeSession.id)
        } else {
            location.assign("/")
        }
    }
    const onClickCreateSession = async event => {
        event.target.classList.add("is-loading")
        event.preventDefault();
        try {
            const resp = await fetch('/api/session', {
                method: 'POST',
                redirect: 'error',
            });
            const body = await resp.json();
            if (body.url) {
                location.assign(body.url);
            } else {
                error = body.error || 'something went wrong';
            }
        } catch (err) {
            console.error(err);
        }
        event.target.classList.remove("is-loading")
    };
</script>

<div>
    {#if activeSession}
        <button class="button" on:click={onClickGoToActiveSession}>Go to active session</button>
    {:else}
        <button class="button" on:click={onClickCreateSession}>Start session</button>
    {/if}
    {#if error}{error}{/if}
</div>
