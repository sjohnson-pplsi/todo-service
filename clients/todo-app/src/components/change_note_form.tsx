"use client";

import { FormInput } from "@cardboardrobots/form";

import { useYupForm } from "@/hooks/use_yup_form";
import { changeNote, Todo } from "@/services/todo_service";
import EditIcon from "@mui/icons-material/Edit";
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
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
  const { control, reset, handleSubmit } = useYupForm(TodoSchema, {
    defaultValues: { note: todo.note },
  });
  const [open, setOpen] = useState(false);

  const handleOpen = useCallback(() => setOpen(true), []);

  const handleClose = useCallback(() => {
    setOpen(false);
    reset();
  }, [reset]);

  const handleCreate = handleSubmit(
    useCallback(
      async ({ note }) => {
        await changeNote(todo.id, note);
        setOpen(false);
        onChange(note);
      },
      [todo.id, onChange]
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
        PaperProps={{
          component: "form",
          onSubmit: handleCreate,
          sx: { width: "80%", maxWidth: "500px" },
        }}
      >
        <DialogTitle id="scroll-dialog-title">Change note</DialogTitle>
        <DialogContent dividers>
          <Box display="flex" flexDirection="column" gap={2}>
            <FormInput control={control} name="note" label="Note" />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button type="submit">Change</Button>
        </DialogActions>
      </Dialog>
    </>
  );
};
