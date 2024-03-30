<script lang="ts">
	import { page } from '$app/stores';
    import { Heading, P } from 'flowbite-svelte';
</script>

<Heading>Houve algum erro durante a autenticação do usuário...</Heading>
<P>
  {$page.status}: {$page.error?.message}
</P>
