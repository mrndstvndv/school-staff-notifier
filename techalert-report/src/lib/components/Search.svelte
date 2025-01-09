<script lang="ts">
	import SearchInputItem from "./SearchInputItem.svelte";

	export let placeholder: string = "";
	export let classNames: string = "";
	export let items: string[] = [];
	export let value = "";
	export let form: HTMLFormElement;
	export let name: string = "";
	export let id: string = "";
	export let required: boolean = false;

	let inputRef: HTMLInputElement;
	let popoverRef: HTMLDialogElement;
	let containerRef: HTMLDivElement;
	let filterText = "";
	let currentFocus = 0;

	// Computed filtered list based on the input text
	$: filteredItems = items.filter((item) =>
		item.toLowerCase().includes(filterText.toLowerCase()),
	);

	function onFocus(event: Event) {
		if (items.length == 0) return;

		if (value !== "") {
			inputRef.select();
		}

		let target = event.target as HTMLInputElement;
		let popover = target.nextElementSibling as HTMLDialogElement;
		openPopover();

		inputRef.addEventListener("keyup", onKeyUp);
	}

	function formKeyDown(event: KeyboardEvent) {
		if (event.key === "Enter") {
			event.preventDefault();
		}
	}

	function onKeyUp(event: KeyboardEvent) {
		if (popoverRef.open == false) return;

		if (form !== undefined) form.addEventListener("keydown", formKeyDown);

		switch (event.key) {
			case "ArrowUp":
				currentFocus = Math.max(0, currentFocus - 1);
				value = filteredItems[currentFocus];
				break;
			case "ArrowDown":
				currentFocus = Math.min(
					filteredItems.length - 1,
					currentFocus + 1,
				);
				value = filteredItems[currentFocus];
				break;
			case "Enter":
				value = filteredItems[currentFocus];
				closePopover();
				break;
		}
	}

	function closePopover() {
		if (form !== undefined)
			form.removeEventListener("keydown", formKeyDown);
		inputRef.removeEventListener("keyup", onKeyUp);
		popoverRef.open = false;
	}

	function onInput(event: Event) {
		let target = event.target as HTMLInputElement;
		filterText = target.value;
		currentFocus = 0;
		openPopover()
	}

	function openPopover() {
		if (filteredItems.length > 0) {
			popoverRef.open = true;
		}
	}

	function onBlur(event: FocusEvent) {
		if (!containerRef.contains(event.relatedTarget as Node)) {
			closePopover();
		}
	}
</script>

<div bind:this={containerRef} class={`relative ${classNames}`}>
	<input
		bind:this={inputRef}
		bind:value
		on:focus={onFocus}
		on:input={onInput}
		on:click={onFocus}
		on:blur={onBlur}
		class="w-full"
		{name}
		{placeholder}
		{id}
		{required}
		aria-autocomplete="none"
	/>

	<dialog
		bind:this={popoverRef}
		tabindex="-1"
		class="absolute left-0 w-full border-[1.5px] border-gray-300 shadow-sm rounded-lg p-2"
	>
		<ul class="grid p-2 item-list" tabindex="-1">
			{#each filteredItems as item, index}
				<SearchInputItem
					label={item}
					focused={index == currentFocus}
					onClick={(label) => {
						value = label;
						closePopover();
					}}
				/>
			{/each}
		</ul>
	</dialog>
</div>

<style>
	ul {
		list-style: none;
		margin: 0;
	}

	dialog {
		margin-top: 5px;
		padding: 0px;
		z-index: 1000;
	}

	input {
		@apply border-[1.5px] border-gray-300 shadow-sm rounded-lg p-2 focus:outline-none focus:ring-[1.5px] focus:ring-green-500 focus:border-green-500;
	}
</style>
