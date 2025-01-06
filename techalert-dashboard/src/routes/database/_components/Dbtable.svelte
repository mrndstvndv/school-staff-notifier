<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import * as Card from "$lib/components/ui/card/index.js";
	import * as Table from "$lib/components/ui/table";
	import * as Tabs from "$lib/components/ui/tabs";
	import { Plus } from "lucide-svelte";

	export let table;
	export let value;
	export let onAdd = () => {};

	let dialog: HTMLDialogElement | null = null;
	$: fields = () => {
		if (table[0] == null) {
			return [];
		}

		return Object.keys(table[0]);
	};

	// @ts-ignore
	function snakeCaseToTitle(snakeCaseString: string): string {
		// Split the snake_case string into words
		const words = snakeCaseString.split("_");

		// Capitalize the first letter of each word and join them with spaces
		const titleCaseString = words
			.map((word) => word.charAt(0).toUpperCase() + word.slice(1))
			.join(" ");

		return titleCaseString;
	}

	function removePostfix(str: string) {
		return str.slice(0, -1);
	}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<dialog
	on:click|self={() => {
		if (dialog !== null) {
			dialog.close();
		}
	}}
	bind:this={dialog}
	class="rounded-lg border shadow-lg border-gray-300 backdrop:bg-background/80 backdrop:fixed backdrop:inset-0 backdrop:backdrop-blur-sm min-w-80"
>
	<Card.Header>
		<Card.Title
			>Add a new {removePostfix(snakeCaseToTitle(value))}</Card.Title
		>
		<!--<Card.Description>Set the server endpoint</Card.Description>-->
	</Card.Header>
	<Card.Content>
		<slot></slot>
	</Card.Content>
	<Card.Footer class="flex justify-between">
		<Button
			on:click={() => {
				dialog?.close();
			}}
			variant="outline">Cancel</Button
		>
		<Button
			on:click={() => {
				onAdd();
				dialog?.close();
			}}>Add</Button
		>
	</Card.Footer>
</dialog>

<Tabs.Content {value}>
	<Table.Root class="border rounded">
		<Table.Header>
			<Table.Row class="hover:bg-black/[0.02]">
				{#each fields() as field}
					{#if field != "id"}
						<Table.Head>{snakeCaseToTitle(field)}</Table.Head>
					{/if}
				{/each}
			</Table.Row>
		</Table.Header>
		<Table.Body>
			{#each table as student, i (i)}
				<Table.Row class="hover:bg-black/[0.02]">
					{#each fields() as field}
						{#if field != "id"}
							<Table.Cell class="font-medium"
								>{student[field]}</Table.Cell
							>
						{/if}
					{/each}
				</Table.Row>
			{/each}
		</Table.Body>
		<Table.Caption>
			<Button
				on:click={() => {
					dialog?.showModal();
				}}
			>
				<Plus class="mr-2 h-4 w-4" />
				Add
			</Button>
		</Table.Caption>
	</Table.Root>
</Tabs.Content>
