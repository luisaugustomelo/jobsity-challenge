import { Component, Input, OnInit } from '@angular/core';
import { ApiService } from '../../services/api-handler.service';
import { Task } from '../../constants/tasks.interface';
import { TaskService } from 'src/app/services/task-handler.service';

@Component({
  selector: 'app-tasks-list',
  templateUrl: './tasks-list.component.html',
})
export class TasksListComponent implements OnInit {
  @Input() tasks: Task[] = [];

  constructor(private apiService: ApiService, private taskService: TaskService) {}

  ngOnInit(): void {
    this.taskService.tasks$.subscribe((tasks: Task[]) => {
      this.tasks = tasks;
    });

    this.taskService.loadTasks();
  }

  removeTask(task: Task) {
    this.apiService.removeTask(task.id).subscribe(
      () => {
        const taskIndex = this.tasks.indexOf(task);
        if (taskIndex !== -1) {
          this.tasks.splice(taskIndex, 1);
        }
      },
      (error) => {
        console.error('Error to remove task:', error);
      }
    );
  }

  editTask(index: number) {
    const task = this.tasks[index];
    task.isEditing = true;
  }
  
  saveTask(task: Task) {
    this.apiService.saveTask(task.id, task.description, task.status).subscribe(
      () => {
        task.isEditing = false;
        task.description = task.description;
        task.status = task.status;
      },
      (error) => console.error('Error to save task:', error)
    );
  }
}
