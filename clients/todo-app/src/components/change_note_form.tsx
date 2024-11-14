"use client";

import { FormInput } from "@/components/form_input";
import { useYupForm } from "@/hooks/useYupForm";
import { changeNote, Todo } from "@/services/todo_service";
import EditIcon from "@mui/icons-material/Edit";
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Icon,
  IconButton,
} from "@mui/material";
import { FC, useCallback, useState } from "react";
import { object, string } from "yup";

const TodoSchema = object({
  note: string().required(),
});

export const ChangeNoteForm: FC<{
  todo: Todo;
  onChange: (note: string) => void;
}> = ({ todo, onChange }) => {
  const { control, handleSubmit } = useYupForm(TodoSchema, {
    defaultValues: { note: todo.note },
  });
  const [open, setOpen] = useState(false);
  const handleOpen = useCallback(() => setOpen(true), []);
  const handleClose = useCallback(() => setOpen(false), []);
  const handleCreate = handleSubmit(
    useCallback(
      async ({ note }) => {
        await changeNote(todo.id, note);
        setOpen(false);
        onChange(note);
      },
      [todo.id, onChange, handleSubmit]
    )
  );

  return (
    <>
      <IconButton onClick={handleOpen}>
        <EditIcon fontSize="small" />
      </IconButton>
      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <form onSubmit={handleCreate}>
          <DialogTitle id="scroll-dialog-title">Change note</DialogTitle>
          <DialogContent dividers>
            <FormInput control={control} name="note" label="Note" />
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button type="submit">Change</Button>
          </DialogActions>
        </form>
      </Dialog>
    </>
  );
};
