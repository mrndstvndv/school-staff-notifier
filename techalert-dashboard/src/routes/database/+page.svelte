<script lang="ts">
	import "../../app.css";
	import csulogo from "$lib/assets/csulogo.png";

	import * as Tabs from "$lib/components/ui/tabs";
	import { Input } from "$lib/components/ui/input/index.js";
	import { Label } from "$lib/components/ui/label/index.js";
	import ArrowLeft from "lucide-svelte/icons/arrow-left";

	import Dbtable from "./_components/Dbtable.svelte";

	import { onMount } from "svelte";

	import {
		store,
		getDbInstance,
		selectFromTable,
		insertToTable,
	} from "$lib/app";
	import type Database from "@tauri-apps/plugin-sql";

	let students: Object[] = [];
	let studentForm: HTMLFormElement;

	let teachers: Object[] = [];
	let teacherForm: HTMLFormElement;

	let subjects: object[] = [];
	let subjectForm: HTMLFormElement;

	let courses: object[] = [];
	let courseForm: HTMLFormElement;

	let db: Database | null = null;

	async function onAdd(
		table: string,
		form: HTMLFormElement,
		refreshCallback: () => {},
	) {
		if (db == null) return;

		let value = Object.fromEntries(new FormData(form));
		await insertToTable(db, value, table);

		refreshCallback();
	}

	async function getTeachers() {
		if (db == null) {
			console.log("Db is not initialized yet");
			return;
		}

		teachers = (await selectFromTable(
			db,
			["first_name", "last_name"],
			"teacher_table",
		)) as Object[];
	}

	async function getStudents() {
		if (db == null) {
			console.log("Db is not initialized yet");
			return;
		}

		students = (await selectFromTable(
			db,
			["first_name", "last_name", "course", "year", "section"],
			"student_table",
		)) as Object[];
	}

	async function getSubjects() {
		if (db == null) {
			console.log("Db is not initialized yet");
			return;
		}

		subjects = (await selectFromTable(
			db,
			[
				"subject_description",
				"subject_code",
				"course",
				"year",
				"teacher_id",
			],
			"subject_table",
		)) as Object[];
	}

	async function getCourses() {
		if (db == null) {
			console.log("Db is not initialized yet");
			return;
		}

		courses = (await selectFromTable(
			db,
			["code"],
			"course_table",
		)) as Object[];
	}

	onMount(async () => {
		const host: string = (await store.get("host")) as string;
		db = await getDbInstance(host);

		getStudents();
		getTeachers();
		getSubjects();
		getCourses();
	});
</script>

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
		class="flex justify-self-end items-center self-center justify-end md:col-start-3 col-start-2 rounded-full px-2 py-1"
	>
		<button
			on:click={() => {
				window.history.back();
			}}
			class="hover:bg-[#4B3329] rounded-full p-2"
		>
			<ArrowLeft class="text-white"></ArrowLeft>
		</button>
	</div>
</header>

<div class="p-5">
	<Tabs.Root value="students">
		<Tabs.List class="bg-black/5">
			<Tabs.Trigger value="students">Students</Tabs.Trigger>
			<Tabs.Trigger value="teachers">Teachers</Tabs.Trigger>
			<Tabs.Trigger value="subjects">Subjects</Tabs.Trigger>
			<Tabs.Trigger value="courses">Courses</Tabs.Trigger>
		</Tabs.List>

		<Dbtable
			value="students"
			bind:table={students}
			onAdd={() => {
				onAdd("student_table", studentForm, getStudents);
			}}
		>
			<form bind:this={studentForm} class="flex flex-col gap-4">
				<div class="flex flex-col space-y-1.5">
					<Label for="first_name">First Name</Label>
					<Input id="first_name" name="first_name" required />
				</div>

				<div class="flex flex-col space-y-1.5">
					<Label for="last_name">Last Name</Label>
					<Input id="last_name" name="last_name" required />
				</div>

				<div class="flex flex-col space-y-1.5">
					<Label for="course">Course</Label>
					<Input id="course" name="course" required />
				</div>

				<div class="flex gap-2">
					<div class="flex flex-col space-y-1.5">
						<Label for="year">Year</Label>
						<Input id="year" name="year" type="number" required />
					</div>

					<div class="flex flex-col space-y-1.5">
						<Label for="section">Section</Label>
						<Input id="section" name="section" required />
					</div>
				</div>
			</form>
		</Dbtable>

		<Dbtable
			value="teachers"
			table={teachers}
			onAdd={() => {
				onAdd("teacher_table", teacherForm, getTeachers);
			}}
		>
			<form bind:this={teacherForm} class="flex flex-col gap-4">
				<div class="flex flex-col space-y-1.5">
					<Label for="first_name">First Name</Label>
					<Input id="first_name" name="first_name" required />
				</div>

				<div class="flex flex-col space-y-1.5">
					<Label for="last_name">Last Name</Label>
					<Input id="last_name" name="last_name" required />
				</div>
			</form>
		</Dbtable>

		<Dbtable
			value="subjects"
			table={subjects}
			onAdd={() => {
				onAdd("subject_table", subjectForm, getSubjects);
			}}
		>
			<form bind:this={subjectForm} class="flex flex-col gap-4">
				<div class="flex flex-col space-y-1.5">
					<Label for="subject_description">Subject Description</Label>
					<Input
						id="subject_description"
						name="subject_description"
						required
					/>
				</div>

				<div class="flex flex-col space-y-1.5">
					<Label for="subject_code">Subject Code</Label>
					<Input id="subject_code" name="subject_code" required />
				</div>

				<div class="flex flex-col space-y-1.5">
					<Label for="course">Course</Label>
					<Input id="course" name="course" required />
				</div>

				<div class="flex flex-col space-y-1.5">
					<Label for="year">Year</Label>
					<Input id="year" name="year" type="number" required />
				</div>

				<div class="flex flex-col space-y-1.5">
					<Label for="teacher_id">Teacher</Label>
					<Input
						id="teacher_id"
						name="teacher_id"
						type="number"
						required
					/>
				</div>
			</form>
		</Dbtable>

		<Dbtable
			value="courses"
			table={courses}
			onAdd={() => {
				onAdd("course_table", courseForm, getCourses);
			}}
		>
			<form bind:this={courseForm}>
				<div class="flex flex-col space-y-1.5">
					<Label for="code">Code</Label>
					<Input id="code" name="code" required />
				</div>
			</form>
		</Dbtable>
	</Tabs.Root>
</div>
