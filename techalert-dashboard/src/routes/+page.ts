import { APIENDPOINT } from "$lib/constants";
import { Issue, IssueList } from "$lib/types/issues";

export type PageData = {
	issues: Issue[];
	error: string;
};

export const load: ({ fetch }: { fetch: any; }) => Promise<PageData> = async ({ fetch }) => {
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
		return {
			issues: issueList.issues,
			error: ""
		}
	} catch (error) {
		console.debug("Failed to report issue:", error);
	}

	return {
		issues: [],
		error: "Failed to get issues"
	};
}
