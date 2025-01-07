import { fetch } from "@tauri-apps/plugin-http";
import { FaultyComponent, Issue } from "$lib/types/issues";

type Server = {
	host: string;
	port: number;
}

// TODO: Make this into a component
export class IssueReporter {
	private form: HTMLFormElement;
	private server: Server;

	constructor(host: string, port: number) {
		this.server = {
			host,
			port,
		}

		console.log('IssueReporter initialized');
		this.form = document.getElementById('issue-form') as HTMLFormElement;
		console.log('Form:', this.form);
		this.initializeForm();
	}

	private initializeForm(): void {
		this.form.addEventListener('submit', this.handleSubmit.bind(this));
	}

	public setServer(host: string, port: number): void {
		this.server.host = host;
		this.server.port = port;
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
			urgency: parseInt((document.getElementById('urgency') as HTMLInputElement).value)
		};
	}

	private async postIssue(issueData: Issue): Promise<void> {
		console.log('Posting issue:', issueData);
		console.log('Encoded issue:', Issue.encode(issueData).finish());

		console.log('Host:', this.server.host);
		console.log('Port:', this.server.port);

		try {
			const response = await fetch(`http://${this.server.host}:${this.server.port}/reportIssue`, {
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
