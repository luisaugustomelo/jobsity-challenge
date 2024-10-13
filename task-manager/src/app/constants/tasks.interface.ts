export interface Task {
    name: string;
    status: string;
    completed: boolean;
    isEditing?: boolean; // Novo campo para controlar o modo de ediçã
  }