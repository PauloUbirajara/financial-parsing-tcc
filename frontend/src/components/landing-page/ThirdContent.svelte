<script lang="ts">
  import { List, Li } from "flowbite-svelte";
  import { fly } from "svelte/transition";

  export let screenHeight: number;
</script>

<div
  class="third-content"
  transition:fly={{ x: -20 }}
  style={`height: ${screenHeight}px; min-height: 300px;`}
>
  <div id="third-content-text" class="bg-custom-primary text-custom-background">
    <p class="text-4xl font-bold">
      Exporte suas transações em formatos convenientes
    </p>
    <p class="text-xl">
      Após adicionar suas transações em nossa plataforma, é possível exportá-las
      em diversos formatos:
    </p>
    <List position="inside">
      <Li>PDF</Li>
      <Li>CSV</Li>
      <Li>HTML</Li>
    </List>
  </div>
  <div id="third-content-img-div">
    <div id="third-content-img"></div>
  </div>
</div>

<style>
  #third-content-img {
    background: url("/landing-page/third-content.webp");
    background-repeat: no-repeat;
    background-attachment: fixed;
    background-size: cover;
    background-position: center;
    height: 100%;
  }

  .third-content {
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  #third-content-text {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    align-items: start;
    justify-content: center;
    padding: 2.5rem;
  }
</style>
