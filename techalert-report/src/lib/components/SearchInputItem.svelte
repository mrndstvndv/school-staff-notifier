<script lang="ts">
	export let focused = false;
	export let label = "";
	export let onClick = (label: string) => {};
</script>

<button
	type="button"
	class:focus={focused}
	tabindex="-1"
	on:click={() => {
		onClick(label);
	}}
>
	{label}
</button>

<style>
	button {
		border-radius: 0;
		padding: 5px;
	}

	button.focus {
		outline: none;
		border-radius: 0;
		@apply bg-gray-200 rounded-md;
	}
</style>
