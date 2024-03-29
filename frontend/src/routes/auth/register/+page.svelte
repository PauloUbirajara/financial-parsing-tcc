<script>
	import { Input, Label, Button, P } from "flowbite-svelte";
</script>

<form>
	<P size="4xl" weight="bold" class="mb-5">Crie sua conta</P>
	<div class="mb-6">
		<Label for="email" class="mb-2">E-mail</Label>
		<Input
			type="email"
			id="email"
			size="md"
			placeholder="john.doe@company.com"
			required
		/>
	</div>
	<div class="mb-6">
		<Label for="password" class="mb-2">Senha</Label>
		<Input
			type="password"
			size="md"
			id="password"
			placeholder="•••••••••"
			required
		/>
	</div>
	<div class="mb-6">
		<Label for="confirm_password" class="mb-2"
			>Confirmar Senha</Label
		>
		<Input
			type="password"
			size="md"
			id="confirm_password"
			placeholder="•••••••••"
			required
		/>
	</div>
	<Button type="submit">Avançar</Button>
</form>
