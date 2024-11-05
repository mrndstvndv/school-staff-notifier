interface Issue {
	studentName: string;
	professorName: string;
	section: string;
	labName: string;
	description: string;
	timestamp: string;
	schoolYear: number;
	pcNumber: number;
}

class IssueReporter {
	private form: HTMLFormElement;

	constructor() {
		this.form = document.getElementById('issueForm') as HTMLFormElement;
		this.initializeForm();
	}

	private initializeForm(): void {
		this.form.addEventListener('submit', this.handleSubmit.bind(this));
	}

	private getFormData(): Issue {
		return {
			studentName: (document.getElementById('studentName') as HTMLInputElement).value,
			professorName: (document.getElementById('professorName') as HTMLInputElement).value,
			section: (document.getElementById('section') as HTMLInputElement).value,
			labName: (document.getElementById('labName') as HTMLInputElement).value,
			description: (document.getElementById('description') as HTMLTextAreaElement).value,
			timestamp: new Date().toISOString(),
			schoolYear: parseInt((document.getElementById('schoolYear') as HTMLInputElement).value),
			pcNumber: parseInt((document.getElementById('pcNumber') as HTMLInputElement).value)
		};
	}

	private async postIssue(issueData: Issue): Promise<void> {
		try {
			const response = await fetch("http://localhost:3333/reportIssue", {
				method: 'POST',
				body: JSON.stringify(issueData)
			});

			if (!response.ok) {
				console.log('Failed to report issue!');
			}

			const json = await response.json();
			console.log(json);
		} catch (error) {
			console.log('Failed to report issue!');
		}
	}

	private async handleSubmit(e: Event): Promise<void> {
		e.preventDefault();

		const issueData = this.getFormData();
		console.log('Issue reported:', issueData);

		// TODO: send issueData to the backend
		await this.postIssue(issueData);
	}
}

// Initialize the issue reporter when the DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
	new IssueReporter();
});
