export interface Task {
	id: string
	user_id: string
	title: string
	is_active: boolean
	description: string
	leetcode_url: string
	notification_count: number
	next_notification: Date
	created_at: Date
}
