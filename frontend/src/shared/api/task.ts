import Api from './api'

export interface Task {
	id: string
	user_id: string
	title: string
	is_active: boolean
	description: string
	leetcode_url: string
	notification_count: number
	next_notification: string
	created_at: string
}

export interface ListTasks {
	list: Task[]
	total_pages: number
	has_prev: boolean
	has_next: boolean
}

export interface CreateTask {
	title: string
	leetcode_url: string
}

export interface UpdateTask {
	title?: string
	description?: string
	is_active?: boolean
	leetcode_url?: string
}

export class TaskApi extends Api {
	createTask<T = Task>(req: CreateTask): Promise<T> {
		return this.post<T>(`/task`, req)
	}

	getTask<T = Task>(taskId: string): Promise<T> {
		return this.get<T>(`/task/${taskId}`)
	}

	listTasks<T = ListTasks>(page: number, size: number): Promise<T> {
		const query = new URLSearchParams([
			['page', String(page)],
			['size', String(size)],
		])

		return this.get<T>(`/tasks?${query.toString()}`)
	}

	updateTask<T = Task>(taskId: string, req: UpdateTask): Promise<T> {
		return this.patch<T>(`/task/${taskId}`, req)
	}

	deleteTask(taskId: string): Promise<void> {
		return this.delete<void>(`/task/${taskId}`)
	}
}
