<script lang="ts">
  import { fly } from "svelte/transition";

  export let screenHeight: number;
</script>

<div
  class="second-content"
  transition:fly={{ x: -20 }}
  style={`height: ${screenHeight}px; min-height: 300px;`}
>
  <div id="second-content-img-div">
    <div id="second-content-img"></div>
  </div>
  <div
    id="second-content-text"
    class="bg-custom-foreground text-custom-background"
  >
    <p class="text-4xl font-bold">
      Gerencie diversas transações em um só lugar
    </p>
    <p class="text-xl">
      Com transações categorizáveis, é possível entender os dados através de
      gráficos para tomar uma melhor decisão durante a gerência de suas
      finanças.
    </p>
  </div>
</div>

<style>
  #second-content-img {
    background: url(./landing-page/second-content.webp);
    background-repeat: no-repeat;
    background-attachment: fixed;
    background-size: cover;
    background-position: center;
    height: 100%;
  }

  .second-content {
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  #second-content-text {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    align-items: start;
    justify-content: center;
    padding: 2.5rem;
  }
</style>
