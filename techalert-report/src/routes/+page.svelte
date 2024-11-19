<script lang="ts">
	import "../styles.css";
	import csulogo from "$lib/assets/csulogo.png";
	import { onMount } from "svelte";
	import { IssueReporter } from "$lib/types/IssueReporter";
	import Search from "$lib/components/Search.svelte";
	import Settings from "lucide-svelte/icons/settings";
	import { LazyStore } from "@tauri-apps/plugin-store";

	let store: LazyStore;

	let host = "localhost";
	let port = 8080;

	const courses = [
		"BSIT",
		"BSCS",
		"BSIS",
		"BSCE",
		"BSME",
		"BSIE",
		"BSEE",
		"BSA",
	];

	const courseToProf = new Map([
		["BSIT", []],
		[
			"BSCS",
			[
				"DELL CASTILLO, YSSANDREA KNERRE Z",
				"CIELOS, ZYBER JOHN G",
				"DUMAYAG, DEXTER P",
				"SEMILLA, MARIA CRISTINA C",
				"ARESGA, ROLLY JAMES L",
				"BRAVO, EDISON D",
				"CARLET, NICOLE",
			],
		],
		["BSIS", []],
		["BSCE", []],
	]);

	let course: string = courses[0];
	$: professors = courseToProf.get(course) ?? [];

	let form: HTMLFormElement;
	let issueReporter: IssueReporter;

	onMount(async () => {
		store = new LazyStore("settings.json");

		host = (await store.get("host")) ?? host;
		port = (await store.get("port")) ?? port;

		issueReporter = new IssueReporter(host, port);
	});

	function submitDialog(event: Event) {
		const target = event.target as HTMLFormElement;

		if (target === null) throw new Error("Event target is null");

		const formData = new FormData(target);

		host = (formData.get("host") as string) ?? host;
		port = Number.parseInt(formData.get("port") as string) ?? port;

		store?.set("host", host);
		store?.set("port", port);

		issueReporter.setServer(host, port);
		store.save();

		settingsDialogRef.close();
	}

	let settingsDialogRef: HTMLDialogElement;
	function onSettingsClick() {
		settingsDialogRef.showModal();
	}

	// TODO: Student names should be searchable based on year/section and course
	// TODO: Separate the student info from the reporting
</script>

