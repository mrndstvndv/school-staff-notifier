import { fetch } from "@tauri-apps/plugin-http";
import { FaultyComponent, Issue } from "$lib/types/issues";

export class IssueReporter {
	private form: HTMLFormElement;

	constructor() {
		console.log('IssueReporter initialized');
		this.form = document.getElementById('issue-form') as HTMLFormElement;
		console.log('Form:', this.form);
		this.initializeForm();
	}

	private initializeForm(): void {
		this.form.addEventListener('submit', this.handleSubmit.bind(this));
	}

	private getFormData(): Issue {
		const issueCheckboxes = document.querySelectorAll('input[type="checkbox"]:checked');
		const faultyComponents: FaultyComponent[] = Array.from(issueCheckboxes).map((checkbox) => ({
			name: (checkbox as HTMLInputElement).nextElementSibling?.textContent || '',
			parentId: 0,
			id: 0,
			fixed: false,
		}));

		return {
			student: {
				firstName: (document.getElementById('first-name') as HTMLInputElement).value,
				lastName: (document.getElementById('last-name') as HTMLInputElement).value,
				year: parseInt((document.getElementById('year') as HTMLInputElement).value),
				section: (document.getElementById('section') as HTMLInputElement).value,
				// WARN: Add course
				course: (document.getElementById('course') as HTMLInputElement).value,
				professor: (document.getElementById('professor') as HTMLInputElement).value
			},
			labRoom: (document.getElementById('lab-room') as HTMLSelectElement).value,
			pcNumber: parseInt((document.getElementById('pc-number') as HTMLInputElement).value),
			faultyComponents: faultyComponents,
			timestamp: Date.now().toString(),
			concern: (document.querySelector('textarea') as HTMLTextAreaElement).value,
			id: 0,
		};
	}

	private async postIssue(issueData: Issue): Promise<void> {
		console.log('Posting issue:', issueData);
		console.log('Encoded issue:', Issue.encode(issueData).finish());

		try {
			const response = await fetch("http://localhost:3333/reportIssue", {
				method: 'POST',
				headers: {
					'Content-Type': 'application/x-protobuf',
				},
				body: Issue.encode(issueData).finish(),
			});

			if (!response.ok) {
				throw new Error('Failed to report issue');
			}
		} catch (error) {
			console.error('Failed to report issue:', error);
		}
	}

	private async handleSubmit(e: Event): Promise<void> {
		e.preventDefault();

		const issueData = this.getFormData();

		await this.postIssue(issueData);
	}
}
