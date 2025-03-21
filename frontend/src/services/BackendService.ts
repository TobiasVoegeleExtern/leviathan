import axios, { AxiosInstance, AxiosResponse } from "axios";

class BackendService {
  private axiosInstance: AxiosInstance;

  constructor() {
    this.axiosInstance = axios.create({
      baseURL: "http://localhost:8080",
      timeout: 5000,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  async getPing(): Promise<string> {
    const response: AxiosResponse<{ message: string }> = await this.axiosInstance.get("/ping");
    return response.data.message;
  }

  
}

export default new BackendService();