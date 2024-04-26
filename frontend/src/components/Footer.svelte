<script lang="ts">
  import { Footer, FooterCopyright, FooterLinkGroup } from "flowbite-svelte";
</script>

<Footer class="p-5">
  <FooterCopyright
    href="/"
    by="Financial Parsing™"
    copyrightMessage="Todos os direitos reservados"
    year={new Date().getFullYear()}
  />
  <FooterLinkGroup
    ulClass="flex flex-wrap items-center mt-3 text-sm text-gray-500 dark:text-gray-400 sm:mt-0"
  />
</Footer>
