import { z } from "zod";

export const JWT_SECRET = process.env.AUTH_SECRET;

export const FormSchemaUser = z.object({
  id: z.string(),
  first_name: z
    .string({
      required_error: "Please enter the first name.",
    })
    .min(1, { message: "First name must be at least 1 character." }),
  last_name: z
    .string({
      required_error: "Please enter the last name.",
    })
    .min(1, { message: "last name must be at least 1 character." }),
  email: z
    .string({
      required_error: "Please enter the email.",
    })
    .min(1, { message: "email must be at least 1 character." }),
});

export const CreateUser = FormSchemaUser.omit({ id: true });
