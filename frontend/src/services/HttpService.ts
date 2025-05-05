class HttpService {
    // Helper method for making POST requests
    async post<T>(url: string, data: any): Promise<T> {
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
  
      if (!response.ok) {
        throw new Error(`Failed to POST data to ${url}: ${response.statusText}`);
      }
  
      return response.json();
    }
  
    // Helper method for making GET requests
    async get<T>(url: string): Promise<T> {
      const response = await fetch(url, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (!response.ok) {
        throw new Error(`Failed to GET data from ${url}: ${response.statusText}`);
      }
  
      return response.json();
    }
  
    // Helper method for making PUT requests
    async put<T>(url: string, data: any): Promise<T> {
      const response = await fetch(url, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
  
      if (!response.ok) {
        throw new Error(`Failed to PUT data to ${url}: ${response.statusText}`);
      }
  
      return response.json();
    }
  
    // Helper method for making DELETE requests
    async delete(url: string): Promise<void> {
      const response = await fetch(url, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (!response.ok) {
        throw new Error(`Failed to DELETE data from ${url}: ${response.statusText}`);
      }
    }
  }
  
  export const httpService = new HttpService();
  