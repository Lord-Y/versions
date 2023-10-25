import axios from 'axios'
import type { AxiosInstance } from 'axios'

const axiosClient: AxiosInstance = axios.create({
  headers: {
    'Content-Type': 'application/json',
  },
})

class axiosService {
  async genericGet(headers: any, url: string, params: any): Promise<any> {
    // console.log('headers', headers, 'url', url, 'params', params)
    return await axiosClient.get(url, {
      headers: headers,
      params: params,
    })
  }
}

export default new axiosService()
