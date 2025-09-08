"use client"

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Table, TableCaption, TableHeader, TableRow, TableHead, TableBody, TableCell } from "@/components/ui/table";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { useForm } from "react-hook-form";
import { useEffect, useState } from "react";
import { Todo } from "@/types/todo";
import { createTodo, getAllTodos } from "@/services/requests";

const todoFormSchema = yup.object().shape({
  title: yup.string().required(),
})

type TodoFormInputs = yup.InferType<typeof todoFormSchema>;

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([]);

  const form = useForm<TodoFormInputs>({
    resolver: yupResolver(todoFormSchema),
  })

  useEffect(() => {
    console.log(process.env.NEXT_PUBLIC_API_URL);
    const fetchTodos = async () => {
      const todos = await getAllTodos();
      setTodos(todos);
    }
    fetchTodos();
  }, []);

  const onSubmit = async (data: TodoFormInputs) => {
    const todo = await createTodo(data.title);
    debugger
    setTodos([todo, ...todos]);
  }

  return (
    <div className="flex flex-col gap-2 m-4">
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <div className="flex gap-2">
          <Input {...form.register("title")} />
          <Button>Adicionar</Button>
        </div>

        <div className="mt-4">
          <Table>
            <TableCaption>A list of your recent invoices.</TableCaption>
            <TableHeader>
              <TableRow>
                <TableHead className="w-[100px]">ID</TableHead>
                <TableHead>Título</TableHead>
                <TableHead>Status</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {
                todos.map(todo => (
                  <TableRow key={todo.id}>
                    <TableCell>{todo.id}</TableCell>
                    <TableCell>{todo.title}</TableCell>
                    <TableCell>{todo.completed ? "Completo" : "Incompleto"}</TableCell>
                  </TableRow>
                ))
              }
            </TableBody>
          </Table>

        </div>

      </form>
    </div >
  );
}
