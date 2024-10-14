export interface Task {
    id: number,
    description: string;
    status: string;
    completed: boolean;
    isEditing?: boolean;
  }