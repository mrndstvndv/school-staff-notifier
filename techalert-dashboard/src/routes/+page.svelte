<script>
	// Mock data for demonstration
	const totalComputers = 30;
	const nonWorkingComputers = 5;

	const computers = Array(totalComputers)
		.fill(undefined, 0, 30)
		.map((_, index) => ({
			id: index + 1,
			status: index < nonWorkingComputers ? "Not Working" : "Working",
			issues:
				index < nonWorkingComputers
					? [
							"Keyboard not working",
							"App out of date",
							"Required IDE not installed",
						][index % 3]
					: null,
		}));

	const recentNotifications = [
		{
			id: 1,
			message: "Computer #3 reported keyboard issue",
			time: "2 hours ago",
		},
		{
			id: 2,
			message: "Computer #7 missing required software",
			time: "4 hours ago",
		},
		{ id: 3, message: "Computer #12 app update needed", time: "1 day ago" },
	];
</script>

<div class="dashboard">
	<header>
		<h1>Computer Lab Dashboard</h1>
	</header>

	<div class="content">
		<div class="top-widgets">
			<div class="card summary-widget">
				<h2>Lab Status Summary</h2>	
				<p class="text-lg">
					<strong>{nonWorkingComputers}</strong> out of
					<strong>{totalComputers}</strong> computers are currently not
					working
				</p>
			</div>

			<div class="card notification-widget">
				<h2>Recent Notifications</h2>
				<ul class="notification-list">
					{#each recentNotifications as notification}
						<li>
							<p>{notification.message}</p>
							<small>{notification.time}</small>
						</li>
					{/each}
				</ul>
			</div>
		</div>

		<main>
			<div class="card computer-grid">
				<h2>Computer Status</h2>
				<div class="grid-container">
					{#each computers as computer}
						<div
							class="computer-item {computer.status ===
							'Not Working'
								? 'not-working'
								: ''}"
						>
							<h3>Computer #{computer.id}</h3>
							<p class="status">{computer.status}</p>
							{#if computer.issues}
								<p class="issues">{computer.issues}</p>
							{/if}
						</div>
					{/each}
				</div>
			</div>
		</main>
	</div>
</div>

<style>
	:root {
		--background: hsl(0 0% 100%);
		--foreground: hsl(222.2 84% 4.9%);
		--card: hsl(0 0% 100%);
		--card-foreground: hsl(222.2 84% 4.9%);
		--popover: hsl(0 0% 100%);
		--popover-foreground: hsl(222.2 84% 4.9%);
		--primary: hsl(222.2 47.4% 11.2%);
		--primary-foreground: hsl(210 40% 98%);
		--secondary: hsl(210 40% 96.1%);
		--secondary-foreground: hsl(222.2 47.4% 11.2%);
		--muted: hsl(210 40% 96.1%);
		--muted-foreground: hsl(215.4 16.3% 46.9%);
		--accent: hsl(210 40% 96.1%);
		--accent-foreground: hsl(222.2 47.4% 11.2%);
		--warning: hsl(38 92% 50%);
		--warning-foreground: hsl(38 92% 10%);
		--border: hsl(214.3 31.8% 91.4%);
		--input: hsl(214.3 31.8% 91.4%);
		--ring: hsl(222.2 84% 4.9%);
		--radius: 0.5rem;
	}

	.dark {
		--background: hsl(222.2 84% 4.9%);
		--foreground: hsl(210 40% 98%);
		--card: hsl(222.2 84% 4.9%);
		--card-foreground: hsl(210 40% 98%);
		--popover: hsl(222.2 84% 4.9%);
		--popover-foreground: hsl(210 40% 98%);
		--primary: hsl(210 40% 98%);
		--primary-foreground: hsl(222.2 47.4% 11.2%);
		--secondary: hsl(217.2 32.6% 17.5%);
		--secondary-foreground: hsl(210 40% 98%);
		--muted: hsl(217.2 32.6% 17.5%);
		--muted-foreground: hsl(215 20.2% 65.1%);
		--accent: hsl(217.2 32.6% 17.5%);
		--accent-foreground: hsl(210 40% 98%);
		--warning: hsl(38 92% 50%);
		--warning-foreground: hsl(38 92% 90%);
		--border: hsl(217.2 32.6% 17.5%);
		--input: hsl(217.2 32.6% 17.5%);
		--ring: hsl(212.7 26.8% 83.9%);
	}

	.dashboard {
		font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
			"Helvetica Neue", Arial, sans-serif;
		background-color: var(--background);
		color: var(--foreground);
		line-height: 1.5;

		max-width: 1200px;
		margin: 0 auto;
		padding: 20px;
	}

	header {
		background-color: var(--primary);
		color: var(--primary-foreground);
		padding: 20px;
		margin-bottom: 20px;
		border-radius: var(--radius);
	}

	h1,
	h2,
	h3 {
		margin: 0;
		font-weight: 600;
	}

	h1 {
		font-size: 1.5rem;
	}

	h2 {
		font-size: 1.25rem;
		margin-bottom: 1rem;
	}

	h3 {
		font-size: 1rem;
	}

	.content {
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	.top-widgets {
		display: flex;
		gap: 20px;
		flex-wrap: wrap;
	}

	.card {
		background-color: var(--card);
		color: var(--card-foreground);
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 20px;
	}

	.summary-widget,
	.notification-widget {
		flex: 1;
		min-width: 250px;
	}

	.summary-widget {
		display: flex;
		flex-direction: column;
		justify-content: center;
		text-align: center;

		background-color: var(--accent);
		color: var(--accent-foreground);
	}

	p {
		margin: 0;
	}

	.text-lg {
		font-size: 1.125rem;
	}

	.grid-container {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
		gap: 1rem;
	}

	.computer-item {
		background-color: var(--secondary);
		color: var(--secondary-foreground);
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 1rem;
		transition: all 0.3s ease;
	}

	.computer-item:hover {
		transform: translateY(-2px);
		box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
	}

	.computer-item.not-working {
		background-color: var(--warning);
		color: var(--warning-foreground);
	}

	.status {
		font-weight: 600;
		margin: 0.5rem 0;
	}

	.issues {
		font-size: 0.875rem;
		color: var(--warning-foreground);
		opacity: 0.8;
	}

	.notification-list {
		list-style-type: none;
		padding: 0;
		margin: 0;
	}

	.notification-list li {
		margin-bottom: 15px;
		border-bottom: 1px solid var(--border);
		padding-bottom: 10px;
	}

	.notification-list li:last-child {
		border-bottom: none;
		margin-bottom: 0;
		padding-bottom: 0;
	}

	small {
		color: var(--muted-foreground);
		display: block;
		margin-top: 5px;
	}

	@media (max-width: 768px) {
		.top-widgets {
			flex-direction: column;
		}

		.summary-widget,
		.notification-widget {
			width: 100%;
		}

		.grid-container {
			grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
		}
	}
</style>

