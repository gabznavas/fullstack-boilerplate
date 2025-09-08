import { Todo } from "@/types/todo";

export const API_URL = process.env.NEXT_PUBLIC_API_URL;


export const createTodo = async (title: string): Promise<Todo> => {
  const response = await fetch(`${API_URL}/todos`, {
    method: "POST",
    body: JSON.stringify({ title }),
  });
  if (!response.ok) {
    throw new Error("Failed to create todo");
  }
  return response.json();
};

export const getAllTodos = async (): Promise<Todo[]> => {
  const response = await fetch(`${API_URL}/todos`, {
    method: "GET",
  });
  if (!response.ok) {
    throw new Error("Failed to get todos");
  }
  return response.json();
};