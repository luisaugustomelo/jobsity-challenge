import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms'; 
import { TasksFormComponent } from './tasks-form.component';
import { By } from '@angular/platform-browser';  

describe('TasksFormComponent', () => {
  let component: TasksFormComponent;
  let fixture: ComponentFixture<TasksFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TasksFormComponent ],
      imports: [ FormsModule ]  
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TasksFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should add a task', () => {
    component.newTask = 'New Task';
    component.newTaskStatus = 'to do';
    
    component.addTask();
    
    expect(component.tasks.length).toBe(1);
    expect(component.tasks[0].name).toBe('New Task');
    expect(component.tasks[0].status).toBe('to do');
  });

  it('should clear input fields after adding a task', () => {
    component.newTask = 'Task to clear';
    component.newTaskStatus = 'doing';

    component.addTask();
    
    expect(component.newTask).toBe('');  
    expect(component.newTaskStatus).toBe('to do'); 
  });

  it('should not add a task if the name is empty', () => {
    component.newTask = '';
    component.newTaskStatus = 'to do';
    
    component.addTask();
    
    expect(component.tasks.length).toBe(0);  
  });

  it('should render added tasks in the template', () => {
    component.newTask = 'Task Render';
    component.newTaskStatus = 'to do';
    component.addTask();

    fixture.detectChanges();  

    const taskElements = fixture.debugElement.queryAll(By.css('li'));
    expect(taskElements.length).toBe(1);
    expect(taskElements[0].nativeElement.textContent).toContain('Task Render');
  });
});
