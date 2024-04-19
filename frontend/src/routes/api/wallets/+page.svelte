<script lang="ts">
  import WalletAddDrawer from "./WalletAddDrawer.svelte";
  import WalletList from "./WalletList.svelte";

  let hideAddDrawer = true;
  function onAdd() {
    hideAddDrawer = false;
  }
</script>

<WalletAddDrawer bind:hideAddDrawer />
<WalletList {onAdd} />
