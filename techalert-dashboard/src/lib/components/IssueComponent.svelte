<script lang="ts">
	import { Issue } from "$lib/types/issues";

	export let issue: Issue;
	export let onComponentToggle: (
		event: Event,
		componentIndex: number,
	) => void;
	$: isFixed = issue.faultyComponents.every((item) => item.fixed === true);
</script>

<div
	class="mx-auto p-6 bg-white rounded-lg shadow-sm border border-gray-200 w-full"
>
	<header class="flex items-center gap-4 mb-6">
		<div
			class="flex items-center justify-center w-12 h-12 {isFixed
				? 'bg-green-100'
				: 'bg-red-100'} rounded-full"
		>
			<!-- TODO: when the status is 1 then it should be blue  -->
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="w-6 h-6 {isFixed ? 'text-green-600' : 'text-red-600'}"
				viewBox="0 0 24 24"
				fill="none"
				stroke="currentColor"
				stroke-width="2"
			>
				<rect x="2" y="3" width="20" height="14" rx="2" ry="2" />
				<line x1="8" y1="21" x2="16" y2="21" />
				<line x1="12" y1="17" x2="12" y2="21" />
			</svg>
		</div>
		<div>
			<h1 class="text-xl font-semibold text-gray-900">
				PC-{issue.pcNumber}
			</h1>
			<time class="text-sm text-gray-500"
				>{new Date(
					Number.parseInt(issue.timestamp, 10),
				).toLocaleString()}</time
			>
		</div>
	</header>

	{#if issue.concern}
		<div class="mb-6">
			<p class="text-gray-700 leading-relaxed">
				{issue.concern}
			</p>
		</div>
	{/if}

	{#if issue.faultyComponents.length > 0}
		<div>
			<h2 class="text-lg font-medium text-gray-900 mb-3">Issues</h2>
			<div class="space-y-2">
				{#each issue.faultyComponents as item, index}
					<label
						class="flex items-center gap-3 p-3 rounded-lg border border-gray-200 hover:bg-gray-50 cursor-pointer transition-colors"
					>
						<input
							type="checkbox"
							checked={item.fixed}
							on:change={(event) => {
								onComponentToggle(event, index);
							}}
							class="w-5 h-5 rounded accent-white text-black"
						/>
						<span class="text-gray-700">{item.name}</span>
					</label>
				{/each}
			</div>
		</div>
	{/if}
</div>

<style>
	/* Custom checkbox styles */
	/*input[type="checkbox"] {*/
	/*	accent-color: rgb(220, 38, 38);*/
	/*}*/
</style>