<div class="min-h-screen">
	<!-- Header -->
	<header class="bg-[#5C4033] p-4 grid md:grid-cols-3 grid-cols-[auto,1fr]">
		<div
			class="container mx-auto flex items-center gap-4 justify-center md:col-start-2"
		>
			<img src={csulogo} alt="University Logo" class="w-16 h-16" />
			<h1 class="text-white text-2xl md:text-3xl font-semibold">
				Cagayan State University
			</h1>
		</div>
		<div class="flex items-center justify-end md:col-start-3 col-start-2">
			<button
				on:click={onSettingsClick}
				class="hover:bg-[#4B3329] rounded-md p-2"
			>
				<Settings class="col-start-3" color="white" />
			</button>
		</div>
	</header>

	<!-- TODO: add functionality to the dialog -->
	<!-- svelte-ignore a11y_click_events_have_key_events, a11y_no_static_element_interactions (because of reasons) -->
	<dialog
		on:click|self={() => settingsDialogRef.close()}
		bind:this={settingsDialogRef}
		class="p-6 rounded-lg border shadow-lg border-gray-300 backdrop:bg-background/80 backdrop:fixed backdrop:inset-0 backdrop:backdrop-blur-sm"
	>
		<h2 class="font-semibold text-lg">Endpoint</h2>
		<p class="text-sm mt-[6px]" style="color: hsl(240 3.8% 46.1%);">
			Set the endpoint for the issue reporting system.
		</p>
		<form class="py-4 grid gap-4" on:submit|preventDefault={submitDialog}>
			<div class="grid grid-cols-4 items-center gap-4">
				<label for="host" class="text-right text-sm font-medium"
					>Host</label
				>
				<input
					value={host}
					required
					class="col-span-3"
					type="text"
					name="host"
					id="host"
					placeholder="192.168.1.1"
				/>
			</div>
			<div class="grid grid-cols-4 items-center gap-4">
				<label for="port" class="text-right text-sm font-medium"
					>Port</label
				>
				<input
					value={port}
					required
					class="col-span-3"
					type="number"
					id="port"
					name="port"
					placeholder="8080"
				/>
			</div>

			<div class="flex justify-end mt-4">
				<button
					class="bg-green-700 text-white px-8 py-2 rounded-md hover:bg-green-800 transition-colors inline-flex items-center gap-2 text-center"
				>
					Save Changes
				</button>
			</div>
		</form>
	</dialog>

	<!-- Main Form -->
	<main class="container mx-auto p-4 md:p-8">
		<div class="bg-white rounded-lg shadow-lg p-6 md:p-8">
			<form
				class="grid md:grid-cols-2 grid-rows-[auto,1fr] gap-8"
				id="issue-form"
				bind:this={form}
			>
				<!-- Left Column - Input Fields -->
				<div class="space-y-5">
					<div class="grid md:grid-cols-3 gap-5">
						<div>
							<label
								for="course"
								class="block text-gray-700 font-semibold mb-2 text-sm"
								>Course</label
							>
							<Search
								items={courses}
								bind:value={course}
								classNames="w-full"
								name="course"
								id="course"
								{form}
								required
							/>
						</div>

						<div>
							<label
								for="year"
								class="block text-gray-700 font-semibold mb-2 text-sm"
								>Year</label
							>
							<select
								required
								class="w-full"
								name="year"
								id="year"
							>
								<option value="1">1</option>
								<option value="2">2</option>
								<option value="3">3</option>
								<option value="4">4</option>
							</select>
						</div>

						<div>
							<label
								for="section"
								class="block text-gray-700 font-semibold mb-2 text-sm"
								>Section</label
							>
							<select
								required
								class="w-full"
								name="section"
								id="section"
							>
								<option value="A">A</option>
								<option value="B">B</option>
								<option value="C">C</option>
								<option value="D">D</option>
							</select>
						</div>
					</div>

					<div class="grid md:grid-cols-2 gap-5">
						<div>
							<label
								for="first-name"
								class="block text-gray-700 font-semibold mb-2 text-sm"
								>First name</label
							>
							<input
								required
								class="w-full"
								type="text"
								id="first-name"
							/>
						</div>

						<div>
							<label
								for="last-name"
								class="block text-gray-700 font-semibold mb-2 text-sm"
								>Last name</label
							>
							<input
								required
								class="w-full"
								type="text"
								id="last-name"
							/>
						</div>
					</div>

					<div>
						<label
							for="professor"
							class="block text-gray-700 font-semibold mb-2 text-sm"
							>Professor</label
						>

						<!-- old proffessor input
						<input
							required
							class="w-full"
							type="text"
							name="professor"
							id="professor"
						/> -->

						<Search
							classNames="w-full"
							items={professors}
							name="professor"
							id="professor"
							value={professors[0]}
							{form}
							required
						/>
					</div>

					<div class="grid grid-cols-2 gap-5">
						<div>
							<label
								for="lab-room"
								class="block text-gray-700 font-semibold mb-2 text-sm"
								>Laboratory Room</label
							>
							<select
								class="w-full"
								name="lab-room"
								id="lab-room"
							>
								<option value="cla">CLA</option>
								<option value="clb">CLB</option>
								<option value="clc">CLC</option>
							</select>
						</div>

						<div>
							<label
								for="pc-number"
								class="block text-gray-700 font-semibold mb-2 text-sm"
								>PC number</label
							>
							<input
								required
								type="number"
								name="pc-number"
								id="pc-number"
								class="h-[40px] w-full"
							/>
						</div>
					</div>
				</div>

				<!-- Right Column - Concern Textarea -->
				<div class="flex flex-col gap-4">
					<div>
						<p
							class="block text-gray-700 font-semibold mb-2 text-sm"
						>
							Issues
						</p>
						<ul class="list-none grid grid-cols-2">
							<!-- TODO: Add guard for when non of the items are selected-->
							<label for="keyboard" class="flex items-center">
								<input id="keyboard" type="checkbox" />
								<span class="ml-2">Keyboard</span>
							</label>

							<label for="mouse" class="flex items-center">
								<input id="mouse" type="checkbox" />
								<span class="ml-2">Mouse</span>
							</label>

							<label for="monitor" class="flex items-center">
								<input id="monitor" type="checkbox" />
								<span class="ml-2">Monitor</span>
							</label>

							<label for="powersupply" class="flex items-center">
								<input id="powersupply" type="checkbox" />
								<span class="ml-2">Power supply</span>
							</label>
						</ul>
					</div>

					<textarea
						placeholder="Type your concern here."
						class="w-full row-span-2 h-[200px] md:h-full p-4 border-[1.5px] border-green-600 rounded-lg resize-none focus:ring-2 focus:ring-green-500 focus:border-green-500"
					></textarea>
				</div>

				<div class="flex justify-end md:col-start-2 md:row-start-2">
					<button
						type="submit"
						class="bg-green-700 text-white px-8 py-2 rounded-md hover:bg-green-800 transition-colors inline-flex items-center gap-2 text-center"
					>
						Send
						<!-- <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
							stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
							class="lucide lucide-send-horizontal">
							<path
								d="M3.714 3.048a.498.498 0 0 0-.683.627l2.843 7.627a2 2 0 0 1 0 1.396l-2.842 7.627a.498.498 0 0 0 .682.627l18-8.5a.5.5 0 0 0 0-.904z" />
							<path d="M6 12h16" />
						</svg> -->
					</button>
				</div>
			</form>
		</div>
	</main>
</div>

<style>
	label,
	h2,
	p {
		user-select: none;
	}

	input[type="text"],
	input[type="number"] {
		@apply border-[1.5px] border-gray-300 shadow-sm rounded-lg p-2 focus:outline-none focus:ring-[1.5px] focus:ring-green-500 focus:border-green-500;
	}

	input[type="checkbox"] {
		@apply w-4 h-4 text-green-600 border-gray-300 rounded focus:ring-blue-500 focus:ring-2;
	}

	button:hover {
		cursor: default;
	}

	select {
		@apply border-[1.5px] border-gray-300 shadow-sm rounded-lg p-2 focus:outline-none focus:ring-[1.5px] focus:ring-green-500 focus:border-green-500 bg-white;
	}
</style>
