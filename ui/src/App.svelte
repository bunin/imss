<script>
  import { Router, Route, Link } from 'svelte-routing';
  import Home from './components/Home.svelte';
  import Photos from './components/Photos.svelte';

  export let url = '';

  /**
   * @name authorized
   * @description Defines if the user is authorized
   */
  let authorized = false;

  (async () => {
    try {
      const resp = await fetch('/api/auth/check');
      const body = await resp.json();
      if (body.url) {
        location.assign(body.url);
      } else {
        authorized = true;
      }
    } catch (err) {
      console.error(err);
    }
  })();
</script>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<main>
  {#if authorized === null}
    Loading...
  {:else if authorized}
    <Router {url}>
      <nav>
        <Link to="/">Home</Link>
      </nav>
      <div>
        <Route path="/" component={Home} />
        <Route path="/photos/:sessionId" component={Photos} />
      </div>
    </Router>
  {:else}Unauthorized. Redirecting to authentication page.{/if}
</main>
