"use client";
import { FC, useCallback, useState } from "react";
import { Checkbox, Divider, ListItem, ListItemText } from "@mui/material";

import { completeTodo, resetTodo, Todo } from "@/services/todo_service";

export const TodoListItem: FC<{
  todo: Todo;
}> = ({ todo }) => {
  const [complete, setComplete] = useState(todo.status);
  const handleChange: React.ChangeEventHandler<HTMLInputElement> = useCallback(
    async (event) => {
      if (event.target.checked) {
        await completeTodo(todo.id);
        setComplete(true);
      } else {
        await resetTodo(todo.id);
        setComplete(false);
      }
    },
    [todo.id]
  );
  return (
    <>
      <ListItem
        key={todo.id}
        secondaryAction={
          <Checkbox
            edge="end"
            onChange={handleChange}
            checked={complete}
            inputProps={{
              "aria-labelledby": `checkbox-list-secondary-label-${todo.id}`,
            }}
          />
        }
      >
        <ListItemText primary={todo.note} secondary="" />
      </ListItem>
      <Divider />
    </>
  );
};
