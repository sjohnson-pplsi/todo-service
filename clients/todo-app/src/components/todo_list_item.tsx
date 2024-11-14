"use client";
import { FC, useCallback, useState } from "react";
import { Box, Checkbox, Divider, ListItem, ListItemText } from "@mui/material";

import { completeTodo, resetTodo, Todo } from "@/services/todo_service";
import { ChangeNoteForm } from "./change_note_form";

export const TodoListItem: FC<{
  todo: Todo;
}> = ({ todo }) => {
  const [complete, setComplete] = useState(todo.status);
  const [note, setNote] = useState(todo.note);

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

  const handleNote = useCallback((note: string) => {
    setNote(note);
  }, []);

  return (
    <>
      <ListItem
        key={todo.id}
        secondaryAction={
          <Box>
            <ChangeNoteForm todo={todo} onChange={handleNote} />
            <Checkbox
              edge="end"
              onChange={handleChange}
              checked={complete}
              inputProps={{
                "aria-labelledby": `checkbox-list-secondary-label-${todo.id}`,
              }}
            />
          </Box>
        }
      >
        <ListItemText primary={note} secondary="" />
      </ListItem>
      <Divider />
    </>
  );
};
