import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Task } from '../constants/tasks.interface';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private apiUrl = `${environment.apiUrl}/task`;

  constructor(private http: HttpClient) {}

  removeTask(taskId: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${taskId}/delete`);
  }

  saveTask(taskId: number, description: string, status: string): Observable<Task> {
    return this.http.patch<Task>(`${this.apiUrl}/${taskId}/update`, { description, status });
  }

  addTask(task: Task): Observable<Task> {
    return this.http.post<Task>(`${this.apiUrl}/create`, {
      description: task.description,
      status: task.status
    });
  }

  getAllTasks(): Observable<Task[]> {
    return this.http.get<Task[]>(`${this.apiUrl}/allTasks`);
  }
}