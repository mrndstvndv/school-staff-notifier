<script lang="ts">
	export let checked: boolean = true;
	export let label: string;
	export let negativeLabel: string = "";
	export let checkedStyle: string = "bg-green-100 border border-green-200";
	export const style = "border";
	export let onToggle = () => {};

	$: bg = () => {
		if (checked == true) {
			return checkedStyle;
		}
		return style;
	};
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
	role="button"
	tabindex="0"
	on:click={() => {
		onToggle();
	}}
	class={`flex class ${bg()} rounded-lg px-3 pr-4 py-1.5 justify-center items-center gap-1 hover:cursor-pointer hover:shadow-sm font-semibold text-sm select-none`}
>
	{#if negativeLabel != "" && checked == false}
		{negativeLabel}
	{:else}
		{label}
	{/if}
</div>
