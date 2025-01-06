<script lang="ts">
	import { onDestroy, onMount } from "svelte";
	import "../app.css";
	import csulogo from "$lib/assets/csulogo.png";
	import type { PageData } from "./+page";
	import { Issue } from "$lib/types/issues";
	import IssueComponent from "$lib/components/IssueComponent.svelte";
	import { invoke } from "@tauri-apps/api/core";
	import WebSocket from "@tauri-apps/plugin-websocket";

	import Settings from "lucide-svelte/icons/settings";
	import Database from "lucide-svelte/icons/database";

	import { Button } from "$lib/components/ui/button/index.js";
	import * as Card from "$lib/components/ui/card/index.js";
	import { Input } from "$lib/components/ui/input/index.js";
	import { Label } from "$lib/components/ui/label/index.js";
	import { fetch } from "@tauri-apps/plugin-http";
	import { store } from "$lib/app";
	import FilterChips from "$lib/components/FilterChips.svelte";

	let ws: WebSocket | null = null;
	export let data: PageData;
	let issues = data.issues;
	let sortDirection: "asc" | "desc" = "desc";

	let showFinished = true;

	$: sortedIssues = issues.sort((a, b) => {
		const timeA = Number.parseInt(a.timestamp, 10);
		const timeB = Number.parseInt(b.timestamp, 10);
		return sortDirection === "asc" ? timeA - timeB : timeB - timeA;
	});

	function sortIssues() {
		console.debug("Sorting issues", issues);
		sortDirection = sortDirection === "asc" ? "desc" : "asc";
	}

	function isFixed(issue: Issue) {
		return issue.faultyComponents.every((item) => item.fixed === true);
	}

	onDestroy(async () => {
		console.debug("Closing websocket");

		if (ws !== null) {
			await ws.disconnect();
		}
	});

	onMount(async () => {
		host = (await store.get("host")) ?? "";
		port = (await store.get("port")) ?? "";

		for (let i = 0; i < issues.length; i++) {
			console.log(issues[i]);
		}

		if (window.WebSocket) {
			ws = await WebSocket.connect(`ws://${host}:${port}/ws`);
			console.log("Connecting to websocket");

			ws.addListener((msg) => {
				const data = msg.data;
				if (Array.isArray(data)) {
					if (data[0] == 1) {
						data.shift();
						let issue = Issue.decode(new Uint8Array(data));
						issues = issues.concat(issue);
						invoke("notify", {
							title: `New Issue: PC-${issue.pcNumber} on ${issue.labRoom}`,
							body: `${issue.concern === "" ? "A new issue has been reported" : issue.concern}`,
						});
					}
				} else {
					console.debug("Received message", msg.data);
				}
			});
		} else {
			console.debug("Websocket not supported");
		}
	});

	const onComponentToggle = async (
		event: Event,
		issueIndex: number,
		componentIndex: number,
	) => {
		const issue = issues[issueIndex];
		const component = issue.faultyComponents[componentIndex];
		const status = !component.fixed;

		console.log(issue.faultyComponents, issue.id);

		// TODO: update database of issue toggle, do a get request
		// TODO: componentId, status[0/1]

		console.log(component.id);
		try {
			const response = await fetch(
				`http://${host}:${port}/updateStatus`,
				{
					method: "POST",
					headers: {
						"Content-Type": "application/x-www-form-urlencoded",
					},
					body: `id=${component.id}&status=${status ? 1 : 0}`,
				},
			);

			if (!response.ok) {
				throw new Error("Failed to report issue");
			}
		} catch (error) {
			console.error("Failed to report issue:", error);
			// TODO: show an error toast

			// INFO: revert checkbox
			const target = event.target as HTMLInputElement;
			target.checked = component.fixed;

			return;
		}

		component.fixed = status;
		issues = issues;
		return;
	};

	let settingsDialogRef: null | HTMLDialogElement = null;
	let endpointFormRef: null | HTMLFormElement = null;

	let host = "";
	let port = "";

	const onSettingsClick = async () => {
		if (settingsDialogRef === null) {
			console.error("settingsDialogRef is null");
			return;
		}

		let hostElement = document.getElementById("host") as HTMLInputElement;
		hostElement.value = host as string;

		let portElement = document.getElementById("port") as HTMLInputElement;
		portElement.value = port as string;

		settingsDialogRef.showModal();
	};

	const onEndpointSave = () => {
		if (endpointFormRef === null || settingsDialogRef === null) {
			console.error("endpointFormRef is null");
			return;
		}

		let formData = new FormData(endpointFormRef);
		host = (formData.get("host") as string) ?? "";
		port = (formData.get("port") as string) ?? "";

		store.set("host", host);
		store.set("port", port);

		store.save();

		settingsDialogRef?.close();
	};
