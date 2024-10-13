import { Component } from '@angular/core';
import { Task } from '../../constants/tasks.interface';

@Component({
  selector: 'app-tasks-form',
  templateUrl: './tasks-form.component.html',
})
export class TasksFormComponent {
  newTask: string = ''; 
  newTaskStatus: string = 'to do'; 

  tasks: Task[] = []; 

  addTask() {
    if (this.newTask.trim()) {
      const newTask: Task = {
        name: this.newTask,
        status: this.newTaskStatus, 
        completed: false,
        isEditing: false 
      };
      this.tasks.push(newTask);
      this.newTask = ''; 
      this.newTaskStatus = 'to do'; 
    }
  }
}
