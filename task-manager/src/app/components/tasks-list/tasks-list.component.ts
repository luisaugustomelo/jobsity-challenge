import { Component, Input, OnInit, ViewChild, ElementRef } from '@angular/core';
import { Task } from '../../constants/tasks.interface';

@Component({
  selector: 'app-tasks-list',
  templateUrl: './tasks-list.component.html',
})
export class TasksListComponent implements OnInit {
  @Input() tasks: Task[];

  @ViewChild('taskInput', { static: false }) taskInput: ElementRef;

  constructor() {}

  ngOnInit(): void {}

  removeTask(task: Task) {
    const taskIndex = this.tasks.indexOf(task);
    if (taskIndex !== -1) {
      this.tasks.splice(taskIndex, 1);
    }
  }

  toggleCompleted(task: Task) {
    task.completed = !task.completed;
  }

  editTask(index: number) {
    const task = this.tasks[index];

    if (task.isEditing) {
      this.saveTaskName(task); 
    } else {
      task.isEditing = true; 
      setTimeout(() => {
        if (this.taskInput) {
          this.taskInput.nativeElement.focus(); 
        }
      }, 0);
    }
  }

  saveTaskName(task: Task) {
    task.isEditing = false; 
  }
}