</script>

<div class="issues-container">
	<!-- Header -->
	<header class="bg-[#5C4033] p-4 grid grid-cols-[auto,1fr]">
		<div
			class="container mx-auto flex items-center gap-4 justify-center md:col-start-2"
		>
			<img src={csulogo} alt="University Logo" class="w-16 h-16" />
			<h1 class="text-white text-2xl md:text-3xl font-semibold">
				Cagayan State University
			</h1>
		</div>
		<div
			class="flex justify-self-end items-center self-center justify-end md:col-start-3 col-start-2 bg-black/10 rounded-full px-2 py-1"
		>
			<a href="/database" class="hover:bg-[#4B3329] rounded-full p-2">
				<Database color="white" />
			</a>
			<button
				on:click={onSettingsClick}
				class="hover:bg-[#4B3329] rounded-full p-2"
			>
				<Settings color="white" />
			</button>
		</div>
	</header>
	<!-- Header -->

	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
	<dialog
		on:click|self={() => {
			if (settingsDialogRef !== null) {
				settingsDialogRef.close();
			}
		}}
		class="rounded-lg border shadow-lg border-gray-300 backdrop:bg-background/80 backdrop:fixed backdrop:inset-0 backdrop:backdrop-blur-sm min-w-80"
		bind:this={settingsDialogRef}
	>
		<Card.Header>
			<Card.Title>Configure server</Card.Title>
			<Card.Description>Set the server endpoint</Card.Description>
		</Card.Header>
		<Card.Content>
			<form bind:this={endpointFormRef}>
				<div class="grid w-full items-center gap-4">
					<div class="flex flex-col space-y-1.5">
						<Label for="host">Host</Label>
						<Input id="host" name="host" placeholder="localhost" />
					</div>
					<div class="flex flex-col space-y-1.5">
						<Label for="port">Port</Label>
						<Input
							id="port"
							name="port"
							type="number"
							placeholder="8080"
						/>
					</div>
				</div>
			</form>
		</Card.Content>
		<Card.Footer class="flex justify-between">
			<Button
				on:click={() => {
					settingsDialogRef?.close();
				}}
				variant="outline">Cancel</Button
			>
			<Button on:click={() => onEndpointSave()}>Save</Button>
		</Card.Footer>
	</dialog>

	{#if data.error !== ""}
		<p class="error">{data.error}</p>
	{:else if issues.length === 0}
		<p class="no-issues">No issues found</p>
	{:else}
		<div class="flex justify-between mt-4 mb-4 mx-8">
			<div>
				<FilterChips bind:checked={showFinished} label="Finished"
				></FilterChips>
			</div>

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
			{#each sortedIssues as issue, issueIndex (issue.id)}
				{#if !isFixed(issue) || showFinished}
					<IssueComponent
						bind:issue
						onComponentToggle={(event, componentIndex) => {
							onComponentToggle(
								event,
								issueIndex,
								componentIndex,
							);
						}}
					/>
				{/if}
			{/each}
		</div>
	{/if}
</div>
