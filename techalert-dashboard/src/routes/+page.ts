import { store } from "$lib/app";
import { Issue, IssueList } from "$lib/types/issues";
import { fetch } from "@tauri-apps/plugin-http";

export type PageData = {
	issues: Issue[];
	error: string;
};

export const load: () => Promise<PageData> = async () => {
	let error = "";

	const host = await store.get("host");
	const port = await store.get("port");

	try {
		let res = await fetch(`http://${host}:${port}/getIssues`, {
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
		return {
			issues: issueList.issues,
			error: ""
		}
	} catch (e) {
		console.debug("Failed to report issue:", e);
		error = `Failed to get issues: ${e}`;
	}

	return {
		issues: [],
		error: error
	};
}
