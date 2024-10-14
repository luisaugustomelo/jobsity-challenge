import { Component, Input, OnInit } from '@angular/core';
import { ApiService } from '../../services/api-handler.service';
import { Task } from '../../constants/tasks.interface';

@Component({
  selector: 'app-tasks-list',
  templateUrl: './tasks-list.component.html',
})
export class TasksListComponent implements OnInit {
  @Input() tasks: Task[] = [];

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {}

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

  toggleCompleted(task: Task) {
    this.apiService.toggleCompleted(task.id).subscribe(
      () => {
        task.completed = !task.completed;
        task.status = task.completed ? 'completed' : 'to do';
      },
      (error) => console.error('Error to update status:', error)
    );
  }

  saveTask(task: Task) {
    this.apiService.saveTask(task.id, task.name, task.status).subscribe(
      (updatedTask: any) => {
        console.log(updatedTask)
        task.isEditing = false;
        task.name = updatedTask.name;
        task.status = updatedTask.status;
      },
      (error) => console.error('Error to save task:', error)
    );
  }
}
