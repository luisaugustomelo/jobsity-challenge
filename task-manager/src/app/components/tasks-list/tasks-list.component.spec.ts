import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { TasksListComponent } from './tasks-list.component';
import { By } from '@angular/platform-browser'; 

describe('TasksListComponent', () => {
  let component: TasksListComponent;
  let fixture: ComponentFixture<TasksListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TasksListComponent ],
      imports: [ FormsModule ] 
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TasksListComponent);
    component = fixture.componentInstance;

    component.tasks = [
      { name: 'Task 1', status: 'to do', completed: false, isEditing: false },
      { name: 'Task 2', status: 'to do', completed: false, isEditing: false },
      { name: 'Task 3', status: 'to do', completed: false, isEditing: false }
    ];

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render tasks', () => {
    const taskElements = fixture.debugElement.queryAll(By.css('li'));
    expect(taskElements.length).toBe(2); 
    expect(taskElements[0].nativeElement.textContent).toContain('Task 1');
    expect(taskElements[1].nativeElement.textContent).toContain('Task 2');
    expect(taskElements[1].nativeElement.textContent).toContain('Task 3');
  });

  it('should enter edit mode when edit icon is clicked', () => {
    const editButton = fixture.debugElement.queryAll(By.css('span'))[2]; 
    editButton.triggerEventHandler('click', null);
    fixture.detectChanges();

    const inputElement = fixture.debugElement.query(By.css('input'));
    expect(inputElement).toBeTruthy(); 
    expect(inputElement.nativeElement.value).toBe('Task 1'); 
  });

  it('should save edited task name on blur', () => {
    const editButton = fixture.debugElement.queryAll(By.css('span'))[2]; 
    editButton.triggerEventHandler('click', null);
    fixture.detectChanges();

    const inputElement = fixture.debugElement.query(By.css('input'));
    inputElement.nativeElement.value = 'Updated Task 1'; 
    inputElement.nativeElement.dispatchEvent(new Event('input')); 
    inputElement.triggerEventHandler('blur', null);
    fixture.detectChanges();

    const taskElement = fixture.debugElement.query(By.css('li'));
    expect(taskElement.nativeElement.textContent).toContain('Updated Task 1'); 
  });

  it('should toggle completed status when toggle button is clicked', () => {
    const toggleButton = fixture.debugElement.queryAll(By.css('span'))[1]; 
    toggleButton.triggerEventHandler('click', null);
    fixture.detectChanges();

    const taskElement = fixture.debugElement.query(By.css('li'));
    expect(taskElement.classes['line-through']).toBeTrue(); 
  });

  it('should remove a task when remove icon is clicked', () => {
    const removeButton = fixture.debugElement.queryAll(By.css('span'))[0]; 
    removeButton.triggerEventHandler('click', null);
    fixture.detectChanges();

    const taskElements = fixture.debugElement.queryAll(By.css('li'));
    expect(taskElements.length).toBe(1);
    expect(taskElements[0].nativeElement.textContent).toContain('Task 2'); 
  });
});
