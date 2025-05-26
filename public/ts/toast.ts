export default class Toast {
	private static container: HTMLDivElement;

	private static initializeContainer() {
		if (!this.container) {
			this.container = document.createElement('div');
			this.container.style.position = 'fixed';
			this.container.style.top = '20px';
			this.container.style.left = '50%';
			this.container.style.transform = 'translateX(-50%)';
			this.container.style.zIndex = '1000';
			this.container.style.display = 'flex';
			this.container.style.flexDirection = 'column';
			this.container.style.alignItems = 'center';
			this.container.style.gap = '10px';
			document.body.appendChild(this.container);
		}
	}

	private static createToastElement(message: string, type: 'success' | 'error') {
		const toast = document.createElement('div');
		toast.textContent = message;
		toast.style.padding = '10px 20px';
		toast.style.borderRadius = '5px';
		toast.style.color = '#fff';
		toast.style.fontSize = '14px';
		toast.style.boxShadow = '0 2px 5px rgba(0, 0, 0, 0.2)';
		toast.style.opacity = '1';
		toast.style.transition = 'opacity 0.5s ease, transform 0.5s ease';
		toast.style.transform = 'translateY(0)';

		if (type === 'success') {
			toast.style.backgroundColor = '#4caf50';
		} else if (type === 'error') {
			toast.style.backgroundColor = '#f44336';
		}

		setTimeout(() => {
			toast.style.opacity = '0';
			toast.style.transform = 'translateY(-20px)';
			setTimeout(() => {
				toast.remove();
			}, 500);
		}, 3000);

		return toast;
	}

	static success(message: string) {
		this.initializeContainer();
		const toast = this.createToastElement(message, 'success');
		this.container.appendChild(toast);
	}

	static error(message: string) {
		this.initializeContainer();
		const toast = this.createToastElement(message, 'error');
		this.container.appendChild(toast);
	}
}