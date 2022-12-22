/* eslint-disable @typescript-eslint/no-explicit-any */
export class FileUtil {




	/**
	 * 下载文本文件
	 * @param filename 
	 */
	public static download(content: string | Blob | any, filename: string): void {
		let blob!: Blob;
		if (content instanceof Blob) {
			blob = content;
		} else if (typeof content === 'string') {
			blob = new Blob([content], { type: 'txt' });
		} else {
			blob = new Blob([JSON.stringify(content)], { type: 'txt' });
		}
		const objurl = URL.createObjectURL(blob);
		const link = document.createElement("a");
		link.download = filename;
		link.href = objurl;
		link.target = '_blank';
		link.click();
	}




	/**
	 * 弹出对话框选择一个本地文件\
	 * 如果取消选择操作将会堵塞线程一段时间(5s)
	 * @param fileformat 文件格式
	 * @result 选择的文件对象，点击取消按钮时返回null
	 */
	public static async selectLocalFile(fileformat: string, multiple = false): Promise<File[]> {
		return new Promise<File[]>((resolve, reject) => {
			const fileInput = document.createElement('input');
			fileInput.multiple = multiple;
			fileInput.type = 'file';
			fileInput.accept = fileformat;
			fileInput.click();
			fileInput.addEventListener('change', () => {
				const files = [];
				if (fileInput.files!.length > 0) {
					for (let i = 0; i < fileInput.files!.length; i++) {
						files.push(fileInput.files![i]);
					}
				}
				resolve(files);
			});
			window.addEventListener('focus', () => {
				setTimeout(() => {
					if (fileInput.files!.length == 0) {
						resolve([]);
					}
				}, 5000);
			}, { once: true })
		});
	}






	/**
	 * 从本地文件系统读取Json对象
	 * @param file 
	 */
	public static async loadJsonFromFile<T>(file: File): Promise<T> {
		return new Promise<T>((resolve, reject) => {
			const reader = new FileReader();
			reader.readAsText(file);
			reader.onload = () => {
				try {
					resolve(JSON.parse(reader.result!.toString()));
				}
				catch (e) {
					reject(new Error('Invalid JSON file format。'));
				}
			};
			reader.onerror = (ex) => {
				reject(ex);
			}
		});
	}


	/**
	 * 从本地文件系统中读取文本数据
	 * @param file 本地文件
	 */
	public static async loadTextFromFile(file: File): Promise<string> {
		return new Promise<string>((resolve, reject) => {
			const reader = new FileReader();
			reader.readAsText(file);
			reader.onload = () => {
				resolve(reader.result!.toString());
			};
			reader.onerror = (ex) => {
				reject(ex);
			}
		});
	}


	/**
	 * 从本地文件系统中读取文本数据
	 * @param file 本地文件
	 */
	public static async loadImageUrlFromFile(file: File): Promise<string> {
		return new Promise<string>((resolve, reject) => {
			const reader = new FileReader();
			reader.readAsDataURL(file);
			reader.onload = () => {
				resolve(reader.result!.toString());
			};
			reader.onerror = (ex) => {
				reject(ex);
			}
		});
	}




	public static async loadBlobFromFile(file: File): Promise<ArrayBuffer> {
		return new Promise<ArrayBuffer>((resolve, reject) => {
			const reader = new FileReader();
			reader.readAsArrayBuffer(file);
			reader.onload = () => {
				resolve(reader.result as ArrayBuffer);
			};
			reader.onerror = (ex) => {
				reject(ex);
			}
		});
	}



}