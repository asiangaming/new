<script>
  import Carousel from 'svelte-carousel'
  let slider = [
    {img:"images/banner/banner_1.png"},
    {img:"images/banner/banner_11.png"},
  ]
</script>
<Carousel
  arrows={false}
  autoplay
  autoplayDuration={5000}
  autoplayProgressVisible
  pauseOnFocus>
  {#each slider as rec}
  <img src="{rec.img}" class="w-full" />
  {/each}
</Carousel>