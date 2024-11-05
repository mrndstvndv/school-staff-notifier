<script lang="ts">
	import "../app.css";
	import { onMount } from "svelte";

	interface Issue {
		studentName: string;
		professorName: string;
		section: string;
		labName: string;
		description: string;
		timestamp: string;
		schoolYear: number;
		pcNumber: number;
	}

	interface Computer {
		id: number;
		name: string;
		status: "operational" | "issue" | "maintenance";
		issues: Issue[] | null;
	}

	let computers: Computer[] = [];
	let filter = "all";

	onMount(() => {
		// Simulating data fetch
		computers = [
			{ id: 1, name: "PC-001", status: "operational", issues: null },
			{
				id: 2,
				name: "PC-002",
				status: "issue",
				issues: [
					{
						studentName: "John Doe",
						professorName: "Dr. Smith",
						section: "A",
						labName: "Lab 1",
						description: "Blue screen",
						timestamp: "2023-10-01T10:00:00Z",
						schoolYear: 2023,
						pcNumber: 2,
					},
				],
			},
			{ id: 3, name: "PC-003", status: "operational", issues: null },
			{
				id: 4,
				name: "PC-004",
				status: "issue",
				issues: [
					{
						studentName: "Jane Doe",
						professorName: "Dr. Brown",
						section: "B",
						labName: "Lab 2",
						description: "No internet connection",
						timestamp: "2023-10-02T11:00:00Z",
						schoolYear: 2023,
						pcNumber: 4,
					},
				],
			},
			{
				id: 5,
				name: "PC-005",
				status: "maintenance",
				issues: [
					{
						studentName: "Admin",
						professorName: "N/A",
						section: "N/A",
						labName: "Lab 3",
						description: "Scheduled update",
						timestamp: "2023-10-03T12:00:00Z",
						schoolYear: 2023,
						pcNumber: 5,
					},
				],
			},
			{ id: 6, name: "PC-006", status: "operational", issues: null },
		];
	});

	$: filteredComputers =
		filter === "all"
			? computers
			: computers.filter((pc) => pc.status === filter);

	function getStatusColor(status: Computer["status"]) {
		switch (status) {
			case "operational":
				return "bg-green-500";
			case "issue":
				return "bg-red-500";
			case "maintenance":
				return "bg-yellow-500";
			default:
				return "bg-gray-500";
		}
	}
</script>

<div class="min-h-screen bg-gray-100">
	<header class="bg-white shadow">
		<div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
			<h1 class="text-3xl font-bold text-gray-900">
				Computer Lab Dashboard
			</h1>
		</div>
	</header>
	<main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
		<div class="px-4 py-6 sm:px-0">
			<div class="mb-4">
				<label for="filter" class="mr-2">Filter by status:</label>
				<select
					id="filter"
					bind:value={filter}
					class="border rounded p-2"
				>
					<option value="all">All</option>
					<option value="operational">Operational</option>
					<option value="issue">Issue</option>
					<option value="maintenance">Maintenance</option>
				</select>
			</div>
			<div
				class="grid grid-cols-1 gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4"
			>
				{#each filteredComputers as pc (pc.id)}
					<div class="bg-white overflow-hidden shadow rounded-lg">
						<div class="px-4 py-5 sm:p-6">
							<div class="flex items-center">
								<div
									class={`flex-shrink-0 rounded-md p-3 ${getStatusColor(pc.status)}`}
								>
									<svg
										class="h-6 w-6 text-white"
										fill="none"
										viewBox="0 0 24 24"
										stroke="currentColor"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											stroke-width="2"
											d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
										/>
									</svg>
								</div>
								<div class="ml-5 w-0 flex-1">
									<dl>
										<dt
											class="text-sm font-medium text-gray-500 truncate"
										>
											{pc.name}
										</dt>
										<dd class="flex items-baseline">
											<div
												class="text-2xl font-semibold text-gray-900"
											>
												{pc.status}
											</div>
										</dd>
									</dl>
								</div>
							</div>
							{#if pc.issues}
								<div class="mt-3 text-sm text-red-600">
									Issue: {pc.issues[0].description}
								</div>
							{/if}
						</div>
					</div>
				{/each}
			</div>
		</div>
	</main>
</div>
