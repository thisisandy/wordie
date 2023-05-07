import axios from 'axios'
interface LoginResp {
  token: string
  code: number
  message: string
  expire: string
}
export class Request {
  client =  axios.create({
    baseURL:  import.meta.env.DEV? 'http://127.0.0.1:8080' : 'http://localhost:8080',
    timeout: 5000,
    headers: {
      'Content-Type': 'application/json',
    },
  })
  constructor() {
    
  }
  login(user: { email: string, password: string })  {
    return this.client.post<LoginResp>('/login', user)
  }
register(user: { email: string, password: string , username: string}) {
    return this.client.post('/register', user)
  }
}
