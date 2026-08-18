type Method = 'GET' | 'POST' | 'PATCH' | 'DELETE' | 'PUT'

const METHODS_WITH_BODY: Set<Method> = new Set(['POST', 'PATCH', 'PUT'])

export class ApiError extends Error {
	readonly status: number
	readonly payload: unknown

	constructor(status: number, payload: unknown) {
		super(`API error ${status}`)
		this.name = 'ApiError'
		this.status = status
		this.payload = payload
	}
}

class Api {
	_addr: string
	_headers: HeadersInit

	constructor(addr?: string) {
		this._addr = addr ?? 'http://localhost:8090'
		this._headers = {
			'Content-Type': 'application/json',
		}
	}

	async request(path: string, method: Method, body?: object): Promise<Response> {
		const req: RequestInit = {
			method: method,
			headers: this._headers,
		}

		if (METHODS_WITH_BODY.has(method) && body !== undefined) {
			req.body = JSON.stringify(body)
		}

		const resp = await fetch(`${this._addr}/${path}`, req)
		if (!resp.ok) {
			let payload: unknown
			try {
				payload = await resp.json()
			} catch (error) {
				console.error("Can't read body of response:", error)
			}

			throw new ApiError(resp.status, payload)
		}

		return resp
	}
}

export default Api
