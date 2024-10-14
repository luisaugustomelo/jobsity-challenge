export interface Task {
    id: number,
    name: string;
    status: string;
    completed: boolean;
    isEditing?: boolean;
  }