<script lang="ts">
	import { onMount } from "svelte";
	import "../app.css";
	import type { PageData } from "./+page";
	import { Issue } from "$lib/types/issues";

	export let data: PageData;
	let issues = data.issues;

	onMount(async () => {
		// TODO: listent for new issues with webscocket
		if (window.WebSocket) {
			let ws = new WebSocket("http://localhost:3333/ws");
			console.log("Connecting to websocket");

			ws.binaryType = "arraybuffer";

			ws.onopen = () => {
				console.debug("Connected to websocket");

				window.onclose = () => {
					console.debug("Closing websocket");
					ws.close();
				};
			};

			ws.onmessage = (evt) => {
				if (evt.data instanceof ArrayBuffer) {
					let arr = new Uint8Array(evt.data);
					issues = issues.concat(Issue.decode(arr));
				} else {
					console.debug("Received message", evt.data);
				}
			};

			ws.onerror = (evt) => {
				console.debug("Websocket error", evt);
			};

			ws.onclose = (evt) => {
				console.debug("Websocket closed", evt);
			};
		} else {
			console.debug("Websocket not supported");
		}
	});

	function getStatusColor(status: number) {
		if (status == 1) {
			return "bg-green-500";
		} else {
			return "bg-red-500";
		}
	}
</script>

<div class="issues-container">
	{#if issues.length === 0}
		<p class="no-issues">No issues found</p>
	{:else}
		{#each issues.reverse() as issue}
			<div class="bg-white overflow-hidden shadow rounded-lg mt-6 mb-6">
				<div class="px-4 py-5 sm:p-6 grid gap-4">
					<div class="flex items-center">
						<div
							class={`flex-shrink-0 rounded-md p-3 ${getStatusColor(issue.status)}`}
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
									PC-{issue.pcNumber}
								</dt>
								<!-- Concern -->
								<div>
									<p>
										{issue.concern}
									</p>
								</div>
							</dl>
						</div>
					</div>

					<!-- Issue List -->
					<div class="flex flex-col gap-2">
						<dt class="font-semibold truncate">Issues</dt>

						<ul>
							{#each issue.issues as i, index}
								<div class="flex items-center mb-4 group">
									<input
										id={`checkbox-${i}-${index}`}
										type="checkbox"
										value=""
										class="peer size-5 bg-white border-[2px] checked:text-green-500 border-gray-300 rounded focus:ring-green-500 focus:ring-2"
									/>
									<label
										for={`checkbox-${i}-${index}`}
										class="ms-2 select-none font-medium peer-checked:text-green-600 text-red-500"
										>{i}</label
									>
								</div>
							{/each}
						</ul>
					</div>
				</div>
			</div>
		{/each}
	{/if}
</div>

<style>
	.issues-container {
		padding: 1rem;
	}

	.loading,
	.no-issues {
		text-align: center;
		padding: 2rem;
		color: #666;
	}

	.issue-card {
		border: 1px solid #ddd;
		border-radius: 8px;
		padding: 1rem;
		margin-bottom: 1rem;
		background-color: white;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
	}

	.student-info {
		margin-top: 1rem;
		padding-top: 1rem;
		border-top: 1px solid #eee;
	}
</style>
