<script lang="ts">
	import { onDestroy, onMount } from "svelte";
	import "../app.css";
	import type { PageData } from "./+page";
	import { Issue } from "$lib/types/issues";
	import IssueComponent from "$lib/components/IssueComponent.svelte";
	import { invoke } from "@tauri-apps/api/core";
	import WebSocket from "@tauri-apps/plugin-websocket";
    import { APIENDPOINT } from "$lib/constants";

	let ws: WebSocket | null = null;
	export let data: PageData;
	let issues = data.issues;
	let sortDirection: "asc" | "desc" = "desc";
	$: sortedIssues = issues.sort((a, b) => {
		const timeA = Number.parseInt(a.timestamp, 10);
		const timeB = Number.parseInt(b.timestamp, 10);
		return sortDirection === "asc" ? timeA - timeB : timeB - timeA;
	});

	function sortIssues() {
		console.debug("Sorting issues", issues);
		sortDirection = sortDirection === "asc" ? "desc" : "asc";
	}

	onDestroy(async () => {
		console.debug("Closing websocket");

		if (ws !== null) {
			await ws.disconnect();
		}
	});

	onMount(async () => {

		for (let i = 0; i < issues.length; i++) {
			console.log(issues[i]);
		}

		if (window.WebSocket) {
			ws = await WebSocket.connect(`ws://${APIENDPOINT}/ws`);
			console.log("Connecting to websocket");

			ws.addListener((msg) => {
				if (Array.isArray(msg.data)) {
					let issue = Issue.decode(new Uint8Array(msg.data));
					issues = issues.concat(issue);
					invoke("notify", {
						title: `New Issue: PC-${issue.pcNumber} on ${issue.labRoom}`,
						body: `${issue.concern === "" ? "A new issue has been reported" : issue.concern}`,
					});
				} else {
					console.debug("Received message", msg.data);
				}
			});
		} else {
			console.debug("Websocket not supported");
		}
	});
</script>

<div class="issues-container">
	<header class="border-b bg-white">
		<div class="container flex h-14 items-center justify-between">
			<h1 class="text-lg font-semibold">My App</h1>
			<p>steven</p>
		</div>
	</header>
	{#if data.error !== ""}
		<p class="error">{data.error}</p>
	{:else if issues.length === 0}
		<p class="no-issues">No issues found</p>
	{:else}
		<!--<h1 class="text-2xl font-semibold mb-4">Issues</h1>-->
		<div class="flex justify-end mt-4 mb-4">
			<button
				on:click={sortIssues}
				class="inline-flex items-center px-4 py-2 bg-gray-100 hover:bg-gray-200 text-gray-800 text-sm font-medium rounded-md"
			>
				Sort by Date
				<svg
					class="ml-2 h-4 w-4"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
				>
					{#if sortDirection === "asc"}
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M3 4h13M3 8h9m-9 4h6m4 0l4-4m0 0l4 4m-4-4v12"
						/>
					{:else}
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M3 4h13M3 8h9m-9 4h9m5-4v12m0 0l-4-4m4 4l4-4"
						/>
					{/if}
				</svg>
			</button>
		</div>

		<div class="grid m-4 gap-4">
			{#each sortedIssues as issue}
				<IssueComponent bind:issue />
			{/each}
		</div>
	{/if}
</div>
