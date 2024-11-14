"use client";

import { FormInput } from "@/components/form_input";
import { useYupForm } from "@/hooks/useYupForm";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
} from "@mui/material";
import { useCallback, useState } from "react";
import { object, string } from "yup";

const MessageSchema = object({
  message: string(),
});

export const MessageForm = () => {
  const { control } = useYupForm(MessageSchema);
  const [open, setOpen] = useState(false);
  const handleOpen = useCallback(() => setOpen(true), []);
  const handleClose = useCallback(() => setOpen(false), []);

  return (
    <>
      <Button onClick={handleOpen}>Open modal</Button>
      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <DialogTitle id="scroll-dialog-title">Subscribe</DialogTitle>
        <DialogContent dividers>
          <FormInput control={control} name="message" label="Message" />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleClose}>Subscribe</Button>
        </DialogActions>
      </Dialog>
    </>
  );
};
