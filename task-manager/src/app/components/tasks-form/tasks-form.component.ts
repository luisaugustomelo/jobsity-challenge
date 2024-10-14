import { Component, EventEmitter, Output } from '@angular/core';
import { Task } from '../../constants/tasks.interface';
import { ApiService } from 'src/app/services/api-handler.service';
import { TaskService } from 'src/app/services/task-handler.service';

@Component({
  selector: 'app-tasks-form',
  templateUrl: './tasks-form.component.html',
})
export class TasksFormComponent {
  newTaskId: 0;
  newTask: string = ''; 
  newTaskStatus: string = 'to do'; 

  @Output() taskAdded = new EventEmitter<Task>();

  tasks: Task[] = []; 

  constructor(private taskService: TaskService) {} 

  addTask() {
    if (this.newTask.trim()) {
      const task: Task = {
        id: this.newTaskId,
        description: this.newTask,
        status: this.newTaskStatus,
        completed: false,
        isEditing: false,
      };

      this.taskService.addTask(task);
    }
  }
}
