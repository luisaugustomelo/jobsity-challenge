import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { Task } from '../constants/tasks.interface';
import { ApiService } from './api-handler.service';

@Injectable({
  providedIn: 'root',
})
export class TaskService {
  private tasksSubject: BehaviorSubject<Task[]> = new BehaviorSubject<Task[]>([]);
  tasks$: Observable<Task[]> = this.tasksSubject.asObservable();

  constructor(private apiService: ApiService) {}

  loadTasks(): void {
    this.apiService.getAllTasks().subscribe((tasks: Task[]) => {
      this.tasksSubject.next(tasks);
    });
  }

  addTask(task: Task): void {
    this.apiService.addTask(task).subscribe((newTask: Task) => {
      const currentTasks = this.tasksSubject.value;
      this.tasksSubject.next([...currentTasks, newTask]);
    });
  }
}
