<script>
  import { navigating } from "$app/stores";
  import { ImagePlaceholder } from "flowbite-svelte";
  import Navbar from "../../components/Navbar.svelte";
</script>

{#if $navigating}
  <div class="container mx-auto my-5">
    <ImagePlaceholder />
  </div>
{:else}
  <Navbar isLogged={false} />
  <slot />
{/if}
