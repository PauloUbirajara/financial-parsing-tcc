<script lang="ts">
  import WalletDetail from "./WalletDetail.svelte";
  import WalletEdit from "./WalletEdit.svelte";

  enum WalletEnum {
    DETAIL,
    EDIT,
  }
  let step = WalletEnum.DETAIL;

  function onEdit() {
    step = WalletEnum.EDIT;
  }
  function onDetail() {
    step = WalletEnum.DETAIL;
  }
</script>

{#if step === WalletEnum.DETAIL}
  <WalletDetail {onEdit} />
{/if}

{#if step === WalletEnum.EDIT}
  <WalletEdit {onDetail} />
{/if}
