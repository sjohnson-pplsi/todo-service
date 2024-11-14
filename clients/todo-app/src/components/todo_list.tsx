"use client";

import { FC, useCallback, useState } from "react";

import { Todo } from "@/services/todo_service";
import { Container, List } from "@mui/material";

import { TodoListItem } from "@/components/todo_list_item";

import { CreateTodoForm } from "./create_todo_form";

export const TodoList: FC<{ todos: Todo[] }> = ({ todos }) => {
  const [todoList, setTodoList] = useState(todos);
  const handleCreate = useCallback(
    (todo: Todo) => {
      setTodoList([...todoList, todo]);
    },
    [todoList]
  );
  return (
    <>
      <Container>
        <CreateTodoForm onCreate={handleCreate} />
      </Container>
      <List>
        {todoList.map((todo) => (
          <TodoListItem key={todo.id} todo={todo} />
        ))}
      </List>
    </>
  );
};
