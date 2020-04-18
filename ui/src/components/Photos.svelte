<script>
  import { onDestroy } from 'svelte';

  export const sessionId = '';

  let photos = [];

  (async () => {
    try {
      const res = await fetch('/api/test');
      // const res = await fetch(`/api/session/${sessionId}`);
      const body = await res.json();
      photos = body.mediaItems;

      const sse = new EventSource('/api/sse');
      sse.addEventListener('message', ({ data }) => {
        if (data.photo) {
          photos.unshift(photo);
        }
      });
      sse.addEventListener('error', err => console.error(err));

      onDestroy(() => {
        sse.close();
      });
      console.log(body);
    } catch (err) {
      console.error(err);
    }
  })();
</script>

<style>
  a {
    display: block;
    margin: 10px;
    padding: 10px;
    float: left;
    width: 200px;
    height: 200px;
    border: 1px solid #808080;
  }

  a img {
    max-width: 100%;
    max-height: 100%;
  }
</style>

<section>
  {#each photos as photo}
    <a href={photo.baseUrl} target="_blank">
      <img src={photo.baseUrl} alt={''} />
    </a>
  {/each}
</section>
