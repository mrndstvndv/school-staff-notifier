<script lang="ts">
	import { Check, X } from "lucide-svelte";

	export let checked: boolean = true;
	export let label: string;
	export let negativeLabel: string = "";

	$: bg = () => {
		if (checked == true) {
			return "bg-green-100 border border-green-200";
		}
		return "border";
	};
</script>

<button
	on:click={() => {
		checked = !checked;
	}}
	class={`flex class ${bg()} rounded-lg px-3 pr-4 py-1.5 justify-center items-center gap-1 hover:cursor-pointer hover:shadow-sm font-semibold text-sm`}
>
	{#if checked}
		<Check class="size-4"></Check>
	{:else}
		<X class="size-4"></X>
	{/if}

	{#if negativeLabel != "" && checked == false}
		{negativeLabel}
	{:else}
		{label}
	{/if}
</button>
