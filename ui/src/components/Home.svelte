<script>
  let error = null;
  const onClickCreateSession = async event => {
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
  };
</script>

<div>
  <button on:click={onClickCreateSession}>Create session</button>
  {#if error}{error}{/if}
</div>
