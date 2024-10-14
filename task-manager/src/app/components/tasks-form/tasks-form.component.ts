import { Component } from '@angular/core';
import { Task } from '../../constants/tasks.interface';
import { ApiService } from 'src/app/services/api-handler.service';

@Component({
  selector: 'app-tasks-form',
  templateUrl: './tasks-form.component.html',
})
export class TasksFormComponent {
  newTaskId: 0;
  newTask: string = ''; 
  newTaskStatus: string = 'to do'; 

  tasks: Task[] = []; 

  constructor(private apiService: ApiService) {} 

  addTask() {
    if (this.newTask.trim()) {
      const task: Task = {
        id: this.newTaskId,
        description: this.newTask,
        status: this.newTaskStatus,
        completed: false,
        isEditing: false,
      };
      

      this.apiService.addTask(task).subscribe(
        (addedTask: any) => {
          const task: Task = {
            id: addedTask.id,
            description: addedTask.description,
            status: addedTask.status, 
            completed: addedTask.status === "completed" ? true : false,
            isEditing: false 
          }
          console.log(task)
          this.tasks.push(task); 
        },
        (error) => console.error('Error to add task:', error)
      );
    }
  }
}
