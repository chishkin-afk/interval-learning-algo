import Api from './api'

export interface User {
	id: string
	username: string
	email: string
	created_at: string
	updated_at: string
}

export interface Register {
	username: string
	email: string
	password: string
}

export interface Login {
	email: string
	password: string
}

export interface UpdateUser {
	username?: string
}

export class AuthApi extends Api {
	register(req: Register): Promise<void> {
		return this.post<void>(`/auth/register`, req)
	}

	login(req: Login): Promise<void> {
		return this.post<void>(`/auth/login`, req)
	}

	// REQUIRES: cookie token
	me<T = User>(): Promise<T> {
		return this.get<T>(`/auth/me`)
	}

	// REQUIRES: cookie token
	updateUser<T = User>(req: UpdateUser): Promise<T> {
		return this.patch(`/auth/user`, req)
	}

	// REQUIRES: cookie token
	logout(): Promise<void> {
		return this.delete<void>(`/auth/logout`)
	}

	// REQUIRES: cookie token
	deleteUser(): Promise<void> {
		return this.delete<void>(`/auth/user`)
	}
}
