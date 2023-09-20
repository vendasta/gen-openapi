import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getListFields(): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/listFields`);
  }

  saveFields(data: any): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/saveFields`, data);
  }
}
