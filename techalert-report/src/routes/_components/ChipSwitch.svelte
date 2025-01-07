<script lang="ts">
	import Chip from "$lib/components/Chip.svelte";

	let urgency = 0;

	let chips = [
		{
			name: "Low",
			bg: "bg-green-100",
		},
		{
			name: "Medium",
			bg: "bg-orange-100",
		},
		{ name: "High", bg: "bg-red-100" },
	];
</script>

<div class="flex gap-6">
	{#each chips as chip, index}
		<Chip
			onToggle={() => {
				urgency = index;
			}}
			label={chip.name}
			checked={index == urgency}
			checkedStyle={`${chip.bg} border border-white/0`}
		></Chip>
	{/each}

	<input id="urgency" name="urgency" type="hidden" value={urgency} />
</div>
