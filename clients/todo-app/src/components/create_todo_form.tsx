"use client";

import { FormInput } from "@/components/form_input";
import { useYupForm } from "@/hooks/use_yup_form";
import { createTodo, Todo } from "@/services/todo_service";
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
} from "@mui/material";
import { FC, useCallback, useState } from "react";
import { object, string } from "yup";

const TodoSchema = object({
  note: string().required(),
});

export const CreateTodoForm: FC<{
  onCreate: (todo: Todo) => void;
}> = ({ onCreate }) => {
  const { control, handleSubmit } = useYupForm(TodoSchema);
  const [open, setOpen] = useState(false);
  const handleOpen = useCallback(() => setOpen(true), []);
  const handleClose = useCallback(() => setOpen(false), []);
  const handleCreate = handleSubmit(
    useCallback(
      async ({ note }) => {
        const id = await createTodo({ note });
        setOpen(false);
        onCreate({
          due: new Date(),
          id,
          note,
          status: false,
          version: 0,
        });
      },
      [onCreate, handleSubmit]
    )
  );

  return (
    <>
      <Box display="flex" justifyContent="flex-end" mt={2}>
        <Button onClick={handleOpen} variant="contained">
          Create todo
        </Button>
      </Box>
      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <form onSubmit={handleCreate}>
          <DialogTitle id="scroll-dialog-title">Create todo</DialogTitle>
          <DialogContent dividers>
            <FormInput control={control} name="note" label="Note" />
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button type="submit">Create</Button>
          </DialogActions>
        </form>
      </Dialog>
    </>
  );
};
