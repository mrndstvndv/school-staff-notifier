import { Issue } from "./issues";

class IssueReporter {
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
		const issues = Array.from(issueCheckboxes).map(checkbox => (checkbox as HTMLInputElement).nextElementSibling?.textContent || '');

		return {
			student: {
				firstName: (document.getElementById('first-name') as HTMLInputElement).value,
				lastName: (document.getElementById('last-name') as HTMLInputElement).value,
				year: parseInt((document.getElementById('year') as HTMLInputElement).value),
				section: (document.getElementById('section') as HTMLInputElement).value,
				// WARN: Add course
				course: "CS",
				professor: (document.getElementById('professor') as HTMLInputElement).value
			},
			labRoom: (document.getElementById('lab-room') as HTMLSelectElement).value,
			pcNumber: parseInt((document.getElementById('pc-number') as HTMLInputElement).value),
			issues: issues,
			timestamp: Date.now().toString(),
			concern: (document.querySelector('textarea') as HTMLTextAreaElement).value,
			id: 0,
			status: 0,
		};
	}

	private async postIssue(issueData: Issue): Promise<void> {
		console.debug('Posting issue:', issueData);
		console.debug('Encoded issue:', Issue.encode(issueData).finish());

		try {
			const response = await fetch("http://localhost:3333/reportIssue", {
				method: 'POST',
				headers: {
					'Content-Type': 'application/octet-stream',
				},
				body: Issue.encode(issueData).finish(),
			});

			if (!response.ok) {
				throw new Error('Failed to report issue');
			}

			//const json = await response.json();
			//console.debug('Issue reported successfully:', json);
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

// Initialize the issue reporter when the DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
	new IssueReporter();
});
