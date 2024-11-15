"use client";

import { FC, useCallback, useState } from "react";
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Fab,
} from "@mui/material";
import AddIcon from "@mui/icons-material/add";
import { object, string } from "yup";

import { FormInput } from "@cardboardrobots/form";
import { useYupForm } from "@/hooks/use_yup_form";
import { createTodo, Todo } from "@/services/todo_service";

const TodoSchema = object({
  note: string().required(),
});

export const CreateTodoForm: FC<{
  onCreate: (todo: Todo) => void;
}> = ({ onCreate }) => {
  const { control, reset, handleSubmit } = useYupForm(TodoSchema);

  const [open, setOpen] = useState(false);

  const handleOpen = useCallback(() => setOpen(true), []);

  const handleClose = useCallback(() => {
    setOpen(false);
    reset();
  }, [reset]);

  const handleCreate = handleSubmit(
    useCallback(
      async ({ note }) => {
        const id = await createTodo({ note });
        setOpen(false);
        reset();
        onCreate({
          due: new Date(),
          id,
          note,
          status: false,
          version: 0,
        });
      },
      [reset, onCreate]
    )
  );

  return (
    <>
      <Box height={56 + 16} />
      <Fab
        color="primary"
        variant="extended"
        onClick={handleOpen}
        sx={{ position: "fixed", right: 0, bottom: 0, margin: 2 }}
      >
        <AddIcon sx={{ mr: 1 }} />
        Create todo
      </Fab>
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
        <DialogTitle id="scroll-dialog-title">Create todo</DialogTitle>
        <DialogContent dividers>
          <Box display="flex" flexDirection="column" gap={2}>
            <FormInput control={control} name="note" label="Note" />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button type="submit">Create</Button>
        </DialogActions>
      </Dialog>
    </>
  );
};
