"use client";

import { FormInput } from "@/components/FormInput";
import { useYupForm } from "@/hooks/useYupForm";
import { object, string } from "yup";

const MessageSchema = object({
  message: string(),
});

export const MessageForm = () => {
  const { control } = useYupForm(MessageSchema);
  return (
    <>
      <FormInput control={control} name="message" label="Message" />
    </>
  );
};
