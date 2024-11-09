<script lang="ts">
	import { onMount } from "svelte";
	import { APIENDPOINT } from "../lib/constants";
	import "../app.css";
	import { Issue, IssueList } from "../lib/types/issues";

	let issues: Issue[];

	onMount(async () => {
		try {
			let res = await fetch(`${APIENDPOINT}/getIssues`, {
				method: "GET",
				headers: {
					"Content-Type": "application/x-protobuf",
				},
			});

			if (!res.ok) {
				throw new Error(`Failed to get issues ${res.status}`);
			}

			let data = await res.arrayBuffer();
			let issueList = IssueList.decode(new Uint8Array(data));
			issues = issueList.issues;
			console.log("Issues:", issues[0]);
		} catch (error) {
			console.debug("Failed to report issue:", error);
		}

		// TODO: listent for new issues with webscocket
		if (window.WebSocket) {
			let ws = new WebSocket("http://localhost:3333/ws");
			ws.onopen = () => {
				console.debug("Connected to websocket");
			};

			ws.onmessage = (msg) => {
				console.debug("Got message:", msg);
			};

			ws.onclose = () => {
				console.debug("Websocket closed");
			};
		} else {
			console.debug("Websocket not supported");
		}
	});
</script>

<div class="issues-container">
	{#if issues === undefined}
		<p class="loading">Loading issues...</p>
	{:else if issues.length === 0}
		<p class="no-issues">No issues found</p>
	{:else}
		{#each issues as issue}
			<div class="issue-card">
				<h3>{issue.student?.firstName} {issue.student?.lastName}</h3>
				<div class="issue-details">
					<!--
					<p><strong>Issue ID:</strong> {issue.id}</p>
					<p><strong>Status:</strong> {issue.status}</p>
					-->
					<p>
						<strong>Created:</strong>
						{new Date(Number.parseInt(issue.timestamp, 10)).toLocaleString()}
					</p>
					{#if issue.concern}
						<p><strong>Description:</strong> {issue.concern}</p>
					{/if}
					{#if issue.student}
						<div class="student-info">
							<p><strong>Student Info:</strong></p>
							<p>
								Name: {issue.student.firstName}
								{issue.student.lastName}
							</p>
							<p>
								Year/Section: {issue.student.year}/{issue
									.student.section}
							</p>
						</div>
					{/if}
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

	.issue-card h3 {
		margin: 0 0 1rem 0;
		color: #333;
	}

	.issue-details {
		display: grid;
		gap: 0.5rem;
	}

	.issue-details p {
		margin: 0;
	}

	.student-info {
		margin-top: 1rem;
		padding-top: 1rem;
		border-top: 1px solid #eee;
	}
</style>
